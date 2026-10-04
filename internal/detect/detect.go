// Package detect finds the technologies and coding agents a repository uses.
package detect

import (
	"bufio"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// Result is what was found in a repository.
type Result struct {
	// Technologies such as ruby, rails, go, docker, kamal, node.
	Technologies []string
	// Doctrines suggested for the technologies found.
	Doctrines []string
	// Agents such as claude-code.
	Agents []string
}

var railsInLockfile = regexp.MustCompile(`(?m)^    rails \(`)
var railsInGemfile = regexp.MustCompile(`(?m)^\s*gem\s+["']rails["']`)

// Detect looks at the repository root and its first-level folders, so a Go
// service next to a Rails app is found too.
func Detect(root string) (Result, error) {
	dirs := []string{root}
	entries, err := os.ReadDir(root)
	if err != nil {
		return Result{}, err
	}
	for _, e := range entries {
		name := e.Name()
		if e.IsDir() && !strings.HasPrefix(name, ".") && name != "vendor" && name != "node_modules" && name != "tmp" {
			dirs = append(dirs, filepath.Join(root, name))
		}
	}

	found := map[string]bool{}
	for _, dir := range dirs {
		if exists(dir, "Gemfile") {
			found["ruby"] = true
			if matches(filepath.Join(dir, "Gemfile.lock"), railsInLockfile) || matches(filepath.Join(dir, "Gemfile"), railsInGemfile) {
				found["rails"] = true
			}
		}
		if exists(dir, "go.mod") {
			found["go"] = true
		}
		if exists(dir, "Dockerfile") {
			found["docker"] = true
		}
		if exists(dir, "config/deploy.yml") {
			found["kamal"] = true
		}
		if exists(dir, "package.json") {
			found["node"] = true
		}
	}

	var r Result
	for _, t := range []string{"ruby", "rails", "go", "docker", "kamal", "node"} {
		if found[t] {
			r.Technologies = append(r.Technologies, t)
		}
	}
	switch {
	case found["rails"]:
		r.Doctrines = append(r.Doctrines, "rails/default")
	case found["ruby"]:
		r.Doctrines = append(r.Doctrines, "ruby/default")
	}
	if found["go"] {
		r.Doctrines = append(r.Doctrines, "go/default")
	}
	if exists(root, "CLAUDE.md") || exists(root, ".claude") {
		r.Agents = append(r.Agents, "claude-code")
	}
	return r, nil
}

// Stacks maps the short names accepted on the command line to doctrines.
var Stacks = map[string]string{
	"ruby":  "ruby/default",
	"rails": "rails/default",
	"go":    "go/default",
}

// DoctrineFor turns "rails" or "rails/default" into "rails/default".
func DoctrineFor(stack string) string {
	if d, ok := Stacks[stack]; ok {
		return d
	}
	return stack
}

func exists(dir, name string) bool {
	_, err := os.Stat(filepath.Join(dir, name))
	return err == nil
}

func matches(path string, re *regexp.Regexp) bool {
	f, err := os.Open(path)
	if err != nil {
		return false
	}
	defer f.Close()
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		if re.MatchString(sc.Text()) {
			return true
		}
	}
	return false
}
