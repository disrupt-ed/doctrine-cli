package dependencies

import (
	"context"
	"fmt"
	"strings"
	"time"
)

// goModule adds Go proxy facts and returns the GitHub "owner/repo", if any.
func (c *Client) goModule(ctx context.Context, r *Report) (string, error) {
	base := orDefault(c.GoProxyURL, "https://proxy.golang.org") + "/" + escapePath(r.Name)

	var latest struct {
		Version string
		Time    time.Time
	}
	if _, err := c.getJSON(ctx, base+"/@latest", false, &latest); err != nil {
		return "", fmt.Errorf("go proxy lookup for %q: %w", r.Name, err)
	}
	now := c.now()
	r.add("latest version", latest.Version)
	r.add("latest release", date(latest.Time, now))

	if body, _, err := c.get(ctx, base+"/@v/list", false); err != nil {
		r.note("Go proxy version list unavailable: %v", err)
	} else {
		r.add("tagged versions", fmt.Sprint(len(strings.Fields(string(body)))))
	}

	if body, _, err := c.get(ctx, base+"/@v/"+escapePath(latest.Version)+".mod", false); err != nil {
		r.note("go.mod of latest version unavailable: %v", err)
	} else {
		goVersion, direct := parseGoMod(string(body))
		if goVersion != "" {
			r.add("go directive", goVersion)
		}
		r.add("direct requirements", fmt.Sprint(direct))
	}

	parts := strings.Split(r.Name, "/")
	if len(parts) >= 3 && parts[0] == "github.com" {
		r.add("source", "https://github.com/"+parts[1]+"/"+parts[2])
		return parts[1] + "/" + parts[2], nil
	}
	return "", nil
}

// escapePath escapes a module path or version for the proxy: each
// uppercase letter becomes "!" followed by its lowercase form.
func escapePath(path string) string {
	var b strings.Builder
	for _, r := range path {
		if r >= 'A' && r <= 'Z' {
			b.WriteByte('!')
			r += 'a' - 'A'
		}
		b.WriteRune(r)
	}
	return b.String()
}

// parseGoMod returns the go directive and the number of requirements
// not marked "// indirect".
func parseGoMod(data string) (goVersion string, direct int) {
	inRequire := false
	for _, line := range strings.Split(data, "\n") {
		line = strings.TrimSpace(line)
		switch {
		case inRequire:
			if line == ")" {
				inRequire = false
			} else if line != "" && !strings.HasPrefix(line, "//") && !strings.Contains(line, "// indirect") {
				direct++
			}
		case strings.HasPrefix(line, "require") && strings.HasSuffix(line, "("):
			inRequire = true
		case strings.HasPrefix(line, "require "):
			if !strings.Contains(line, "// indirect") {
				direct++
			}
		case strings.HasPrefix(line, "go "):
			goVersion = strings.Fields(line)[1]
		}
	}
	return goVersion, direct
}
