> Grep `{"pattern":"func Test","path":"[redacted]","output_mode":"content"}`

> Glob `{"pattern":"internal/reader/opml/*.go","path":"[redacted]"}`

> Read `{"file_path":"[redacted]","offset":230,"limit":55}`

> Read `{"file_path":"[redacted]"}`

> Read `{"file_path":"[redacted]"}`

> Bash `{"command":"cat >> internal/reader/opml/parser_test.go <<'EOF'\n\nfunc TestParseMalformedOpml(t *testing.T) {\n\ttests := []struct {\n\t\tname string\n\t\tdata string\n\t}{\n\t\t{\"empty input\", ``},\n\t\t{\"whitespace only\", \" \\n\\t \"},\n\t\t{\"wrong root element\", `<?xml version=\"1.0\"?><rs…`

> Read `{"file_path":"[redacted]"}`

> Read `{"file_path":"[redacted]","offset":279}`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"func TestParseOpmlWithMinifluxSettings(t *te…`

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

