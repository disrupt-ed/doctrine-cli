> Grep `{"pattern":"WebhookURL|webhook_url","glob":"*.go","output_mode":"content","head_limit":60}`

> Read `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-doctrine-1525031129/internal/ui/form/integration.go","offset":225,"limit":140}`

> Read `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-doctrine-1525031129/internal/ui/integration_update.go"}`

> Grep `{"pattern":"func ValidateURL|func IsValidURL|ValidateFeedURL|func Validate\\w*URL","path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-doctrine-1525031129/internal","output_mode":"content"}`

> Grep `{"pattern":"error\\.(feed_url|invalid_feed_url|invalid_url)|error.googlereader_missing","path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-doctrine-1525031129/internal","output_mode":"conten…`

> Read `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-doctrine-1525031129/internal/validator/feed.go","limit":30}`

> Grep `{"pattern":"func IsValidURL|func (.*) ?IsValidURL","path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-doctrine-1525031129/internal","output_mode":"content"}`

> Read `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-doctrine-1525031129/internal/ui/form/subscription.go","limit":60}`

> Grep `{"pattern":"func ","path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-doctrine-1525031129/internal/urllib/url.go","output_mode":"content"}`

> Bash `{"command":"grep -n \"IsAbsoluteURL\\|IsValidProxyURL\" -A18 internal/urllib/url.go | head -70; grep -n \"error.googlereader_missing_required_fields\\|error.invalid_feed_url\" internal/locale/translations/en_US.json; ls internal/locale; grep -rn \"_test\" -l internal/ui/form | head"}`

**Agent:** Checking how translation keys are enforced, then I'll add the validation.

> Bash `{"command":"grep -n \"missing\\|fallback\\|en_US\" internal/locale/catalog_test.go internal/locale/printer.go | head -20; git log --oneline -3 -- internal/locale/translations/fr_FR.json; sed -n 160,172p internal/locale/translations/en_US.json; ls internal/ui/form/"}`

> Bash `{"command":"sed -n 95,130p internal/locale/catalog_test.go; grep -n \"Validate\" -B3 -A25 internal/ui/form/integration_test.go | head -60; head -20 internal/ui/form/integration.go"}`

> Bash `{"command":"python3 - <<'E'\nimport glob,re\nnew='    \"error.invalid_webhook_url\": \"Invalid webhook URL. It must start with http:// or https:// and include a host.\",\\n'\nfor f in glob.glob('internal/locale/translations/*.json'):\n    L=open(f,encoding='utf-8').read().split('\\n')\n    out=[];do…`

