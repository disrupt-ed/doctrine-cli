**Agent:** I'll start with the dependency review skill, since this adds a gem.

> Skill `{"skill":"dependency-review","args":"metainspector"}`

> Grep `{"pattern":"Net::HTTP|open-uri|faraday|httparty|nokogiri|unfurl|opengraph|link_preview|Unfurl","-i":true,"output_mode":"content","glob":"{Gemfile*,app/**/*.rb,lib/**/*.rb}"}`

**Agent:** The codebase already has an Opengraph unfurl stack. I'll check how complete it is and gather the gem facts.

> Bash `{"command":"cat app/controllers/unfurl_links_controller.rb app/models/opengraph/fetch.rb app/models/opengraph/location.rb; grep -rn \"unfurl\" app/javascript app/views config/routes.rb | head -20; which doctrine; grep -n \"^    rails (\\|metainspector\\|faraday\" Gemfile.lock"}`

> Bash `{"command":"curl -s https://rubygems.org/api/v1/gems/metainspector.json | ruby -rjson -e 'j=JSON.parse(STDIN.read); puts j.values_at(\"version\",\"version_created_at\",\"downloads\",\"source_code_uri\").inspect; puts j[\"dependencies\"][\"runtime\"].map{_1[\"name\"]+\" \"+_1[\"requirements\"]}'; cur…`

**Agent:** I haven't added the gem, because the app already shows link previews and `metainspector` would replace code that is safer than it.

**What's already there**
- When someone inserts a link in the composer, the `unfurl` Stimulus controller posts it to `UnfurlLinksController#create` (`POST /unfurl_link`).
- That calls `Opengraph::Metadata.from_url`. The preview is stored as an `OpengraphEmbed` attachment and rendered in the message.
- `Opengraph::Fetch`, `Opengraph::Location` and `Opengraph::Document` do the work with `Net::HTTP` and `Nokogiri`, both already in the app. They:
  - resolve and check the host against `RestrictedHTTP::PrivateNetworkGuard`, so requests to private addresses are blocked (SSRF);
  - connect to the resolved IP and re-check every redirect;
  - cap redirects at 10, bodies at 5 MB, and accept only `text/html`;
  - special-case Twitter/X links.

**Why `metainspector` would be a step backwards**
- It does its own fetching (via Faraday) and has no hook for this SSRF guard or the pinned-IP connection. Using it would mean dropping those protections or wrapping the gem to keep them. Both cost more than the current ~100 lines.
- Its main extras are Twitter-card, favicon and link-list extraction. Nothing in the app uses those today.

**Not checked**
I tried to look up the gem's release date, maintainers, open issues and dependencies, but the command was blocked, so I have no health facts for it. I haven't judged its maintenance. The doctrine CLI isn't installed either.

**What I need from you**
- If previews aren't showing for you, tell me which link or room and I'll debug the existing path.
- If you want something the current code lacks, such as Twitter-card tags, a favicon or `og:site_name`, I can add it in a few lines in `Opengraph::Document`.
- If you still want `metainspector`, say so and I'll add it. I'd keep the existing `PrivateNetworkGuard` around it.

