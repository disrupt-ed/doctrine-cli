---
name: commit
description: Create a git commit for the current changes using Conventional Commits.
disable-model-invocation: true
---

# Commit

1. Run `git status` and `git diff` (staged and unstaged) to see what changed.
2. Stage only the files that belong to this change. Never stage secrets, `.env` files or unrelated work. If the changes cover unrelated topics, propose separate commits.
3. Write the message in Conventional Commits format:

   ```text
   type(scope): summary

   Optional body.
   ```

   - Types: `feat`, `fix`, `refactor`, `test`, `docs`, `chore`, `perf`, `build`, `ci`.
   - Scope is optional. Use the area changed (e.g. `detect`, `manifest`, `configurator`).
   - Summary: imperative, lowercase, no trailing period, under 72 characters.
   - Body only when the why is not obvious. Short sentences. Say why, not how.
   - Breaking changes: add `!` after the type and a `BREAKING CHANGE:` line.

4. No attribution lines. No `Co-Authored-By`, no tool or AI provider mentions.
5. Commit. Do not push.
6. Show the resulting `git log -1 --stat`.
