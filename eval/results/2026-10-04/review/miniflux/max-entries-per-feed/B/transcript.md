**Agent:** I'll start by looking at how Miniflux config options and entry cleanup work today.

> Grep `{"pattern":"ArchiveReadInterval|ARCHIVE_READ_DAYS|CleanupArchiveReadDays|CleanupArchiveUnreadDays","output_mode":"files_with_matches"}`

> Bash `{"command":"head -5 go.mod; git log --oneline -1"}`

> Read `{"file_path":"[redacted]"}`

> Grep `{"pattern":"ArchiveReadInterval|ArchiveUnreadInterval|ARCHIVE_UNREAD","output_mode":"content","path":"[redacted]"}`

> Read `{"file_path":"[redacted]","offset":100,"limit":80}`

> Grep `{"pattern":"func \\(s \\*Storage\\) (ArchiveEntries|RefreshFeedEntries)","output_mode":"content","path":"[redacted]","-A":4…`

> Read `{"file_path":"[redacted]","offset":655,"limit":30}`

> Bash `{"command":"grep -n \"CLEANUP_ARCHIVE_BATCH_SIZE\" -r . --include=* -I | grep -v \"^./.git/\"; sed -n 195,225p miniflux.1; sed -n 1190,1235p internal/config/options_parsing_test.go; grep -n \"ArchiveEntries\" -r internal --include=*_test.go | head"}`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"\t\t\t\"CLEANUP_REMOVE_SESSIONS_DAYS\": {","new_strin…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"func (c *configOptions) CleanupRemoveSessionsInterval…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"func (s *Storage) ArchiveEntries(","new_string":"// Ar…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"\tif nbIcons, err","new_string":"\tif rowsAffected…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":".TP\n.B CLEANUP_REMOVE_SESSIONS_DAYS","new_string":".TP\n.B CLEANUP_M…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"func TestCleanupArchiveUnreadIntervalOpt…`

> Bash `{"command":"gofmt -l internal; go vet ./internal/config ./internal/storage ./internal/cli && go test ./internal/config 2>&1 | tail -5"}`

**Agent:** I added a new config option, `CLEANUP_MAX_ENTRIES_PER_FEED`, that caps how many entries each feed keeps. `gofmt`, `go vet` and the config tests pass. The new SQL has not been run against a database.

- **Default:** `0`, which disables the cap. Negative values are rejected.
- **When it runs:** as part of the existing daily cleanup job, in `internal/cli/cleanup_tasks.go`. It deletes at most `CLEANUP_ARCHIVE_BATCH_SIZE` entries per run, so a large backlog is trimmed over several runs.
- **What gets deleted:** for each feed with more than N entries, the oldest ones by `published_at`. I chose `published_at` because the feed's own ordering seemed the more natural meaning of "oldest".
- **What is protected:** starred and shared entries are never deleted, but they count toward N. A feed with N starred entries therefore keeps only those.
- **Tombstones:** deleted entries get a tombstone, as with the existing archive options, so the next feed refresh doesn't re-import them.
- **Code:** the query is `ArchiveEntriesOverFeedLimit` in `internal/storage/entry.go`. I also added the man page entry in `miniflux.1` and a parsing test in `internal/config/options_parsing_test.go`.

The ranking query uses a window function over the whole `entries` table. That is fine once a day, but on a very large database the first run may be slow.

The new storage query has no automated test, because `internal/storage` has no database tests. Run it against a Postgres instance before relying on it.

I did not add a per-feed or per-user setting, only the global one. If you want per-feed limits, that would need a schema change.

