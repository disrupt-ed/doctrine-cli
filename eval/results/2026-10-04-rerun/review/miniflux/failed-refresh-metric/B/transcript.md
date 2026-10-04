> Grep `{"pattern":"prometheus\\.|NewGaugeVec|NewCounter|NewHistogram","path":"[redacted]","output_mode":"files_with_matches","glob":"*.go"}`

> Read `{"file_path":"[redacted]"}`

> Grep `{"pattern":"metric\\.","path":"[redacted]","output_mode":"content","glob":"*.go"}`

> Read `{"file_path":"[redacted]","offset":35,"limit":40}`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"\tScraperRequestDuration = prometheus.NewHistogramVec…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"\tprometheus.MustRegister(ScraperRequestDuration)","n…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"\t\t\t\tstatus = metric.StatusError\n\t\t\t}","new_st…`

> Bash `{"command":"go build ./internal/... && echo ok"}`

**Agent:** I added a counter, `miniflux_background_feed_refresh_failures_total`, and `go build ./internal/...` passes. I haven't run any tests or checked the metric on a running instance.

- **Definition:** it's declared and registered in `internal/metric/metric.go`.
- **Increment:** the background worker bumps it in `internal/worker/worker.go` whenever `RefreshFeed` returns an error. That code only runs when the metrics collector is enabled, same as the existing refresh-duration histogram.

The existing `miniflux_background_feed_refresh_duration` histogram already records failures as `..._count{status="error"}`. The new counter gives you a simpler series to alert on, but it's redundant with that one. If you'd rather not carry both, I can drop the counter.

