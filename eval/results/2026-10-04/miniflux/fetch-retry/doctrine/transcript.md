> Grep `{"pattern":"func \\(r \\*RequestBuilder\\) ExecuteRequest|ExecuteRequest\\(","path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-fetch-retry-doctrine-2244850815","output_mode":"content","glob":"*.go"}`

> Bash `{"command":"head -5 go.mod && ls internal/reader/fetcher && grep -rn \"ErrBadRequest\\|ErrNotFound\\|ErrUnexpected\\|ErrServer\\|ErrTooMany\" --include=*.go internal | head -20"}`

> Read `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-fetch-retry-doctrine-2244850815/internal/reader/fetcher/request_builder.go"}`

> Read `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-fetch-retry-doctrine-2244850815/internal/reader/handler/handler.go","offset":215,"limit":90}`

> Read `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-fetch-retry-doctrine-2244850815/internal/reader/fetcher/response_handler.go"}`

> Read `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-fetch-retry-doctrine-2244850815/internal/reader/handler/handler.go","offset":115,"limit":30}`

> Write `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-fetch-retry-doctrine-2244850815/internal/reader/fetcher/retry.go","content":"// SPDX-FileCopyrightText: Copyright The Miniflux Authors. All rights reserved.…`

> Bash `{"command":"sed -i 's/requestBuilder.ExecuteRequest(feedCreationRequest.FeedURL)/requestBuilder.ExecuteRequestWithRetry(feedCreationRequest.FeedURL)/; s/requestBuilder.ExecuteRequest(originalFeed.FeedURL)/requestBuilder.ExecuteRequestWithRetry(originalFeed.FeedURL)/' internal/reader/handler/handler.…`

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-fetch-retry-doctrine-2244850815/internal/reader/handler/handler.go","old_string":"requestBuilder.ExecuteRequest(feedCreationRequest.Feed…`

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-fetch-retry-doctrine-2244850815/internal/reader/handler/handler.go","old_string":"requestBuilder.ExecuteRequest(originalFeed.FeedURL)","…`

> Write `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-fetch-retry-doctrine-2244850815/internal/reader/fetcher/retry_test.go","content":"// SPDX-FileCopyrightText: Copyright The Miniflux Authors. All rights rese…`

**Agent:** Feed fetches now retry on temporary failures with exponential backoff, but I haven't compiled it or run any tests. The shell was denied in this session, so `gofmt`, `go vet` and `go test` are all still to do.

**What I changed**
- **New method:** `ExecuteRequestWithRetry` in `internal/reader/fetcher/retry.go`. It makes up to 3 attempts, waiting 1s and then 2s between them.
- **What counts as temporary:** timeouts, connection errors, a dropped connection (EOF), and 502, 503 and 504 responses.
- **What is not retried:** TLS errors, the private-network refusal, 404 and other 4xx responses, and 429. The 429 case is left alone because the refresh path already reschedules rate-limited feeds using `Retry-After`.
- **Where it's used:** the two feed fetches in `internal/reader/handler/handler.go`, feed creation and feed refresh. Icon, scraper and media requests still fail on the first error.
- **Proxy rotation:** the proxy is chosen inside `ExecuteRequest`, so each retry can use a different proxy from the rotator.
- **Tests:** `internal/reader/fetcher/retry_test.go` covers recovery after 503/502, giving up after 3 attempts, no retry on 404 or 429, and a refused connection. It sets the base delay to 1ms.

**Things to know**
- Backoff uses `time.Sleep`, so a worker is blocked for up to 3s per failing feed. That is fine at small scale. With many failing feeds it ties up the worker pool.
- Retry-After on a 503 is not honored. Say if you want that.
- A 500 is not retried, since it is often a real bug on the server. It is easy to add to the list in `isTemporaryFailure`.

Please run `gofmt -l internal/ && go vet ./internal/reader/... && go test ./internal/reader/fetcher/ ./internal/reader/handler/`, or allow me to run it.

