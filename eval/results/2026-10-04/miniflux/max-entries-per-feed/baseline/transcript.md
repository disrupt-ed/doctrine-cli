**Agent:** I'll start by looking at how config options and entry cleanup work in this codebase.

> Grep `{"pattern":"ArchiveReadDays|POLLING_PARSING_ERROR_LIMIT|PollingParsingErrorLimit","output_mode":"content","glob":"**/*.go","head_limit":60}`

> Grep `{"pattern":"POLLING_PARSING_ERROR_LIMIT","output_mode":"files_with_matches","head_limit":20}`

> Grep `{"pattern":"ARCHIVE_READ_DAYS|ArchiveReadDays|func.*ArchiveEntries|CleanOldEntries|RemoveEntries","output_mode":"content","glob":"!*_test.go"}`

> Read `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-max-entries-per-feed-baseline-3210720700/internal/config/options.go","offset":480,"limit":60}`

> Read `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-max-entries-per-feed-baseline-3210720700/internal/config/options.go","offset":120,"limit":60}`

> Grep `{"pattern":"func \\(c \\*configOptions\\) (Cleanup|Polling|Media|MaintenanceMode)","path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-max-entries-per-feed-baseline-3210720700/internal/config/options.go","output_mo…`

> Read `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-max-entries-per-feed-baseline-3210720700/internal/storage/entry.go","offset":340,"limit":90}`

> Grep `{"pattern":"ArchiveEntries|DeleteEntries|RemoveUnreadEntries|CleanupEntries","output_mode":"content","path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-max-entries-per-feed-baseline-3210720700/internal"}`

> Read `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-max-entries-per-feed-baseline-3210720700/internal/cli/cleanup_tasks.go"}`

> Read `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-max-entries-per-feed-baseline-3210720700/internal/config/options.go","offset":655,"limit":30}`

> Grep `{"pattern":"CLEANUP_ARCHIVE_UNREAD_DAYS","output_mode":"content","-C":6,"path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-max-entries-per-feed-baseline-3210720700/miniflux.1"}`

> Grep `{"pattern":"CLEANUP_ARCHIVE_BATCH_SIZE|CleanupArchiveBatchSize","output_mode":"content","path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-max-entries-per-feed-baseline-3210720700/internal/config","glob":"*_test.g…`

> Read `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-max-entries-per-feed-baseline-3210720700/miniflux.1","offset":190,"limit":50}`

> Read `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-max-entries-per-feed-baseline-3210720700/internal/config/options_parsing_test.go","offset":275,"limit":60}`

**Agent:** Now the edits: config option, accessor, storage method, cleanup task, docs, test.

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-max-entries-per-feed-baseline-3210720700/internal/config/options.go","old_string":"\t\t\t\"CLEANUP_REMOVE_SESSIONS_DAYS\": {","new_strin…`

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-max-entries-per-feed-baseline-3210720700/internal/config/options.go","old_string":"func (c *configOptions) CleanupRemoveSessionsInterval…`

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-max-entries-per-feed-baseline-3210720700/internal/storage/entry.go","old_string":"// SetEntriesStatus update the status of the given lis…`

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-max-entries-per-feed-baseline-3210720700/internal/cli/cleanup_tasks.go","old_string":"\tif nbIcons, err","new_string":"\tif maxEntries :…`

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-max-entries-per-feed-baseline-3210720700/miniflux.1","old_string":".B CLEANUP_REMOVE_SESSIONS_DAYS","new_string":".B CLEANUP_MAX_ENTRIES…`

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-max-entries-per-feed-baseline-3210720700/internal/config/options_parsing_test.go","old_string":"func TestCreateAdminOptionParsing(t *tes…`

> Bash `{"command":"go build ./... && go test ./internal/config/ 2>&1 | tail -15","description":"Build and run config tests"}`

**Agent:** I added a `CLEANUP_MAX_ENTRIES_PER_FEED` option that caps how many entries each feed keeps. The build passes and the config tests pass. I didn't run the new SQL against a database.

**How it behaves**
- **Default:** `0`, which disables the cap, so nothing changes unless you set it. Negative values are rejected.
- **When it runs:** as part of the regular cleanup job, after the read and unread archiving steps. The job's `CLEANUP_ARCHIVE_BATCH_SIZE` limits how many entries it removes per run.
- **What gets deleted:** for any feed over the cap, the oldest entries by publication date (ties broken by id).
- **Tombstones:** deleted entries get a tombstone, so the next refresh doesn't bring them back. This is the same mechanism the existing archiving uses.
- **Protected entries:** starred and shared entries are never deleted, but they do count toward the cap. A feed with many starred entries can therefore stay above the cap.
- **Unread entries:** unread entries over the cap are deleted too, not only read ones.

**Where the changes are**
- **Option and accessor:** `internal/config/options.go`.
- **Deletion query:** `ArchiveEntriesOverFeedLimit` in `internal/storage/entry.go`.
- **Wiring:** `internal/cli/cleanup_tasks.go`.
- **Docs:** `miniflux.1`.
- **Test:** the option test in `internal/config/options_parsing_test.go`.

The query ranks every entry in the table with a window function on each cleanup run. That may be slow on a very large database. If it is, I can limit the ranking to feeds that are over the cap.

I haven't committed anything.