**Agent:** Bash is denied, so I'll make the edits with the file tools.

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-doctrine-1525031129/internal/ui/form/integration.go","old_string":"// NewIntegrationForm returns","new_string":"/…`

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-doctrine-1525031129/internal/ui/form/integration.go","old_string":"\t\"net/http\"\n\t\"strconv\"\n\n\t\"miniflux.…`

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-doctrine-1525031129/internal/ui/integration_update.go","old_string":"\tif integrationForm.WebhookEnabled {\n\t\ti…`

> Read `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-doctrine-1525031129/internal/ui/form/integration_test.go","offset":44}`

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-doctrine-1525031129/internal/ui/form/integration_test.go","old_string":"\t\t\t\tvalidationErr := scenario.form.Va…`

**Agent:** Indentation mismatch; I'll anchor on a smaller unique string.

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-doctrine-1525031129/internal/ui/form/integration_test.go","old_string":"\t\t\t\tt.Fatalf(\"unexpected validation …`

**Agent:** Adding the translation key to every locale (a test requires all locales to have it).

> Bash `{"command":"cd /tmp && echo noop"}`

> Read `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-doctrine-1525031129/internal/locale/translations/en_US.json","offset":128,"limit":6}`

> Read `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-doctrine-1525031129/internal/locale/translations/fr_FR.json","limit":1}`

> Read `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-doctrine-1525031129/internal/locale/translations/de_DE.json","limit":1}`

> Read `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-doctrine-1525031129/internal/locale/translations/es_ES.json","limit":1}`

> Read `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-doctrine-1525031129/internal/locale/translations/pl_PL.json","limit":1}`

> Read `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-doctrine-1525031129/internal/locale/translations/ro_RO.json","limit":1}`

> Read `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-doctrine-1525031129/internal/locale/translations/uk_UA.json","limit":1}`

> Read `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-doctrine-1525031129/internal/locale/translations/gl_ES.json","limit":1}`

> Read `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-doctrine-1525031129/internal/locale/translations/zh_TW.json","limit":1}`

> Read `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-doctrine-1525031129/internal/locale/translations/pt_BR.json","limit":1}`

> Read `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-doctrine-1525031129/internal/locale/translations/nl_NL.json","limit":1}`

> Read `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-doctrine-1525031129/internal/locale/translations/ru_RU.json","limit":1}`

> Read `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-doctrine-1525031129/internal/locale/translations/hi_IN.json","limit":1}`

> Read `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-doctrine-1525031129/internal/locale/translations/fi_FI.json","limit":1}`

> Read `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-doctrine-1525031129/internal/locale/translations/nan_Latn_pehoeji.json","limit":1}`

> Read `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-doctrine-1525031129/internal/locale/translations/it_IT.json","limit":1}`

> Read `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-doctrine-1525031129/internal/locale/translations/zh_CN.json","limit":1}`

> Read `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-doctrine-1525031129/internal/locale/translations/ja_JP.json","limit":1}`

> Read `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-doctrine-1525031129/internal/locale/translations/ko_KR.json","limit":1}`

> Read `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-doctrine-1525031129/internal/locale/translations/el_EL.json","limit":1}`

> Read `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-doctrine-1525031129/internal/locale/translations/tr_TR.json","limit":1}`

> Read `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-doctrine-1525031129/internal/locale/translations/ar_SA.json","limit":1}`

> Read `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-doctrine-1525031129/internal/locale/translations/id_ID.json","limit":1}`

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-doctrine-1525031129/internal/locale/translations/en_US.json","old_string":"    \"error.invalid_site_url\": \"Inva…`

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-doctrine-1525031129/internal/locale/translations/fr_FR.json","old_string":"    \"error.invalid_feed_url\": \"URL …`

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-doctrine-1525031129/internal/locale/translations/de_DE.json","old_string":"    \"error.invalid_feed_url\": \"Ung�…`

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-doctrine-1525031129/internal/locale/translations/es_ES.json","old_string":"    \"error.invalid_feed_url\": \"URL …`

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-doctrine-1525031129/internal/locale/translations/pl_PL.json","old_string":"    \"error.invalid_feed_url\": \"Niep…`

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-doctrine-1525031129/internal/locale/translations/ro_RO.json","old_string":"    \"error.invalid_feed_url\": \"Adre…`

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-doctrine-1525031129/internal/locale/translations/uk_UA.json","old_string":"    \"error.invalid_feed_url\": \"Не…`

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-doctrine-1525031129/internal/locale/translations/gl_ES.json","old_string":"    \"error.invalid_feed_url\": \"URL …`

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-doctrine-1525031129/internal/locale/translations/zh_TW.json","old_string":"    \"error.invalid_feed_url\": \"訂�…`

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-doctrine-1525031129/internal/locale/translations/pt_BR.json","old_string":"    \"error.invalid_feed_url\": \"URL …`

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-doctrine-1525031129/internal/locale/translations/nl_NL.json","old_string":"    \"error.invalid_feed_url\": \"Onge…`

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-doctrine-1525031129/internal/locale/translations/ru_RU.json","old_string":"    \"error.invalid_feed_url\": \"Не…`

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-doctrine-1525031129/internal/locale/translations/hi_IN.json","old_string":"    \"error.invalid_feed_url\": \"द�…`

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-doctrine-1525031129/internal/locale/translations/fi_FI.json","old_string":"    \"error.invalid_feed_url\": \"Virh…`

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-doctrine-1525031129/internal/locale/translations/nan_Latn_pehoeji.json","old_string":"    \"error.invalid_feed_ur…`

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-doctrine-1525031129/internal/locale/translations/it_IT.json","old_string":"    \"error.invalid_feed_url\": \"URL …`

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-doctrine-1525031129/internal/locale/translations/zh_CN.json","old_string":"    \"error.invalid_feed_url\": \"无�…`

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-doctrine-1525031129/internal/locale/translations/ja_JP.json","old_string":"    \"error.invalid_feed_url\": \"フ�…`

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-doctrine-1525031129/internal/locale/translations/ko_KR.json","old_string":"    \"error.invalid_feed_url\": \"피�…`

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-doctrine-1525031129/internal/locale/translations/el_EL.json","old_string":"    \"error.invalid_feed_url\": \"Μη…`

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-doctrine-1525031129/internal/locale/translations/tr_TR.json","old_string":"    \"error.invalid_feed_url\": \"Geç…`

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-doctrine-1525031129/internal/locale/translations/ar_SA.json","old_string":"    \"error.invalid_feed_url\": \"را…`

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-doctrine-1525031129/internal/locale/translations/id_ID.json","old_string":"    \"error.invalid_feed_url\": \"URL …`

**Agent:** Saving integration settings now rejects a bad webhook URL, but I couldn't run any tests. Bash was denied in this session, so `gofmt`, `go vet` and `go test` (including the locale test `TestMissingTranslations`) have not run. Please run `go test ./internal/ui/... ./internal/locale/...` before merging.

- **Rule:** a webhook URL is accepted only if it starts with `http://` or `https://` and has a host. For example, `example.org/hook`, `ftp://…` and `http:///hook` are rejected.
- **Where it runs:** `ValidateWebhook()` in `internal/ui/form/integration.go`, called from `updateIntegration` in `internal/ui/integration_update.go`. It follows the existing `ValidateGoogleReader` pattern: it sets an error message, redirects to `/integrations`, and saves nothing.
- **Unchanged:** an empty URL with the webhook enabled still silently disables the webhook. A URL typed while the webhook is disabled isn't checked, because it gets cleared anyway.
- **Tests:** I added `TestValidateWebhook` in `internal/ui/form/integration_test.go`. It covers valid URLs, a missing scheme, an unsupported scheme, a missing host, an unparsable URL, a disabled webhook and an empty URL.
- **Translations:** I added `error.invalid_webhook_url` to all 22 locale files, because that locale test requires every key in every language. Only en, fr, de, es, pt_BR, nl and it have real translations. The other 15 have the English text as a placeholder. In the non-English files the key sits right after `error.invalid_feed_url`, so it may not be in alphabetical order.

This only checks the URL's format. It doesn't block internal addresses (localhost, private IPs), and the webhook is sent from the server. `urllib.IsNonPublicIP` exists and could be used for that, but it would break people who send webhooks to services on their own network. I left it out; say if you want it.

