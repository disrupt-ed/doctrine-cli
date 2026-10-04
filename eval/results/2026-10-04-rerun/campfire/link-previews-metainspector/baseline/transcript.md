**Agent:** I'll start by looking at how messages and the Gemfile are set up.

> Bash `{"command":"git status | head -3; grep -n -i \"gem \" Gemfile | head -60; ls app/models app/models/message app/jobs app/views/messages; grep -rn -i \"unfurl\\|opengraph\\|link_preview\\|metainspector\\|ssrf\\|RestrictedHTTP\" --include=*.rb --include=*.erb --include=Gemfile* . -l | head"}`

> Bash `{"command":"ls app/models/opengraph app/models/message; cat app/models/opengraph/*.rb app/models/message.rb app/models/message/broadcasts.rb; grep -rln \"RestrictedHTTP\\|Opengraph\" app lib config test | head -30"}`

**Agent:** Link previews already exist here. I'll check how they work before changing anything.

> Bash `{"command":"cat app/models/opengraph/metadata/fetching.rb app/controllers/unfurl_links_controller.rb lib/restricted_http/private_network_guard.rb; sed -n 1,40p app/javascript/controllers/unfurl_controller.js"}`

**Agent:** I haven't made any changes. Link previews already exist in this app, and adding MetaInspector would work against how they're built.

**What's there now**
- When someone pastes a link into the composer, `unfurl_controller.js` POSTs it to `/unfurl_link` (`UnfurlLinksController`).
- `Opengraph::Metadata.from_url` fetches the page and parses the OG tags. The result is inserted as an `application/vnd.actiontext.opengraph-embed` attachment, so it's stored in the message body.
- It has special handling for twitter/x.com links via fxtwitter.
- It has tests: `test/models/opengraph/*`, `test/system/unfurling_links_test.rb`, and others.

**Why MetaInspector is a poor fit**
- **SSRF protection:** `Opengraph::Fetch` pins the connection to an IP that `RestrictedHTTP::PrivateNetworkGuard` (surfguard) has already checked. It re-checks every redirect hop, limits the body to 5 MB, and only accepts `text/html`. MetaInspector fetches through its own Faraday stack, so none of that applies. I believe you could get the guard back with a custom Faraday adapter, but that would be more code than the current fetcher.
- **Duplication:** It would do the same job as `Opengraph::Document` and `Opengraph::Fetch`. You'd end up either keeping both or rewriting a working, tested feature.

**Options**
1. **Do nothing.** Previews already work, so if they're broken somewhere, tell me where and I'll fix that.
2. **Swap the parsing to MetaInspector.** I'd keep `Opengraph::Fetch` as the guarded fetcher and pass the HTML to `MetaInspector.new(url, document: html)`. That gets you its fallbacks, such as `<title>`, meta description and favicon, without losing the SSRF protection.
3. **Replace everything with MetaInspector.** I'd do this only if you accept the SSRF trade-off, or want me to build a guarded Faraday adapter.

Which one do you want? Or did you mean something different, such as previews rendered server-side for messages that don't have an embed yet, or for messages posted by bots or the API?

