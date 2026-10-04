package dependencies

import (
	"context"
	"fmt"
	"net/url"
	"strings"
	"time"
)

type gemInfo struct {
	Version       string   `json:"version"`
	Licenses      []string `json:"licenses"`
	Downloads     int64    `json:"downloads"`
	SourceCodeURI string   `json:"source_code_uri"`
	HomepageURI   string   `json:"homepage_uri"`
	Dependencies  struct {
		Runtime []struct{} `json:"runtime"`
	} `json:"dependencies"`
}

type gemVersion struct {
	CreatedAt   time.Time `json:"created_at"`
	RubyVersion string    `json:"ruby_version"`
}

// rubyGems adds registry facts and returns the GitHub "owner/repo", if any.
func (c *Client) rubyGems(ctx context.Context, r *Report) (string, error) {
	base := orDefault(c.RubyGemsURL, "https://rubygems.org")
	name := url.PathEscape(r.Name)

	var gem gemInfo
	if _, err := c.getJSON(ctx, base+"/api/v1/gems/"+name+".json", false, &gem); err != nil {
		return "", fmt.Errorf("rubygems lookup for %q: %w", r.Name, err)
	}

	r.add("latest version", gem.Version)

	var versions []gemVersion
	if _, err := c.getJSON(ctx, base+"/api/v1/versions/"+name+".json", false, &versions); err != nil {
		r.note("RubyGems release history unavailable: %v", err)
	} else if len(versions) > 0 {
		now := c.now()
		yearAgo := now.AddDate(-1, 0, 0)
		first := versions[0].CreatedAt
		recent := 0
		for _, v := range versions {
			if v.CreatedAt.After(yearAgo) {
				recent++
			}
			if v.CreatedAt.Before(first) {
				first = v.CreatedAt
			}
		}
		r.add("latest release", date(versions[0].CreatedAt, now))
		r.add("first release", date(first, now))
		r.add("releases in last 12 months", fmt.Sprint(recent))
		r.add("total releases", fmt.Sprint(len(versions)))
		if versions[0].RubyVersion != "" {
			r.add("required ruby version", versions[0].RubyVersion)
		}
	}

	r.add("runtime dependencies", fmt.Sprint(len(gem.Dependencies.Runtime)))
	license := strings.Join(gem.Licenses, ", ")
	if license == "" {
		license = "none declared"
	}
	r.add("license", license)
	r.add("total downloads", fmt.Sprint(gem.Downloads))

	source := gem.SourceCodeURI
	if source == "" {
		source = gem.HomepageURI
	}
	if source != "" {
		r.add("source", source)
	}
	repo := gitHubRepoFromURL(gem.SourceCodeURI)
	if repo == "" {
		repo = gitHubRepoFromURL(gem.HomepageURI)
	}
	return repo, nil
}

// gitHubRepoFromURL returns "owner/repo" for a github.com URL, else "".
func gitHubRepoFromURL(raw string) string {
	u, err := url.Parse(raw)
	if err != nil || (u.Host != "github.com" && u.Host != "www.github.com") {
		return ""
	}
	parts := strings.Split(strings.Trim(u.Path, "/"), "/")
	if len(parts) < 2 || parts[0] == "" || parts[1] == "" {
		return ""
	}
	return parts[0] + "/" + strings.TrimSuffix(parts[1], ".git")
}
