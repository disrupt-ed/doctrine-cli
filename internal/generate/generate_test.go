package generate

import (
	"os"
	"path/filepath"
	"testing"
)

func TestWriteReplacesOnlyOwnedFiles(t *testing.T) {
	root := t.TempDir()
	owned := filepath.Join(".claude", "rules", "doctrine")
	stale := filepath.Join(root, owned, "rails.md")
	mine := filepath.Join(root, ".claude", "rules", "mine.md")
	for _, p := range []string{stale, mine} {
		os.MkdirAll(filepath.Dir(p), 0o755)
		os.WriteFile(p, []byte("old"), 0o644)
	}

	files := map[string][]byte{filepath.Join(owned, "go.md"): []byte("go")}
	written, err := Write(root, files, []string{owned})
	if err != nil {
		t.Fatal(err)
	}
	if len(written) != 1 {
		t.Fatalf("written = %v", written)
	}
	if _, err := os.Stat(stale); !os.IsNotExist(err) {
		t.Error("stale generated file was kept")
	}
	if _, err := os.Stat(mine); err != nil {
		t.Error("file outside the owned directory was removed")
	}
	if got, _ := os.ReadFile(filepath.Join(root, owned, "go.md")); string(got) != "go" {
		t.Errorf("go.md = %q", got)
	}
}

func TestWriteIntoEmptyProject(t *testing.T) {
	root := t.TempDir()
	if _, err := Write(root, map[string][]byte{"a/b.md": []byte("x")}, []string{"a"}); err != nil {
		t.Fatal(err)
	}
}
