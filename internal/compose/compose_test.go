package compose

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/disrupt-ed/doctrine-cli/doctrines"
	"github.com/disrupt-ed/doctrine-cli/internal/manifest"
)

func TestComposeWithoutOverrides(t *testing.T) {
	m := manifest.Manifest{Version: 1, Doctrines: "0.1", Extends: []string{"go/default"}}
	e, err := Compose(doctrines.FS, m, t.TempDir())
	if err != nil {
		t.Fatal(err)
	}
	if len(e.Doctrines) != 2 || e.HasOverrides() {
		t.Fatalf("got %d doctrines, overrides %v", len(e.Doctrines), e.HasOverrides())
	}
	if strings.Contains(e.Markdown(), "Project Overrides") {
		t.Fatal("unexpected overrides section")
	}
}

func TestComposeOverridesComeLastAndWin(t *testing.T) {
	root := t.TempDir()
	local := filepath.Join(root, manifest.LocalDir)
	os.MkdirAll(local, 0o755)
	os.WriteFile(filepath.Join(local, "rails.md"), []byte("# Rails Overrides\n\nService objects are fine under app/services/payments/.\n"), 0o644)
	os.WriteFile(filepath.Join(local, "a-first.md"), []byte("# First\n"), 0o644)
	os.WriteFile(filepath.Join(local, "notes.txt"), []byte("zz-not-markdown-zz"), 0o644)

	m := manifest.Manifest{Version: 1, Doctrines: "0.1", Extends: []string{"rails/default"},
		Engineering: map[string]string{"dependencies": "open", "kiss": "normal"}}
	e, err := Compose(doctrines.FS, m, root)
	if err != nil {
		t.Fatal(err)
	}
	md := e.Markdown()
	project := strings.Index(md, "# Project Overrides")
	if project < 0 || project < strings.LastIndex(md, "# Data and Migrations") {
		t.Fatal("project overrides must come after the doctrines")
	}
	for _, want := range []string{"they win", "Simplicity: normal", "Dependencies: open", "Service objects are fine", "# First"} {
		if !strings.Contains(md[project:], want) {
			t.Errorf("missing %q", want)
		}
	}
	if strings.Contains(md, "zz-not-markdown-zz") {
		t.Error("non-Markdown file was included")
	}
	if strings.Index(md, "Simplicity") > strings.Index(md, "Dependencies: open") {
		t.Error("preferences out of order")
	}
	if strings.Index(md, "# First") > strings.Index(md, "Service objects") {
		t.Error("override files must be in alphabetical order")
	}
}

func TestEveryPreferenceHasText(t *testing.T) {
	for key, values := range manifest.Preferences {
		for _, v := range values {
			if preferenceText[key+"="+v] == "" {
				t.Errorf("no text for %s: %s", key, v)
			}
		}
	}
}
