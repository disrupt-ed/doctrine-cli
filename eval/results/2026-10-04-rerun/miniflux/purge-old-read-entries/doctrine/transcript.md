**Agent:** Miniflux likely already has cleanup jobs, so I'll check those first.

> Grep `{"pattern":"ArchiveEntries|CleanOldSessions|cleanupJobs|CLEANUP_ARCHIVE","output_mode":"content","glob":"*.go","head_limit":60}`

> Read `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-purge-old-read-entries-doctrine-914331929/internal/cli/cleanup_tasks.go"}`

> Read `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-purge-old-read-entries-doctrine-914331929/internal/storage/entry.go","offset":340,"limit":80}`

> Read `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-purge-old-read-entries-doctrine-914331929/internal/config/options.go","offset":120,"limit":30}`

> Grep `{"pattern":"runCleanupTasks|CleanupFrequency","output_mode":"content","path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-purge-old-read-entries-doctrine-914331929/internal"}`

**Agent:** I didn't change any code, because Miniflux already does this.

- **Job:** `runCleanupTasks` in `internal/cli/cleanup_tasks.go:26` calls `store.ArchiveEntries(model.EntryStatusRead, …)`.
- **Schedule:** the scheduler runs it in `internal/cli/scheduler.go:55`, every `CLEANUP_FREQUENCY_HOURS` (default 24). It also runs from `cli.go:244`, which I didn't open, so I didn't check how that path is triggered.
- **Age:** `CLEANUP_ARCHIVE_READ_DAYS` defaults to 60 (`internal/config/options.go:134`).
- **What it deletes:** read entries older than that, in batches of `CLEANUP_ARCHIVE_BATCH_SIZE` (default 10000). It skips starred and shared entries and records tombstones so feeds don't re-ingest the deleted entries.

Two things to know:

- **Age is measured from `created_at`.** The cutoff counts from when the entry was created, not from when it was marked read. An entry that was only just read can be purged on the next run if it is already 60 days old.
- **Not tested.** I only read the code and didn't run any tests.

If that doesn't cover what you meant, tell me what's different. I could, for example, measure age from when the entry was marked read, delete starred entries too, or skip tombstones. I'd rather not build a second job beside the existing one.

