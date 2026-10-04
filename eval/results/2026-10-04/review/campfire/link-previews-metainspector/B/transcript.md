**Agent:** I'll start with the dependency review, since you named a specific gem.

> Skill `{"skill":"dependency-review","args":"metainspector"}`

> Grep `{"pattern":"unfurl|link_preview|Net::HTTP|faraday|httparty|nokogiri|ssrf|private_ip|opengraph|Opengraph","-i":true,"glob":"{Gemfile,Gemfile.lock,app/**/*.rb,lib/**/*.rb}","output_mode":"content","head_limit":60}`

> Bash `{"command":"which [redacted]; grep -n metainspector Gemfile Gemfile.lock; ls app/models/opengraph lib/restricted_http; grep -rn \"unfurl\" config/routes.rb app/javascript -il | head; curl -s https://rubygems.org/api/v1/gems/metainspector.json | ruby -rjson -e 'j=JSON.parse(STDIN.read); puts j.values_a…`

> Read `{"file_path":"[redacted]"}`

> Read `{"file_path":"[redacted]"}`

**Agent:** I haven't added metainspector. The app already has link previews, and I think the gem would be a step backwards.

**What's already there**
- `UnfurlLinksController#create` calls `Opengraph::Metadata.from_url`, which fetches the page, parses the OpenGraph tags and returns JSON.
- `lib/rails_ext/actiontext_opengraph_embeds.rb` stores the preview as an Action Text attachment in the message.
- `ContentFilters::RemoveSoloUnfurledLinkText` hides the raw URL when the preview is the whole message.
- `Opengraph::Fetch` is a hardened fetcher built on `Net::HTTP`. It has:
  - an SSRF guard (`RestrictedHTTP::PrivateNetworkGuard`, backed by the `surfguard` gem);
  - IP pinning, so the address that was checked is the one connected to;
  - re-validation of each redirect, capped at 10;
  - a 5 MB body cap and a `text/html`-only check.
- It also has a twitter/x fallback through fxtwitter.

**Why metainspector would hurt**
- It would duplicate the whole feature: a second way to fetch and parse metadata, next to the existing one.
- It does its own HTTP fetching (via faraday), so it would bypass the SSRF guard unless you hand it pre-fetched HTML. Previews are fetched from user-supplied URLs, so I'd treat that as a security regression.
- The existing code is a handful of small files, which is cheap to own.

**What I couldn't check**
Shell access was denied in this session. So I haven't looked up metainspector's last release, maintainers, open issues or Rails/Ruby support, and I haven't confirmed it's absent from the Gemfile. The grep found no mention of it.

**Recommendation**
If something specific about the current previews is lacking, I'd extend `Opengraph::Document` or `Opengraph::Metadata` for that. Examples would be missing fields, a particular site that renders badly, or the Twitter fallback. What's the problem you're seeing?

If you still want metainspector, say so and I'll add it. I'd feed it HTML fetched through `Opengraph::Fetch` so the SSRF protections stay in place.

