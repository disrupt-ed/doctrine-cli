> Grep `{"pattern":"WebhookURL|webhook_url","glob":"*.go","output_mode":"content","head_limit":60}`

> Read `{"file_path":"[redacted]"}`

> Grep `{"pattern":"func Validate|validator\\.|ValidateFeedURL|IsValidURL","path":"[redacted]","output_mode":"content","glob":"*.go","hea…`

> Grep `{"pattern":"func IsValid|func Validate","path":"[redacted]","output_mode":"content"}`

> Read `{"file_path":"[redacted]","offset":1,"limit":20}`

> Grep `{"pattern":"ValidateGoogleReader","path":"[redacted]","output_mode":"content","-A":20}`

> Grep `{"pattern":"ValidateFeedURL|IsValidURL|func.*URL","path":"[redacted]","output_mode":"content"}`

> Read `{"file_path":"[redacted]"}`

> Grep `{"pattern":"googlereader_missing_required_fields|error.feed_invalid_proxy_url|error.invalid_feed_url","path":"[redacted]…`

> Grep `{"pattern":"func IsAbsoluteURL","path":"[redacted]","output_mode":"content","-A":12}`

> Bash `{"command":"ls internal/locale/translations && grep -n \"googlereader_missing_required_fields\" internal/locale/translations/*.json | head -30"}`

**Agent:** Adding a webhook URL check to the integration save handler. I'm checking how new translation keys get added across locales first.

> Bash `{"command":"grep -n \"linktaco_missing_required_fields\" internal/locale/translations/*.json | head -30; sed -n 55,60p internal/urllib/url.go; sed -n 71,90p internal/urllib/url.go"}`

> Bash `{"command":"for f in internal/locale/translations/*.json; do\nsed -i '' '/\"error.invalid_feed_url\":/{\np\ns/\"error.invalid_feed_url\": .*/\"error.invalid_webhook_url\": \"Invalid webhook URL. It must be an absolute HTTP or HTTPS URL.\",/\n}' $f; done; grep -n \"invalid_webhook_url\" internal/loca…`

> Grep `{"pattern":"en_US|defaultLanguage|fallback","path":"[redacted]","glob":"*.go","output_mode":"content"}`

> Grep `{"pattern":"\"error.invalid_site_url\"|\"error.invalid_feed_url\"","path":"[redacted]","output_mode":"content…`

> Grep `{"pattern":"\"error.invalid_theme\"","path":"[redacted]","output_mode":"content","-A":1,"glob":"{en_US,fr_FR,…`

> Grep `{"pattern":"\"error.invalid_timezone\"","path":"[redacted]","output_mode":"content","-A":1,"glob":"en_US.json…`

**Agent:** Next key is `error.local_auth_disabled`, so I'll insert before it in all 23 locale files, then add the validation code.

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"    \"error.local_auth_disabled\":","n…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"// NewIntegrationForm returns","new_string":"/…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"\t\"net/http\"\n\t\"strconv\"","new_string":"\…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"\tif integrationForm.WebhookEnabled {\n\t\ti…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"    \"error.local_auth_disabled\":","n…`

> Grep `{"pattern":"\"error.local_auth_disabled\"","path":"[redacted]","output_mode":"count"}`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"    \"error.local_auth_disabled\":","n…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"    \"error.local_auth_disabled\":","n…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"    \"error.local_auth_disabled\":","n…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"    \"error.local_auth_disabled\":","n…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"    \"error.local_auth_disabled\":","n…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"    \"error.local_auth_disabled\":","n…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"    \"error.local_auth_disabled\":","n…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"    \"error.local_auth_disabled\":","n…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"    \"error.local_auth_disabled\":","n…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"    \"error.local_auth_disabled\":","n…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"    \"error.local_auth_disabled\":","n…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"    \"error.local_auth_disabled\":","n…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"    \"error.local_auth_disa…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"    \"error.local_auth_disabled\":","n…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"    \"error.local_auth_disabled\":","n…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"    \"error.local_auth_disabled\":","n…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"    \"error.local_auth_disabled\":","n…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"    \"error.local_auth_disabled\":","n…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"    \"error.local_auth_disabled\":","n…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"    \"error.local_auth_disabled\":","n…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"    \"error.local_auth_disabled\":","n…`

**Agent:** Saving integration settings now rejects a webhook URL that isn't an absolute `http`/`https` URL with a host. I haven't compiled or tested any of it, because Bash was denied in this session.

- **Validation:** `ValidateWebhook()` in `internal/ui/form/integration.go` parses the URL and returns `error.invalid_webhook_url` if it fails. It only checks when the webhook integration is enabled and a URL is entered, so an empty URL still disables the integration as before.
- **Handler:** `updateIntegration` in `internal/ui/integration_update.go` calls it before touching the webhook settings. On failure it shows the error and redirects to `/integrations` without saving, like the Google Reader and LinkTaco checks.
- **Translations:** I added `error.invalid_webhook_url` to all 23 locale files, placed alphabetically. Only English and French are translated. The other 21 contain the English text as a placeholder.

I didn't add a test. The per-feed webhook URL (feed edit) is not validated either, and I left it alone since you asked about integration settings.

Please run `go build ./... && go test ./internal/ui/... ./internal/locale/...` to check it.

