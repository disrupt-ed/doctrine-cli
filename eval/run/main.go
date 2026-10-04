// Command run executes the evaluation tasks with Claude Code, with and
// without Doctrine, and prepares blind review packets.
//
//	go run ./eval/run tasks [-app NAME] [-task ID] [-arm baseline|doctrine] [-parallel N]
//	go run ./eval/run blind -results DIR
package main

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"math/rand/v2"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
	"time"

	"go.yaml.in/yaml/v3"
)

type app struct {
	repo   string
	commit string
	env    []string // extra environment for the agent, e.g. the Ruby version
}

var apps = map[string]app{
	"campfire": {
		repo:   "https://github.com/basecamp/once-campfire.git",
		commit: "90b330024dec3e757c79b6a7e6568f93da8e3148",
		env:    []string{"MISE_RUBY_VERSION=3.4.10"},
	},
	"miniflux": {
		repo:   "https://github.com/miniflux/v2.git",
		commit: "c52bdef6e9811f7c20cebc034a6c0ec6799dc6df",
	},
}

var arms = []string{"baseline", "doctrine"}

type task struct {
	ID       string   `yaml:"id"`
	Prompt   string   `yaml:"prompt"`
	Followup string   `yaml:"followup"`
	Watch    []string `yaml:"watch"`
}

type run struct {
	app  string
	task task
	arm  string
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "usage: run tasks|blind [flags]")
		os.Exit(2)
	}
	var err error
	switch os.Args[1] {
	case "tasks":
		err = runTasks(os.Args[2:])
	case "blind":
		err = blind(os.Args[2:])
	default:
		err = fmt.Errorf("unknown command %q", os.Args[1])
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}

func runTasks(args []string) error {
	fs := flag.NewFlagSet("tasks", flag.ExitOnError)
	appName := fs.String("app", "", "only this app")
	taskID := fs.String("task", "", "only this task")
	armName := fs.String("arm", "", "only this arm (baseline or doctrine)")
	model := fs.String("model", "claude-sonnet-5-5", "Claude model")
	parallel := fs.Int("parallel", 4, "runs at the same time")
	results := fs.String("results", filepath.Join("eval", "results", time.Now().Format("2006-01-02")), "results directory")
	work := fs.String("work", filepath.Join(os.TempDir(), "doctrine-eval"), "directory for checkouts")
	output := fs.String("output", filepath.Join("eval", "claude-output"), "Doctrine output per app, used by the doctrine arm")
	fs.Parse(args)

	var runs []run
	for name := range apps {
		if *appName != "" && name != *appName {
			continue
		}
		tasks, err := loadTasks(filepath.Join("eval", "tasks", name+".yml"))
		if err != nil {
			return err
		}
		for _, t := range tasks {
			if *taskID != "" && t.ID != *taskID {
				continue
			}
			for _, a := range arms {
				if *armName != "" && a != *armName {
					continue
				}
				runs = append(runs, run{app: name, task: t, arm: a})
			}
		}
	}
	if len(runs) == 0 {
		return fmt.Errorf("no runs match")
	}

	for name := range apps {
		if *appName != "" && name != *appName {
			continue
		}
		if err := ensureCache(*work, name); err != nil {
			return err
		}
	}

	r := runner{model: *model, results: *results, work: *work, output: *output}
	sem := make(chan struct{}, *parallel)
	var wg sync.WaitGroup
	var mu sync.Mutex
	var failed []string
	for _, rn := range runs {
		wg.Add(1)
		sem <- struct{}{}
		go func() {
			defer wg.Done()
			defer func() { <-sem }()
			start := time.Now()
			err := r.do(rn)
			mu.Lock()
			defer mu.Unlock()
			label := fmt.Sprintf("%s/%s/%s", rn.app, rn.task.ID, rn.arm)
			if err != nil {
				failed = append(failed, label)
				fmt.Printf("FAIL %s: %v\n", label, err)
				return
			}
			fmt.Printf("done %s (%s)\n", label, time.Since(start).Round(time.Second))
		}()
	}
	wg.Wait()
	if len(failed) > 0 {
		return fmt.Errorf("%d runs failed: %s", len(failed), strings.Join(failed, ", "))
	}
	return nil
}

