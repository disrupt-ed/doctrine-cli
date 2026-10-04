**Agent:** I'll start by reading how the existing background jobs are set up.

> Grep `{"pattern":"ArchiveEntries|CleanOldSessions|func.*Cleanup","output_mode":"content","glob":"*.go"}`

> Read `{"file_path":"[redacted]"}`

> Read `{"file_path":"[redacted]","offset":355,"limit":60}`

> Grep `{"pattern":"CLEANUP_ARCHIVE_READ_DAYS|CleanupArchiveReadInterval","output_mode":"content","-n":true}`

**Agent:** I didn't change anything, because Miniflux already does this.

- **The job:** `runCleanupTasks` in `internal/cli/cleanup_tasks.go:26` calls `store.ArchiveEntries(model.EntryStatusRead, config.Opts.CleanupArchiveReadInterval(), ...)`.
- **The 60 days:** `CLEANUP_ARCHIVE_READ_DAYS` defaults to 60. `TestCleanupArchiveReadIntervalOptionParsing` checks that default.
- **How it runs:** `ArchiveEntries` (`internal/storage/entry.go:368`) deletes read entries older than the interval in batches. It skips starred entries and entries with a share code. It also records tombstones so deleted entries aren't re-ingested from the feed.
- **Schedule:** `runCleanupTasks` is the existing cleanup task. The job's run frequency comes from `CLEANUP_FREQUENCY_HOURS`. I didn't confirm that setting's default or how the task is scheduled.

If the job isn't running in your setup, check `CLEANUP_ARCHIVE_READ_DAYS` and `CLEANUP_FREQUENCY_HOURS`, and the batch size (`CLEANUP_ARCHIVE_BATCH_SIZE`). A negative interval or a batch size of 0 makes it do nothing.

If you meant something different, tell me which. Possible differences are that starred entries should also be deleted, tombstones shouldn't be recorded, or the job should run on its own schedule. I can implement any of those.

