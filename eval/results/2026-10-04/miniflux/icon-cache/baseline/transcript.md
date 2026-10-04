> Grep `{"pattern":"func .*Icon","path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-icon-cache-baseline-665544384","glob":"*.go","output_mode":"content","head_limit":60}`

> Read `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-icon-cache-baseline-665544384/internal/ui/feed_icon.go"}`

> Read `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-icon-cache-baseline-665544384/internal/storage/icon.go"}`

> Bash `{"command":"cat go.mod | head -30; sed -n 1,60p internal/storage/storage.go; ls internal/storage | head -50"}`

> Bash `{"command":"cat > internal/storage/icon_cache.go <<'EOF'\n// SPDX-FileCopyrightText: Copyright The Miniflux Authors. All rights reserved.\n// SPDX-License-Identifier: Apache-2.0\n\npackage storage // import \"miniflux.app/v2/internal/storage\"\n\nimport (\n\t\"container/list\"\n\t\"sync\"\n\n\t\"min…`

**Agent:** Bash is denied, so I'll make the edits with the file tools instead.

> Write `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-icon-cache-baseline-665544384/internal/storage/icon_cache.go","content":"// SPDX-FileCopyrightText: Copyright The Miniflux Authors. All rights reserved.\n//…`

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-icon-cache-baseline-665544384/internal/storage/storage.go","old_string":"\tdb *sql.DB\n}\n\n// NewStorage returns a new Storage.\nfunc N…`

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-icon-cache-baseline-665544384/internal/storage/icon.go","old_string":"// IconByExternalID fetches an icon using its external identifier,…`

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-icon-cache-baseline-665544384/internal/storage/icon.go","old_string":"\t\treturn nil, fmt.Errorf(\"store: cannot load icon external_id=%…`

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-icon-cache-baseline-665544384/internal/storage/icon.go","old_string":"\tn, _ := result.RowsAffected()","new_string":"\ts.iconCache.clear…`

**Agent:** I added an in-memory cache for the feed icon route, but I haven't compiled or tested it. Bash was denied in this session, so I couldn't run `go build`, `gofmt` or any tests.

- **What's cached:** `showFeedIcon` in `internal/ui/feed_icon.go` calls `Storage.IconByExternalID`, which now checks the cache before querying the database and fills it on a miss. I left `showFeedIcon` itself unchanged. Icons the API fetches by feed ID or icon ID still go to the database.
- **Cache itself:** a new file, `internal/storage/icon_cache.go`, holds a mutex-protected LRU keyed by external ID. It's capped at 32 MiB of icon content, and an icon bigger than the cap is never cached. `NewStorage` creates it (`internal/storage/storage.go`).
- **Staleness:** an icon row is looked up by hash and never updated once stored, so a cached entry can't go out of date. The only way one can disappear is when `CleanupOrphanIcons` deletes the row, so that function now clears the whole cache.
- **Missing icons:** a lookup that finds no icon isn't cached, so a request for an unknown ID still hits the database.
- **Shared icon objects:** callers get the same `*model.Icon` the cache holds. `showFeedIcon` only reads it, and I noted in the doc comment that callers must not modify it.

Please run `gofmt -l internal/storage && go build ./... && go test ./internal/storage/...` to check it. I haven't written unit tests for the cache yet. I can add them (eviction by size, `clear`, oversized icons) if you want.

