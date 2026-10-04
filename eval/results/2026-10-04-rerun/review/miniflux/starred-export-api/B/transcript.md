**Agent:** I'll start by looking at how the existing API handles entries and exports.

> Grep `{"pattern":"starred|/export|Export","path":"[redacted]","output_mode":"content","-i":true}`

> Read `{"file_path":"[redacted]","offset":1,"limit":210}`

> Read `{"file_path":"[redacted]","offset":40,"limit":40}`

> Grep `{"pattern":"func \\(e \\*EntryQueryBuilder\\) (WithLimit|WithStarred|GetEntries|WithSorting|WithEnclosures)|limit > 0|func ValidateRange","path":"[redacted]…`

> Bash `{"command":"head -5 go.mod && grep -n \"miniflux \\\"\" internal/api/api_integration_test.go | head -3; grep -n \"func JSON\\b\\|func JSON(\" -A12 internal/http/response/response.go; sed -n 2287,2325p internal/api/api_integration_test.go; sed -n 95,130p internal/api/api_test.go; grep -rn \"v1/export…`

**Agent:** Now the client's `Export` and `Entries` methods, to see how a test would call this.

> Bash `{"command":"grep -n \"func (c \\*Client) Export\" -A22 client/client.go; grep -n \"func (c \\*Client) Entries(\" -A25 client/client.go; grep -n \"Content-Disposition\" -r internal | head -5; sed -n 1,40p internal/http/response/json.go 2>/dev/null; ls docs 2>/dev/null"}`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"func (h *handler) findEntries(","new_string":"// ex…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"\tmux.HandleFunc(\"POST /v1/import\", ","new_string":"\tmux.Ha…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"// Import imports an OPML file.\nfunc (c *Client) Import(","new_s…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"func TestImportEndpoint(t *testing.T) {","new…`

> Bash `{"command":"gofmt -l internal client; go vet ./internal/api ./client && go test ./internal/api ./client 2>&1 | tail -15"}`

> Bash `{"command":"python3 - <<'E'\np='internal/api/api_integration_test.go'\ns=open(p).read()\na=s.index('func TestExportStarredHandlerRequiresAuthentication')\nb=s.index('func TestImportEndpoint')\nblk=s[a:b]\nopen(p,'w').write(s[:a]+s[b:])\nopen('internal/api/api_test.go','a').write('\\n'+blk.rstrip()+'…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"func TestExportStarredHandlerRequiresAuthenti…`

> Read `{"file_path":"[redacted]","offset":106,"limit":4}`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"func TestGetEntryIDsHandlerRejectsInvalidStarredParam(","…`

> Bash `{"command":"gofmt -d internal/api/api.go | head -20; go vet ./internal/api ./client && go test ./internal/api ./client 2>&1 | tail -15"}`

**Agent:** My edit dropped a space in `api.go`, so I'm fixing that formatting now.

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"\"POST /v1/import\",handler","new_string":"\"POST /v1/import\"…`

> Bash `{"command":"gofmt -l internal client; go build ./... && go test ./internal/api 2>&1 | tail -3"}`

**Agent:** I added `GET /v1/export/starred`, which returns every starred entry for the authenticated user as JSON. The build, `gofmt`, `go vet` and the `internal/api` and `client` tests pass. The new integration test `TestExportStarredEndpoint` skips without a configured Miniflux test server, so I haven't run it.

**What I did and why**
- **Existing API:** `GET /v1/entries?starred=true` already lists starred entries, but it's paginated and capped by `model.MaxEntryLimit`. A dedicated endpoint is the only way to get everything in one call.
- **Handler:** `exportStarredEntriesHandler` is in `internal/api/entry_handlers.go`. It skips `WithLimit`, so the query has no `LIMIT`, and it rewrites media proxy URLs the same way `findEntries` does.
- **Response shape:** it returns `{"total": N, "entries": [...]}`, the same format as `/v1/entries`.
- **Route:** I put it next to the OPML export at `/v1/export/starred`. It goes through the normal authentication.
- **Client:** I added `Client.ExportStarred()` and `ExportStarredContext()`.

**Tests**
- The unauthenticated-access test, `TestExportStarredHandlerRequiresAuthentication`, ran and passes.
- The integration test checks that a new user gets an empty result and that one starred entry comes back as total 1.

**Tradeoffs**
- All starred entries are loaded into memory before the response is written. That's fine for typical use, but a user with tens of thousands of starred entries would produce a large response. Streaming the output would need a new query path, so I held off.
- I didn't add a `Content-Disposition: attachment` header, so browsers display the JSON instead of downloading it. I can add it if you want download behavior.

