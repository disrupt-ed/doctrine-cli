> Grep `{"pattern":"WebhookURL|webhook_url","glob":"*.go","output_mode":"content","head_limit":60}`

> Read `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-baseline-1497758937/internal/ui/integration_update.go"}`

> Grep `{"pattern":"func Validate|validator\\.|ValidateFeedURL|IsValidURL","path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-baseline-1497758937/internal","output_mode":"content","glob":"*.go","hea…`

> Grep `{"pattern":"func IsValid|func Validate","path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-baseline-1497758937/internal/validator/validator.go","output_mode":"content"}`

> Read `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-baseline-1497758937/internal/ui/form/integration.go","offset":1,"limit":20}`

> Grep `{"pattern":"ValidateGoogleReader","path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-baseline-1497758937/internal/ui/form/integration.go","output_mode":"content","-A":20}`

> Grep `{"pattern":"ValidateFeedURL|IsValidURL|func.*URL","path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-baseline-1497758937/internal/validator","output_mode":"content"}`

> Read `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-baseline-1497758937/internal/validator/feed.go"}`

> Grep `{"pattern":"googlereader_missing_required_fields|error.feed_invalid_proxy_url|error.invalid_feed_url","path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-baseline-1497758937/internal/locale/t…`

> Grep `{"pattern":"func IsAbsoluteURL","path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-baseline-1497758937/internal/urllib","output_mode":"content","-A":12}`

> Bash `{"command":"ls internal/locale/translations && grep -n \"googlereader_missing_required_fields\" internal/locale/translations/*.json | head -30"}`

**Agent:** Adding a webhook URL check to the integration save handler. I'm checking how new translation keys get added across locales first.

> Bash `{"command":"grep -n \"linktaco_missing_required_fields\" internal/locale/translations/*.json | head -30; sed -n 55,60p internal/urllib/url.go; sed -n 71,90p internal/urllib/url.go"}`

> Bash `{"command":"for f in internal/locale/translations/*.json; do\nsed -i '' '/\"error.invalid_feed_url\":/{\np\ns/\"error.invalid_feed_url\": .*/\"error.invalid_webhook_url\": \"Invalid webhook URL. It must be an absolute HTTP or HTTPS URL.\",/\n}' $f; done; grep -n \"invalid_webhook_url\" internal/loca…`

> Grep `{"pattern":"en_US|defaultLanguage|fallback","path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-baseline-1497758937/internal/locale","glob":"*.go","output_mode":"content"}`

> Grep `{"pattern":"\"error.invalid_site_url\"|\"error.invalid_feed_url\"","path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-baseline-1497758937/internal/locale/translations","output_mode":"content…`

> Grep `{"pattern":"\"error.invalid_theme\"","path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-baseline-1497758937/internal/locale/translations","output_mode":"content","-A":1,"glob":"{en_US,fr_FR,…`

> Grep `{"pattern":"\"error.invalid_timezone\"","path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-baseline-1497758937/internal/locale/translations","output_mode":"content","-A":1,"glob":"en_US.json…`

