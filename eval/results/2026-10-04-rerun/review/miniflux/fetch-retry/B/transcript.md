> Grep `{"pattern":"func \\(r \\*RequestBuilder\\)|func.*ErrorKind|type LocalizedErrorWrapper","path":"[redacted]","output_mode":"content","h…`

> Bash `{"command":"ls internal/reader internal/reader/fetcher && grep -rn \"NewRequestBuilder\\|\\.ExecuteRequest\" --include=*.go internal | grep -v _test | head -30"}`

> Read `{"file_path":"[redacted]"}`

> Bash `{"command":"grep -rn \"ExecuteRequest\" --include='*.go' . | grep -v _test | head -30"}`

> Read `{"file_path":"[redacted]"}`

> Read `{"file_path":"[redacted]","offset":100,"limit":190}`

> Write `{"file_path":"[redacted]","content":"// SPDX-FileCopyrightText: Copyright The Miniflux Authors. All rights reserved.…`

> Write `{"file_path":"[redacted]","content":"// SPDX-FileCopyrightText: Copyright The Miniflux Authors. All rights rese…`

> Bash `{"command":"sed -i 's/NewResponseHandler(requestBuilder.ExecuteRequest(feedCreationRequest.FeedURL))/NewResponseHandler(requestBuilder.ExecuteRequestWithRetry(feedCreationRequest.FeedURL))/; s/NewResponseHandler(requestBuilder.ExecuteRequest(originalFeed.FeedURL))/NewResponseHandler(requestBuilder.E…`

**Agent:** Bash was denied, so I'll make the two call-site edits with Edit instead.

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"requestBuilder.ExecuteRequest(feedCreationRequest.Feed…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"requestBuilder.ExecuteRequest(originalFeed.FeedURL)","…`

**Agent:** Feed fetches now retry on temporary errors with exponential backoff. I haven't compiled it or run any tests, because Bash was denied in this session. Please run `go build ./... && go test ./internal/reader/fetcher/ ./internal/reader/handler/`.

**What changed**
- **Retry logic:** a new `ExecuteRequestWithRetry` on `RequestBuilder` in `internal/reader/fetcher/retry.go` wraps the existing `ExecuteRequest`.
- **Call sites:** `CreateFeed` and `RefreshFeed` in `internal/reader/handler/handler.go` now call it. Other fetches (icons, scraper, media proxy, and so on) are unchanged.
- **Backoff:** up to 3 attempts, with a 1s delay, then 2s, capped at 10s. These are constants in `retry.go`. After the last attempt, the final response or error goes back to the caller, so the existing error handling is unchanged.
- **What counts as temporary:**
  - timeouts, EOF and connection-level errors (`net.OpError`)
  - DNS errors, except "not found"
  - HTTP 502, 503 and 504
- **What doesn't retry:**
  - TLS errors
  - the private-network refusal (`ErrPrivateNetworkHost`)
  - 429, because `RefreshFeed` already schedules around `Retry-After`
  - 500 and 4xx responses

Failed responses have their bodies closed before the next attempt. Each attempt builds a fresh client, so the proxy rotator will pick the next proxy on a retry.

I added `retry_test.go` covering backoff timing, giving up, no retry on permanent statuses, and connection errors.

The worst case for one fetch is now about 3× the timeout plus 3s of backoff. For `CreateFeed` that is time the user waits on the page. Say so if you'd rather skip retries there or use fewer attempts.

