**Agent:** I'll start by reading how the API is set up.

> Bash `{"command":"head -5 go.mod; ls internal; ls internal/api | head -50; grep -rn -i \"rate\\|limit\" internal/config/options.go | head -30; grep -n \"Middleware\\|Use(\\|Handle\" internal/api/api.go | head -30"}`

> Read `{"file_path":"[redacted]"}`

> Read `{"file_path":"[redacted]","limit":25}`

> Bash `{"command":"grep -rn -i \"ratelimit\\|rate limit\\|429\\|TooManyRequests\" --include=*.go internal | head; sed -n 480,515p internal/config/options.go; grep -n \"PollingLimitPerHost\\|POLLING_LIMIT_PER_HOST\" -r . --include=* -l | head"}`

> Bash `{"command":"sed -n 100,140p internal/api/api.go; grep -rn \"POLLING_LIMIT_PER_HOST\\|PollingLimitPerHost\" . -l | head; grep -n \"func.*Duration\\|func.*Int()\" internal/config/options.go | head -20; grep -n \"TooManyRequests\\|func JSON\\|func NewBuilder\" -r internal/http/response | head -30; ls i…`

> Bash `{"command":"sed -n 60,100p internal/api/api.go; sed -n 95,120p internal/http/response/json.go; sed -n 1,45p internal/api/api_test.go; sed -n 90,110p internal/api/api_test.go; grep -n \"HTTP_SERVER_TIMEOUT\\|HTTPServerTimeout\\|HTTP_CLIENT_TIMEOUT\" -r internal miniflux.1 | head -20; grep -n \"contex…`

> Bash `{"command":"cat go.mod | sed -n 5,40p; sed -n 270,300p internal/config/options.go; sed -n 783,800p internal/config/options.go; grep -n \"intType\\|minuteType\\|secondType\" internal/config/*.go | grep -v options.go: | head; grep -n \"POLLING_LIMIT_PER_HOST\" -B2 -A6 miniflux.1; grep -n \"func.*Strin…`

> Bash `{"command":"grep -n '^\t\t\t\"[A-Z_0-9]*\": {' internal/config/options.go | sed -n 1,12p; grep -n \"^func (c \\*configOptions)\" internal/config/options.go | sed -n 1,8p; grep -n \"^\\.B \" miniflux.1 | sed -n 1,12p; grep -n \"ADMIN_USERNAME\" -A8 miniflux.1 | head -20"}`

**Agent:** Adding a per-key rate limit with a stdlib token bucket, a new config option, and a 429 response. I'm reading the remaining pieces I need to match.

> Bash `{"command":"sed -n 118,142p internal/http/response/json.go; sed -n 60,70p internal/config/options.go; sed -n 620,626p internal/config/options.go; sed -n 140,150p miniflux.1"}`

**Agent:** Now the edits: config option, 429 response, limiter, and middleware wiring.

> Bash `{"command":"python3 - <<'EOF'\np='internal/config/options.go'\ns=open(p).read()\ns=s.replace('''\t\t\t\"AUTH_PROXY_HEADER\": {''','''\t\t\t\"API_KEY_RATE_LIMIT\": {\n\t\t\t\tparsedIntValue: 0,\n\t\t\t\trawValue:       \"0\",\n\t\t\t\tvalueType:      intType,\n\t\t\t\tvalidator: func(rawValue string)…`

**Agent:** That shell command was denied, so I'll make the same changes with the Read, Edit and Write tools instead.

> Read `{"file_path":"[redacted]","offset":90,"limit":12}`

> Read `{"file_path":"[redacted]","limit":16}`

> Read `{"file_path":"[redacted]","offset":155,"limit":10}`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"\t\t\t\"AUTH_PROXY_HEADER\": {","new_string":"\t\t\t\"API_K…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"func (c *configOptions) AuthProxyHeader() string {","new_st…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":".B AUTH_PROXY_HEADER","new_string":".B API_KEY_RATE_LIMIT\nMaximum number o…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"\t\"net/http\"\n\n\t\"miniflux","new_string":"\t\"net/h…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"// JSONNotFound sends","new_string":"// JSONTooManyRequ…`

> Write `{"file_path":"[redacted]","content":"// SPDX-FileCopyrightText: Copyright The Miniflux Authors. All rights reserved.\n//…`

> Write `{"file_path":"[redacted]","content":"// SPDX-FileCopyrightText: Copyright The Miniflux Authors. All rights reserved…`

> Bash `{"command":"grep -n \"Opts\" internal/config/config.go | head -5"}`

**Agent:** `config.Opts` can be nil when the handler is built in tests, so I'll read the limit at request time and pass it to `allow`.

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"// Each bucket holds up to perMinute tokens and refills cont…`

> Bash `{"command":"sed -i 's/newRateLimiter(60) \\/\\/ 1 token per second, burst of 60/newRateLimiter()\\n\\tconst perMinute = 60 \\/\\/ 1 token per second, burst of 60/; s/l.allow(\"\\([ab]\\)\", now/l.allow(\"\\1\", perMinute, now/' internal/api/ratelimit_test.go; grep -n \"allow\\|newRate\" internal/api…`

**Agent:** Bash is now denied, so I'll finish with Edit and Write.

> Write `{"file_path":"[redacted]","content":"// SPDX-FileCopyrightText: Copyright The Miniflux Authors. All rights reserved…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"type middleware struct {\n\tstore *storage.Storage\n}\n\nfu…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"\t\t\tresponse.JSONUnauthorized(w, r)\n\t\t\treturn\n\t\t}\…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"\t\"net/http\"\n\n\t\"miniflux","new_string":"\t\"net/http\…`

**Agent:** API requests authenticated with an API key are now rate limited per key, but I haven't built or tested any of it. Bash was denied partway through, so I couldn't run `go build`, `go vet`, `gofmt` or `go test ./internal/api/... ./internal/config/...`. Please run those before merging.

- **Config option:** `API_KEY_RATE_LIMIT` is the maximum requests per minute for each key. It defaults to `0`, which means no limit, so existing clients see no change unless you set it. I added it to `internal/config/options.go` and documented it in `miniflux.1`.
- **Limiter:** `internal/api/ratelimit.go` is a token bucket held in memory, one bucket per key, using only the standard library (no new dependency). A key can burst up to the full per-minute limit, then refills steadily.
- **Where it applies:** the check runs in `validateAPIKeyAuth` in `internal/api/middleware.go`, after the key is confirmed valid. Because only valid keys get a bucket, an attacker sending random keys can't grow the map. Requests using Basic auth are not limited.
- **Rejected requests:** these get a JSON `429` with a `Retry-After` header, through a new `response.JSONTooManyRequests` in `internal/http/response/json.go`. The rejection is logged at warn level.
- **Test:** `internal/api/ratelimit_test.go` covers the burst, the rejection and `Retry-After`, refill over time, and separate buckets for different keys. I did not add a middleware-level test or a config-parsing test.

Counts reset when the server restarts, and each instance counts separately. That is fine for a single Miniflux process; running several behind a load balancer would need shared state. I didn't build that.

Decision for you: should Basic auth be limited too? I left it out because the request was specifically per API key, but it would be a small addition.

