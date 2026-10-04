// Package dependencies gathers raw facts about a Ruby gem or a Go module
// from its registry and, when the source lives on GitHub, from GitHub.
// It reports facts only. Judging them is left to the reader.
package dependencies

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

// Ecosystem names a package registry.
type Ecosystem string

const (
	RubyGems Ecosystem = "rubygems"
	Go       Ecosystem = "go"
)

// Guess returns Go for names that look like module paths
// (a dot before the first slash, e.g. github.com/x/y), else RubyGems.
func Guess(name string) Ecosystem {
	host, _, found := strings.Cut(name, "/")
	if found && strings.Contains(host, ".") {
		return Go
	}
	return RubyGems
}

// TokenFromEnv returns GITHUB_TOKEN, then GH_TOKEN, or "".
func TokenFromEnv() string {
	if t := os.Getenv("GITHUB_TOKEN"); t != "" {
		return t
	}
	return os.Getenv("GH_TOKEN")
}

// Client fetches facts. Zero values fall back to the public services.
type Client struct {
	HTTP        *http.Client
	RubyGemsURL string // default https://rubygems.org
	GoProxyURL  string // default https://proxy.golang.org
	GitHubURL   string // default https://api.github.com
	GitHubToken string
	Now         func() time.Time // default time.Now
}

// Report holds the facts gathered about one dependency.
type Report struct {
	Name      string
	Ecosystem Ecosystem
	Facts     []Fact
	Notes     []string
}

// Fact is one labelled raw value.
type Fact struct {
	Label, Value string
}

func (r *Report) add(label, value string) {
	r.Facts = append(r.Facts, Fact{Label: label, Value: value})
}

func (r *Report) note(format string, args ...any) {
	r.Notes = append(r.Notes, fmt.Sprintf(format, args...))
}

// String renders aligned "label:  value" lines, then notes.
func (r Report) String() string {
	var b strings.Builder
	fmt.Fprintf(&b, "%s (%s)\n", r.Name, r.Ecosystem)
	width := 0
	for _, f := range r.Facts {
		width = max(width, len(f.Label)+1)
	}
	for _, f := range r.Facts {
		fmt.Fprintf(&b, "  %-*s  %s\n", width, f.Label+":", f.Value)
	}
	if len(r.Notes) > 0 {
		b.WriteString("\nNotes:\n")
		for _, n := range r.Notes {
			fmt.Fprintf(&b, "  - %s\n", n)
		}
	}
	return b.String()
}

// Inspect gathers facts about name. It returns an error only when the
// registry lookup itself fails; other failures become notes.
func (c *Client) Inspect(ctx context.Context, eco Ecosystem, name string) (Report, error) {
	r := Report{Name: name, Ecosystem: eco}
	var repo string // "owner/repo" on GitHub, or ""
	var err error
	switch eco {
	case RubyGems:
		repo, err = c.rubyGems(ctx, &r)
	case Go:
		repo, err = c.goModule(ctx, &r)
	default:
		return r, fmt.Errorf("unknown ecosystem %q", eco)
	}
	if err != nil {
		return r, err
	}
	if c.GitHubToken == "" {
		r.note("No GitHub token: GitHub limits unauthenticated requests to 60 per hour.")
	}
	if repo == "" {
		r.note("Source repository not found on GitHub: repository facts not gathered.")
	} else {
		c.gitHubRepo(ctx, &r, repo)
	}
	c.gitHubAdvisories(ctx, &r, eco, name)
	return r, nil
}

func (c *Client) now() time.Time {
	if c.Now != nil {
		return c.Now()
	}
	return time.Now()
}

func orDefault(v, def string) string {
	if v == "" {
		return def
	}
	return strings.TrimSuffix(v, "/")
}

// get fetches url and returns the body and headers. GitHub requests carry
// the token when one is set.
func (c *Client) get(ctx context.Context, url string, gitHub bool) ([]byte, http.Header, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, nil, fmt.Errorf("build request: %w", err)
	}
	req.Header.Set("User-Agent", "doctrine-cli")
	if gitHub {
		req.Header.Set("Accept", "application/vnd.github+json")
		if c.GitHubToken != "" {
			req.Header.Set("Authorization", "Bearer "+c.GitHubToken)
		}
	}
	client := c.HTTP
	if client == nil {
		client = &http.Client{Timeout: 20 * time.Second}
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, nil, fmt.Errorf("GET %s: %w", url, err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(io.LimitReader(resp.Body, 10<<20))
	if err != nil {
		return nil, nil, fmt.Errorf("read %s: %w", url, err)
	}
	if resp.StatusCode != http.StatusOK {
		msg := resp.Status
		if (resp.StatusCode == http.StatusForbidden || resp.StatusCode == http.StatusTooManyRequests) &&
			resp.Header.Get("X-RateLimit-Remaining") == "0" {
			msg += " (rate limit reached)"
		}
		return nil, nil, fmt.Errorf("GET %s: %s", url, msg)
	}
	return body, resp.Header, nil
}

func (c *Client) getJSON(ctx context.Context, url string, gitHub bool, v any) (http.Header, error) {
	body, header, err := c.get(ctx, url, gitHub)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, v); err != nil {
		return nil, fmt.Errorf("decode %s: %w", url, err)
	}
	return header, nil
}

// date renders t as "YYYY-MM-DD (age)".
func date(t, now time.Time) string {
	return t.Format("2006-01-02") + " (" + age(t, now) + ")"
}

// age describes how long ago t was, in days, months or years.
func age(t, now time.Time) string {
	days := int(now.Sub(t).Hours() / 24)
	switch {
	case days < 1:
		return "today"
	case days < 31:
		return plural(days, "day")
	case days < 365:
		return plural(days/30, "month")
	default:
		return plural(days/365, "year")
	}
}

func plural(n int, unit string) string {
	if n == 1 {
		return "1 " + unit + " ago"
	}
	return fmt.Sprintf("%d %ss ago", n, unit)
}

// countAtLeast renders n, with a "+" when the page was full and more exist.
func countAtLeast(n int, more bool) string {
	if more {
		return fmt.Sprintf("%d+", n)
	}
	return fmt.Sprint(n)
}
