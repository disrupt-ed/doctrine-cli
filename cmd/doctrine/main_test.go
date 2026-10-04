package main

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"context"
	"encoding/json"
	"io"
	"io/fs"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/disrupt-ed/doctrine-cli/doctrines"
	"github.com/disrupt-ed/doctrine-cli/internal/doctrine"
	"github.com/disrupt-ed/doctrine-cli/internal/manifest"
)

// setup gives each test its own project, home and config, with stats off
// unless the test turns them on.
func setup(t *testing.T, files map[string]string) string {
	t.Helper()
	root := t.TempDir()
	for name, content := range files {
		path := filepath.Join(root, name)
		os.MkdirAll(filepath.Dir(path), 0o755)
		os.WriteFile(path, []byte(content), 0o644)
	}
	t.Chdir(root)
	home := t.TempDir()
	t.Setenv("HOME", home)
	t.Setenv("XDG_CONFIG_HOME", filepath.Join(home, ".config"))
	t.Setenv("XDG_CACHE_HOME", filepath.Join(home, ".cache"))
	t.Setenv("AppData", filepath.Join(home, "AppData", "Roaming"))
	t.Setenv("LocalAppData", filepath.Join(home, "AppData", "Local"))
	t.Setenv("DOCTRINE_TELEMETRY", "0")
	t.Setenv("DO_NOT_TRACK", "")
	return root
}

func doctrineRun(t *testing.T, stdin string, args ...string) (string, error) {
	t.Helper()
	var out bytes.Buffer
	err := run(context.Background(), args, strings.NewReader(stdin), &out)
	return out.String(), err
}

func mustRun(t *testing.T, args ...string) string {
	t.Helper()
	out, err := doctrineRun(t, "", args...)
	if err != nil {
		t.Fatalf("doctrine %s: %v\n%s", strings.Join(args, " "), err, out)
	}
	return out
}

func exists(root, path string) bool {
	_, err := os.Stat(filepath.Join(root, path))
	return err == nil
}

func TestInitDetectsRails(t *testing.T) {
	root := setup(t, map[string]string{"Gemfile": "gem \"rails\"\n", "CLAUDE.md": "# Mine\n"})
	mustRun(t, "init")
	for _, f := range []string{".doctrine/doctrine.yml", ".claude/rules/doctrine/engineering.md", ".claude/rules/doctrine/ruby.md", ".claude/rules/doctrine/rails.md", ".claude/skills/dependency-review/SKILL.md"} {
		if !exists(root, f) {
			t.Errorf("missing %s", f)
		}
	}
	if got, _ := os.ReadFile(filepath.Join(root, "CLAUDE.md")); string(got) != "# Mine\n" {
		t.Error("CLAUDE.md was changed")
	}
	if _, err := doctrineRun(t, "", "init"); err == nil {
		t.Error("second init should fail")
	}
}

func TestInitWithStacks(t *testing.T) {
	root := setup(t, nil)
	if _, err := doctrineRun(t, "", "init"); err == nil {
		t.Fatal("init without stacks in an empty repo should fail")
	}
	mustRun(t, "init", "ruby", "rails", "go")
	m, _, err := manifest.Read(root)
	if err != nil {
		t.Fatal(err)
	}
	if strings.Join(m.Extends, ",") != "rails/default,go/default" {
		t.Fatalf("Extends = %v", m.Extends)
	}
}

func TestAddAndRemove(t *testing.T) {
	root := setup(t, map[string]string{"go.mod": "module x\n"})
	mustRun(t, "init")
	if exists(root, ".claude/rules/doctrine/rails.md") {
		t.Fatal("rails before add")
	}
	mustRun(t, "add", "rails")
	if !exists(root, ".claude/rules/doctrine/rails.md") {
		t.Fatal("rails.md missing after add")
	}
	mustRun(t, "remove", "rails")
	if exists(root, ".claude/rules/doctrine/rails.md") || exists(root, ".claude/rules/doctrine/ruby.md") {
		t.Fatal("rails files left after remove")
	}
	if _, err := doctrineRun(t, "", "remove", "go"); err == nil {
		t.Fatal("removing the last doctrine should fail")
	}
	if _, err := doctrineRun(t, "", "add", "cobol"); err == nil {
		t.Fatal("unknown doctrine should fail")
	}
}

