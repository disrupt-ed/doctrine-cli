---
name: pr
description: Open a GitHub pull request for the current branch with a user- and product-focused description.
disable-model-invocation: true
---

# Pull Request

1. Check the state: `git status`, current branch, and `git log` / `git diff` against the base branch (usually `main`).
2. If on `main`, create a branch first (`type/short-description`, e.g. `feat/rails-detection`).
3. If there are uncommitted changes, ask whether to commit them first (use the `/commit` conventions).
4. Push the branch with `-u` if it has no upstream.
5. Create the PR with `gh pr create`.

## Title

Conventional Commits format, like a commit summary: `type(scope): summary`. Under 72 characters.

## Description

Write for the user and the product, not for the diff. Short sentences. Straight to the point.

```markdown
## Why

What problem this solves, for whom. One to three sentences.

## What changes

What the user or developer can now do, or what behaves differently.
A few bullets at most.

## How to verify

Steps or commands to see it working.
```

- Skip a section if it adds nothing.
- No file-by-file walkthroughs. The diff already shows that.
- Mention breaking changes and follow-up work explicitly, in one line each.
- No attribution lines. No tool or AI provider mentions.

6. Return the PR URL.
