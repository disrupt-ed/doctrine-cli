package dependencies

import (
	"context"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

var testNow = time.Date(2026, 10, 4, 12, 0, 0, 0, time.UTC)

// fakeServer serves the fixtures for every endpoint. GitHub endpoints
// answer with gitHubStatus; their Authorization headers are recorded.
type fakeServer struct {
	gitHubStatus int
	authHeaders  []string
}

func (f *fakeServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	routes := map[string]string{
		"/api/v1/gems/rack-attack.json":          "gem.json",
		"/api/v1/versions/rack-attack.json":      "versions.json",
		"/github.com/!example/mux/@latest":       "latest.json",
		"/github.com/!example/mux/@v/list":       "list.txt",
		"/github.com/!example/mux/@v/v1.8.1.mod": "mod.txt",
		"/github/repos/rack/rack-attack":         "repo.json",
		"/github/repos/rack/rack-attack/commits": "commits.json",
		"/github/repos/Example/mux":              "repo.json",
		"/github/repos/Example/mux/commits":      "commits.json",
		"/github/advisories":                     "advisories.json",
	}
	if strings.HasPrefix(r.URL.Path, "/github/") {
		f.authHeaders = append(f.authHeaders, r.Header.Get("Authorization"))
		if f.gitHubStatus != http.StatusOK {
			w.Header().Set("X-RateLimit-Remaining", "0")
			w.WriteHeader(f.gitHubStatus)
			return
		}
	}
	file, ok := routes[r.URL.Path]
	if !ok {
		http.NotFound(w, r)
		return
	}
	data, err := os.ReadFile(filepath.Join("testdata", file))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(data)
}

func newTestClient(t *testing.T, f *fakeServer, token string) *Client {
	t.Helper()
	srv := httptest.NewServer(f)
	t.Cleanup(srv.Close)
	return &Client{
		HTTP:        srv.Client(),
		RubyGemsURL: srv.URL,
		GoProxyURL:  srv.URL,
		GitHubURL:   srv.URL + "/github",
		GitHubToken: token,
		Now:         func() time.Time { return testNow },
	}
}

func factMap(r Report) map[string]string {
	m := map[string]string{}
	for _, f := range r.Facts {
		m[f.Label] = f.Value
	}
	return m
}

func checkFacts(t *testing.T, r Report, want map[string]string) {
	t.Helper()
	got := factMap(r)
	for label, value := range want {
		if got[label] != value {
			t.Errorf("%s = %q, want %q", label, got[label], value)
		}
	}
}

func TestInspectRubyGems(t *testing.T) {
	c := newTestClient(t, &fakeServer{gitHubStatus: http.StatusOK}, "secret")
	r, err := c.Inspect(context.Background(), RubyGems, "rack-attack")
	if err != nil {
		t.Fatal(err)
	}
	checkFacts(t, r, map[string]string{
		"latest version":                "6.7.0",
		"latest release":                "2026-07-01 (3 months ago)",
		"first release":                 "2012-07-20 (14 years ago)",
		"releases in last 12 months":    "2",
		"total releases":                "4",
		"required ruby version":         ">= 2.4",
		"runtime dependencies":          "1",
		"license":                       "MIT",
		"total downloads":               "123456789",
		"source":                        "https://github.com/rack/rack-attack/tree/v6.7.0",
		"repository archived":           "false",
		"last push":                     "2026-09-20 (14 days ago)",
		"stars":                         "5600",
		"open issues and pull requests": "42",
		"default branch":                "main",
		"commits in last 12 months":     "4",
		"contributors in last 12 months (from up to 100 commits)": "3",
		"published security advisories":                           "3",
		"high or critical advisories (not withdrawn)":             "1",
	})
	if len(r.Notes) != 0 {
		t.Errorf("unexpected notes: %v", r.Notes)
	}
}

func TestInspectGo(t *testing.T) {
	c := newTestClient(t, &fakeServer{gitHubStatus: http.StatusOK}, "secret")
	r, err := c.Inspect(context.Background(), Go, "github.com/Example/mux")
	if err != nil {
		t.Fatal(err)
	}
	checkFacts(t, r, map[string]string{
		"latest version":      "v1.8.1",
		"latest release":      "2023-11-05 (2 years ago)",
		"tagged versions":     "3",
		"go directive":        "1.20",
		"direct requirements": "3",
		"source":              "https://github.com/Example/mux",
		"stars":               "5600",
	})
}

func TestInspectGitHubFailureKeepsRegistryFacts(t *testing.T) {
	c := newTestClient(t, &fakeServer{gitHubStatus: http.StatusForbidden}, "")
	r, err := c.Inspect(context.Background(), RubyGems, "rack-attack")
	if err != nil {
		t.Fatal(err)
	}
	checkFacts(t, r, map[string]string{"latest version": "6.7.0"})
	if _, ok := factMap(r)["stars"]; ok {
		t.Error("stars reported despite GitHub failure")
	}
	notes := strings.Join(r.Notes, "\n")
	for _, want := range []string{"No GitHub token", "GitHub repository facts unavailable", "rate limit reached"} {
		if !strings.Contains(notes, want) {
			t.Errorf("notes missing %q:\n%s", want, notes)
		}
	}
}

func TestInspectRegistryFailure(t *testing.T) {
	c := newTestClient(t, &fakeServer{gitHubStatus: http.StatusOK}, "")
	if _, err := c.Inspect(context.Background(), RubyGems, "no-such-gem"); err == nil {
		t.Fatal("expected error for unknown gem")
	}
}

func TestAuthorizationHeader(t *testing.T) {
	tests := []struct {
		token, want string
	}{
		{"", ""},
		{"secret", "Bearer secret"},
	}
	for _, tt := range tests {
		f := &fakeServer{gitHubStatus: http.StatusOK}
		c := newTestClient(t, f, tt.token)
		if _, err := c.Inspect(context.Background(), RubyGems, "rack-attack"); err != nil {
			t.Fatal(err)
		}
		if len(f.authHeaders) == 0 {
			t.Fatal("no GitHub requests made")
		}
		for _, got := range f.authHeaders {
			if got != tt.want {
				t.Errorf("token %q: Authorization = %q, want %q", tt.token, got, tt.want)
			}
		}
	}
}

func TestGuess(t *testing.T) {
	tests := []struct {
		name string
		want Ecosystem
	}{
		{"rack-attack", RubyGems},
		{"rails", RubyGems},
		{"github.com/gorilla/mux", Go},
		{"golang.org/x/net", Go},
		{"go.yaml.in/yaml/v3", Go},
		{"foo/bar", RubyGems},
	}
	for _, tt := range tests {
		if got := Guess(tt.name); got != tt.want {
			t.Errorf("Guess(%q) = %q, want %q", tt.name, got, tt.want)
		}
	}
}

func TestEscapePath(t *testing.T) {
	tests := []struct{ in, want string }{
		{"github.com/gorilla/mux", "github.com/gorilla/mux"},
		{"github.com/BurntSushi/toml", "github.com/!burnt!sushi/toml"},
		{"v1.0.0-RC1", "v1.0.0-!r!c1"},
	}
	for _, tt := range tests {
		if got := escapePath(tt.in); got != tt.want {
			t.Errorf("escapePath(%q) = %q, want %q", tt.in, got, tt.want)
		}
	}
}

func TestAge(t *testing.T) {
	tests := []struct {
		ago  time.Duration
		want string
	}{
		{time.Hour, "today"},
		{24 * time.Hour, "1 day ago"},
		{10 * 24 * time.Hour, "10 days ago"},
		{45 * 24 * time.Hour, "1 month ago"},
		{200 * 24 * time.Hour, "6 months ago"},
		{400 * 24 * time.Hour, "1 year ago"},
		{800 * 24 * time.Hour, "2 years ago"},
	}
	for _, tt := range tests {
		if got := age(testNow.Add(-tt.ago), testNow); got != tt.want {
			t.Errorf("age(-%v) = %q, want %q", tt.ago, got, tt.want)
		}
	}
}

func TestParseGoMod(t *testing.T) {
	data, err := os.ReadFile(filepath.Join("testdata", "mod.txt"))
	if err != nil {
		t.Fatal(err)
	}
	goVersion, direct := parseGoMod(string(data))
	if goVersion != "1.20" || direct != 3 {
		t.Errorf("parseGoMod = %q, %d; want 1.20, 3", goVersion, direct)
	}
}

func TestGitHubRepoFromURL(t *testing.T) {
	tests := []struct{ in, want string }{
		{"https://github.com/rack/rack-attack", "rack/rack-attack"},
		{"https://github.com/rack/rack-attack.git", "rack/rack-attack"},
		{"https://github.com/rack/rack-attack/tree/main", "rack/rack-attack"},
		{"https://gitlab.com/a/b", ""},
		{"https://github.com/rack", ""},
		{"", ""},
	}
	for _, tt := range tests {
		if got := gitHubRepoFromURL(tt.in); got != tt.want {
			t.Errorf("gitHubRepoFromURL(%q) = %q, want %q", tt.in, got, tt.want)
		}
	}
}

func TestReportString(t *testing.T) {
	r := Report{Name: "x", Ecosystem: RubyGems,
		Facts: []Fact{{"a", "1"}, {"longer", "2"}},
		Notes: []string{"n"}}
	want := "x (rubygems)\n  a:       1\n  longer:  2\n\nNotes:\n  - n\n"
	if got := r.String(); got != want {
		t.Errorf("String() =\n%q\nwant\n%q", got, want)
	}
}
