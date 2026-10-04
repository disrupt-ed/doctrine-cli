> Grep `{"pattern":"func Test","path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-opml-malformed-tests-doctrine-1092124467/internal/reader/opml","output_mode":"content"}`

> Read `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-opml-malformed-tests-doctrine-1092124467/internal/reader/opml/parser_test.go","offset":230,"limit":50}`

> Bash `{"command":"cat internal/reader/opml/parser.go; grep -rn \"opml.Parse\\|func.*Import\" internal/reader/opml/*.go internal/ui/opml_import.go internal/api/*.go | head -20; ls internal/reader/opml"}`

> Bash `{"command":"python3 - <<'EOF'\np='internal/reader/opml/parser_test.go'\ns=open(p).read()\nmarker=\"func TestParseOpmlWithMinifluxSettings\"\nnew='''func TestParseMalformedOPML(t *testing.T) {\n\ttestCases := []struct {\n\t\tname string\n\t\tdata string\n\t}{\n\t\t{\"empty input\", ``},\n\t\t{\"white…`

> Read `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-opml-malformed-tests-doctrine-1092124467/internal/reader/opml/opml.go"}`

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-opml-malformed-tests-doctrine-1092124467/internal/reader/opml/parser_test.go","old_string":"func TestParseOpmlWithMinifluxSettings(","ne…`

**Agent:** I added the tests to `internal/reader/opml/parser_test.go`, but I haven't run them. Bash was denied in this session, so `go test` and `gofmt` have not been checked.

- **`TestParseMalformedOPML`** is a table-driven test. Each case should make `parse` return an error and a nil slice:
  - empty input
  - whitespace-only input
  - a non-OPML root element (`<rss>`)
  - a document truncated before `<body>`
  - a document truncated inside an attribute
  - an invalid `miniflux:crawler="maybe"` boolean
- **`TestParseOpmlSkipsOutlinesWithoutFeedURL`** checks that an outline with no `xmlUrl` is skipped and a valid sibling is still imported.

I left out an unsupported-encoding case. I wasn't sure how `encoding.CharsetReader` handles it, and without a test run I couldn't check.

Please run `go test ./internal/reader/opml` and `gofmt -l internal/reader/opml`. I'm not certain how the non-strict decoder treats the truncated cases, so those are the ones most likely to need adjusting. If you allow Bash, I'll run both and fix any failures.

