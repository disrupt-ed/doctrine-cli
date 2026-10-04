// Package manifest reads and edits a project's .doctrine/doctrine.yml.
package manifest

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"go.yaml.in/yaml/v3"
)

// Path is where the manifest lives, relative to the project root.
var Path = filepath.Join(".doctrine", "doctrine.yml")

// LocalDir holds the project's own override files.
var LocalDir = filepath.Join(".doctrine", "local")

// Manifest is the parsed content of .doctrine/doctrine.yml.
type Manifest struct {
	Version     int               `yaml:"version"`
	Doctrines   Release           `yaml:"doctrines"`
	Extends     []string          `yaml:"extends"`
	Engineering map[string]string `yaml:"engineering,omitempty"`
}

// Release is a doctrine release such as "0.1". It keeps the text as written,
// so 0.10 doesn't become 0.1.
type Release string

func (r *Release) UnmarshalYAML(n *yaml.Node) error {
	if n.Kind != yaml.ScalarNode {
		return fmt.Errorf("doctrines: expected a version like 0.1")
	}
	*r = Release(n.Value)
	return nil
}

// Preferences lists the allowed engineering keys and values. The first value
// is the default.
var Preferences = map[string][]string{
	"kiss":         {"strong", "normal"},
	"yagni":        {"strong", "normal"},
	"dependencies": {"conservative", "balanced", "open"},
	"abstractions": {"late", "balanced"},
}

// PreferenceOrder is the order preferences are written in.
var PreferenceOrder = []string{"kiss", "yagni", "dependencies", "abstractions"}

// Parse parses and validates a manifest.
func Parse(data []byte) (Manifest, error) {
	var m Manifest
	if err := yaml.Unmarshal(data, &m); err != nil {
		return m, fmt.Errorf("parse manifest: %w", err)
	}
	if m.Version != 1 {
		return m, fmt.Errorf("manifest: unsupported version %d (expected 1)", m.Version)
	}
	if m.Doctrines == "" {
		return m, fmt.Errorf("manifest: doctrines release is missing")
	}
	if len(m.Extends) == 0 {
		return m, fmt.Errorf("manifest: extends is empty")
	}
	for k, v := range m.Engineering {
		values, ok := Preferences[k]
		if !ok {
			return m, fmt.Errorf("manifest: unknown engineering preference %q", k)
		}
		if !slices.Contains(values, v) {
			return m, fmt.Errorf("manifest: engineering.%s must be one of %s, got %q", k, strings.Join(values, ", "), v)
		}
	}
	return m, nil
}

// Read reads the manifest of the project in dir.
func Read(dir string) (Manifest, []byte, error) {
	data, err := os.ReadFile(filepath.Join(dir, Path))
	if err != nil {
		return Manifest{}, nil, err
	}
	m, err := Parse(data)
	return m, data, err
}

// Marshal writes a new manifest in its canonical form.
func (m Manifest) Marshal() []byte {
	var b bytes.Buffer
	fmt.Fprintf(&b, "version: %d\ndoctrines: %s\n\nextends:\n", m.Version, m.Doctrines)
	for _, e := range m.Extends {
		fmt.Fprintf(&b, "  - %s\n", e)
	}
	if len(m.Engineering) > 0 {
		b.WriteString("\nengineering:\n")
		for _, k := range PreferenceOrder {
			if v, ok := m.Engineering[k]; ok {
				fmt.Fprintf(&b, "  %s: %s\n", k, v)
			}
		}
	}
	return b.Bytes()
}

// SetExtends replaces the extends list in an existing manifest, keeping the
// rest of the file and its comments.
func SetExtends(data []byte, extends []string) ([]byte, error) {
	return edit(data, "extends", func(n *yaml.Node) {
		n.Kind = yaml.SequenceNode
		n.Tag = "!!seq"
		n.Value = ""
		n.Style = 0
		n.Content = nil
		for _, e := range extends {
			n.Content = append(n.Content, &yaml.Node{Kind: yaml.ScalarNode, Tag: "!!str", Value: e})
		}
	})
}

// SetRelease changes the pinned doctrine release in an existing manifest.
func SetRelease(data []byte, release string) ([]byte, error) {
	return edit(data, "doctrines", func(n *yaml.Node) {
		n.Kind = yaml.ScalarNode
		n.Tag = "!!float"
		n.Value = release
		n.Content = nil
	})
}

func edit(data []byte, key string, change func(*yaml.Node)) ([]byte, error) {
	var doc yaml.Node
	if err := yaml.Unmarshal(data, &doc); err != nil {
		return nil, fmt.Errorf("parse manifest: %w", err)
	}
	if len(doc.Content) == 0 || doc.Content[0].Kind != yaml.MappingNode {
		return nil, fmt.Errorf("manifest: expected a mapping")
	}
	root := doc.Content[0]
	found := false
	for i := 0; i+1 < len(root.Content); i += 2 {
		if root.Content[i].Value == key {
			change(root.Content[i+1])
			found = true
		}
	}
	if !found {
		return nil, fmt.Errorf("manifest: %s is missing", key)
	}
	var b bytes.Buffer
	enc := yaml.NewEncoder(&b)
	enc.SetIndent(2)
	if err := enc.Encode(&doc); err != nil {
		return nil, fmt.Errorf("write manifest: %w", err)
	}
	if err := enc.Close(); err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
