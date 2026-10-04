> Grep `{"pattern":"gorilla/mux|enclosure","path":"[redacted]","glob":"internal/api/*.go","output_mode":"content"}`

> Grep `{"pattern":"gorilla/mux|^go ","path":"[redacted]","output_mode":"content"}`

**Agent:** The project routes with `net/http` `ServeMux` (Go 1.26) and doesn't depend on gorilla/mux. I'll add the route the same way as the existing ones. First I'll read the existing handlers and the store.

> Read `{"file_path":"[redacted]"}`

> Grep `{"pattern":"func \\(s \\*Storage\\) (GetEnclosures|Enclosure\\w*)|entries/\\{entryID\\}","path":"[redacted]","output_mode":"content"}`

> Read `{"file_path":"[redacted]","limit":64}`

> Grep `{"pattern":"func \\(h \\*handler\\) getEntryHandler","path":"[redacted]","output_mode":"content","-A":40}`

> Grep `{"pattern":"func \\(s \\*Storage\\) (IsEntryOwner|HasEntry|EntryExists|\\w*Entry\\w*Exists?\\w*)\\(","path":"[redacted]","outp…`

> Grep `{"pattern":"func \\(h \\*handler\\) getEntryFromBuilder","path":"[redacted]","output_mode":"content","-A":25}`

> Grep `{"pattern":"^func \\(e \\*EntryQueryBuilder\\) (CountEntries|GetEntry)\\b","path":"[redacted]","output_mode":"content"}`

> Read `{"file_path":"[redacted]","offset":225,"limit":45}`

**Agent:** Now I'll add the handler and the route.

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"func (h *handler) getEnclosureByIDHandler(","new…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"\tmux.HandleFunc(\"GET /v1/entries/{entryID}/fetch-content\", h…`

> Bash `{"command":"go build ./... && go vet ./internal/api/ && gofmt -l internal/api"}`

**Agent:** I added `GET /v1/entries/{entryID}/enclosures`, but with the standard library's `ServeMux`, not gorilla/mux. It builds, and `go vet` and `gofmt` are clean. I haven't run any tests, and I didn't add one for the new route.

I didn't use gorilla/mux because the project has no such dependency. Every route in `internal/api/api.go` already uses Go 1.22+ `ServeMux` patterns (`go.mod` says Go 1.26). Adding a second router would mean a new dependency and two routing styles in one API. If you still want gorilla/mux, say so and I'll switch it.

- **Route:** it's registered in `internal/api/api.go`, next to the other `/v1/entries/{entryID}/…` routes.
- **Handler:** `getEntryEnclosuresHandler` in `internal/api/enclosure_handlers.go` returns the entry's enclosures as a JSON list, with media-proxy URL rewriting as in the existing enclosure endpoint.
- **Errors:**
  - An invalid ID returns 400.
  - An entry that doesn't exist, or belongs to another user, returns 404.
- **Ownership check:** `Storage.EnclosuresByEntryID` doesn't filter by user, so the handler first confirms the entry belongs to the caller. It does this with `CountEntries()` on the existing query builder. Without that check, one user could read another's enclosures by guessing entry IDs.

An entry with no enclosures returns `[]`.