func loadTasks(path string) ([]task, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var f struct {
		Tasks []task `yaml:"tasks"`
	}
	if err := yaml.Unmarshal(data, &f); err != nil {
		return nil, fmt.Errorf("parse %s: %w", path, err)
	}
	return f.Tasks, nil
}

func ensureCache(work, name string) error {
	dir := filepath.Join(work, "cache", name)
	if _, err := os.Stat(dir); err == nil {
		return nil
	}
	a := apps[name]
	if err := git("", "clone", "--quiet", a.repo, dir); err != nil {
		return err
	}
	return git(dir, "checkout", "--quiet", a.commit)
}

type runner struct {
	model, results, work, output string
}

func (r runner) do(rn run) error {
	a := apps[rn.app]
	out := filepath.Join(r.results, rn.app, rn.task.ID, rn.arm)
	if err := os.MkdirAll(out, 0o755); err != nil {
		return err
	}

	dir, err := os.MkdirTemp(filepath.Join(r.work), rn.app+"-"+rn.task.ID+"-"+rn.arm+"-")
	if err != nil {
		return err
	}
	if err := git("", "clone", "--quiet", "--shared", filepath.Join(r.work, "cache", rn.app), dir); err != nil {
		return err
	}
	if err := git(dir, "checkout", "--quiet", a.commit); err != nil {
		return err
	}
	if rn.arm == "doctrine" {
		if err := os.CopyFS(dir, os.DirFS(filepath.Join(r.output, rn.app))); err != nil {
			return fmt.Errorf("copy doctrine output: %w", err)
		}
		if err := git(dir, "add", "-A"); err != nil {
			return err
		}
		if err := git(dir, "-c", "user.name=eval", "-c", "user.email=eval@localhost", "commit", "--quiet", "-m", "eval setup"); err != nil {
			return err
		}
	}

	transcript, err := os.Create(filepath.Join(out, "transcript.jsonl"))
	if err != nil {
		return err
	}
	defer transcript.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 45*time.Minute)
	defer cancel()

	res, err := r.claude(ctx, dir, a.env, rn.task.Prompt, "", transcript)
	if err != nil {
		return err
	}
	total := res
	if rn.task.Followup != "" {
		res2, err := r.claude(ctx, dir, a.env, rn.task.Followup, res.SessionID, transcript)
		if err != nil {
			return err
		}
		total.CostUSD += res2.CostUSD
		total.DurationMS += res2.DurationMS
		total.NumTurns += res2.NumTurns
	}

	if err := git(dir, "add", "-A"); err != nil {
		return err
	}
	diff, err := gitOutput(dir, "diff", "--cached", "HEAD")
	if err != nil {
		return err
	}
	if err := os.WriteFile(filepath.Join(out, "diff.patch"), diff, 0o644); err != nil {
		return err
	}
	if err := writeTranscriptMarkdown(filepath.Join(out, "transcript.jsonl"), filepath.Join(out, "transcript.md")); err != nil {
		return err
	}
	m, err := measure(dir)
	if err != nil {
		return err
	}
	m.CostUSD = total.CostUSD
	m.DurationSeconds = total.DurationMS / 1000
	m.Turns = total.NumTurns
	data, err := yaml.Marshal(m)
	if err != nil {
		return err
	}
	if err := os.WriteFile(filepath.Join(out, "metrics.yml"), data, 0o644); err != nil {
		return err
	}
	return os.RemoveAll(dir)
}

type result struct {
	SessionID  string  `json:"session_id"`
	CostUSD    float64 `json:"total_cost_usd"`
	DurationMS int     `json:"duration_ms"`
	NumTurns   int     `json:"num_turns"`
	IsError    bool    `json:"is_error"`
}

var allowedTools = []string{
	"Read", "Edit", "Write", "Glob", "Grep", "Skill", "TodoWrite",
	"Bash(git status:*)", "Bash(git diff:*)", "Bash(git log:*)", "Bash(git show:*)",
	"Bash(ls:*)", "Bash(cat:*)", "Bash(head:*)", "Bash(grep:*)", "Bash(find:*)",
	"Bash(go build:*)", "Bash(go test:*)", "Bash(go vet:*)", "Bash(gofmt:*)",
	"Bash(bin/rails test:*)", "Bash(bin/rails generate:*)", "Bash(bin/rails db:migrate:*)",
	"Bash(bin/rails db:prepare:*)", "Bash(bin/rubocop:*)",
}

