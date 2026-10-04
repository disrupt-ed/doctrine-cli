> Grep `{"pattern":"func \\(r \\*RequestBuilder\\)|func .*Fetch\\(|ErrBadGateway|ErrServerError|ErrTooManyRequests","path":"[redacted]","output_mode…`

> Read `{"file_path":"[redacted]","offset":1,"limit":300}`

> Grep `{"pattern":"ExecuteRequest\\(","path":"[redacted]","output_mode":"content"}`

> Read `{"file_path":"[redacted]","offset":296}`

> Read `{"file_path":"[redacted]","offset":100,"limit":200}`

> Read `{"file_path":"[redacted]"}`

> Write `{"file_path":"[redacted]","content":"// SPDX-FileCopyrightText: Copyright The Miniflux Authors. All rights reserved.…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"requestBuilder.ExecuteRequest(feedCreationRequest.Feed…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"requestBuilder.ExecuteRequest(originalFeed.FeedURL))",…`

> Read `{"file_path":"[redacted]","limit":40}`

> Write `{"file_path":"[redacted]","content":"// SPDX-FileCopyrightText: Copyright The Miniflux Authors. All rights rese…`

> Bash `{"command":"gofmt -l internal/reader && go vet ./internal/reader/... && go test ./internal/reader/fetcher/ ./internal/reader/handler/ 2>&1 | tail -20"}`

**Agent:** Feed fetches now retry temporary failures with exponential backoff. `gofmt` and `go vet` are clean and the fetcher tests pass. The `handler` package has no tests, so the two call sites I changed aren't covered by them.

**What I added**
- **New method:** `RequestBuilder.ExecuteRequestWithRetry` in `internal/reader/fetcher/retry.go`. It's a plain loop with `time.Sleep`, no new dependency, and wraps the existing `ExecuteRequest`.
- **Attempts:** up to 3 in total, waiting 1s and then 2s between them.
- **What gets retried:**
  - 502, 503 and 504 responses.
  - Timeouts, EOF and dial or connection errors.
  - Temporary DNS failures.
- **What doesn't get retried:**
  - TLS errors and the private-network block (`ErrPrivateNetworkHost`).
  - DNS "not found" and invalid URLs.
  - 4xx responses, including 429. The existing rate-limit handling in `RefreshFeed` still schedules the next check from `Retry-After`.
  - 500 responses, since those are often deterministic.
- **Call sites:** feed creation and feed refresh in `handler.go`. Icon, scraper, media proxy and other requests are unchanged.
- **Tests:** `retry_test.go` covers retried statuses, giving up after 3 attempts, no retry for 404 and 429, a dropped connection that succeeds on retry, and no retry for an invalid URL.

**Things to know**
- A refresh against a failing feed now holds a worker for up to about 3s of sleeping, plus up to 3 request timeouts (`HTTPClientTimeout`, 20s by default). If that matters for your worker pool, lowering the attempt count is a one-line change (`maxFetchAttempts`).
- The sleep isn't cancellable because `ExecuteRequest` doesn't take a context. Adding one would mean changing its signature, so I left it.
- With a proxy rotator, each retry takes the next proxy, which helps when one proxy is the problem.
- I haven't added a user-facing config option for the retry count or delay, since nobody asked for one.

