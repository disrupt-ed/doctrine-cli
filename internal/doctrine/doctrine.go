// Package doctrine loads official doctrines and resolves their parents.
package doctrine

import (
	"fmt"
	"io/fs"
	"path"
	"slices"
	"strings"

	"go.yaml.in/yaml/v3"
)

// Doctrine is one official doctrine, such as rails/default.
type Doctrine struct {
	Name        string   `yaml:"name"`
	Description string   `yaml:"description"`
	Parent      string   `yaml:"parent"`
	Files       []string `yaml:"files"`

	// Body is the doctrine's Markdown files joined in order.
	Body string `yaml:"-"`
}

// Topic is the part of the name before the slash: "rails" for rails/default.
func (d Doctrine) Topic() string {
	topic, _, _ := strings.Cut(d.Name, "/")
	return topic
}

// Load reads one doctrine from a release.
func Load(release fs.FS, name string) (Doctrine, error) {
	var d Doctrine
	data, err := fs.ReadFile(release, path.Join(name, "doctrine.yml"))
	if err != nil {
		return d, fmt.Errorf("unknown doctrine %q", name)
	}
	if err := yaml.Unmarshal(data, &d); err != nil {
		return d, fmt.Errorf("parse %s/doctrine.yml: %w", name, err)
	}
	if d.Name != name {
		return d, fmt.Errorf("%s/doctrine.yml: name is %q", name, d.Name)
	}
	parts := make([]string, 0, len(d.Files))
	for _, f := range d.Files {
		body, err := fs.ReadFile(release, path.Join(name, f))
		if err != nil {
			return d, fmt.Errorf("read %s/%s: %w", name, f, err)
		}
		parts = append(parts, string(body))
	}
	d.Body = strings.Join(parts, "\n")
	return d, nil
}

// Resolve loads the given doctrines with their parents, parents first and
// without duplicates.
func Resolve(release fs.FS, extends []string) ([]Doctrine, error) {
	var out []Doctrine
	seen := map[string]bool{}
	for _, name := range extends {
		var chain []Doctrine
		for n := name; n != ""; {
			if slices.ContainsFunc(chain, func(d Doctrine) bool { return d.Name == n }) {
				return nil, fmt.Errorf("doctrine %s: parent cycle", name)
			}
			d, err := Load(release, n)
			if err != nil {
				return nil, err
			}
			chain = append(chain, d)
			n = d.Parent
		}
		for i := len(chain) - 1; i >= 0; i-- {
			if !seen[chain[i].Name] {
				seen[chain[i].Name] = true
				out = append(out, chain[i])
			}
		}
	}
	return out, nil
}

// All lists every doctrine in a release, sorted by name.
func All(release fs.FS) ([]Doctrine, error) {
	matches, err := fs.Glob(release, "*/*/doctrine.yml")
	if err != nil {
		return nil, err
	}
	var out []Doctrine
	for _, m := range matches {
		d, err := Load(release, path.Dir(m))
		if err != nil {
			return nil, err
		}
		out = append(out, d)
	}
	return out, nil
}

// TopLevel drops doctrines that another one in the list already brings in
// as a parent, and duplicates. [ruby/default rails/default] gives [rails/default].
func TopLevel(release fs.FS, names []string) ([]string, error) {
	ancestors := map[string]bool{}
	for _, name := range names {
		d, err := Load(release, name)
		if err != nil {
			return nil, err
		}
		for p := d.Parent; p != "" && !ancestors[p]; {
			ancestors[p] = true
			pd, err := Load(release, p)
			if err != nil {
				return nil, err
			}
			p = pd.Parent
		}
	}
	var out []string
	for _, name := range names {
		if !ancestors[name] && !slices.Contains(out, name) {
			out = append(out, name)
		}
	}
	return out, nil
}
