// Package compose builds a project's effective doctrine: the official
// doctrines it extends, then its own preferences and overrides.
package compose

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/disrupt-ed/doctrine-cli/internal/doctrine"
	"github.com/disrupt-ed/doctrine-cli/internal/manifest"
)

// Effective is everything that influences the agent, in order.
type Effective struct {
	Doctrines   []doctrine.Doctrine
	Preferences []string   // sentences, one per preference set in the manifest
	Overrides   []Override // files from .doctrine/local/
}

// Override is one of the project's own Markdown files.
type Override struct {
	Name string
	Body string
}

// Compose resolves the manifest against a release and reads the project's
// overrides from root.
func Compose(release fs.FS, m manifest.Manifest, root string) (Effective, error) {
	var e Effective
	docs, err := doctrine.Resolve(release, m.Extends)
	if err != nil {
		return e, err
	}
	e.Doctrines = docs
	for _, k := range manifest.PreferenceOrder {
		if v, ok := m.Engineering[k]; ok {
			e.Preferences = append(e.Preferences, preferenceText[k+"="+v])
		}
	}
	e.Overrides, err = readOverrides(filepath.Join(root, manifest.LocalDir))
	return e, err
}

var preferenceText = map[string]string{
	"kiss=strong":               "Simplicity: strong. Pick the simplest solution that works, even when a more general one looks nicer.",
	"kiss=normal":               "Simplicity: normal. Prefer simple solutions, but some generality is fine when it clearly helps.",
	"yagni=strong":              "Build for today: strong. Build only what the task needs now.",
	"yagni=normal":              "Build for today: normal. Small, cheap preparations for a clearly coming need are fine. Say what they are.",
	"dependencies=conservative": "Dependencies: conservative. Add a new dependency only when writing the code yourself would clearly be worse, and say why first.",
	"dependencies=balanced":     "Dependencies: balanced. Add a well-maintained dependency when it saves real work. Say why.",
	"dependencies=open":         "Dependencies: open. Use well-maintained dependencies freely. List new ones in your summary.",
	"abstractions=late":         "Abstractions: late. Wait for two or three real cases before abstracting.",
	"abstractions=balanced":     "Abstractions: balanced. Abstract when it makes the code clearly easier to read, even with one case.",
}

func readOverrides(dir string) ([]Override, error) {
	entries, err := os.ReadDir(dir)
	if errors.Is(err, fs.ErrNotExist) {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("read overrides: %w", err)
	}
	var out []Override
	for _, e := range entries {
		if e.IsDir() || filepath.Ext(e.Name()) != ".md" {
			continue
		}
		body, err := os.ReadFile(filepath.Join(dir, e.Name()))
		if err != nil {
			return nil, fmt.Errorf("read override: %w", err)
		}
		out = append(out, Override{Name: e.Name(), Body: string(body)})
	}
	slices.SortFunc(out, func(a, b Override) int { return strings.Compare(a.Name, b.Name) })
	return out, nil
}

// HasOverrides reports whether the project adds anything of its own.
func (e Effective) HasOverrides() bool {
	return len(e.Preferences) > 0 || len(e.Overrides) > 0
}

// ProjectMarkdown is the project's preferences and overrides, under a heading
// that tells the agent they win.
func (e Effective) ProjectMarkdown() string {
	var b strings.Builder
	b.WriteString("# Project Overrides\n\nThese come from this project. Where they disagree with the doctrine, they win.\n")
	if len(e.Preferences) > 0 {
		b.WriteString("\n## Preferences\n\n")
		for _, p := range e.Preferences {
			b.WriteString("- " + p + "\n")
		}
	}
	for _, o := range e.Overrides {
		b.WriteString("\n" + strings.TrimRight(o.Body, "\n") + "\n")
	}
	return b.String()
}

// Markdown is the whole effective doctrine as one document, for reading.
func (e Effective) Markdown() string {
	var parts []string
	for _, d := range e.Doctrines {
		parts = append(parts, d.Body)
	}
	if e.HasOverrides() {
		parts = append(parts, e.ProjectMarkdown())
	}
	return strings.Join(parts, "\n")
}
