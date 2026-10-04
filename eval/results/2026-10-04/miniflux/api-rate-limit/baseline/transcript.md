**Agent:** I'll start by looking at how the API is wired up.

> Grep `{"pattern":"APIKey|api_key|RateLimit|rate_limit","glob":"internal/**/*.go","output_mode":"files_with_matches"}`

> Glob `{"pattern":"internal/api/*.go"}`

> Read `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-api-rate-limit-baseline-1249941355/internal/api/middleware.go"}`

> Read `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-api-rate-limit-baseline-1249941355/internal/api/api.go","limit":70}`

> Grep `{"pattern":"func \\(.*\\) (Is|Get)[A-Za-z]*\\(\\)|POLLING_LIMIT|PollingLimit|APIKey","path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-api-rate-limit-baseline-1249941355/internal/config/options.go","output_mode":…`

> Read `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-api-rate-limit-baseline-1249941355/internal/api/api.go","offset":70}`

> Read `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-api-rate-limit-baseline-1249941355/internal/config/options.go","offset":485,"limit":30}`

> Grep `{"pattern":"func ClientIP|func JSONTooMany|func JSONUnauthorized|func JSONErrorResponse|func JSONBadRequest","path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-api-rate-limit-baseline-1249941355/internal/http","ou…`

> Read `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-api-rate-limit-baseline-1249941355/internal/http/response/json.go","offset":20,"limit":110}`

> Grep `{"pattern":"POLLING_LIMIT_PER_HOST|PollingLimitPerHost","output_mode":"content","glob":"!internal/config/options.go"}`

> Bash `{"command":"grep -n \"GetUserIDContextKey\\|func UserID\\|func IsAuthenticated\" -r internal/http/request/context.go; grep -n \"go 1\\|golang.org/x/time\" go.mod; grep -n \"POLLING_LIMIT_PER_HOST\" -B2 -A8 miniflux.1"}`

> Bash `{"command":"sed -n 119,175p internal/http/response/json.go; sed -n 925,940p internal/config/options.go; grep -n '\"API\\|\"AUTH_PROXY\\|\"BATCH_SIZE\\|\"BASE_URL' internal/config/options.go; grep -n \"^\\.B BATCH_SIZE\\|^\\.B BASE_URL\\|^\\.B CLEANUP_ARCHIVE_BATCH_SIZE\" -B3 miniflux.1; grep -n \"fu…`

> Bash `{"command":"sed -n 80,100p internal/config/options.go; sed -n 640,655p internal/config/options.go; sed -n 160,176p miniflux.1; sed -n 1,16p internal/api/api_test.go"}`

**Agent:** Exploration is done. I'm now adding the config option, a token-bucket limiter, a 429 response helper and the middleware wiring.

> Bash `{"command":"python3 - <<'EOF'\np='internal/config/options.go'\ns=open(p).read()\ns=s.replace('''\t\t\t\"AUTH_PROXY_HEADER\": {''','''\t\t\t\"API_RATE_LIMIT\": {\n\t\t\t\tparsedIntValue: 0,\n\t\t\t\trawValue:       \"0\",\n\t\t\t\tvalueType:      intType,\n\t\t\t\tvalidator: func(rawValue string) err…`

**Agent:** Bash was denied, so I'll make the same changes with the Edit tool.

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-api-rate-limit-baseline-1249941355/internal/config/options.go","old_string":"\t\t\t\"AUTH_PROXY_HEADER\": {","new_string":"\t\t\t\"API_R…`

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-api-rate-limit-baseline-1249941355/internal/config/options.go","old_string":"func (c *configOptions) BaseURL() string {","new_string":"f…`

> Read `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-api-rate-limit-baseline-1249941355/internal/http/response/json.go","limit":20}`

> Read `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-api-rate-limit-baseline-1249941355/miniflux.1","offset":160,"limit":10}`

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-api-rate-limit-baseline-1249941355/miniflux.1","old_string":".TP\n.B AUTH_PROXY_HEADER\n","new_string":".TP\n.B API_RATE_LIMIT\nMaximum …`

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-api-rate-limit-baseline-1249941355/internal/http/response/json.go","old_string":"\t\"log/slog\"\n\t\"net/http\"\n","new_string":"\t\"log…`

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-api-rate-limit-baseline-1249941355/internal/http/response/json.go","old_string":"// JSONNotFound sends","new_string":"// JSONTooManyRequ…`

**Agent:** Now the limiter itself.

> Write `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-api-rate-limit-baseline-1249941355/internal/api/rate_limiter.go","content":"// SPDX-FileCopyrightText: Copyright The Miniflux Authors. All rights reserved.\…`

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-api-rate-limit-baseline-1249941355/internal/api/rate_limiter.go","old_string":"// cleanup drops buckets that have been idle long enough …`

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-api-rate-limit-baseline-1249941355/internal/api/rate_limiter.go","old_string":"\tl.cleanup(now, limitPerMinute)","new_string":"\tl.clean…`

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-api-rate-limit-baseline-1249941355/internal/api/middleware.go","old_string":"type middleware struct {\n\tstore *storage.Storage\n}\n\nfu…`

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-api-rate-limit-baseline-1249941355/internal/api/middleware.go","old_string":"\t\tslog.Info(\"[API] User authenticated successfully with …`

**Agent:** Now tests for the limiter and the config option.

> Write `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-api-rate-limit-baseline-1249941355/internal/api/rate_limiter_test.go","content":"// SPDX-FileCopyrightText: Copyright The Miniflux Authors. All rights reser…`

> Read `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-api-rate-limit-baseline-1249941355/internal/config/options_parsing_test.go","offset":1040,"limit":25}`

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-api-rate-limit-baseline-1249941355/internal/config/options_parsing_test.go","old_string":"func TestPollingLimitPerHostOptionParsing(t *t…`

> Bash `{"command":"go build ./... && go vet ./internal/api ./internal/config ./internal/http/... && go test ./internal/api ./internal/config ./internal/http/... 2>&1 | tail -20"}`

**Agent:** The API is now rate limited per API key, and the build, vet and tests pass. I haven't run a request through the real middleware against a database.

- **Config:** a new `API_RATE_LIMIT` option sets the maximum requests per minute for each key. It defaults to `0`, which means disabled, so existing installs behave as before. Negative values are rejected, and the option is documented in `miniflux.1`.
- **Limiter:** `internal/api/rate_limiter.go` is a token bucket per API key. A key can burst up to the limit, and the bucket refills at limit/60 tokens per second. Idle buckets are removed so memory stays bounded. I wrote it by hand because the repo doesn't depend on `golang.org/x/time`.
- **Middleware:** the check is in `validateAPIKeyAuth`, after the key is validated. Unknown tokens can't create buckets, and rejected requests skip the `SetLastLogin` and `SetAPIKeyUsedTimestamp` writes. Over the limit, the client gets a `429` with `{"error_message":"too many requests"}` and a `Retry-After` header.
- **Tests:** unit tests cover the limiter and the config parsing.

Basic-auth requests, which don't use an API key, aren't limited. If you want a limit there too, it would need a different key, such as the username.

The new files are `internal/api/rate_limiter.go` and `internal/api/rate_limiter_test.go`. I also edited `internal/api/middleware.go`, `internal/config/options.go`, `internal/http/response/json.go` and `miniflux.1`. Nothing is committed.

