> Grep `{"pattern":"WebhookURL|webhook_url","glob":"**/*.go","output_mode":"files_with_matches"}`

> Read `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-baseline-1444693949/internal/ui/form/integration.go"}`

> Read `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-baseline-1444693949/internal/ui/integration_update.go"}`

> Grep `{"pattern":"func IsValidURL|func ValidateURL|googlereader_missing_required_fields|func IsValidWebhook","output_mode":"content","glob":"!*.po","head_limit":30}`

> Grep `{"pattern":"func .*(URL|Url)\\w*\\(","path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-baseline-1444693949/internal/validator","output_mode":"content"}`

> Grep `{"pattern":"error.invalid_feed_url|error.invalid_url|error.feed_url_not_empty","path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-baseline-1444693949/internal/locale/translations/en_US.json"…`

> Grep `{"pattern":"IsValidURL|ValidateURL|url\\.Parse|ParseRequestURI","path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-baseline-1444693949/internal","glob":"{validator,ui/form,api}/*.go","output…`

> Grep `{"pattern":"googlereader_missing_required_fields","path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-baseline-1444693949/internal","glob":"*.go","output_mode":"content"}`

**Agent:** I'm looking for an existing URL-validation helper to reuse for the webhook check. Next I'll check how the feed validator does it.

> Grep `{"pattern":"ProxyURL|IsHTTPURL|func .*Scheme","path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-baseline-1444693949/internal/validator","output_mode":"content","glob":"!*_test.go"}`

> Grep `{"pattern":"func ","path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-baseline-1444693949/internal/urllib/url.go","output_mode":"content"}`

> Read `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-baseline-1444693949/internal/urllib/url.go","limit":85,"offset":45}`

> Bash `{"command":"grep -n \"error.invalid_feed_url\\|error.http_client\" internal/locale/translations/en_US.json internal/locale/translations/fr_FR.json | head; grep -rn \"en_US\\|missing\" internal/locale/*_test.go | head"}`

> Bash `{"command":"sed -n 95,135p internal/locale/catalog_test.go; sed -n 128,132p internal/locale/translations/en_US.json; ls internal/ui/form/*_test.go"}`

> Bash `{"command":"grep -h -A1 '\"error.invalid_' internal/locale/translations/en_US.json | tail -6; grep -c 'error.invalid_language\"' internal/locale/translations/*.json | grep -v ':1'"}`

**Agent:** I'm adding an `error.invalid_webhook_url` key to every locale file. The catalog test requires every key to exist in all languages. After that I'll add the validator and wire it into the handler.

> Bash `{"command":"cd /private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-baseline-1444693949/internal/locale/translations && for f in *.json; do\ncase $f in fr_FR.json) msg=\"URL de webhook non valide.\";;…`