// claude runs one headless Claude Code turn. User-level settings and
// instructions are excluded so both arms see only the project.
func (r runner) claude(ctx context.Context, dir string, env []string, prompt, resume string, transcript *os.File) (result, error) {
	args := []string{
		"-p", prompt,
		"--model", r.model,
		"--setting-sources", "project,local",
		"--output-format", "stream-json", "--verbose",
		// Only file tools and the commands needed to build and test.
		// Anything else is denied without asking. Same for both arms.
		"--permission-mode", "dontAsk",
		"--allowedTools", strings.Join(allowedTools, ","),
	}
	if resume != "" {
		args = append(args, "--resume", resume)
	}
	cmd := exec.CommandContext(ctx, "claude", args...)
	cmd.Dir = dir
	cmd.Env = append(os.Environ(), env...)
	cmd.Stderr = os.Stderr
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return result{}, err
	}
	if err := cmd.Start(); err != nil {
		return result{}, err
	}
	var res result
	sc := bufio.NewScanner(stdout)
	sc.Buffer(make([]byte, 1024*1024), 64*1024*1024)
	for sc.Scan() {
		line := sc.Bytes()
		transcript.Write(line)
		transcript.Write([]byte("\n"))
		var ev struct {
			Type string `json:"type"`
		}
		if json.Unmarshal(line, &ev) == nil && ev.Type == "result" {
			json.Unmarshal(line, &res)
		}
	}
	if err := cmd.Wait(); err != nil {
		return res, fmt.Errorf("claude: %w", err)
	}
	if res.SessionID == "" {
		return res, fmt.Errorf("claude: no result event")
	}
	return res, nil
}

type metrics struct {
	DependenciesAdded   []string `yaml:"dependencies_added"`
	LinesAdded          int      `yaml:"lines_added"`
	LinesRemoved        int      `yaml:"lines_removed"`
	FilesAdded          int      `yaml:"files_added"`
	FilesChanged        int      `yaml:"files_changed"`
	TestFilesChanged    int      `yaml:"test_files_changed"`
	CostUSD             float64  `yaml:"cost_usd"`
	DurationSeconds     int      `yaml:"duration_seconds"`
	Turns               int      `yaml:"turns"`
	AbandonedDependency *bool    `yaml:"abandoned_dependency_chosen"`
	FrameworkNative     *bool    `yaml:"framework_native"`
	FollowedPatterns    *bool    `yaml:"followed_existing_patterns"`
	ExplainedTradeoff   *bool    `yaml:"asked_or_explained_tradeoff"`
	DeferredToHuman     *bool    `yaml:"deferred_to_human"`
	TestsPass           *bool    `yaml:"tests_pass"`
	Correct             *bool    `yaml:"correct"`
}

var (
	gemLine       = regexp.MustCompile(`^\+\s*gem\s+["']([^"']+)["']`)
	goRequireLine = regexp.MustCompile(`^\+\s*(?:require\s+)?([a-z0-9.\-]+\.[a-z]{2,}/[^\s]+)\s+v[0-9]`)
)

// measure counts what can be counted from the staged diff. The judgment
// fields stay empty for the reviewer.
func measure(dir string) (metrics, error) {
	var m metrics
	numstat, err := gitOutput(dir, "diff", "--cached", "--numstat", "HEAD")
	if err != nil {
		return m, err
	}
	for _, line := range strings.Split(strings.TrimSpace(string(numstat)), "\n") {
		f := strings.Fields(line)
		if len(f) < 3 {
			continue
		}
		var add, del int
		fmt.Sscan(f[0], &add)
		fmt.Sscan(f[1], &del)
		m.LinesAdded += add
		m.LinesRemoved += del
		m.FilesChanged++
		if isTestFile(f[2]) {
			m.TestFilesChanged++
		}
	}
	status, err := gitOutput(dir, "diff", "--cached", "--name-status", "HEAD")
	if err != nil {
		return m, err
	}
	for _, line := range strings.Split(string(status), "\n") {
		if strings.HasPrefix(line, "A\t") {
			m.FilesAdded++
		}
	}
	deps, err := gitOutput(dir, "diff", "--cached", "HEAD", "--", "Gemfile", "go.mod")
	if err != nil {
		return m, err
	}
	for _, line := range strings.Split(string(deps), "\n") {
		if strings.Contains(line, "// indirect") {
			continue
		}
		if s := gemLine.FindStringSubmatch(line); s != nil {
			m.DependenciesAdded = append(m.DependenciesAdded, s[1])
		} else if s := goRequireLine.FindStringSubmatch(line); s != nil {
			m.DependenciesAdded = append(m.DependenciesAdded, s[1])
		}
	}
	return m, nil
}