func TestGenerateKeepsOverridesAndUserFiles(t *testing.T) {
	root := setup(t, map[string]string{"go.mod": "module x\n", ".claude/rules/mine.md": "mine"})
	mustRun(t, "init")
	os.MkdirAll(filepath.Join(root, ".doctrine/local"), 0o755)
	os.WriteFile(filepath.Join(root, ".doctrine/local/go.md"), []byte("# Go Overrides\n\nUse chi, we already do.\n"), 0o644)
	mustRun(t, "generate", "claude")
	project, err := os.ReadFile(filepath.Join(root, ".claude/rules/doctrine/project.md"))
	if err != nil || !strings.Contains(string(project), "Use chi") {
		t.Fatalf("project.md: %v %s", err, project)
	}
	if !exists(root, ".claude/rules/mine.md") {
		t.Fatal("user rule was removed")
	}
	if _, err := doctrineRun(t, "", "generate", "cursor"); err == nil {
		t.Fatal("unknown agent should fail")
	}
}

func TestInspect(t *testing.T) {
	setup(t, map[string]string{"go.mod": "module x\n"})
	if _, err := doctrineRun(t, "", "inspect"); err == nil || !strings.Contains(err.Error(), "doctrine init") {
		t.Fatalf("inspect before init: %v", err)
	}
	mustRun(t, "init")
	out := mustRun(t, "inspect")
	if !strings.Contains(out, "# Engineering Decisions") || !strings.Contains(out, "# Go Already Does This") {
		t.Fatal("inspect is missing doctrines")
	}
	out = mustRun(t, "inspect", "rails")
	if !strings.HasPrefix(out, "# The Rails Way") || strings.Contains(out, "# Engineering Decisions") {
		t.Fatal("inspect rails should print only rails/default")
	}
	if !strings.Contains(mustRun(t, "list"), "rails/default") {
		t.Fatal("list is missing rails/default")
	}
}

func TestUpdate(t *testing.T) {
	root := setup(t, map[string]string{"go.mod": "module x\n"})
	if !strings.Contains(mustRun(t, "init"), "go/default") {
		t.Fatal("init")
	}
	if !strings.Contains(mustRun(t, "update"), "Already on") {
		t.Fatal("expected already on latest")
	}

	// Pin the project to an older release, served by a fake repository
	// where one line of go.md differs.
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(releaseArchive(t, "Write the obvious code.", "Write clever code."))
	}))
	defer srv.Close()
	old := doctrine.ArchiveURL
	doctrine.ArchiveURL = srv.URL + "/%s"
	defer func() { doctrine.ArchiveURL = old }()

	data, _ := os.ReadFile(filepath.Join(root, manifest.Path))
	data, _ = manifest.SetRelease(data, "0.0.9")
	os.WriteFile(filepath.Join(root, manifest.Path), data, 0o644)

	out, err := doctrineRun(t, "n\n", "update")
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(out, "-- Write clever code.") || !strings.Contains(out, "+- Write the obvious code.") {
		t.Fatalf("diff missing:\n%s", out)
	}
	if m, _, _ := manifest.Read(root); m.Doctrines != "0.0.9" {
		t.Fatal("declining must not move the pin")
	}

	if _, err := doctrineRun(t, "y\n", "update"); err != nil {
		t.Fatal(err)
	}
	if m, _, _ := manifest.Read(root); string(m.Doctrines) != doctrines.Release {
		t.Fatalf("pin = %s after accepting", m.Doctrines)
	}
}

