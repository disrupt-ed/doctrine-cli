> Grep `{"pattern":"func \\(r \\*ResponseHandler\\)|func NewRequestBuilder|func \\(r \\*RequestBuilder\\) Request|ErrBadGateway|IsModified","path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-fetch-retry-baseline-329099680…`

> Read `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-fetch-retry-baseline-3290996806/internal/reader/fetcher/request_builder.go"}`

> Read `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-fetch-retry-baseline-3290996806/internal/reader/fetcher/response_handler.go"}`

**Agent:** Now let me see how callers use these.

> Grep `{"pattern":"ExecuteRequest\\(","path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-fetch-retry-baseline-3290996806/internal","output_mode":"content"}`

> Read `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-fetch-retry-baseline-3290996806/internal/reader/handler/handler.go","offset":100,"limit":200}`

> Write `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-fetch-retry-baseline-3290996806/internal/reader/fetcher/retry.go","content":"// SPDX-FileCopyrightText: Copyright The Miniflux Authors. All rights reserved.…`

> Write `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-fetch-retry-baseline-3290996806/internal/reader/fetcher/retry_test.go","content":"// SPDX-FileCopyrightText: Copyright The Miniflux Authors. All rights rese…`

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-fetch-retry-baseline-3290996806/internal/reader/handler/handler.go","old_string":"requestBuilder.ExecuteRequest(originalFeed.FeedURL))",…`

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-fetch-retry-baseline-3290996806/internal/reader/handler/handler.go","old_string":"requestBuilder.ExecuteRequest(feedCreationRequest.Feed…`

> Bash `{"command":"go build ./... && go vet ./internal/reader/... && go test ./internal/reader/fetcher/ ./internal/reader/handler/ 2>&1 | tail -20"}`

**Agent:** Feed creation and feed refresh now retry temporary failures with exponential backoff. The build and `go vet` pass, and the fetcher tests pass, including four new ones.

**What counts as temporary**
- Responses with status 502, 503 or 504.
- Request timeouts, connection reset, refused or aborted, and unexpected EOF.

These are not retried:
- TLS errors, hostname resolution failures and private-network refusals.
- 4xx responses. 429 is skipped because refresh already reschedules those using `Retry-After`.
- 500 responses, since they are more often a persistent server bug than a blip.

**Behaviour**
- Each fetch makes up to 3 attempts, waiting 1s after the first failure and 2s after the second.
- If every attempt fails, the last response or error goes back to the caller unchanged, so the existing error handling and messages still apply.
- The delays are fixed constants in `internal/reader/fetcher/retry.go`, not configuration options. Say if you want them configurable.
- A feed that keeps failing can now hold a refresh worker for about 3 seconds of waiting, plus up to 3 request timeouts, per refresh.

**Changes**
- `internal/reader/fetcher/retry.go` adds `ExecuteRequestWithRetry` and the temporary-failure check.
- `internal/reader/handler/handler.go` uses it in `CreateFeed` and `RefreshFeed`.
- `internal/reader/fetcher/retry_test.go` covers four cases: success after retries, giving up after the maximum attempts, no retry on permanent statuses, and retry on timeout.

The scraper, icon finder, subscription discovery and other fetches still use the plain request, since you asked about feed fetches. I haven't committed anything.

