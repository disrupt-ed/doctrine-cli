**Agent:** I'll start by looking at how existing integrations are structured.

> Bash `{"command":"ls internal/integration && grep -rli ntfy --include=* . 2>/dev/null | grep -v '^./.git/' | head -50; head -5 go.mod"}`

> Bash `{"command":"grep -rli ntfy . --exclude-dir=.git --exclude-dir=node_modules | head -50; cat internal/integration/ntfy/ntfy.go"}`

> Bash `{"command":"grep -rniE 'ntfy' internal client README.md --include=*.go --include=*.html --include=*.md --exclude-dir=ntfy | grep -v translations/ ; grep -n ntfy internal/locale/translations/en_US.json internal/locale/translations/fr_FR.json; grep -n -i ntfy internal/database/migrations.go | head; ls…`

> Bash `{"command":"grep -n -i -B3 -A25 'func PushEntries' internal/integration/integration.go; grep -rniE 'pushover' internal client README.md --exclude-dir=translations --exclude-dir=pushover | grep -v '^internal/integration/integration.go'; tail -30 internal/database/migrations.go; ls internal/database; …`

> Bash `{"command":"grep -n -B2 -A28 'feed.PushoverEnabled' internal/integration/integration.go; grep -n -B3 -A25 'feed.NtfyEnabled' internal/integration/integration.go | head -5; cat internal/integration/pushover/pushover.go; grep -n -i 'pushover' -B2 -A2 internal/model/feed.go internal/ui/form/feed.go | h…`

> Bash `{"command":"sed -n 115,175p internal/model/integration.go; sed -n 205,260p internal/storage/integration.go; sed -n 335,420p internal/storage/integration.go; sed -n 470,520p internal/storage/integration.go; sed -n 600,700p internal/storage/integration.go"}`

> Bash `{"command":"sed -n 360,440p internal/storage/feed.go; sed -n 170,200p internal/storage/feed_query_builder.go; sed -n 255,280p internal/storage/feed_query_builder.go; sed -n 560,600p internal/integration/integration.go"}`

**Agent:** Following the ntfy/pushover pattern: per-user Gotify settings (URL, token, priority) plus a per-feed toggle. Now wiring it up.

> Bash `{"command":"python3 - <<'EOF'\nimport re\ndef sub(path, old, new, count=1):\n    s=open(path).read()\n    assert s.count(old)>=1, (path, old)\n    s=s.replace(old,new,count)\n    open(path,'w').write(s)\n\n# migration\nsub('internal/database/migrations.go', \"\t\t\t\tON enclosures (user_id, entry_id…`

**Agent:** Bash is now denied, so I'll make the same edits with the Edit tool. I can't run the build or tests this way, and I'll say so at the end.

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"\t\t\t\tON enclosures (user_id, entry_id, sha256(u…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"\tArchiveorgEnabled                bool\n}","new_str…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"\tPushoverEnabled             bool      `json:\"pushover_en…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"\tPushoverEnabled             bool      `json:\"pushover_enabled\"…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"\t\t\tarchiveorg_enabled\n\t\tFROM","new_string":"…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"\t\t&integration.ArchiveorgEnabled,\n\t)","new_str…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"\t\t\treadeck_push_enabled=$121\n\t\tWHERE\n\t\t\t…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"\t\tintegration.ReadeckPushEnabled,\n\t\tintegrati…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"\t\t\tlanguage=$40\n\t\tWHERE\n\t\t\tid=$41 AND user_id=$…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"\t\tfeed.Language,\n\t\tfeed.ID,\n\t\tfeed.UserID,","new_…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"\t\t\tf.ignore_entry_updates\n\t\tFROM","ne…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"\t\t\t&feed.IgnoreEntryUpdates,\n\t\t)","ne…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"\tPushoverEnabled  bool\n","new_string":"\tPushoverEnable…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"\tfeed.PushoverEnabled = f.PushoverEnabled\n","new_string…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"\t\tPushoverEnabled:             r.FormValue(\"pushover_e…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"\t\tPushoverEnabled:             feed.PushoverEnabled,\n"…`

> Grep `{"pattern":"Archiveorg|Priority|strconv|ntfy","path":"[redacted]","glob":"{form/integration.go,integration_show.go}","output_mode"…`

> Read `{"file_path":"[redacted]","offset":255,"limit":160}`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"\t\tArchiveorgEnabled:                r.FormValue(…`

**Agent:** I mangled the indentation; redoing the edits precisely.

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"\t\tArchiveorgEnabled:                r.FormValue(…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"func NewIntegrationForm(r *http.Request) *Integrat…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"\tintegration.ArchiveorgEnabled = i.ArchiveorgEnab…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"\tArchiveorgEnabled                bool\n","new_st…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"\t\tArchiveorgEnabled:                integration.…`

> Write `{"file_path":"[redacted]","content":"// SPDX-FileCopyrightText: Copyright The Miniflux Authors. All righ…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"\tif userIntegrations.PushoverEnabled && feed.…`

> Grep `{"pattern":"integration/(espial|karakeep)\"","path":"[redacted]","output_mode":"content"}`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"\t\"miniflux.app/v2/internal/integration/karak…`

> Read `{"file_path":"[redacted]","offset":225,"limit":40}`

