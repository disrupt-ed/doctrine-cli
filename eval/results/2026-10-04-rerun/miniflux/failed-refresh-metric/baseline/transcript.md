> Grep `{"pattern":"prometheus\\.|NewGaugeVec|NewCounter|NewHistogram","path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-failed-refresh-metric-baseline-1748362290","output_mode":"files_with_matches","glob":"*.go"}`

> Read `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-failed-refresh-metric-baseline-1748362290/internal/metric/metric.go"}`

> Grep `{"pattern":"metric\\.","path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-failed-refresh-metric-baseline-1748362290","output_mode":"content","glob":"*.go"}`

> Read `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-failed-refresh-metric-baseline-1748362290/internal/worker/worker.go","offset":35,"limit":40}`

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-failed-refresh-metric-baseline-1748362290/internal/metric/metric.go","old_string":"\tScraperRequestDuration = prometheus.NewHistogramVec…`

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-failed-refresh-metric-baseline-1748362290/internal/metric/metric.go","old_string":"\tprometheus.MustRegister(ScraperRequestDuration)","n…`

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-failed-refresh-metric-baseline-1748362290/internal/worker/worker.go","old_string":"\t\t\t\tstatus = metric.StatusError\n\t\t\t}","new_st…`

> Bash `{"command":"go build ./internal/... && echo ok"}`

**Agent:** I added a counter, `miniflux_background_feed_refresh_failures_total`, and `go build ./internal/...` passes. I haven't run any tests or checked the metric on a running instance.

- **Definition:** it's declared and registered in `internal/metric/metric.go`.
- **Increment:** the background worker bumps it in `internal/worker/worker.go` whenever `RefreshFeed` returns an error. That code only runs when the metrics collector is enabled, same as the existing refresh-duration histogram.

The existing `miniflux_background_feed_refresh_duration` histogram already records failures as `..._count{status="error"}`. The new counter gives you a simpler series to alert on, but it's redundant with that one. If you'd rather not carry both, I can drop the counter.

