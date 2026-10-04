> Grep `{"pattern":"opml","-i":true,"glob":"**/*_test.go","output_mode":"files_with_matches"}`

> Glob `{"pattern":"internal/reader/opml/*.go"}`

> Read `{"file_path":"[redacted]"}`

> Read `{"file_path":"[redacted]"}`

> Read `{"file_path":"[redacted]"}`

> Bash `{"command":"cat >> internal/reader/opml/parser_test.go <<'EOF'\n\nfunc TestParseMalformedOpmlReturnsError(t *testing.T) {\n\ttestCases := []struct {\n\t\tname string\n\t\tdata string\n\t}{\n\t\t{\"empty document\", ``},\n\t\t{\"whitespace only\", \" \\n\\t \"},\n\t\t{\"truncated document\", `<?xml v…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"\tif subscriptions[0] != expected {\n\t\tt.Er…`

**Agent:** I added malformed-OPML tests to `internal/reader/opml/parser_test.go`, but I haven't run them. Bash was denied in this session, so `go test ./internal/reader/opml/` is still to do. Some expectations may need adjusting once it runs, because the parser decodes with `Strict = false`.

The new tests:
- **`TestParseMalformedOpmlReturnsError`** is a table test. Each case must return an error and a nil result:
  - an empty document;
  - whitespace-only input;
  - a truncated document;
  - an unclosed `<body>`;
  - a mismatched closing tag;
  - a non-OPML root element (`<rss>`);
  - an invalid `miniflux:crawler="maybe"` boolean.
- **`TestParseOpmlWithoutBody`** checks that an OPML with only a `<head>` parses to zero subscriptions without an error.
- **`TestParseOpmlIgnoresOutlinesWithoutFeedURL`** checks that outlines with no `xmlUrl`, a blank `xmlUrl`, or no children are skipped, and only the valid feed is returned.

I left out an "unknown encoding" case. I wasn't sure how `encoding.CharsetReader` handles an unknown charset, so I didn't want to guess.

Can you allow `go test` so I can run these and fix any wrong expectations?

