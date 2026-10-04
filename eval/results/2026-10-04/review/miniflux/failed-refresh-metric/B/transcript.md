> Grep `{"pattern":"prometheus\\.|NewGaugeVec|NewCounter|NewHistogram","path":"[redacted]","glob":"*.go","output_mode":"content","head_limit":60}`

> Read `{"file_path":"[redacted]","limit":90}`

> Grep `{"pattern":"BackgroundFeedRefreshDuration","output_mode":"content","-C":12,"glob":"!internal/metric/*"}`

**Agent:** I haven't added anything, because this metric already exists.

`BackgroundFeedRefreshDuration` in `internal/metric/metric.go:24` is a histogram with a `status` label (`success` or `error`). The worker records it after every background refresh (`internal/worker/worker.go:49-54`). A failed refresh therefore increments the histogram's count series:

```
miniflux_background_feed_refresh_duration_count{status="error"}
```

In PromQL, the failure count is `increase(miniflux_background_feed_refresh_duration_count{status="error"}[1h])`, and the failure ratio is the `error` count divided by the sum over all statuses.

- **Existing metric:** it only records when the metrics collector is enabled (`config.Opts.HasMetricsCollector()`), like the other metrics. It covers background-worker refreshes only. Manual refreshes from the UI or API aren't counted.
- **Related gauge:** `miniflux_broken_feeds` is a different measure. It is the current number of feeds in an error state, not a running count of failures.

If you want a dedicated counter anyway, for example `miniflux_feed_refresh_failures_total`, say so. It would be about 10 lines: a `prometheus.NewCounter` in `metric.go`, a `MustRegister` call, and an `Inc()` in the worker's error branch. I'd only do it if you want a counter that doesn't depend on the histogram, or one that includes manual refreshes. For manual refreshes the counter would have to be incremented inside `RefreshFeed` or at its other call sites.

