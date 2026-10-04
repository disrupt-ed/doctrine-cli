> Grep `{"pattern":"prometheus\\.|NewGaugeVec|NewCounter","output_mode":"files_with_matches","glob":"*.go"}`

> Read `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-failed-refresh-metric-baseline-182864093/internal/metric/metric.go"}`

> Grep `{"pattern":"metric\\.","output_mode":"content","glob":"*.go","path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-failed-refresh-metric-baseline-182864093/internal"}`

> Read `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-failed-refresh-metric-baseline-182864093/internal/worker/worker.go","offset":30,"limit":50}`

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-failed-refresh-metric-baseline-182864093/internal/metric/metric.go","old_string":"\tScraperRequestDuration = prometheus.NewHistogramVec(…`

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-failed-refresh-metric-baseline-182864093/internal/metric/metric.go","old_string":"\tprometheus.MustRegister(ScraperRequestDuration)","ne…`

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-failed-refresh-metric-baseline-182864093/internal/worker/worker.go","old_string":"\t\t\t\tstatus = metric.StatusError\n\t\t\t}","new_str…`

> Bash `{"command":"go build ./internal/... && git status --short"}`

**Agent:** I added a counter, `miniflux_background_feed_refresh_failures_total`. It goes up by one each time a background worker's feed refresh returns an error. The build passes; I didn't run any tests or look at the `/metrics` output.

- **Definition:** the counter is `BackgroundFeedRefreshFailures` in `internal/metric/metric.go`, and it's registered in `NewCollector`.
- **Where it's incremented:** `internal/worker/worker.go`, in the existing error branch. That branch only runs when metrics are enabled, so the counter follows the same setting.
- **Existing metric:** the `miniflux_background_feed_refresh_duration` histogram already has a `status="error"` label. Its `_count{status="error"}` series gives the same number, so the new counter is somewhat redundant. I added it because you asked for a dedicated metric, and it's simpler to query and alert on.
- **Scope:** it only counts refreshes triggered by the background workers, as the histogram does. Manual refreshes from the UI or API aren't counted.

