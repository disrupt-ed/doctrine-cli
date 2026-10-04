# Agent Doctrine

**Teach your AI how you build software.**

> Status: early development. Nothing here is released yet.

AI coding agents can write code. What they don't consistently have is engineering judgment. They add dependencies they don't need, invent abstractions, ignore what the framework already provides and write more code than necessary.

Agent Doctrine gives your coding agent stack-aware engineering judgment: Rails-aware for a Rails app, Go-aware for a Go service, both for a repository with both. It influences decisions. It never blocks them. You always have the final word.

## How it works

```text
Your repository ──► doctrine init ──► .doctrine/doctrine.yml ──► agent configuration
```

1. `doctrine` detects your stack (Ruby, Rails, Go).
2. It picks strong default doctrines for it.
3. It generates configuration for your coding agent (Claude Code first).

Doctrines are plain Markdown and YAML. You can read every word that influences your AI.

## Quick start

```bash
brew install disrupt-ed/tap/doctrine
cd your-project
doctrine init
```

That's it. `doctrine init` detects your stack, writes `.doctrine/doctrine.yml` and generates Claude Code configuration in `.claude/rules/doctrine/` and `.claude/skills/dependency-review/`. Commit those files. Your `CLAUDE.md` is never touched.

No Homebrew? Download a binary from [GitHub Releases](https://github.com/disrupt-ed/doctrine-cli/releases), or `go install github.com/disrupt-ed/doctrine-cli/cmd/doctrine@latest`.

### Install with Claude Code

Paste this into Claude Code in your repository:

```text
Install Agent Doctrine in this repository: install the doctrine CLI
(brew install disrupt-ed/tap/doctrine, or go install github.com/disrupt-ed/doctrine-cli/cmd/doctrine@latest),
run `doctrine init`, show me `doctrine detect` and the files it generated, and don't commit.
```

## Commands

```bash
doctrine detect                    # show detected technologies and suggested doctrines
doctrine init [stacks...]          # set up Doctrine in this repository
doctrine install HANDLE            # install a configuration made on doctrine.codedynamic.com
doctrine inspect [doctrine...]     # print the full effective Doctrine, or one doctrine
doctrine list                      # list official doctrines
doctrine generate [claude]         # regenerate agent configuration
doctrine add rails | remove go     # change which doctrines apply
doctrine update                    # review and accept a newer doctrine release
doctrine dependency inspect NAME   # report facts about a gem or Go module
```

## Official doctrines

```text
engineering/default   KISS, YAGNI, minimal dependencies, late abstraction
ruby/default          Ruby idioms and standard library first
rails/default         "Can Rails already do this?" before "which gem?"
go/default            Standard library, explicit code, small interfaces
```

They live in [`doctrines/`](doctrines/). Read them, fork them, open a pull request.

## Customizing

Defaults are opinionated. Your project can disagree.

```text
.doctrine/
├── doctrine.yml     # which doctrines apply, and your preferences
└── local/
    └── rails.md     # your project's overrides
```

```markdown
# Rails Overrides

Service objects are appropriate for external payment integrations under app/services/payments/.
```

Project overrides always win. Updates never touch your files and never change your AI's behavior without your review.

## Principles

- **Influence, don't block.** The agent explains concerns. You decide.
- **Transparent.** No hidden prompts. Every doctrine is readable.
- **Stack-aware.** Good Rails is different from good Go.
- **Works offline.** No account, no server required. doctrine.codedynamic.com is optional.
- **Agent-independent.** Claude Code first. Codex, Cursor and others later.

## Privacy

The CLI can send anonymous usage stats (language, framework, OS and agent). It asks first and shows exactly what would be sent. Off by default. `DO_NOT_TRACK` and `DOCTRINE_TELEMETRY=0` are always respected. It never sends code or anything that identifies your repository.

## Development

```bash
go test ./...
go vet ./...
go run ./cmd/doctrine help
```

The format is specified in [`specification/`](specification/). The evaluation lives in [`eval/`](eval/).

### Releasing

- **CLI:** push a tag like `v0.1.0`. GitHub Actions builds macOS, Linux and Windows binaries, publishes the release and updates the Homebrew formula in `disrupt-ed/homebrew-tap` (needs the `HOMEBREW_TAP_TOKEN` secret).
- **Doctrines:** when `doctrines.Release` changes, also push a tag `doctrines-v<release>`, like `doctrines-v0.1`. Projects pinned to an older release download it from that tag.

See [`CLAUDE.md`](CLAUDE.md) for the engineering conventions this codebase follows. It eats its own dog food.

## License

To be decided.
