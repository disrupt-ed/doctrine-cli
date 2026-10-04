package dependencies

import (
	"context"
	"fmt"
	"net/url"
	"strings"
	"time"
)

// gitHubRepo adds repository, commit and contributor facts. Failures
// become notes.
func (c *Client) gitHubRepo(ctx context.Context, r *Report, repo string) {
	base := orDefault(c.GitHubURL, "https://api.github.com") + "/repos/" + repo
	now := c.now()

	var info struct {
		Archived      bool      `json:"archived"`
		PushedAt      time.Time `json:"pushed_at"`
		Stars         int       `json:"stargazers_count"`
		OpenIssues    int       `json:"open_issues_count"`
		DefaultBranch string    `json:"default_branch"`
	}
	if _, err := c.getJSON(ctx, base, true, &info); err != nil {
		r.note("GitHub repository facts unavailable: %v", err)
	} else {
		r.add("repository archived", fmt.Sprint(info.Archived))
		r.add("last push", date(info.PushedAt, now))
		r.add("stars", fmt.Sprint(info.Stars))
		r.add("open issues and pull requests", fmt.Sprint(info.OpenIssues))
		r.add("default branch", info.DefaultBranch)
	}

	q := url.Values{}
	q.Set("since", now.AddDate(-1, 0, 0).UTC().Format(time.RFC3339))
	q.Set("per_page", "100")
	var commits []struct {
		Author *struct {
			Login string `json:"login"`
		} `json:"author"`
		Commit struct {
			Author struct {
				Email string `json:"email"`
			} `json:"author"`
		} `json:"commit"`
	}
	header, err := c.getJSON(ctx, base+"/commits?"+q.Encode(), true, &commits)
	if err != nil {
		r.note("GitHub commit history unavailable: %v", err)
		return
	}
	authors := map[string]bool{}
	for _, cm := range commits {
		if cm.Author != nil && cm.Author.Login != "" {
			authors[cm.Author.Login] = true
		} else {
			authors[cm.Commit.Author.Email] = true
		}
	}
	more := len(commits) == 100 && strings.Contains(header.Get("Link"), `rel="next"`)
	r.add("commits in last 12 months", countAtLeast(len(commits), more))
	r.add("contributors in last 12 months (from up to 100 commits)", fmt.Sprint(len(authors)))
}

// gitHubAdvisories adds security advisory counts from the GitHub
// advisory database. Failures become notes.
func (c *Client) gitHubAdvisories(ctx context.Context, r *Report, eco Ecosystem, name string) {
	q := url.Values{}
	q.Set("ecosystem", string(eco))
	q.Set("affects", name)
	q.Set("per_page", "100")
	u := orDefault(c.GitHubURL, "https://api.github.com") + "/advisories?" + q.Encode()

	var advisories []struct {
		Severity    string     `json:"severity"`
		WithdrawnAt *time.Time `json:"withdrawn_at"`
	}
	header, err := c.getJSON(ctx, u, true, &advisories)
	if err != nil {
		r.note("GitHub security advisories unavailable: %v", err)
		return
	}
	severe := 0
	for _, a := range advisories {
		if a.WithdrawnAt == nil && (a.Severity == "high" || a.Severity == "critical") {
			severe++
		}
	}
	more := len(advisories) == 100 && strings.Contains(header.Get("Link"), `rel="next"`)
	r.add("published security advisories", countAtLeast(len(advisories), more))
	r.add("high or critical advisories (not withdrawn)", fmt.Sprint(severe))
}
