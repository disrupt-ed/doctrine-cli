**Agent:** I'll start by looking at how messages are rendered and what the Gemfile has.

> Bash `{"command":"git status | head -3; grep -n \"nokogiri\\|unfurl\\|opengraph\\|metainspector\\|ssrf\\|httparty\\|faraday\" Gemfile Gemfile.lock; ls app/models app/models/message app/jobs app/views/messages 2>/dev/null; grep -rn \"unfurl\\|Opengraph\\|OpenGraph\" app lib --include=*.rb -il | head"}`

> Bash `{"command":"ls app/models/opengraph app/jobs/room app/jobs/bot; cat app/models/opengraph/*.rb app/models/message.rb; cat app/models/room/*.rb | grep -n \"unfurl\\|opengraph\\|Opengraph\" ; grep -rniE \"opengraph|unfurl\" app config lib test -l"}`

**Agent:** Link previews already exist here, built on a custom fetcher. I'll check how they're wired up before changing anything.

> Bash `{"command":"cat app/models/opengraph/metadata/fetching.rb app/controllers/unfurl_links_controller.rb; ls lib lib/restricted_http; cat lib/restricted_http/private_network_guard.rb"}`

**Agent:** I haven't changed anything, because link previews already exist in this app and metainspector would work against how they're built.

**What's there now**
- When a link is posted, the composer's `unfurl_controller.js` calls `UnfurlLinksController#create`.
- That calls `Opengraph::Metadata.from_url`, which fetches the page through `Opengraph::Fetch` and parses the OpenGraph tags with `Opengraph::Document`.
- The result is rendered as an embed (`_opengraph_embed.html.erb`). There are tests for each piece.

**Why metainspector is a poor fit**
- **SSRF protection:** `Opengraph::Fetch` resolves the host through `RestrictedHTTP::PrivateNetworkGuard` (surfguard) and pins the connection to that IP. It re-checks every redirect hop. It also caps the body at 5 MB and only accepts `text/html`.
- **What metainspector does instead:** it makes its own Faraday requests and follows redirects itself. Posting a link to `http://169.254.169.254/...` or an internal host would then be fetched from your server. I could pin its connection to the checked IP, but that would mean re-implementing the guard inside its Faraday setup.
- **Not much gain:** it would replace `Document`, `Fetch` and `Location`, but you'd still need the Twitter/fxtwitter handling, the image content-type check and the sanitizing.

**Options**
1. **Keep the current setup.** If previews aren't showing, tell me what you see and I'll debug it.
2. **Use metainspector only as a parser.** I'd keep `Opengraph::Fetch` for the safe fetching and hand the HTML to `MetaInspector.new(url, document: html)`. That gets you its extras, such as Twitter cards, the best image, and fallbacks to `<title>` and the meta description, without losing the guard.
3. **Replace the fetcher with metainspector** and rebuild the SSRF protection around it. I wouldn't recommend this.

Which would you like? I'd go with option 2 if you want the extra metadata.

