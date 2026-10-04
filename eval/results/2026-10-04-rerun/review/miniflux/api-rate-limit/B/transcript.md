**Agent:** I'll start by looking at how the API is wired up.

> Grep `{"pattern":"APIKey|api_key|RateLimit|rate_limit","glob":"internal/{api,http/request,http/server,config,storage}/*.go","output_mode":"content","head_limit":60}`

> Glob `{"pattern":"internal/api/*.go"}`

> Read `{"file_path":"[redacted]"}`

> Read `{"file_path":"[redacted]","limit":80}`

> Grep `{"pattern":"TooManyRequests|ratelimit|rate.limit|x/time","-i":true,"glob":"*.{go,mod}","output_mode":"content","head_limit":20}`

> Bash `{"command":"cat go.mod | head -30; ls vendor 2>/dev/null | head -2; grep -n \"func JSON\\|func Too\\|func NewBuilder\" internal/http/response/*.go; grep -n -i \"BatchSize\\|BATCH_SIZE\" internal/config/*.go | head -30; ls internal/api/*_test.go; sed -n 1,40p internal/api/api_test.go"}`

> Bash `{"command":"sed -n 95,140p internal/config/options.go; sed -n 95,120p internal/http/response/json.go; sed -n 640,660p internal/config/options.go; grep -n \"BATCH_SIZE\" -r --include=*.1 --include=*.md . | head; grep -n \"func.*String\\|func.*Int\\|func.*Bool\" internal/config/parser.go | head"}`

> Bash `{"command":"grep -rln \"CLEANUP_ARCHIVE_BATCH_SIZE\" . ; grep -n -B1 -A8 \"CLEANUP_ARCHIVE_BATCH_SIZE\" miniflux.1 | head -30; grep -n \"AUTH_PROXY_USER_CREATION\\|\\\"API\\|\\\"BASE_URL\" internal/config/options.go | head; sed -n 1,30p internal/http/response/builder.go | head -5; grep -n \"func (b …`

> Bash `{"command":"sed -n 60,96p internal/config/options.go; sed -n 625,640p internal/config/options.go; grep -n \"^\\.B \" miniflux.1 | head -8; grep -n \"JSONBadRequest\" -A22 internal/http/response/json.go | sed -n 1,5p"}`

**Agent:** Adding the config option, a limiter, and wiring it into the API key middleware.

> Bash `{"command":"python3 - <<'EOF'\np='internal/config/options.go'\ns=open(p).read()\ns=s.replace('''\t\t\t\"AUTH_PROXY_HEADER\": {''','''\t\t\t\"API_RATE_LIMIT\": {\n\t\t\t\tparsedIntValue: 0,\n\t\t\t\trawValue:       \"0\",\n\t\t\t\tvalueType:      intType,\n\t\t\t\tvalidator: func(rawValue string) err…`

**Agent:** Bash is now denied, so I'll make the edits with the Edit tool instead.

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"\t\t\t\"AUTH_PROXY_HEADER\": {","new_string":"\t\t\t\"API_RA…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"func (c *configOptions) AuthProxyHeader() string {","new_str…`

> Read `{"file_path":"[redacted]","limit":17}`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"\t\"log/slog\"\n\t\"net/http\"\n\n\t\"miniflux.app/v2/in…`

> Write `{"file_path":"[redacted]","content":"// SPDX-FileCopyrightText: Copyright The Miniflux Authors. All rights reserved.\n…`

> Grep `{"pattern":"config\\.Opts\\b","path":"[redacted]","output_mode":"content","-n":true}`

> Grep `{"pattern":"Opts\\s*(=|\\*)|var Opts","path":"[redacted]","output_mode":"content","-n":true,"glob":"!*_test.go"}`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"type middleware struct {\n\tstore *storage.Storage\n}\n\nfun…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"\t\tslog.Info(\"[API] User authenticated successfully with t…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"\t\"net/http\"\n\n\t\"miniflux.app/v2/internal/config\"","ne…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"\tmiddleware := newMiddleware(store)","new_string":"\trateLimit := …`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"\t\"miniflux.app/v2/internal/storage\"","new_string":"\t\"miniflux.…`

> Grep `{"pattern":"^\\.B AUTH_PROXY_HEADER","path":"[redacted]","output_mode":"content","-n":true,"-B":2}`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":".TP\n.B AUTH_PROXY_HEADER","new_string":".TP\n.B API_RATE_LIMIT\nMaximum num…`

> Write `{"file_path":"[redacted]","content":"// SPDX-FileCopyrightText: Copyright The Miniflux Authors. All rights reserv…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"\tif retryAfter != 20*time.Second {","new_string":"\t…`

**Agent:** I've added a per-API-key rate limit to the API, but I haven't compiled it or run any tests. Bash was denied partway through, so `go build` and `go test ./internal/api/... ./internal/config/...` still need running.

- **Setting:** a new `API_RATE_LIMIT` option, in requests per minute for each API key. It defaults to `0`, which means disabled, so nothing changes for existing installs.
- **Behaviour:** each key can burst up to the limit, and tokens refill at limit per 60 seconds. A request over the limit gets `429 Too Many Requests` with a `Retry-After` header and a JSON error body.
- **Where it applies:** only to requests authenticated with the `X-Auth-Token` header. Basic-auth requests aren't limited, since you asked for a limit per API key.
- **Order of checks:** the limit applies after the key is validated, so a bad key still gets a `401`. Only real keys get a counter, so a flood of random tokens can't grow memory.
- **Cleanup:** counters for idle keys are dropped about once a minute.

The limiter is in `internal/api/rate_limiter.go`. It's wired into `internal/api/middleware.go` and `internal/api/api.go`, with the option in `internal/config/options.go`, a new `JSONTooManyRequests` response helper in `internal/http/response/json.go`, and a `miniflux.1` entry. I added unit tests in `internal/api/rate_limiter_test.go` for burst, per-key isolation, refill and cleanup. There's no test yet for the middleware returning 429, or for the config option parsing.

The counters live in memory, so they reset on restart and aren't shared between multiple instances.

