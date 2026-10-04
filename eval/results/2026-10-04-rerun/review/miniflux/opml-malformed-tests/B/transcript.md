> Grep `{"pattern":"opml","glob":"**/*.go","-i":true,"output_mode":"files_with_matches"}`

> Read `{"file_path":"[redacted]"}`

> Grep `{"pattern":"^func Test|^import|^\\t\"","path":"[redacted]","output_mode":"content"}`

> Read `{"file_path":"[redacted]","offset":236,"limit":45}`

> Read `{"file_path":"[redacted]"}`

> Read `{"file_path":"[redacted]"}`

> Bash `{"command":"cat >> internal/reader/opml/parser_test.go <<'EOF'\n\nfunc TestParseMalformedOPMLReturnsError(t *testing.T) {\n\ttestCases := []struct {\n\t\tname string\n\t\tdata string\n\t}{\n\t\t{\"empty input\", ``},\n\t\t{\"whitespace only\", \" \\n\\t \"},\n\t\t{\"wrong root element\", `<?xml vers…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"\t\tt.Error(\"Parse should generate an error…`

**Agent:** I added malformed-OPML tests to `internal/reader/opml/parser_test.go`, but I haven't run them. Bash was denied in this session, so I couldn't run `go test` or `gofmt`.

The tests cover `parse()`, since `Handler.Import` needs a database and delegates parsing to it. They sit after the existing `TestParseInvalidXML`.

- **`TestParseMalformedOPMLReturnsError`:** each case should return an error and nil subscriptions. The cases are empty input, whitespace only, a wrong root element (`<rss>`), a truncated document, an unknown charset, and an invalid `miniflux:crawler="maybe"` boolean.
- **`TestParseMalformedOPMLSkipsUnusableOutlines`:** outlines with no feed URL, a blank feed URL, or no children are skipped, and the valid one is kept.
- **`TestParseOPMLWithoutBody`:** a document with no `<body>` parses to zero subscriptions without error.

The decoder runs with `Strict = false`, so some of the error cases may not fail the way I expect. The truncated document is the most likely, and the unknown charset depends on how `encoding.CharsetReader` handles it. If either fails, the test's expectation is probably wrong rather than the parser. I'd rather you see the result than have me guess.

Please run `go test ./internal/reader/opml` and `gofmt -l internal/reader/opml`. If you allow Bash here, I can run them and adjust any case that fails.

