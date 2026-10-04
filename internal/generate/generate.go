// Package generate writes an adapter's files into a project.
package generate

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"slices"
)

// Write writes files under root. Files left in the owned directories that are
// no longer generated are removed. Nothing outside the owned directories is
// removed. It returns the paths it wrote, sorted.
func Write(root string, files map[string][]byte, owned []string) ([]string, error) {
	for _, dir := range owned {
		err := filepath.WalkDir(filepath.Join(root, dir), func(path string, d fs.DirEntry, err error) error {
			if err != nil || d.IsDir() {
				return err
			}
			rel, err := filepath.Rel(root, path)
			if err != nil {
				return err
			}
			if _, keep := files[rel]; !keep {
				return os.Remove(path)
			}
			return nil
		})
		if err != nil && !errors.Is(err, fs.ErrNotExist) {
			return nil, fmt.Errorf("clean %s: %w", dir, err)
		}
	}

	var written []string
	for rel, content := range files {
		path := filepath.Join(root, rel)
		if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
			return nil, err
		}
		if err := os.WriteFile(path, content, 0o644); err != nil {
			return nil, fmt.Errorf("write %s: %w", rel, err)
		}
		written = append(written, rel)
	}
	slices.Sort(written)
	return written, nil
}
