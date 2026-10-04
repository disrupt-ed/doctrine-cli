**Agent:** I'll start by reading how the API is wired up and what already exists for rate limiting.

> Grep `{"pattern":"(?i)rate.?limit|x/time|ratelimit","glob":"*.{go,mod}","output_mode":"files_with_matches"}`

> Bash `{"command":"head -5 go.mod; ls internal/api internal/http/server; grep -n \"x/\" go.mod"}`

> Read `{"file_path":"[redacted]"}`

> Read `{"file_path":"[redacted]"}`

> Bash `{"command":"grep -rn \"BATCH_SIZE\\|BatchSize\\|batchSize\" internal/config | head -30; grep -n \"^func \" internal/http/response/json.go internal/http/response/response.go | head -40; grep -rn \"TooMany\" internal --include=*.go | head"}`

> Bash `{"command":"sed -n 95,165p internal/http/response/json.go; sed -n 95,135p internal/config/options.go; sed -n 640,660p internal/config/options.go; grep -n \"type option\\b\" -A12 internal/config/*.go | head -30; grep -rn \"BATCH_SIZE\" --include=*.1 --include=*.md . | head; ls internal/config"}`

> Bash `{"command":"grep -n \"validateGreaterOrEqualThan\\|func validate\" internal/config/validators.go; grep -n '\"POLLING_LIMIT_PER_HOST\"' -B2 -A10 internal/config/options.go | head -40; grep -n \"POLLING_LIMIT_PER_HOST\" -B2 -A8 miniflux.1 | head -30; grep -n \"PollingLimitPerHost\" -B2 -A3 internal/co…`

> Bash `{"command":"grep -n \"config.Opts\" internal/api/*_test.go | head; sed -n 60,100p internal/api/api_test.go; sed -n 28,52p internal/config/validators.go; grep -n \"API_KEY\\|DISABLE_LOCAL_AUTH\" internal/config/options.go miniflux.1"}`

> Bash `{"command":"sed -n 78,95p internal/config/options.go; grep -n \"func (c \\*configOptions) A\" internal/config/options.go | head -3; grep -n \"^\\.B A\" miniflux.1 | head -3"}`

> Bash `{"command":"sed -n 150,162p miniflux.1; sed -n 626,634p internal/config/options.go"}`

**Agent:** Now the edits: config option, accessor, man page, limiter, middleware.

> Bash `{"command":"python3 - <<'EOF'\np='internal/config/options.go'\ns=open(p).read()\ns=s.replace('''\t\t\t\"AUTH_PROXY_HEADER\": {''','''\t\t\t\"API_KEY_RATE_LIMIT\": {\n\t\t\t\tparsedIntValue: 0,\n\t\t\t\trawValue:       \"0\",\n\t\t\t\tvalueType:      intType,\n\t\t\t\tvalidator: func(rawValue string)…`

**Agent:** The shell was denied, so I'll make the same changes with the Edit and Write tools. Those need each file read first.

> Read `{"file_path":"[redacted]","offset":84,"limit":12}`

> Read `{"file_path":"[redacted]","offset":150,"limit":14}`

> Read `{"file_path":"[redacted]","limit":20}`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"\t\t\t\"AUTH_PROXY_HEADER\": {","new_string":"\t\t\t\"API_K…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"func (c *configOptions) AuthProxyHeader() string {","new_st…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":".TP\n.B AUTH_PROXY_HEADER","new_string":".TP\n.B API_KEY_RATE_LIMIT\nMaximu…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"func generateJSONError(","new_string":"// JSONTooManyRe…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"\t\"log/slog\"\n\t\"net/http\"\n","new_string":"\t\"log…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"\tstore *storage.Storage\n}\n\nfunc newMiddleware(s *storag…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"\t\tm.store.SetLastLogin(user.ID)\n\t\tm.store.SetAPIKeyUse…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"\t\"net/http\"\n","new_string":"\t\"net/http\"\n\t\"time\"\…`

> Write `{"file_path":"[redacted]","content":"// SPDX-FileCopyrightText: Copyright The Miniflux Authors. All rights reserved.\n//…`

> Write `{"file_path":"[redacted]","content":"// SPDX-FileCopyrightText: Copyright The Miniflux Authors. All rights reserved…`

> Bash `{"command":"gofmt -l internal/ && go vet ./internal/api/ ./internal/config/ ./internal/http/response/ && go test ./internal/api/ ./internal/config/ ./internal/http/response/ 2>&1 | tail -15"}`

**Agent:** The API now rate limits each API key. It's off by default, and you turn it on with `API_KEY_RATE_LIMIT=<requests per minute>`. `gofmt`, `go vet` and the tests for `api`, `config` and `response` pass. I tested the limiter on its own but didn't send a real request through the middleware.

- **Limiter:** `internal/api/ratelimit.go` is a small in-memory token bucket per key. Each key can burst up to the limit and then refills steadily. I wrote it myself instead of adding `golang.org/x/time/rate`, since it's about 40 lines and avoids a new dependency.
- **Where it runs:** in `validateAPIKeyAuth` (`internal/api/middleware.go`), after the key is checked against the database. Unknown tokens never get a bucket, so memory is bounded by the number of real keys. A throttled request doesn't update `last_login` or the key's "last used" time.
- **Over the limit:** the client gets a `429` with a JSON error and a `Retry-After` header. This uses a new `response.JSONTooManyRequests` that follows the existing error helpers.
- **Config and docs:** the new option is `API_KEY_RATE_LIMIT` (integer, 0 or more). It's documented in `miniflux.1`, and 0 means disabled, like `POLLING_LIMIT_PER_HOST`.
- **Tests:** `internal/api/ratelimit_test.go` checks blocking after the limit, the retry delay, refill over time, and that keys don't affect each other.

Limits to know about:
- **Basic auth:** requests using username and password aren't limited, only requests with an API key. The request was per API key, so I stopped there.
- **Multiple instances:** counters live in memory and reset on restart. Each instance counts separately, so several Miniflux instances behind a load balancer each allow the full limit.
- **Default:** I left it off so existing API clients don't start getting 429s after an upgrade.

The first shell command was denied, so I made the edits with the file tools instead. After that, plain `gofmt`, `go vet` and `go test` ran fine.