func isTestFile(path string) bool {
	return strings.HasSuffix(path, "_test.go") || strings.HasPrefix(path, "test/") || strings.HasPrefix(path, "spec/")
}

// writeTranscriptMarkdown turns the stream-json events into a readable
// conversation: prompts, the agent's text and the tools it called.
func writeTranscriptMarkdown(jsonl, md string) error {
	data, err := os.ReadFile(jsonl)
	if err != nil {
		return err
	}
	var b strings.Builder
	for _, line := range bytes.Split(data, []byte("\n")) {
		var ev struct {
			Type    string `json:"type"`
			Message struct {
				Content json.RawMessage `json:"content"`
			} `json:"message"`
		}
		if json.Unmarshal(line, &ev) != nil {
			continue
		}
		var blocks []struct {
			Type  string          `json:"type"`
			Text  string          `json:"text"`
			Name  string          `json:"name"`
			Input json.RawMessage `json:"input"`
		}
		json.Unmarshal(ev.Message.Content, &blocks)
		for _, bl := range blocks {
			switch {
			case ev.Type == "assistant" && bl.Type == "text":
				fmt.Fprintf(&b, "**Agent:** %s\n\n", bl.Text)
			case ev.Type == "assistant" && bl.Type == "tool_use":
				fmt.Fprintf(&b, "> %s `%s`\n\n", bl.Name, truncate(string(bl.Input), 300))
			}
		}
	}
	return os.WriteFile(md, []byte(b.String()), 0o644)
}

func truncate(s string, n int) string {
	if len(s) <= n {
		return s
	}
	return s[:n] + "…"
}

// tell matches what would reveal the doctrine arm to a reviewer.
var tell = regexp.MustCompile(`(?i)[\w./-]*doctrine[\w./-]*`)

// blind copies each task's two runs into A and B folders in random order,
// and writes the key separately.
func blind(args []string) error {
	fs := flag.NewFlagSet("blind", flag.ExitOnError)
	results := fs.String("results", "", "results directory")
	fs.Parse(args)
	if *results == "" {
		return fmt.Errorf("-results is required")
	}
	key := map[string]map[string]string{}
	appDirs, err := os.ReadDir(*results)
	if err != nil {
		return err
	}
	for _, ad := range appDirs {
		if !ad.IsDir() || ad.Name() == "review" {
			continue
		}
		taskDirs, err := os.ReadDir(filepath.Join(*results, ad.Name()))
		if err != nil {
			return err
		}
		for _, td := range taskDirs {
			order := []string{"baseline", "doctrine"}
			if rand.IntN(2) == 1 {
				order[0], order[1] = order[1], order[0]
			}
			id := ad.Name() + "/" + td.Name()
			key[id] = map[string]string{"A": order[0], "B": order[1]}
			for i, label := range []string{"A", "B"} {
				src := filepath.Join(*results, ad.Name(), td.Name(), order[i])
				dst := filepath.Join(*results, "review", ad.Name(), td.Name(), label)
				if err := os.MkdirAll(dst, 0o755); err != nil {
					return err
				}
				for _, f := range []string{"diff.patch", "transcript.md"} {
					data, err := os.ReadFile(filepath.Join(src, f))
					if err != nil {
						return err
					}
					data = tell.ReplaceAll(data, []byte("[redacted]"))
					if err := os.WriteFile(filepath.Join(dst, f), data, 0o644); err != nil {
						return err
					}
				}
			}
		}
	}
	data, err := yaml.Marshal(key)
	if err != nil {
		return err
	}
	return os.WriteFile(filepath.Join(*results, "review-key.yml"), data, 0o644)
}

func git(dir string, args ...string) error {
	_, err := gitOutput(dir, args...)
	return err
}

func gitOutput(dir string, args ...string) ([]byte, error) {
	cmd := exec.Command("git", args...)
	cmd.Dir = dir
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	out, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("git %s: %w: %s", strings.Join(args, " "), err, strings.TrimSpace(stderr.String()))
	}
	return out, nil
}
