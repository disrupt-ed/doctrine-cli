> Grep `{"pattern":"ArchiveEntries|CleanOldSessions","output_mode":"content","glob":"*.go"}`

> Read `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-purge-old-read-entries-baseline-1357839123/internal/cli/cleanup_tasks.go"}`

> Read `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-purge-old-read-entries-baseline-1357839123/internal/storage/entry.go","offset":350,"limit":70}`

> Grep `{"pattern":"CleanupArchiveRead|CLEANUP_ARCHIVE_READ","output_mode":"content"}`

**Agent:** You don't need to add anything. Miniflux already has this job, and its default is 60 days.

- **The job:** `runCleanupTasks` in `internal/cli/cleanup_tasks.go:26` calls `store.ArchiveEntries(model.EntryStatusRead, config.Opts.CleanupArchiveReadInterval(), ...)`. It runs in the background scheduler.
- **What it deletes:** `internal/storage/entry.go:368` deletes read entries older than the interval. It skips starred entries and shared entries. It also records tombstones so deleted entries aren't re-fetched.
- **The 60-day default:** the `CLEANUP_ARCHIVE_READ_DAYS` option controls the interval, and `internal/config/options_parsing_test.go:1204` asserts that its default is 60 days.

To change the age, set `CLEANUP_ARCHIVE_READ_DAYS`. A negative value disables the job.

I haven't checked the other cleanup settings, such as the batch size and how often the scheduler runs the cleanup. Do you want a different behavior, such as also deleting starred entries, skipping the tombstones, or running on its own schedule? If so, tell me which and I'll make the change.