> Read `{"file_path":"[redacted]","offset":236,"limit":36}`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"            <details {{ if .form.…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"    <details {{ if .form.Pusho…`

> Grep `{"pattern":"form.feed.label.ntfy_activate|form.integration.ntfy_activate","path":"[redacted]","glob":"*.go","output_mode":"con…`

> Grep `{"pattern":"en_US|fallback|missing","path":"[redacted]","glob":"*.go","output_mode":"content","head_limit":20}`

> Grep `{"pattern":"func Test","path":"[redacted]","output_mode":"content"}`

> Read `{"file_path":"[redacted]","offset":180,"limit":12}`

> Read `{"file_path":"[redacted]","offset":255,"limit":12}`

> Read `{"file_path":"[redacted]","offset":70,"limit":60}`

**Agent:** Continuing: the locale test requires every key in all 27 locale files, so I'm adding the new keys to each one. First I'll find the sorted insertion points.

> Grep `{"pattern":"\"form\\.integration\\.(googlereader_activate|fever_[a-z_]+|karakeep_activate|linkace_activate|instapaper_activate)\"|\"form\\.feed\\.label\\.hide_globally\"","path":"[redacted]…`

> Glob `{"pattern":"*.json","path":"[redacted]"}`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"    \"form.feed.label.hide_globally\":","n…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"    \"form.integration.googlereader_activa…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"    \"form.feed.label.hide_globally\":","n…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"    \"form.integration.googlereader_activa…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"    \"form.feed.label.hide_globally\":","n…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"    \"form.integration.googlereader_activa…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"    \"form.feed.label.hide_globally\":","n…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"    \"form.integration.googlereader_activa…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"    \"form.feed.label.hide_globally\":","n…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"    \"form.integration.googlereader_activa…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"    \"form.feed.label.hide_globally\":","n…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"    \"form.integration.googlereader_activa…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"    \"form.feed.label.hide_globally\":","n…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"    \"form.integration.googlereader_activa…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"    \"form.feed.label.hide_globally\":","n…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"    \"form.integration.googlereader_activa…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"    \"form.feed.label.hide_globally\":","n…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"    \"form.integration.googlereader_activa…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"    \"form.feed.label.hide_globally\":","n…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"    \"form.integration.googlereader_activa…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"    \"form.feed.label.hide_globally\":","n…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"    \"form.integration.googlereader_activa…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"    \"form.feed.label.hide_globally\":","n…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"    \"form.integration.googlereader_activa…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"    \"form.feed.label.hide_globally\":","n…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"    \"form.integration.googlereader_activa…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"    \"form.feed.label.hide_glob…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"    \"form.integration.googlere…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"    \"form.feed.label.hide_globally\":","n…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"    \"form.integration.googlereader_activa…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"    \"form.feed.label.hide_globally\":","n…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"    \"form.integration.googlereader_activa…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"    \"form.feed.label.hide_globally\":","n…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"    \"form.integration.googlereader_activa…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"    \"form.feed.label.hide_globally\":","n…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"    \"form.integration.googlereader_activa…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"    \"form.feed.label.hide_globally\":","n…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"    \"form.integration.googlereader_activa…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"    \"form.feed.label.hide_globally\":","n…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"    \"form.integration.googlereader_activa…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"    \"form.feed.label.hide_globally\":","n…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"    \"form.integration.googlereader_activa…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"    \"form.feed.label.hide_globally\":","n…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"    \"form.integration.googlereader_activa…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"    \"form.feed.label.hide_globally\":","n…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"    \"form.integration.googlereader_activa…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"[Espial](https://github.com/jonschoning/espial), [Instapaper]","new_stri…`

**Agent:** The Gotify integration is written but unverified. Bash was denied partway through, so I couldn't run `go build`, `go vet`, `gofmt` or the tests.

**What it does**
- **Settings:** on the Integrations page, each user sets the Gotify server URL, an application token and a default priority (0–10, default 5).
- **Per-feed switch:** each feed gets a "Push entries to Gotify" checkbox, the same way ntfy and Pushover work.
- **Notifications:** each new entry sends one notification, with the feed title as the title and the entry title as the message. Tapping it opens the entry URL (`extras["client::notification"]`).
- **Client:** `internal/integration/gotify/gotify.go` posts to `{url}/message` with the `X-Gotify-Key` header. It uses the existing HTTP client, so the private-network blocking applies. It's called from `PushEntries` in `internal/integration/integration.go`, and failures are logged as warnings without stopping the refresh.
- **Wiring:** the settings are carried through the models, storage, forms, templates and the API `Feed` model. A new migration adds `gotify_enabled`, `gotify_url`, `gotify_token` and `gotify_priority` to `integrations`, and `gotify_enabled` to `feeds`. I also added Gotify to the README integration list.

**Things to check**
- **Build and tests:** please run `go build ./... && go vet ./... && go test ./internal/...`. The SQL parameter renumbering in `UpdateIntegration` and `UpdateFeed`, and the `Scan` column order, are the most likely places for a mistake.
- **Translations:** the locale test requires every key in every language file. I translated the five new strings into French and used the English text in the other 21 languages, so those need real translations.
- **No client test:** the HTTP client reads `config.Opts` and blocks private addresses, so an `httptest` server on localhost would need extra setup. I skipped it rather than add something I couldn't run.
- **Priority is per user, not per feed.** That keeps the feed change down to the one checkbox. Say if you want it per feed like ntfy.

