**Agent:** I'll start by looking at how existing config options and entry cleanup work.

> Grep `{"pattern":"POLLING_LIMIT_PER_HOST|PollingLimitPerHost","output_mode":"files_with_matches"}`

> Grep `{"pattern":"func \\(s \\*Storage\\) (ArchiveEntries|RefreshFeedEntries|DeleteEntries|cleanupEntries)","output_mode":"content","path":"internal/storage"}`

> Bash `{"command":"head -5 go.mod"}`

> Read `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-max-entries-per-feed-doctrine-1543860708/internal/storage/entry.go","offset":300,"limit":120}`

> Grep `{"pattern":"ArchiveEntries|CleanupEntriesDays|CLEANUP_ARCHIVE_UNREAD|CleanupArchiveUnread|cleanupArchiveUnread","output_mode":"content"}`

> Read `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-max-entries-per-feed-doctrine-1543860708/internal/cli/cleanup_tasks.go"}`

> Read `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-max-entries-per-feed-doctrine-1543860708/internal/config/options.go","offset":100,"limit":90}`

> Grep `{"pattern":"CleanupArchiveBatchSize|CLEANUP_ARCHIVE_BATCH_SIZE","output_mode":"content","-n":true}`

> Read `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-max-entries-per-feed-doctrine-1543860708/internal/config/options.go","offset":658,"limit":25}`

> Read `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-max-entries-per-feed-doctrine-1543860708/miniflux.1","offset":194,"limit":36}`

> Grep `{"pattern":"func Test.*Entries|func newTestStorage|func .*Setup","path":"internal/storage","output_mode":"content"}`

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-max-entries-per-feed-doctrine-1543860708/internal/storage/entry.go","old_string":"// SetEntriesStatus update the status of the given lis…`

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-max-entries-per-feed-doctrine-1543860708/internal/config/options.go","old_string":"\t\t\t\"CLEANUP_REMOVE_SESSIONS_DAYS\": {","new_strin…`

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-max-entries-per-feed-doctrine-1543860708/internal/config/options.go","old_string":"func (c *configOptions) CleanupRemoveSessionsInterval…`

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-max-entries-per-feed-doctrine-1543860708/internal/cli/cleanup_tasks.go","old_string":"\tif nbIcons, err","new_string":"\tif config.Opts.…`

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-max-entries-per-feed-doctrine-1543860708/miniflux.1","old_string":".B CLEANUP_REMOVE_SESSIONS_DAYS","new_string":".B CLEANUP_MAX_ENTRIES…`

> Read `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-max-entries-per-feed-doctrine-1543860708/internal/config/options_parsing_test.go","offset":276,"limit":20}`

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-max-entries-per-feed-doctrine-1543860708/internal/config/options_parsing_test.go","old_string":"func TestCreateAdminOptionParsing(t *tes…`

> Bash `{"command":"gofmt -l internal; go vet ./internal/config ./internal/storage ./internal/cli && go test ./internal/config 2>&1 | tail -5"}`

**Agent:** I added `CLEANUP_MAX_ENTRIES_PER_FEED` to cap how many entries each feed keeps. `gofmt` and `go vet` are clean and the config tests pass, but I haven't run the new SQL against a real Postgres. There are no storage tests in the repo to extend.

- **Config:** the option defaults to `0`, which means no cap. It rejects negative values (`internal/config/options.go`), and I documented it in `miniflux.1`.
- **Cleanup job:** when the cap is above 0, the daily cleanup deletes the oldest entries of any feed over it (`internal/cli/cleanup_tasks.go`). Age is by `published_at`, and it deletes at most `CLEANUP_ARCHIVE_BATCH_SIZE` rows per run.
- **Storage:** the query is the new `ArchiveEntriesOverFeedLimit` in `internal/storage/entry.go`. Like `ArchiveEntries`, it leaves starred and shared entries alone and writes a tombstone for each deleted entry, so the next refresh doesn't bring it back.
- **Test:** I added one for parsing the option, covering the default, a valid value and a negative value.

**Behaviours to know about:**
- The cap applies on the cleanup schedule (default every 24 hours, set by `CLEANUP_FREQUENCY_HOURS`), not at the moment a feed refreshes, so a feed can sit above the cap between runs.
- Starred and shared entries still count toward a feed's total but are never deleted, so a feed can end up above the cap if it has many of them.
- Read and unread entries are treated the same. If you'd rather it only trim read entries, that's a small change to the query.

I didn't add an API or UI setting. The cap is global and set through the environment or config file.

