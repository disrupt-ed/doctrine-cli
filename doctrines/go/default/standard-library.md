# Go Already Does This

Check this list before adding a module. The version is when the feature arrived. Check it against the `go` line in `go.mod`.

| Need | Standard library | Since |
|---|---|---|
| HTTP routing with methods and path parameters | `net/http` `ServeMux` patterns (`"GET /items/{id}"`, `r.PathValue`) | 1.22 |
| HTTP server and client, timeouts | `net/http` | |
| Test HTTP handlers and clients | `net/http/httptest` | |
| Structured logging | `log/slog` | 1.21 |
| JSON | `encoding/json` | |
| SQL with any driver | `database/sql` | |
| Command-line flags | `flag` | |
| Environment config | `os.Getenv`, `os.LookupEnv` | |
| HTML and text templates | `html/template`, `text/template` | |
| Embed files in the binary | `embed` | 1.16 |
| Generic helpers for slices and maps | `slices`, `maps` | 1.21 |
| Min and max | `min`, `max` builtins | 1.21 |
| Iterators | range-over-func, `iter` | 1.23 |
| Cancellation and deadlines | `context` | |
| Wait for goroutines | `sync.WaitGroup` | |
| Run something once, lazily | `sync.Once`, `sync.OnceValue` | 1.21 |
| Concurrent map | `sync.Map`, or a map with `sync.Mutex` | |
| Tickers and timers | `time` | |
| Random tokens | `crypto/rand` | |
| Hashing, HMAC | `crypto/sha256`, `crypto/hmac` | |
| Password hashing | `golang.org/x/crypto/bcrypt` or `argon2` (x/ module) | |
| Error wrapping and joining | `fmt.Errorf` with `%w`, `errors.Join` | 1.13, 1.20 |
| Profiling | `net/http/pprof`, `runtime/pprof` | |
| Fuzzing | `testing` fuzz targets | 1.18 |

## Modules that are often not needed

- **Routers** (gorilla/mux, chi, gin, echo): `ServeMux` handles methods and path parameters since Go 1.22. Keep the router the project already uses.
- **Loggers** (zap, zerolog, logrus): `log/slog` is structured and fast enough for most services.
- **Assertion libraries** (testify): plain `if` checks with `t.Errorf` are the Go norm.
- **Config libraries** (viper, envconfig): `flag` and `os.Getenv` cover most services.
- **CLI frameworks** (cobra, urfave/cli): `flag` with a subcommand switch is enough for a handful of commands.
- **Retry and backoff libraries**: a loop with `time.Sleep` or a timer and a context.
- **Utility collections** (lo, funk): `slices` and `maps`.

Keep these if the project already uses them.
