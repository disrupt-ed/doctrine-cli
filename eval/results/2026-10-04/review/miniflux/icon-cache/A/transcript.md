> Grep `{"pattern":"Icon","glob":"internal/{storage,ui,http}/**/*.go","output_mode":"files_with_matches","path":"[redacted]"}`

> Read `{"file_path":"[redacted]"}`

> Read `{"file_path":"[redacted]"}`

> Grep `{"pattern":"IconByExternalID|CleanupOrphanIcons|lru|sync\\.Map|expirable","path":"[redacted]","output_mode":"content","glob":"*.{go,mod}"}`

> Bash `{"command":"head -5 go.mod && sed -n 1,60p internal/storage/storage.go && ls internal/storage | head -50"}`

> Write `{"file_path":"[redacted]","content":"// SPDX-FileCopyrightText: Copyright The Miniflux Authors. All rights reserved.\n//…`

> Write `{"file_path":"[redacted]","content":"// SPDX-FileCopyrightText: Copyright The Miniflux Authors. All rights reserved…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"\tdb *sql.DB\n}\n\n// NewStorage returns a new Storage.\nfunc N…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"// IconByExternalID fetches an icon using its external identifier,…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"\t\treturn nil, fmt.Errorf(\"store: cannot load icon external_id=%…`

> Bash `{"command":"gofmt -l internal/storage; go vet ./internal/storage ./internal/ui && go test -race ./internal/storage 2>&1 | tail"}`

**Agent:** I added an in-memory cache for feed icons. `go vet` is clean and the storage tests pass with `-race`. I didn't run the app, so the cache hit path isn't checked against a real database.

**What I cached.** The hot path is `showFeedIcon`, which calls `Storage.IconByExternalID` on every request. The cache sits inside that method, so the handler and its callers are unchanged. It uses only the standard library: a mutex-guarded map, with no new dependency.

**Files changed:**
- `internal/storage/icon_cache.go` (new): the `iconCache` type, a small mutex-guarded map.
- `internal/storage/storage.go`: `Storage` gets an `iconCache` field. `NewStorage` now builds the struct with a named field.
- `internal/storage/icon.go`: `IconByExternalID` checks the cache first and stores icons it finds in the database.
- `internal/storage/icon_cache_test.go` (new): tests for hit and miss, expiry, and the size bound.

**Choices to check:**
- **Why caching is safe:** An icon's content is fixed for a given external ID, and `StoreFeedIcon` only adds new rows. So the cache never serves wrong content.
- **Staleness:** `CleanupOrphanIcons` can delete an icon that is still cached. Entries expire after 1 hour, so a deleted icon can be served for up to that long. The browser already caches these responses for 72 hours, so this seemed fine.
- **Size limit:** The cache holds at most 1000 icons, and the whole map is cleared when it fills. That keeps memory bounded and the code short. If you want true LRU eviction, it's a bigger change.
- **Unknown IDs:** I don't cache lookups that find nothing. Otherwise requests for random IDs could fill the cache.
- **Multiple instances:** The cache is per process. Each instance keeps its own copy, which is fine for icons.

I left `IconByFeedID`, `IconByUserAndIconID` and `Icons` uncached. They aren't on the per-request icon path.

