package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
	"time"

	"github.com/disrupt-ed/doctrine-cli/internal/detect"
	"github.com/disrupt-ed/doctrine-cli/internal/manifest"
)

// siteURL is doctrine.codedynamic.com. DOCTRINE_URL points the CLI at another server,
// for development.
func siteURL() string {
	if u := os.Getenv("DOCTRINE_URL"); u != "" {
		return strings.TrimRight(u, "/")
	}
	return "https://doctrine.codedynamic.com"
}

var handlePattern = regexp.MustCompile(`^[a-z0-9]{4,32}$`)

// installCmd fetches a manifest made on doctrine.codedynamic.com and sets the project up
// with it. The handle isn't written anywhere.
func installCmd(ctx context.Context, root string, args []string, in io.Reader, out io.Writer) error {
	if len(args) != 1 || !handlePattern.MatchString(args[0]) {
		return fmt.Errorf("usage: doctrine install HANDLE (from doctrine.codedynamic.com)")
	}
	if _, err := os.Stat(filepath.Join(root, manifest.Path)); err == nil {
		return fmt.Errorf("%s already exists. Remove it first to install another configuration", manifest.Path)
	}

	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, siteURL()+"/d/"+args[0], nil)
	if err != nil {
		return err
	}
	req.Header.Set("Accept", "application/yaml")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("fetch configuration: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusNotFound {
		return fmt.Errorf("no configuration %q on %s", args[0], siteURL())
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("fetch configuration: %s", resp.Status)
	}
	data, err := io.ReadAll(io.LimitReader(resp.Body, 64*1024))
	if err != nil {
		return err
	}
	m, err := manifest.Parse(data)
	if err != nil {
		return err
	}
	if err := writeManifest(root, m.Marshal()); err != nil {
		return err
	}
	fmt.Fprintf(out, "Wrote %s (%s)\n", manifest.Path, strings.Join(m.Extends, ", "))
	if err := regenerate(ctx, root, out); err != nil {
		return err
	}
	r, err := detect.Detect(root)
	if err == nil {
		askStats(root, r, in, out)
	}
	return nil
}

// stat is everything an anonymous usage stat contains.
type stat struct {
	Languages  []string `json:"languages"`
	Frameworks []string `json:"frameworks"`
	OS         string   `json:"os"`
	Arch       string   `json:"arch"`
	Agent      string   `json:"agent"`
}

func newStat(r detect.Result) stat {
	s := stat{OS: runtime.GOOS, Arch: runtime.GOARCH, Agent: "claude-code", Languages: []string{}, Frameworks: []string{}}
	for _, t := range r.Technologies {
		switch t {
		case "ruby", "go":
			s.Languages = append(s.Languages, t)
		case "rails":
			s.Frameworks = append(s.Frameworks, t)
		}
	}
	return s
}

// statsDisabled reports whether the environment turns stats off.
func statsDisabled() bool {
	if v := os.Getenv("DO_NOT_TRACK"); v != "" && v != "0" {
		return true
	}
	return os.Getenv("DOCTRINE_TELEMETRY") == "0"
}

// consentPath stores the answer, once per machine.
func consentPath() (string, error) {
	dir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(dir, "doctrine", "stats"), nil
}

// askStats asks once whether to share anonymous usage stats, showing exactly
// what would be sent, then sends if the answer is yes. It never fails the
// command and never asks when there is no terminal.
func askStats(root string, r detect.Result, in io.Reader, out io.Writer) {
	if statsDisabled() {
		return
	}
	path, err := consentPath()
	if err != nil {
		return
	}
	s := newStat(r)
	answer, err := os.ReadFile(path)
	if errors.Is(err, os.ErrNotExist) {
		if !interactive(in) {
			return
		}
		fmt.Fprintf(out, "\nShare anonymous usage stats with doctrine.codedynamic.com? This is what we'd send:\n\n")
		fmt.Fprintf(out, "  language:  %s\n  framework: %s\n  os/arch:   %s/%s\n  agent:     %s\n\n", orNone(s.Languages), orNone(s.Frameworks), s.OS, s.Arch, s.Agent)
		fmt.Fprintln(out, "No code, repository name, path or dependencies. Turn off anytime with DOCTRINE_TELEMETRY=0.")
		yes := confirm(in, out, "")
		answer = []byte("no")
		if yes {
			answer = []byte("yes")
		}
		if os.MkdirAll(filepath.Dir(path), 0o755) == nil {
			os.WriteFile(path, answer, 0o644)
		}
	}
	if strings.TrimSpace(string(answer)) == "yes" {
		sendStat(s)
	}
}

func sendStat(s stat) {
	body, err := json.Marshal(s)
	if err != nil {
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, siteURL()+"/stats", bytes.NewReader(body))
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	resp.Body.Close()
}

// interactive reports whether in is a terminal the developer can answer from.
var interactive = func(in io.Reader) bool {
	f, ok := in.(*os.File)
	if !ok {
		return false
	}
	info, err := f.Stat()
	return err == nil && info.Mode()&os.ModeCharDevice != 0
}