**Agent:** Bash is denied in this mode, so I'll make the edits with the Edit tool. I'll start with the validator in the form.

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-baseline-1444693949/internal/ui/form/integration.go","old_string":"// NewIntegrationForm returns a new Integratio…`

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-baseline-1444693949/internal/ui/form/integration.go","old_string":"\t\"miniflux.app/v2/internal/model\"\n)","new_…`

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-baseline-1444693949/internal/urllib/url.go","old_string":"// resolveToAbsoluteURL resolves a relative URL using a…`

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-baseline-1444693949/internal/ui/integration_update.go","old_string":"\tif integrationForm.WebhookEnabled {\n\t\ti…`

**Agent:** Now the locale files. I'll read each one first, since Edit requires that.

> Glob `{"pattern":"internal/locale/translations/*.json","path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-baseline-1444693949"}`

> Grep `{"pattern":"\"error.invalid_timezone\"","path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-baseline-1444693949/internal/locale/translations","output_mode":"content"}`

> Read `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-baseline-1444693949/internal/locale/translations/en_US.json","offset":135,"limit":1}`

> Read `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-baseline-1444693949/internal/locale/translations/fr_FR.json","offset":133,"limit":1}`

> Read `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-baseline-1444693949/internal/locale/translations/pl_PL.json","offset":136,"limit":1}`

> Read `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-baseline-1444693949/internal/locale/translations/ro_RO.json","offset":136,"limit":1}`

> Read `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-baseline-1444693949/internal/locale/translations/uk_UA.json","offset":136,"limit":1}`

> Read `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-baseline-1444693949/internal/locale/translations/gl_ES.json","offset":135,"limit":1}`

> Read `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-baseline-1444693949/internal/locale/translations/zh_TW.json","offset":130,"limit":1}`

> Read `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-baseline-1444693949/internal/locale/translations/pt_BR.json","offset":133,"limit":1}`

> Read `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-baseline-1444693949/internal/locale/translations/de_DE.json","offset":133,"limit":1}`

> Read `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-baseline-1444693949/internal/locale/translations/nl_NL.json","offset":133,"limit":1}`

> Read `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-baseline-1444693949/internal/locale/translations/ru_RU.json","offset":136,"limit":1}`

> Read `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-baseline-1444693949/internal/locale/translations/hi_IN.json","offset":133,"limit":1}`

> Read `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-baseline-1444693949/internal/locale/translations/ko_KR.json","offset":130,"limit":1}`

> Read `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-baseline-1444693949/internal/locale/translations/nan_Latn_pehoeji.json","offset":130,"limit":1}`

> Read `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-baseline-1444693949/internal/locale/translations/fi_FI.json","offset":133,"limit":1}`

> Read `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-baseline-1444693949/internal/locale/translations/es_ES.json","offset":133,"limit":1}`

> Read `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-baseline-1444693949/internal/locale/translations/it_IT.json","offset":133,"limit":1}`

> Read `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-baseline-1444693949/internal/locale/translations/ja_JP.json","offset":130,"limit":1}`

> Read `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-baseline-1444693949/internal/locale/translations/id_ID.json","offset":130,"limit":1}`

> Read `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-baseline-1444693949/internal/locale/translations/zh_CN.json","offset":130,"limit":1}`

> Read `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-baseline-1444693949/internal/locale/translations/ar_SA.json","offset":147,"limit":1}`

> Read `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-baseline-1444693949/internal/locale/translations/tr_TR.json","offset":133,"limit":1}`

> Read `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-baseline-1444693949/internal/locale/translations/el_EL.json","offset":133,"limit":1}`

**Agent:** Now the edits for all 23 locale files.

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-baseline-1444693949/internal/locale/translations/en_US.json","old_string":"    \"error.invalid_timezone\": \"Inva…`

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-baseline-1444693949/internal/locale/translations/fr_FR.json","old_string":"    \"error.invalid_timezone\": \"Fuse…`

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-baseline-1444693949/internal/locale/translations/pl_PL.json","old_string":"    \"error.invalid_timezone\": \"Niep…`

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-baseline-1444693949/internal/locale/translations/ro_RO.json","old_string":"    \"error.invalid_timezone\": \"Dat�…`

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-baseline-1444693949/internal/locale/translations/uk_UA.json","old_string":"    \"error.invalid_timezone\": \"Не…`

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-baseline-1444693949/internal/locale/translations/gl_ES.json","old_string":"    \"error.invalid_timezone\": \"Zona…`

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-baseline-1444693949/internal/locale/translations/zh_TW.json","old_string":"    \"error.invalid_timezone\": \"無�…`

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-baseline-1444693949/internal/locale/translations/pt_BR.json","old_string":"    \"error.invalid_timezone\": \"Fuso…`

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-baseline-1444693949/internal/locale/translations/de_DE.json","old_string":"    \"error.invalid_timezone\": \"Ung�…`

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-baseline-1444693949/internal/locale/translations/nl_NL.json","old_string":"    \"error.invalid_timezone\": \"Onge…`

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-baseline-1444693949/internal/locale/translations/ru_RU.json","old_string":"    \"error.invalid_timezone\": \"Не…`

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-baseline-1444693949/internal/locale/translations/hi_IN.json","old_string":"    \"error.invalid_timezone\": \"अ�…`

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-baseline-1444693949/internal/locale/translations/ko_KR.json","old_string":"    \"error.invalid_timezone\": \"시�…`

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-baseline-1444693949/internal/locale/translations/nan_Latn_pehoeji.json","old_string":"    \"error.invalid_timezon…`

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-baseline-1444693949/internal/locale/translations/fi_FI.json","old_string":"    \"error.invalid_timezone\": \"Virh…`

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-baseline-1444693949/internal/locale/translations/es_ES.json","old_string":"    \"error.invalid_timezone\": \"Zona…`

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-baseline-1444693949/internal/locale/translations/it_IT.json","old_string":"    \"error.invalid_timezone\": \"Fuso…`

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-baseline-1444693949/internal/locale/translations/ja_JP.json","old_string":"    \"error.invalid_timezone\": \"タ�…`

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-baseline-1444693949/internal/locale/translations/id_ID.json","old_string":"    \"error.invalid_timezone\": \"Zona…`

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-baseline-1444693949/internal/locale/translations/zh_CN.json","old_string":"    \"error.invalid_timezone\": \"无�…`

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-baseline-1444693949/internal/locale/translations/ar_SA.json","old_string":"    \"error.invalid_timezone\": \"ال…`

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-baseline-1444693949/internal/locale/translations/tr_TR.json","old_string":"    \"error.invalid_timezone\": \"Geç…`

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-baseline-1444693949/internal/locale/translations/el_EL.json","old_string":"    \"error.invalid_timezone\": \"Μη…`

**Agent:** Webhook URLs are now validated when integration settings are saved, but I haven't compiled or tested any of it. Bash was denied in this session, so `go build` and `go test` never ran.

- **Validation rule:** `urllib.IsValidWebhookURL` in `internal/urllib/url.go` accepts only absolute `http://` or `https://` URLs that have a host.
- **Form check:** `IntegrationForm.ValidateWebhook()` in `internal/ui/form/integration.go` applies that rule only when the webhook is enabled and a URL was entered. A blank URL still disables the webhook, as before.
- **Handler:** `updateIntegration` in `internal/ui/integration_update.go` now runs the check before the webhook handling. A bad URL shows an error message, redirects to `/integrations`, and saves nothing.
- **Error text:** I added `error.invalid_webhook_url` to all 23 locale files, because `TestMissingTranslations` requires every key in every language. `en_US` and `fr_FR` are translated. The other 21 files have the English text as a placeholder, so they need real translations.

I didn't add unit tests for `IsValidWebhookURL` or `ValidateWebhook`. Please run `go build ./... && go test ./internal/urllib ./internal/ui/... ./internal/locale`, or allow Bash and I'll run it.

