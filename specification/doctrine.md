# Doctrine Format

## Folder

An official doctrine is a folder named `<topic>/<variant>`, for example `rails/default`:

```text
rails/default/
├── doctrine.yml
├── the-rails-way.md
└── ...
```

`doctrine.yml`:

```yaml
name: rails/default
description: Can Rails already do this? before which gem?
parent: ruby/default     # optional
files:                   # Markdown files, in order
  - the-rails-way.md
```

Markdown files are written for the agent: plain instructions, short.

## Releases

All official doctrines are released together under one version, for example `0.1`. The CLI ships the release it was built with. A project pinned to another release gets it from the public repository (tag `doctrines-v<version>`) and caches it in the user cache directory.

## Composition

1. Start from the manifest's `extends`.
2. Add each doctrine's parents, parents first.
3. Remove duplicates, keeping the first occurrence.

`extends: [rails/default, go/default]` gives:

```text
engineering/default
ruby/default
rails/default
go/default
```

Each doctrine's files are joined in the order of its `files` list, separated by one blank line.

## Overrides

The project's preferences (`engineering` in the manifest) and the files in `.doctrine/local/` (alphabetical order) come last, under a "Project Overrides" heading that tells the agent they win over the doctrine.

There is no merging. Overrides are appended, and the agent is told which part wins.

## Generated output (Claude Code)

| File | Content |
|---|---|
| `.claude/rules/doctrine/<topic>.md` | One file per doctrine, e.g. `engineering.md`, `rails.md`. |
| `.claude/rules/doctrine/project.md` | Project overrides, only when there are any. |
| `.claude/skills/dependency-review/SKILL.md` | Dependency review skill. |

Claude Code loads `.claude/rules/` automatically, so the project's `CLAUDE.md` is never touched. The CLI owns `.claude/rules/doctrine/` and `.claude/skills/dependency-review/` and replaces their contents on every `generate`. Generation is deterministic: the same manifest and overrides always give the same files.
