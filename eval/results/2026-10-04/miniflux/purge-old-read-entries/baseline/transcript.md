> Grep `{"pattern":"ArchiveEntries|func.*CleanOld|RunScheduler|archive_read","output_mode":"content","glob":"*.go","head_limit":40}`

**Agent:** Read-entry archiving already exists. I'll check how it's wired up and configured before adding anything.

> Read `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-purge-old-read-entries-baseline-2360161944/internal/cli/cleanup_tasks.go"}`

> Grep `{"pattern":"CleanupArchiveReadDays|CLEANUP_ARCHIVE_READ_DAYS","output_mode":"content","path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-purge-old-read-entries-baseline-2360161944/internal/config"}`

> Grep `{"pattern":"runCleanupTasks|CleanupFrequency","output_mode":"content","path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-purge-old-read-entries-baseline-2360161944/internal","glob":"!*_test.go"}`

**Agent:** I haven't changed anything, because this job already exists in the codebase.

- **The job:** `runCleanupTasks` in `internal/cli/cleanup_tasks.go:26` calls `store.ArchiveEntries(model.EntryStatusRead, config.Opts.CleanupArchiveReadInterval(), ...)`. That deletes read entries older than the interval and records tombstones so feed refreshes don't re-ingest them.
- **The 60-day default:** `CLEANUP_ARCHIVE_READ_DAYS` defaults to 60 (`internal/config/options.go:134`). A test at `options_parsing_test.go:1208` checks that default.
- **The schedule:** the background scheduler (`internal/cli/scheduler.go:29,55`) runs `runCleanupTasks` every `CleanupFrequency()`. It also runs from the CLI at `cli.go:244`.

Because the existing job archives entries by deleting them and leaving tombstones, a second job would duplicate it.

To change the behavior, tell me which of these you want:
- A different age, which you can set today with `CLEANUP_ARCHIVE_READ_DAYS=N`.
- A hard delete with no tombstones. Feeds that still contain those entries would then re-import them on the next refresh, so I'd advise against it.
- A job that skips starred entries, if the current job doesn't already.

