> Grep `{"pattern":"func Test","path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-opml-malformed-tests-doctrine-4065262728/internal/reader/opml","output_mode":"content"}`

> Glob `{"pattern":"internal/reader/opml/*.go","path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-opml-malformed-tests-doctrine-4065262728"}`

> Read `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-opml-malformed-tests-doctrine-4065262728/internal/reader/opml/parser_test.go","offset":230,"limit":55}`

> Read `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-opml-malformed-tests-doctrine-4065262728/internal/reader/opml/parser.go"}`

> Read `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-opml-malformed-tests-doctrine-4065262728/internal/reader/opml/handler.go"}`

> Bash `{"command":"cat >> internal/reader/opml/parser_test.go <<'EOF'\n\nfunc TestParseMalformedOpml(t *testing.T) {\n\ttests := []struct {\n\t\tname string\n\t\tdata string\n\t}{\n\t\t{\"empty input\", ``},\n\t\t{\"whitespace only\", \" \\n\\t \"},\n\t\t{\"wrong root element\", `<?xml version=\"1.0\"?><rs…`

> Read `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-opml-malformed-tests-doctrine-4065262728/internal/reader/opml/opml.go"}`

> Read `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-opml-malformed-tests-doctrine-4065262728/internal/reader/opml/parser_test.go","offset":279}`

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-opml-malformed-tests-doctrine-4065262728/internal/reader/opml/parser_test.go","old_string":"func TestParseOpmlWithMinifluxSettings(t *te…`

**Agent:** I added two tests to `internal/reader/opml/parser_test.go`, but I haven't run them. Bash was denied in this session, so `go test` never ran and they may not pass yet.

- **`TestParseMalformedOpml`:** each of these cases should make `parse` return an error and a nil result:
  - empty input
  - whitespace only
  - a wrong root element (`<rss>`)
  - truncation before the body
  - truncation inside an outline
  - an invalid `miniflux:crawler="maybe"` boolean
- **`TestParseOpmlWithoutSubscriptions`:** an empty `<body>` and an outline with no `xmlUrl` should parse without error and return zero subscriptions.

I tested at the `parse` level because `Handler.Import` needs a database. A parse error is returned from `Import` before any storage call.

I left out a few cases I couldn't confirm. `parse` runs the XML decoder in non-strict mode, so I couldn't tell whether mismatched or unclosed tags and an unknown charset return errors. Add them once you can run the tests.

Please run `go test ./internal/reader/opml/`. If a case fails, the parser's behaviour differs from what I assumed from the code, so check that case first.

