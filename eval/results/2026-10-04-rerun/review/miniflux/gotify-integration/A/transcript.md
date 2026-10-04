**Agent:** I'll start by looking at how existing integrations are wired in.

> Grep `{"pattern":"(?i)ntfy","output_mode":"count","glob":"!*.po"}`

> Read `{"file_path":"[redacted]"}`

> Grep `{"pattern":"(?i)ntfy","path":"[redacted]","output_mode":"content","glob":"{internal/integration/integration.go,internal/model/integration.go,int…`

> Read `{"file_path":"[redacted]","offset":535,"limit":80}`

> Bash `{"command":"tail -30 internal/database/migrations.go && grep -n \"ntfy\\|Ntfy\" internal/model/feed.go internal/storage/feed.go internal/storage/feed_query_builder.go internal/ui/form/feed.go internal/ui/feed_edit.go; sed -n 255,300p internal/storage/integration.go; sed -n 455,470p internal/storage/…`

> Bash `{"command":"F=internal/storage/integration.go; grep -n \"archiveorg\\|Archiveorg\\|WHERE\\|user_id=\\$\" $F; sed -n 255,262p $F >/dev/null; sed -n 236,250p $F; sed -n 345,380p $F; sed -n 500,520p $F; sed -n 625,700p $F; grep -n -i archiveorg internal/model/integration.go internal/ui/form/integration…`

> Bash `{"command":"sed -n 228,236p internal/storage/integration.go; sed -n 360,368p internal/storage/integration.go; sed -n 494,504p internal/storage/integration.go; sed -n 126,140p internal/model/integration.go; sed -n 255,265p internal/ui/form/integration.go; sed -n 390,400p internal/ui/form/integration.…`

> Bash `{"command":"grep -n \"ntfy_url\\|archiveorg_activate\" internal/locale/translations/fr_FR.json; grep -n \"func Test\" internal/locale/*_test.go | head; sed -n 285,292p internal/locale/translations/en_US.json; grep -n \"Pushover\\|pushover\" internal/integration/integration.go | head -3; sed -n 1,15p…`

**Agent:** I've mapped how ntfy is wired in. I'm now adding Gotify as a user-level integration covering the migration, model, storage, form, template, client and locales.

> Bash `{"command":"python3 - <<'EOF'\nimport re,glob\ndef sub(path, old, new, count=1):\n    s=open(path).read()\n    assert old in s, (path, old)\n    s=s.replace(old,new,count)\n    open(path,'w').write(s)\n\n# migration\nsub('internal/database/migrations.go', '''\t\t\t\tON enclosures (user_id, entry_id,…`

**Agent:** The scripted edit was denied, so nothing was applied. I'll make the same changes with the Edit and Write tools, starting with the Gotify client.

> Write `{"file_path":"[redacted]","content":"// SPDX-FileCopyrightText: Copyright The Miniflux Authors. All rights…`

> Read `{"file_path":"[redacted]","offset":125,"limit":10}`

> Read `{"file_path":"[redacted]","offset":130,"limit":10}`

> Read `{"file_path":"[redacted]","offset":262,"limit":50}`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"\tArchiveorgEnabled                bool\n}","new_strin…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"\tArchiveorgEnabled                bool\n}","new_str…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"\tintegration.ArchiveorgEnabled = i.ArchiveorgEnable…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"func NewIntegrationForm(r *http.Request) *Integratio…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"\t\tArchiveorgEnabled:                r.FormValue(\"…`

> Read `{"file_path":"[redacted]","limit":20}`

**Agent:** `strconv` is already imported. Next up are the show handler, the storage layer, the migration, the dispatcher and the template.

> Read `{"file_path":"[redacted]","offset":144,"limit":6}`

> Read `{"file_path":"[redacted]","offset":228,"limit":8}`

