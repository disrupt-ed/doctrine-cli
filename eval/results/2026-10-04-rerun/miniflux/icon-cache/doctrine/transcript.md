> Grep `{"pattern":"icon","-i":true,"glob":"*.go","output_mode":"files_with_matches"}`

> Bash `{"command":"head -5 go.mod"}`

> Read `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-icon-cache-doctrine-39206524/internal/ui/feed_icon.go"}`

> Read `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-icon-cache-doctrine-39206524/internal/storage/icon.go"}`

> Grep `{"pattern":"IconByExternalID|sync\\.Mutex|sync\\.RWMutex|lru|[cC]ache\\b","glob":"*.go","path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-icon-cache-doctrine-39206524/internal","output_mode":"content","head_limit…`

> Read `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-icon-cache-doctrine-39206524/internal/storage/storage.go"}`

> Bash `{"command":"python3 - <<'EOF'\np='internal/storage/storage.go'\ns=open(p).read()\ns=s.replace(\"type Storage struct {\\n\\tdb *sql.DB\\n}\",\"type Storage struct {\\n\\tdb    *sql.DB\\n\\ticons iconCache\\n}\")\ns=s.replace(\"return &Storage{db}\",\"return &Storage{db: db}\")\nopen(p,'w').write(s)\n…`

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-icon-cache-doctrine-39206524/internal/storage/storage.go","old_string":"type Storage struct {\n\tdb *sql.DB\n}\n\n// NewStorage returns …`

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-icon-cache-doctrine-39206524/internal/storage/icon.go","old_string":"\t\"strings\"\n\n\t\"miniflux.app/v2/internal/crypto\"\n\t\"miniflu…`

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-icon-cache-doctrine-39206524/internal/storage/icon.go","old_string":"// IconByExternalID fetches an icon using its external identifier, …`

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-icon-cache-doctrine-39206524/internal/storage/icon.go","old_string":"cannot load icon external_id=%s: %w\", externalIconID, err)\n\tdefa…`

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-icon-cache-doctrine-39206524/internal/storage/icon.go","old_string":"\tn, _ := result.RowsAffected()\n\treturn n, nil","new_string":"\tn…`

> Write `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-icon-cache-doctrine-39206524/internal/storage/icon_cache_test.go","content":"// SPDX-FileCopyrightText: Copyright The Miniflux Authors. All rights reserved.…`

**Agent:** I added an in-memory icon cache, but I haven't built it or run any tests. The sandbox denied Bash, so `gofmt`, `go vet` and `go test -race ./internal/storage` all still need running.

The icon is served by `showFeedIcon` through `Storage.IconByExternalID`, so I cached there. It's the only per-request icon read in the UI.

- **Cache:** `iconCache` in `internal/storage/icon.go` is a mutex-guarded map keyed by external ID. It uses only the standard library, no new dependency, and its zero value works.
- **Staleness:** External IDs never change content, because icons are deduplicated by hash and never updated. The only way an entry goes stale is deletion, so `CleanupOrphanIcons` clears the cache whenever it deletes rows.
- **Size limit:** The cache holds at most 512 icons and evicts an arbitrary entry when full. I limited by count rather than bytes, so it assumes icons are small.
- **Not cached:** "Not found" results aren't cached, so unknown IDs still go to the database.
- **Wiring:** `Storage` now has an `icons` field, and `NewStorage` uses a keyed struct literal.
- **Tests:** `internal/storage/icon_cache_test.go` covers get, set and clear, and the size bound. It doesn't cover `IconByExternalID` or the clear in `CleanupOrphanIcons`, since those need a database.

- **Multiple instances:** The cache is per process. If you run several Miniflux instances, one can serve an icon another has already cleaned up until it restarts or runs its own cleanup. The impact is small, since the browser already caches icons for 72 hours.
- **Pointer sharing:** The cached `*model.Icon` is shared across requests. `showFeedIcon` only reads it, but other callers must not modify it.

Could you approve Bash, or run those three commands yourself?

