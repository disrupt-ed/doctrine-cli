package detect

import (
	"os"
	"path/filepath"
	"slices"
	"testing"
)

func write(t *testing.T, root string, files map[string]string) {
	t.Helper()
	for name, content := range files {
		path := filepath.Join(root, name)
		if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
			t.Fatal(err)
		}
	}
}

func TestDetect(t *testing.T) {
	tests := []struct {
		name                       string
		files                      map[string]string
		technologies, docs, agents []string
	}{
		{"rails from lockfile", map[string]string{"Gemfile": "source 'x'\n", "Gemfile.lock": "GEM\n  specs:\n    rails (8.0.1)\n"},
			[]string{"ruby", "rails"}, []string{"rails/default"}, nil},
		{"rails from Gemfile", map[string]string{"Gemfile": "gem \"rails\", \"~> 8.0\"\n"},
			[]string{"ruby", "rails"}, []string{"rails/default"}, nil},
		{"plain ruby", map[string]string{"Gemfile": "gem 'rack'\n", "Gemfile.lock": "    rack (3.0)\n      railsish (1)\n"},
			[]string{"ruby"}, []string{"ruby/default"}, nil},
		{"go", map[string]string{"go.mod": "module x\n"},
			[]string{"go"}, []string{"go/default"}, nil},
		{"rails with go service in a folder", map[string]string{"Gemfile": "gem 'rails'\n", "service/go.mod": "module x\n", "Dockerfile": "", "config/deploy.yml": ""},
			[]string{"ruby", "rails", "go", "docker", "kamal"}, []string{"rails/default", "go/default"}, nil},
		{"vendored go ignored", map[string]string{"vendor/go.mod": "module x\n"}, nil, nil, nil},
		{"claude", map[string]string{"CLAUDE.md": "", "go.mod": ""},
			[]string{"go"}, []string{"go/default"}, []string{"claude-code"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := t.TempDir()
			write(t, root, tt.files)
			r, err := Detect(root)
			if err != nil {
				t.Fatal(err)
			}
			if !slices.Equal(r.Technologies, tt.technologies) || !slices.Equal(r.Doctrines, tt.docs) || !slices.Equal(r.Agents, tt.agents) {
				t.Errorf("got %v %v %v, want %v %v %v", r.Technologies, r.Doctrines, r.Agents, tt.technologies, tt.docs, tt.agents)
			}
		})
	}
}

func TestDoctrineFor(t *testing.T) {
	for in, want := range map[string]string{"rails": "rails/default", "go": "go/default", "rails/default": "rails/default", "x/y": "x/y"} {
		if got := DoctrineFor(in); got != want {
			t.Errorf("DoctrineFor(%q) = %q, want %q", in, got, want)
		}
	}
}
