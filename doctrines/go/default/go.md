# Go

Go values clarity over cleverness. A little copying is better than a little dependency. Write the Go an experienced Go developer would expect.

## Check the version first

Read the `go` line in `go.mod` before using a newer feature (generics, range-over-func, `ServeMux` patterns, `min` and `max`).

## Standard library first

Go's standard library covers most of what a service needs: HTTP servers and clients, routing, JSON, SQL, logging, templates, crypto, testing. Check it before adding a module. See the standard library list in this doctrine.

`golang.org/x/...` modules are maintained by the Go team and are the next place to look. They are still a dependency, so say why.

## Explicit code

- Write the obvious code. Plain loops, plain `if err != nil`, plain structs.
- Handle every error. Wrap it with context: `fmt.Errorf("load config: %w", err)`. Use `errors.Is` and `errors.As` to check.
- Don't panic in normal flow. Return errors.
- Avoid reflection and code generation unless the codebase already uses them for that job.
- Zero values should be useful. Avoid constructors that only set defaults.

## Interfaces

- Define interfaces where they are used, not where they are implemented.
- Keep them small: one or two methods.
- Don't add an interface for a single implementation just for testing. Test the real thing, or use a small fake where the interface is consumed.
- Accept interfaces, return concrete types.

## Packages

- Name packages for what they provide, not `util`, `common` or `helpers`.
- Keep package boundaries simple. Don't split a package until it is hard to work with.
- Use `internal/` to keep code private to the module.
- Follow the layout the project already has.

## Concurrency

- Don't add goroutines unless they make a real difference.
- Every goroutine needs a clear way to stop. Pass a `context.Context` through anything that does I/O.
- Prefer `sync.WaitGroup`, `errgroup` (if already used) and channels with clear ownership. Protect shared state with a mutex. Keep it simple.
- Run tests with `-race` when you touch concurrent code.

## Tests

- Use the standard `testing` package. Table-driven tests where they fit.
- Use `t.Helper()`, `t.Cleanup()`, `t.TempDir()` and `httptest`.
- Don't add an assertion library if the codebase doesn't use one.
- Use golden files for large outputs.

## Style

`gofmt` and `go vet` must pass. Follow the project's linter if it has one. Short names for short scopes, clearer names for wider ones.
