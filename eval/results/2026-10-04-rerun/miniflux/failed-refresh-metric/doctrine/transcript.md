> Grep `{"pattern":"prometheus\\.|Counter|ErrorCount|error_count","path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-failed-refresh-metric-doctrine-1628013424/internal/metric","output_mode":"content"}`

> Read `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-failed-refresh-metric-doctrine-1628013424/internal/metric/metric.go","limit":70}`

> Grep `{"pattern":"BackgroundFeedRefreshDuration|ScraperRequestDuration","path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-failed-refresh-metric-doctrine-1628013424","output_mode":"content","glob":"*.go"}`

> Read `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-failed-refresh-metric-doctrine-1628013424/internal/worker/worker.go","offset":30,"limit":60}`

**Agent:** I didn't add anything, because the metric you want already exists.

`miniflux_background_feed_refresh_duration` is a histogram with a `status` label of `success` or `error`. `internal/worker/worker.go:49-55` sets `error` whenever `RefreshFeed` returns an error. Prometheus exposes a `_count` series for every histogram, so the number of failed refreshes is:

```
miniflux_background_feed_refresh_duration_count{status="error"}
```

For a rate over time, use `rate(miniflux_background_feed_refresh_duration_count{status="error"}[5m])`.

A separate counter would only repeat that series and add a second metric to maintain. It would also need the same `HasMetricsCollector()` check in the worker.

This covers background worker refreshes only. A manual refresh from the UI or API doesn't go through the worker, so it isn't counted. If you want those included, say so and I'll look at where to record them.

If you still want a dedicated `..._failed_total` counter, I can add it. It's about 10 lines in `internal/metric/metric.go` plus one increment in `worker.go`.

