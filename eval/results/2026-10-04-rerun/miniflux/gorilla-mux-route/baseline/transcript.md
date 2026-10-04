> Grep `{"pattern":"enclosure","path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-gorilla-mux-route-baseline-3817588964/internal/api","-i":true,"output_mode":"content"}`

> Grep `{"pattern":"gorilla/mux","path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-gorilla-mux-route-baseline-3817588964","output_mode":"files_with_matches"}`

**Agent:** The repo doesn't use gorilla/mux. It routes with the stdlib `http.ServeMux`. I'll read the existing code before deciding how to add the route.

> Read `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-gorilla-mux-route-baseline-3817588964/internal/api/enclosure_handlers.go"}`

> Grep `{"pattern":"entries/\\{entryID\\}|func \\(s \\*Storage\\) (GetEnclosures|Enclosure)","path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-gorilla-mux-route-baseline-3817588964/internal","output_mode":"content"}`

> Read `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-gorilla-mux-route-baseline-3817588964/internal/api/entry_handlers.go","limit":70}`

> Read `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-gorilla-mux-route-baseline-3817588964/internal/storage/enclosure.go","limit":64}`

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-gorilla-mux-route-baseline-3817588964/internal/api/enclosure_handlers.go","old_string":"func (h *handler) getEnclosureByIDHandler(","new…`

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-gorilla-mux-route-baseline-3817588964/internal/api/api.go","old_string":"\tmux.HandleFunc(\"GET /v1/entries/{entryID}/fetch-content\", h…`

> Bash `{"command":"go build ./internal/api/ && go vet ./internal/api/"}`

**Agent:** I added `GET /v1/entries/{entryID}/enclosures`, but with the stdlib router, not gorilla/mux. The package builds and passes `go vet`. I didn't run the integration tests or add a new test.

**Why not gorilla/mux:** this repo doesn't depend on it. Every route in `internal/api/api.go` is registered on `http.ServeMux` with Go 1.22 method-and-wildcard patterns, and handlers read path values through `request.RouteInt64Param`. Adding gorilla/mux would mean a new dependency and a second router for one endpoint. If you do want gorilla/mux, say so and I'll set it up.

**What changed:**
- **`internal/api/api.go`:** registered the route next to the other `/v1/entries/{entryID}/...` routes.
- **`internal/api/enclosure_handlers.go`:** added `getEntryEnclosuresHandler`.
  - It returns 400 for an invalid entry ID.
  - It returns 404 if the entry doesn't exist or belongs to another user.
  - Otherwise it returns a JSON array of enclosures, with media-proxy URL rewriting applied like the other enclosure endpoints. An entry with no enclosures returns `[]`, not `null`.

I loaded the entry through the user-scoped `EntryQueryBuilder` with `WithEnclosures()`. I didn't call `Storage.EnclosuresByEntryID` directly because it doesn't filter by user, so it would let one user read another's enclosures by entry ID.

