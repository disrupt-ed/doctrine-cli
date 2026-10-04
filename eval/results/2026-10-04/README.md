# First run, 2026-10-04

Model: `claude-sonnet-5-5`, Claude Code 2.1.289. 26 tasks, each run once without Doctrine (baseline) and once with it. Doctrine release 0.1 as first drafted.

## Blind pre-review by a model, not the real review

The plan's review is by an experienced developer. That hasn't happened yet. To get an early signal, a Claude agent reviewed each pair blind (A/B from `review/`, never the key) and picked the one it would rather merge. Its notes are in `model-prereview-*.yml`.

| App | Doctrine preferred | Baseline preferred | No difference |
|---|---|---|---|
| Campfire (Rails) | 7 | 1 | 5 |
| Miniflux (Go) | 7 | 4 | 2 |
| **Total** | **14** | **5** | **7** |

## Counted

| | Campfire baseline | Campfire doctrine | Miniflux baseline | Miniflux doctrine |
|---|---|---|---|---|
| Dependencies added | 1 (faraday, insisted) | 2 (faraday, insisted; csv, needed on Ruby 3.4) | 0 | 0 |
| Lines added | 524 | 469 | 1256 | 1250 |
| Cost (USD) | 2.25 | 2.72 | 3.66 | 3.86 |

The baseline was already good at avoiding unneeded dependencies: neither arm added kaminari, metainspector, gorilla/mux, a retry library or a feature-flag gem. The differences are in judgment and care, not in dependency counts.

## Where Doctrine helped

- Used the standard library correctly where the baseline hand-rolled it (CSV quoting).
- Found root causes before patching (sidebar N+1; an existing Prometheus series instead of a new counter).
- Raised the convention before adding testify, then did what the developer insisted on.
- Caught subtle issues (an admin link inside a shared cache fragment, a broken Stimulus action string, SSRF on webhook URLs).

## Where Doctrine did not help

- **Too cautious once.** Asked for pagination with kaminari, it found the app's existing pagination and only asked a question. The baseline made the two-line fix with the existing tool and said so.
- **Skipped verification.** Twice the doctrine run didn't build or test its code. One left a test that doesn't compile, the other has a retry check that is too broad.
- **Smaller misses** on tasks where both were close: forgot to evict idle rate-limit buckets, dropped a doc comment.

## Changed after this run

`engineering/default`, `working-with-the-human.md`:

- When the simpler path clearly meets the goal, take it, say so, and offer the original. Ask only when the choice changes the result.
- Build and run the tests you touched before saying you're done.

The re-run with these changes is in `../2026-10-04-rerun/`.

## Limits of this run

- One run per task and arm. Model output varies between runs.
- The agents couldn't use the network, `bundle` or `go get` (see the eval README). This blocked the dependency-review skill from checking facts online in both arms.
- Some Campfire tests already fail at the pinned commit, the same for both arms.
