package doctrine

import (
	"archive/tar"
	"compress/gzip"
	"context"
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/disrupt-ed/doctrine-cli/doctrines"
)

// ArchiveURL is where tagged releases of the public repository are
// downloaded from. %s is the doctrine release.
var ArchiveURL = "https://codeload.github.com/disrupt-ed/doctrine-cli/tar.gz/refs/tags/doctrines-v%s"

// Release returns the files of a doctrine release. The release this binary
// was built with needs no network. Others are downloaded once and cached.
func Release(ctx context.Context, version string) (fs.FS, error) {
	if version == doctrines.Release {
		return doctrines.FS, nil
	}
	cache, err := os.UserCacheDir()
	if err != nil {
		return nil, fmt.Errorf("find cache directory: %w", err)
	}
	dir := filepath.Join(cache, "doctrine", "doctrines", version)
	if _, err := os.Stat(dir); err == nil {
		return os.DirFS(dir), nil
	}
	if err := download(ctx, version, dir); err != nil {
		return nil, fmt.Errorf("doctrine release %s isn't built in (this binary has %s) and couldn't be downloaded: %w", version, doctrines.Release, err)
	}
	return os.DirFS(dir), nil
}

// download extracts the doctrines/ folder of a tagged release into dir.
func download(ctx context.Context, version, dir string) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf(ArchiveURL, version), nil)
	if err != nil {
		return err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("download: %s", resp.Status)
	}
	gz, err := gzip.NewReader(resp.Body)
	if err != nil {
		return err
	}
	if err := os.MkdirAll(filepath.Dir(dir), 0o755); err != nil {
		return err
	}
	tmp, err := os.MkdirTemp(filepath.Dir(dir), ".download-")
	if err != nil {
		return err
	}
	defer os.RemoveAll(tmp)

	tr := tar.NewReader(gz)
	found := false
	for {
		h, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		// Archive paths look like doctrine-doctrines-v0.1/doctrines/rails/default/x.md.
		_, rest, _ := strings.Cut(h.Name, "/")
		rel, ok := strings.CutPrefix(rest, "doctrines/")
		if !ok || h.Typeflag != tar.TypeReg || !filepath.IsLocal(rel) {
			continue
		}
		target := filepath.Join(tmp, rel)
		if err := os.MkdirAll(filepath.Dir(target), 0o755); err != nil {
			return err
		}
		f, err := os.Create(target)
		if err != nil {
			return err
		}
		_, err = io.Copy(f, io.LimitReader(tr, 1<<20))
		f.Close()
		if err != nil {
			return err
		}
		found = true
	}
	if !found {
		return fmt.Errorf("release has no doctrines")
	}
	return os.Rename(tmp, dir)
}