// releaseArchive builds a tagged-release tarball of the embedded doctrines,
// with one change.
func releaseArchive(t *testing.T, from, to string) []byte {
	t.Helper()
	var buf bytes.Buffer
	gz := gzip.NewWriter(&buf)
	tw := tar.NewWriter(gz)
	fs.WalkDir(doctrines.FS, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil || d.IsDir() {
			return err
		}
		data, _ := fs.ReadFile(doctrines.FS, path)
		data = bytes.ReplaceAll(data, []byte(from), []byte(to))
		tw.WriteHeader(&tar.Header{Name: "doctrine-old/doctrines/" + path, Mode: 0o644, Size: int64(len(data)), Typeflag: tar.TypeReg})
		tw.Write(data)
		return nil
	})
	tw.Close()
	gz.Close()
	return buf.Bytes()
}

func TestInstall(t *testing.T) {
	root := setup(t, nil)
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/d/abcd1234" {
			http.NotFound(w, r)
			return
		}
		io.WriteString(w, "version: 1\ndoctrines: "+doctrines.Release+"\nextends:\n  - rails/default\n  - go/default\nengineering:\n  dependencies: balanced\n")
	}))
	defer srv.Close()
	t.Setenv("DOCTRINE_URL", srv.URL)

	if _, err := doctrineRun(t, "", "install", "nope0000"); err == nil || !strings.Contains(err.Error(), "no configuration") {
		t.Fatalf("missing handle: %v", err)
	}
	if _, err := doctrineRun(t, "", "install", "../etc"); err == nil {
		t.Fatal("invalid handle should fail")
	}
	mustRun(t, "install", "abcd1234")
	for _, f := range []string{".claude/rules/doctrine/rails.md", ".claude/rules/doctrine/go.md", ".claude/rules/doctrine/project.md"} {
		if !exists(root, f) {
			t.Errorf("missing %s", f)
		}
	}
	data, _ := os.ReadFile(filepath.Join(root, manifest.Path))
	if strings.Contains(string(data), "abcd1234") {
		t.Error("the handle must not be written into the repository")
	}
}

func TestStats(t *testing.T) {
	var got []map[string]any
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var body map[string]any
		json.NewDecoder(r.Body).Decode(&body)
		got = append(got, body)
	}))
	defer srv.Close()

	old := interactive
	interactive = func(io.Reader) bool { return true }
	defer func() { interactive = old }()

	tests := []struct {
		name, stdin, doNotTrack, telemetry string
		sent                               int
	}{
		{"declined", "n\n", "", "", 0},
		{"accepted", "y\n", "", "", 1},
		{"do not track", "y\n", "1", "", 0},
		{"telemetry off", "y\n", "", "0", 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got = nil
			setup(t, map[string]string{"Gemfile": "gem 'rails'\n"})
			t.Setenv("DOCTRINE_URL", srv.URL)
			t.Setenv("DO_NOT_TRACK", tt.doNotTrack)
			t.Setenv("DOCTRINE_TELEMETRY", tt.telemetry)
			out, err := doctrineRun(t, tt.stdin, "init")
			if err != nil {
				t.Fatal(err)
			}
			if len(got) != tt.sent {
				t.Fatalf("sent %d stats, want %d\n%s", len(got), tt.sent, out)
			}
			if tt.sent == 1 {
				if !strings.Contains(out, "This is what we'd send") {
					t.Error("the prompt must show what is sent")
				}
				for k := range got[0] {
					if !strings.Contains(" languages frameworks os arch agent ", " "+k+" ") {
						t.Errorf("unexpected field %q sent", k)
					}
				}
			}
		})
	}
}

func TestDiff(t *testing.T) {
	if diff("f", "a\nb\n", "a\nb\n") != "" {
		t.Fatal("no change should give no diff")
	}
	got := diff("f", "a\nb\nc\nd\ne\n", "a\nb\nX\nd\ne\n")
	want := "--- f\n+++ f\n@@\n b\n-c\n+X\n d\n\n"
	if got != want {
		t.Fatalf("got:\n%q\nwant:\n%q", got, want)
	}
}
