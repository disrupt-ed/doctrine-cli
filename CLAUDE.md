# Doctrine CLI

Open-source Go CLI. It detects a repository's stack, composes the Doctrine and generates agent configuration (Claude Code first).

Product and technical definitions live in `../.doc/` (`definition.md`, `tech-stack.md`). Read them before making product or architecture decisions.

## Eat our own dog food

Doctrine teaches AI agents to make good engineering decisions. This codebase must show those decisions.

- Simplest solution that works. Complexity has to earn its place.
- Build only what is needed today (YAGNI).
- Standard library first. Ask "can Go already do this?" before adding a module.
- Every new dependency needs a stated reason. Check its health first: last release, last commit, contributors, open issues.
- Avoid abstraction until real complexity or multiple implementations justify it.
- Small interfaces, defined where they are used.
- Explicit code over clever code.
- Follow existing patterns in the codebase before introducing new ones.
- Explain significant tradeoffs. The human decides.

## Go conventions

- Layout: `cmd/doctrine/` for the entry point, `internal/` for everything else.
- Packages: `detect`, `doctrine`, `manifest`, `compose`, `generate`, `dependencies`, `git`, `adapters`.
- Official doctrines live in `doctrines/`, as YAML + Markdown.
- Return errors, wrap them with context (`fmt.Errorf("...: %w", err)`). No panics in normal flow.
- Use `flag` from the standard library until it is clearly not enough.
- Keep concurrency straightforward. Only add it when it makes a real difference.
- `gofmt` and `go vet` must pass.

## Testing

- Standard `testing` package only. No test frameworks.
- Table-driven tests where they fit.
- Golden-file tests for generated agent configuration.
- Cover detection, resolution, inheritance, overrides, manifest parsing, generation and the Claude adapter.
- Run `go test ./...` before considering work done.

## Product rules the code must respect

- Works fully offline. `doctrine init` never needs doctrine.dev.
- Never silently change what influences the AI. Updates show a diff and need the developer's acceptance.
- `.doctrine/` only contains files the developer owns. Never overwrite them.
- The CLI is the only generator. The website calls this binary to build its ZIP.
- Dependency inspection reports raw facts. It does not judge whether activity is "meaningful".
- Usage stats are opt-in. Respect `DO_NOT_TRACK` and `DOCTRINE_TELEMETRY=0`. Never send code or anything identifying the repository. Never block or fail a command because of stats.