> Read `{"file_path":"[redacted]","offset":1100,"limit":5}`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"\t\tArchiveorgEnabled:                integration.Ar…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"\t\t\tarchiveorg_enabled\n\t\tFROM","new_string":"\t…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"\t\t&integration.ArchiveorgEnabled,\n\t)","new_strin…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"\t\t\treadeck_push_enabled=$121\n\t\tWHERE\n\t\t\tus…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"\t\tintegration.ReadeckPushEnabled,\n\t\tintegration…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"\t\t\t\tON enclosures (user_id, entry_id, sha256(url…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"\tif userIntegrations.AppriseEnabled {\n\t\tslog…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"\t\"miniflux.app/v2/internal/integration/espial\…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"    <details {{ if .form.NunuxKe…`

> Read `{"file_path":"[redacted]","offset":70,"limit":60}`

> Grep `{"pattern":"form.integration.instapaper_activate","path":"[redacted]","output_mode":"count"}`

**Agent:** The test requires every locale to have the new keys, so I'm adding them to all 23 files. I'll read each file first, since Edit needs that.

> Read `{"file_path":"[redacted]","limit":1}`

> Read `{"file_path":"[redacted]","limit":1}`

> Read `{"file_path":"[redacted]","limit":1}`

> Read `{"file_path":"[redacted]","limit":1}`

> Read `{"file_path":"[redacted]","limit":1}`

> Read `{"file_path":"[redacted]","limit":1}`

> Read `{"file_path":"[redacted]","limit":1}`

> Read `{"file_path":"[redacted]","limit":1}`

> Read `{"file_path":"[redacted]","limit":1}`

> Read `{"file_path":"[redacted]","limit":1}`

> Read `{"file_path":"[redacted]","limit":1}`

> Read `{"file_path":"[redacted]","limit":1}`

> Read `{"file_path":"[redacted]","limit":1}`

> Read `{"file_path":"[redacted]","limit":1}`

> Read `{"file_path":"[redacted]","limit":1}`

> Read `{"file_path":"[redacted]","limit":1}`

> Read `{"file_path":"[redacted]","limit":1}`

> Read `{"file_path":"[redacted]","limit":1}`

> Read `{"file_path":"[redacted]","limit":1}`

> Read `{"file_path":"[redacted]","limit":1}`

> Read `{"file_path":"[redacted]","limit":1}`

> Read `{"file_path":"[redacted]","limit":1}`

> Read `{"file_path":"[redacted]","limit":1}`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"    \"form.integration.instapaper_activate\"…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"    \"form.integration.instapaper_activate\"…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"    \"form.integration.instapaper_activate\"…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"    \"form.integration.instapaper_activate\"…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"    \"form.integration.instapaper_activate\"…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"    \"form.integration.instapaper_activate\"…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"    \"form.integration.instapaper_activate\"…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"    \"form.integration.instapaper_activate\"…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"    \"form.integration.instapaper_activate\"…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"    \"form.integration.instapaper_activate\"…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"    \"form.integration.instapaper…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"    \"form.integration.instapaper_activate\"…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"    \"form.integration.instapaper_activate\"…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"    \"form.integration.instapaper_activate\"…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"    \"form.integration.instapaper_activate\"…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"    \"form.integration.instapaper_activate\"…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"    \"form.integration.instapaper_activate\"…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"    \"form.integration.instapaper_activate\"…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"    \"form.integration.instapaper_activate\"…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"    \"form.integration.instapaper_activate\"…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"    \"form.integration.instapaper_activate\"…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"    \"form.integration.instapaper_activate\"…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"    \"form.integration.instapaper_activate\"…`

**Agent:** Now I'll build and run the tests that touch the changed packages.

> Bash `{"command":"gofmt -l internal && go build ./... && go vet ./internal/integration/... ./internal/ui/... ./internal/storage/... && go test ./internal/locale/... ./internal/ui/... ./internal/integration/... ./internal/template/... 2>&1 | tail -20","description":"Format check, build, vet and test change…`

> Bash `{"command":"go test ./internal/locale/... ./internal/ui/... ./internal/database/... 2>&1 | grep -v \"no test files\"; git status --short | head -40","description":"Run locale/ui/database tests and show changes"}`

**Agent:** I added a Gotify integration that sends a notification for each new entry. It builds, `gofmt` is clean, and the locale, form and template tests pass. I didn't test it against a real Gotify server or a database. There are no unit tests for the new client.

Settings are per user, on the Integrations page, with a Gotify section. They are a server URL, an application token and a priority from 0 to 10 (default 5).

- **Client:** it lives in `internal/integration/gotify/gotify.go` and posts to `<url>/message`. The token goes in an `X-Gotify-Key` header, so it stays out of URLs and logs. It uses the same private-network blocking as ntfy. The notification title is the feed title, the body is the entry title, and tapping it opens the entry URL.
- **Wiring:** a new migration adds four columns to `integrations`. The model, form, storage and integration dispatcher were updated to match.
- **Locales:** I added the four new keys to all locale files because `TestMissingTranslations` requires them. Only English and French are actually translated. The other 21 locales carry the English text.

Decisions for you:
- **Every feed notifies.** Unlike ntfy, Gotify has no per-feed toggle. Once it's enabled, every new entry from every feed triggers a push. A per-feed opt-in would also touch the feed model, storage, edit form, feed templates and API client. I left it out to keep this change small, but it's the obvious next step if that volume is too noisy.
- **README not updated.** I didn't add Gotify to the integrations list in the README, since I only saw that line truncated and didn't want to edit it blind. I can add it if you want.
- **Not committed.** The changes are in the working tree only.

One thing that went sideways: my first attempt used a Python script to batch the edits, and that was denied. I redid them with the regular edit tools, and nothing from the script was applied.

