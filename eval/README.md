# Evaluation

Does Claude Code + Doctrine make better engineering decisions than Claude Code alone?

The Miniflux tasks were written before the doctrines, so the doctrines couldn't be tuned to them. The Campfire tasks replaced the Lobsters ones after the doctrines existed; they mirror the Lobsters tasks one for one, adapted to Campfire's code.

## Apps

| App | Stack | Repository | Pinned commit |
|---|---|---|---|
| Campfire | Rails main (8.2), SQLite, Resque, Minitest | https://github.com/basecamp/once-campfire | `90b330024dec3e757c79b6a7e6568f93da8e3148` |
| Miniflux | Go 1.26, PostgreSQL | https://github.com/miniflux/v2 | `c52bdef6e9811f7c20cebc034a6c0ec6799dc6df` |

Real, maintained open-source apps. Small invented apps don't have the existing code and conventions that the doctrines are meant to make the agent respect.

Lobsters was the first Rails pick. Its `AGENTS.md` forbids any LLM work on the project, so we don't use it.

## Tasks

- [`tasks/campfire.yml`](tasks/campfire.yml)
- [`tasks/miniflux.yml`](tasks/miniflux.yml)

Each task has a `prompt` written the way a developer would type it. Some have a `followup`: send it as the next message, whatever the agent answered. `watch` lists what the reviewer should look at. It describes the task, not the doctrine.

Some prompts are traps on purpose. They ask for a gem or module the app doesn't need, or for something the app already does.

## Running a task

Each task runs twice from a clean checkout of the pinned commit:

1. **baseline**: Claude Code with the app's own files, nothing else.
2. **doctrine**: same, plus the Doctrine output for that app from [`claude-output/`](claude-output/). That folder is also the golden output for `doctrine generate`.

The runner does this headless, from `cli/`:

```bash
go run ./eval/run tasks                          # every app, task and arm
go run ./eval/run tasks -app miniflux -task fetch-retry -arm doctrine
go run ./eval/run blind -results eval/results/2026-10-04
```

Every run uses the same model (`-model`, default `claude-sonnet-5-5`) and `--setting-sources project,local`, so the user's own `~/.claude` instructions don't leak into either arm. Each agent can read and edit files, read git, build, test and run generators. Everything else is denied without asking: no network, no `bundle install`, no `go get`. This keeps unattended runs safe. It also means a run can import a new module without adding it to `go.mod`, and can't check a dependency's health online. Both arms have the same limits.

Campfire needs Ruby 3.4.10 (`mise install ruby@3.4.10`, then `bundle install` once in a checkout).

Each run is saved under `results/<date>/<app>/<task-id>/<baseline|doctrine>/`:

- `diff.patch`: the change, new files included;
- `transcript.md` and `transcript.jsonl`: the conversation;
- `metrics.yml`: see below.

## Metrics

The runner counts what can be counted. The reviewer fills in the rest.

```yaml
dependencies_added: []               # gems or Go modules added to Gemfile or go.mod
lines_added: 0
lines_removed: 0
files_added: 0
files_changed: 0
test_files_changed: 0
cost_usd: 0
duration_seconds: 0
turns: 0
abandoned_dependency_chosen: null    # reviewer
framework_native: null               # reviewer: used what the language, framework or app already provides
followed_existing_patterns: null     # reviewer
asked_or_explained_tradeoff: null    # reviewer
deferred_to_human: null              # reviewer, follow-up tasks only
tests_pass: null                     # reviewer
correct: null                        # reviewer
```

## Blind review

An experienced Rails and Go developer gets both diffs and transcripts for each task, labelled A and B in random order (`run blind` writes them to `review/`, with words that would give the arm away redacted, and the key to `review-key.yml`). For each task they answer:

- Which one would you rather merge? (A, B, or no difference)
- Why, in one or two sentences.

Unblind only after all tasks are reviewed.

## Go / no-go

Go if the reviewer prefers the doctrine run on clearly more tasks than the baseline, and the doctrine run never refuses what the developer insisted on. Otherwise, read the transcripts where it lost and change the doctrines.
