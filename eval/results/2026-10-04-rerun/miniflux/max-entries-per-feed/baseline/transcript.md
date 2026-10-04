> Grep `{"pattern":"CleanupArchiveReadDays|CLEANUP_ARCHIVE_READ_DAYS|ArchiveReadEntries","output_mode":"files_with_matches"}`

> Grep `{"pattern":"CleanupArchiveUnreadDays|CLEANUP_ARCHIVE_UNREAD_DAYS|ArchiveUnreadEntries|func.*Cleanup","output_mode":"content","glob":"**/*.go"}`

> Read `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-max-entries-per-feed-baseline-2340984474/internal/cli/cleanup_tasks.go"}`

> Read `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-max-entries-per-feed-baseline-2340984474/internal/config/options.go","offset":100,"limit":80}`

> Grep `{"pattern":"func \\(s \\*Storage\\) (ArchiveEntries|DeleteEntries|RefreshFeedEntries)","path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-max-entries-per-feed-baseline-2340984474/internal/storage","output_mode":"c…`

> Read `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-max-entries-per-feed-baseline-2340984474/internal/storage/entry.go","offset":300,"limit":130}`

> Grep `{"pattern":"CLEANUP_ARCHIVE_UNREAD_DAYS","path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-max-entries-per-feed-baseline-2340984474/miniflux.1","output_mode":"content","-A":6}`

> Read `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-max-entries-per-feed-baseline-2340984474/miniflux.1","offset":195,"limit":40}`

> Grep `{"pattern":"validateGreaterOrEqualThan\\(rawValue, 0\\)","path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-max-entries-per-feed-baseline-2340984474/internal/config/options.go","output_mode":"count"}`

> Read `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-max-entries-per-feed-baseline-2340984474/internal/config/options.go","offset":658,"limit":12}`

**Agent:** Adding the config option, a storage method, and wiring it into the cleanup job.

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-max-entries-per-feed-baseline-2340984474/internal/config/options.go","old_string":"\t\t\t\"CLEANUP_REMOVE_SESSIONS_DAYS\": {","new_strin…`

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-max-entries-per-feed-baseline-2340984474/internal/config/options.go","old_string":"func (c *configOptions) CleanupRemoveSessionsInterval…`

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-max-entries-per-feed-baseline-2340984474/miniflux.1","old_string":".B CLEANUP_REMOVE_SESSIONS_DAYS","new_string":".B CLEANUP_MAX_ENTRIES…`

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-max-entries-per-feed-baseline-2340984474/internal/storage/entry.go","old_string":"// SetEntriesStatus update the status","new_string":"/…`

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-max-entries-per-feed-baseline-2340984474/internal/cli/cleanup_tasks.go","old_string":"\tif nbIcons, err :=","new_string":"\tif maxEntrie…`

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-max-entries-per-feed-baseline-2340984474/internal/storage/entry.go","old_string":"\t\t\t\tranked.share_code='' AND\n\t\t\t\tentries.id I…`

> Read `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-max-entries-per-feed-baseline-2340984474/internal/config/options_parsing_test.go","offset":1218,"limit":18}`

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-max-entries-per-feed-baseline-2340984474/internal/config/options_parsing_test.go","old_string":"func TestCleanupFrequencyOptionParsing("…`

> Bash `{"command":"go build ./... && go test ./internal/config/ 2>&1 | tail -15"}`

**Agent:** I added `CLEANUP_MAX_ENTRIES_PER_FEED`. The build and the config tests pass, but I haven't run the new SQL against a database, so the storage query is untested.

- **Config:** the default is `0`, which means no limit. Negative values are rejected. The option is in `internal/config/options.go`, with a `CleanupMaxEntriesPerFeed()` getter, a parsing test and a `miniflux.1` entry.
- **Storage:** the new `ArchiveEntriesOverFeedLimit(maxEntries, limit)` in `internal/storage/entry.go` ranks each feed's entries by `published_at` (newest first). It deletes those past the limit, up to the batch size.
- **Tombstones:** it records a tombstone for each deleted entry, like `ArchiveEntries`, so a later refresh doesn't bring them back.
- **Protected entries:** starred and shared entries are never deleted, but they still count toward the limit. A feed full of bookmarks can therefore stay above the cap.
- **Cleanup job:** `runCleanupTasks` calls it after the read and unread archiving steps. It only runs when the option is above 0, and it shares `CLEANUP_ARCHIVE_BATCH_SIZE` as its per-run cap.
- **Timing:** the cap is enforced by the periodic cleanup job (`CLEANUP_FREQUENCY_HOURS`, 24 hours by default). A feed can exceed it between runs.

Nothing is committed.

