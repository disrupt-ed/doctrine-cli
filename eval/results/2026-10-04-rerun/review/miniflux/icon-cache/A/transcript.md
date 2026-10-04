> Grep `{"pattern":"icon","-i":true,"path":"[redacted]","glob":"internal/ui/*.go","output_mode":"files_with_matches"}`

> Read `{"file_path":"[redacted]"}`

> Grep `{"pattern":"func \\(s \\*Storage\\) (IconByFeedID|IconByHash|.*Icon.*)\\(","path":"[redacted]","output_mode":"content"}`

> Read `{"file_path":"[redacted]"}`

> Grep `{"pattern":"IconByExternalID|IconByFeedID|IconByUserAndIconID","path":"[redacted]","output_mode":"content"}`

> Bash `{"command":"sed -n 1,60p internal/storage/storage.go && grep -rli \"lru\\|sync.Map\" internal --include=*.go | head; ls internal/storage | head -50"}`

> Bash `{"command":"cat > internal/storage/icon_cache.go <<'EOF'\n// SPDX-FileCopyrightText: Copyright The Miniflux Authors. All rights reserved.\n// SPDX-License-Identifier: Apache-2.0\n\npackage storage // import \"miniflux.app/v2/internal/storage\"\n\nimport (\n\t\"sync\"\n\t\"time\"\n\n\t\"miniflux.app/…`

**Agent:** Bash is denied, so I'll make the edits with the Write and Edit tools instead.

> Write `{"file_path":"[redacted]","content":"// SPDX-FileCopyrightText: Copyright The Miniflux Authors. All rights reserved.\n/…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"type Storage struct {\n\tdb *sql.DB\n}\n\n// NewStorage return…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"func (s *Storage) IconByExternalID(externalIconID string) (*model…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"external_id=%s: %w\", externalIconID, err)\n\tdefault:\n\t\tretur…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"\tn, _ := result.RowsAffected()","new_string":"\ts.iconsCache.cle…`

**Agent:** I added an in-memory icon cache, but I haven't compiled it or run any tests. The shell was denied in this session, so I made the edits with the file tools and `go build` never ran.

**What changed**
- **Cache lookup:** `IconByExternalID` in `internal/storage/icon.go` now checks the cache before the database. This is the lookup behind `/feed-icon/{id}`, the route that runs on every icon request. A miss falls through to the database and stores the result.
- **Cache itself:** the new `internal/storage/icon_cache.go` is a mutex-guarded map keyed by external ID. Entries expire after 1 hour. Total content is capped at 16 MiB, and when it's full, arbitrary entries are evicted rather than the least recently used. An icon larger than 16 MiB is never cached.
- **Wiring:** `Storage` in `storage.go` now holds the cache, and `NewStorage` creates it.
- **Invalidation:** icon rows never change once created, so I only clear the cache in `CleanupOrphanIcons`, the one place rows are deleted.
- **Not found:** missing icons aren't cached, so a 404 still hits the database.

Cached icons are shared pointers, so callers must not modify them. The only caller, `showFeedIcon`, just reads.

I left `IconByFeedID` and `IconByUserAndIconID` uncached. They back the API, not the per-request UI path.

I haven't added tests for the cache either. Run `go build ./... && go test ./internal/storage/` to check it. If you allow Bash I can run those and add a unit test for expiry and eviction.