**Agent:** Next key is `error.local_auth_disabled`, so I'll insert before it in all 23 locale files, then add the validation code.

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-baseline-1497758937/internal/locale/translations/en_US.json","old_string":"    \"error.local_auth_disabled\":","n…`

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-baseline-1497758937/internal/ui/form/integration.go","old_string":"// NewIntegrationForm returns","new_string":"/…`

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-baseline-1497758937/internal/ui/form/integration.go","old_string":"\t\"net/http\"\n\t\"strconv\"","new_string":"\…`

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-baseline-1497758937/internal/ui/integration_update.go","old_string":"\tif integrationForm.WebhookEnabled {\n\t\ti…`

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-baseline-1497758937/internal/locale/translations/fr_FR.json","old_string":"    \"error.local_auth_disabled\":","n…`

> Grep `{"pattern":"\"error.local_auth_disabled\"","path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-baseline-1497758937/internal/locale/translations","output_mode":"count"}`

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-baseline-1497758937/internal/locale/translations/pl_PL.json","old_string":"    \"error.local_auth_disabled\":","n…`

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-baseline-1497758937/internal/locale/translations/ro_RO.json","old_string":"    \"error.local_auth_disabled\":","n…`

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-baseline-1497758937/internal/locale/translations/uk_UA.json","old_string":"    \"error.local_auth_disabled\":","n…`

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-baseline-1497758937/internal/locale/translations/gl_ES.json","old_string":"    \"error.local_auth_disabled\":","n…`

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-baseline-1497758937/internal/locale/translations/zh_TW.json","old_string":"    \"error.local_auth_disabled\":","n…`

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-baseline-1497758937/internal/locale/translations/pt_BR.json","old_string":"    \"error.local_auth_disabled\":","n…`

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-baseline-1497758937/internal/locale/translations/de_DE.json","old_string":"    \"error.local_auth_disabled\":","n…`

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-baseline-1497758937/internal/locale/translations/nl_NL.json","old_string":"    \"error.local_auth_disabled\":","n…`

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-baseline-1497758937/internal/locale/translations/fi_FI.json","old_string":"    \"error.local_auth_disabled\":","n…`

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-baseline-1497758937/internal/locale/translations/ja_JP.json","old_string":"    \"error.local_auth_disabled\":","n…`

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-baseline-1497758937/internal/locale/translations/hi_IN.json","old_string":"    \"error.local_auth_disabled\":","n…`

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-baseline-1497758937/internal/locale/translations/ko_KR.json","old_string":"    \"error.local_auth_disabled\":","n…`

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-baseline-1497758937/internal/locale/translations/nan_Latn_pehoeji.json","old_string":"    \"error.local_auth_disa…`

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-baseline-1497758937/internal/locale/translations/zh_CN.json","old_string":"    \"error.local_auth_disabled\":","n…`

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-baseline-1497758937/internal/locale/translations/es_ES.json","old_string":"    \"error.local_auth_disabled\":","n…`

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-baseline-1497758937/internal/locale/translations/el_EL.json","old_string":"    \"error.local_auth_disabled\":","n…`

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-baseline-1497758937/internal/locale/translations/tr_TR.json","old_string":"    \"error.local_auth_disabled\":","n…`

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-baseline-1497758937/internal/locale/translations/it_IT.json","old_string":"    \"error.local_auth_disabled\":","n…`

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-baseline-1497758937/internal/locale/translations/ar_SA.json","old_string":"    \"error.local_auth_disabled\":","n…`

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-baseline-1497758937/internal/locale/translations/id_ID.json","old_string":"    \"error.local_auth_disabled\":","n…`

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-baseline-1497758937/internal/locale/translations/ru_RU.json","old_string":"    \"error.local_auth_disabled\":","n…`

**Agent:** Saving integration settings now rejects a webhook URL that isn't an absolute `http`/`https` URL with a host. I haven't compiled or tested any of it, because Bash was denied in this session.

- **Validation:** `ValidateWebhook()` in `internal/ui/form/integration.go` parses the URL and returns `error.invalid_webhook_url` if it fails. It only checks when the webhook integration is enabled and a URL is entered, so an empty URL still disables the integration as before.
- **Handler:** `updateIntegration` in `internal/ui/integration_update.go` calls it before touching the webhook settings. On failure it shows the error and redirects to `/integrations` without saving, like the Google Reader and LinkTaco checks.
- **Translations:** I added `error.invalid_webhook_url` to all 23 locale files, placed alphabetically. Only English and French are translated. The other 21 contain the English text as a placeholder.

I didn't add a test. The per-feed webhook URL (feed edit) is not validated either, and I left it alone since you asked about integration settings.

Please run `go build ./... && go test ./internal/ui/... ./internal/locale/...` to check it.

