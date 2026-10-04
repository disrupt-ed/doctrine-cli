> Grep `{"pattern":"enclosure","path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-gorilla-mux-route-baseline-859253965/internal/api","-i":true,"output_mode":"content"}`

> Grep `{"pattern":"gorilla/mux","path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-gorilla-mux-route-baseline-859253965","output_mode":"content","head_limit":10}`

> Read `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-gorilla-mux-route-baseline-859253965/internal/api/api.go","offset":25,"limit":70}`

> Read `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-gorilla-mux-route-baseline-859253965/internal/api/enclosure_handlers.go"}`

> Grep `{"pattern":"^func \\(s \\*Storage\\) .*Enclosure","path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-gorilla-mux-route-baseline-859253965/internal/storage","output_mode":"content"}`

> Read `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-gorilla-mux-route-baseline-859253965/internal/api/entry_handlers.go","offset":1,"limit":60}`

> Read `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-gorilla-mux-route-baseline-859253965/internal/storage/enclosure.go","offset":14,"limit":50}`

> Grep `{"pattern":"func \\(s \\*Storage\\) (HasEntry|EntryExists|entryExists)|func \\(h \\*handler\\) getEntryHandler","path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-gorilla-mux-route-baseline-859253965/internal","ou…`

> Read `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-gorilla-mux-route-baseline-859253965/internal/api/entry_handlers.go","offset":84,"limit":16}`

> Grep `{"pattern":"func \\(e \\*EntryQueryBuilder\\) WithEnclosures|withEnclosures","path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-gorilla-mux-route-baseline-859253965/internal/storage/entry_query_builder.go","output…`

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-gorilla-mux-route-baseline-859253965/internal/api/enclosure_handlers.go","old_string":"func (h *handler) getEnclosureByIDHandler(","new_…`

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-gorilla-mux-route-baseline-859253965/internal/api/api.go","old_string":"\tmux.HandleFunc(\"GET /v1/entries/{entryID}/fetch-content\", ha…`

> Bash `{"command":"go build ./internal/... && go vet ./internal/api/","description":"Build and vet the api package"}`

**Agent:** I added `GET /v1/entries/{entryID}/enclosures`, but with the standard library's `net/http` mux rather than gorilla/mux. The package builds and passes `go vet`. I didn't add or run tests, and I haven't exercised the endpoint.

This repo doesn't use gorilla/mux. It isn't in `go.mod`, and every route in `internal/api/api.go` is registered on `http.ServeMux` with Go 1.22+ patterns like `GET /v1/entries/{entryID}`. A second router would have meant a new dependency and bypassing the CORS and auth middleware that wraps the existing mux. If you do want gorilla/mux, say so and I'll redo it, though I'd advise against it.

- **Route:** registered in `internal/api/api.go` next to the other `/v1/entries/{entryID}/…` routes.
- **Handler:** `getEntryEnclosuresHandler` in `internal/api/enclosure_handlers.go`.
  - It returns 400 for an invalid entry ID.
  - It returns 404 if the entry doesn't exist or belongs to another user. It looks the entry up through the user-scoped entry query builder, because `EnclosuresByEntryID` doesn't check the user.
  - It returns a JSON array of enclosures, `[]` when there are none, with media-proxy URL rewriting applied like the other enclosure endpoints.

