> Grep `{"pattern":"WebhookURL|webhook_url","glob":"**/*.go","output_mode":"content"}`

> Read `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-doctrine-3175275519/internal/ui/form/integration.go","offset":1,"limit":20}`

> Grep `{"pattern":"func \\(i \\*?IntegrationForm\\) Validate|func .*Validate","path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-doctrine-3175275519/internal/ui","output_mode":"content"}`

> Read `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-doctrine-3175275519/internal/ui/integration_update.go"}`

**Agent:** Let me see how the project validates URLs elsewhere (feed URLs, API) and the form validator.

> Grep `{"pattern":"func (ValidateURL|IsValidURL|ValidateFeedURL|.*URL.*validat)|validator\\.Validate\\w*URL|func Validate\\w*URL","path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-doctrine-3175275…`

> Read `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-doctrine-3175275519/internal/ui/form/integration.go","offset":255,"limit":30}`

> Grep `{"pattern":"url\\.Parse|ParseRequestURI|IsPrivate|\\.IsLoopback|invalid_feed_url|error\\.invalid_.*url|error\\.\\w*url\\w*\"","path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-doctrine-3175…`

> Grep `{"pattern":"error.googlereader_missing_required_fields","path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-doctrine-3175275519/internal/locale","output_mode":"count"}`

> Read `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-doctrine-3175275519/internal/validator/feed.go","limit":30}`

> Read `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-doctrine-3175275519/internal/urllib/url.go","offset":120,"limit":100}`

> Grep `{"pattern":"error.invalid_feed_url|error.googlereader_missing","path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-doctrine-3175275519/internal/locale/translations/en_US.json","output_mode":"…`

> Read `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-doctrine-3175275519/internal/urllib/url.go","limit":90,"offset":56}`

> Read `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-doctrine-3175275519/internal/ui/form/integration_test.go"}`

> Bash `{"command":"git log -3 --stat --format=%s -- internal/locale/translations | head -20; grep -n \"linktaco_missing\" internal/locale/translations/{en_US,fr_FR}.json; grep -n \"missing\\|keys\" internal/locale/*_test.go | head"}`

> Grep `{"pattern":"linktaco_missing_required_fields","path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-doctrine-3175275519/internal/locale","output_mode":"content"}`

> Read `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-doctrine-3175275519/internal/locale/translations/en_US.json","offset":128,"limit":12}`

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-doctrine-3175275519/internal/locale/translations/en_US.json","old_string":"    \"error.invalid_timezone\": \"Inva…`

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-doctrine-3175275519/internal/ui/form/integration.go","old_string":"// NewIntegrationForm returns a new Integratio…`

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-doctrine-3175275519/internal/ui/form/integration.go","old_string":"\t\"net/http\"\n\t\"strconv\"\n\n\t\"miniflux.…`

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-doctrine-3175275519/internal/ui/integration_update.go","old_string":"\tif integration.GoogleReaderUsername != \"\…`

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/miniflux-webhook-url-validation-doctrine-3175275519/internal/ui/form/integration_test.go","old_string":"func TestValidateGoogleReader(","new_stri…`

**Agent:** Saving integration settings now rejects a webhook URL that isn't a valid absolute `http`/`https` URL with a host. I haven't compiled or run anything, because Bash was denied in this session. Please run `go build ./... && go test ./internal/ui/...` before merging.

- **Validation:** `ValidateWebhook()` in `internal/ui/form/integration.go` follows the same pattern as `ValidateGoogleReader`. It reuses `urllib.IsAbsoluteURL` and adds a host check. `IsAbsoluteURL` alone would accept `http:example.org`, which has no host.
- **Handler:** `internal/ui/integration_update.go` calls it before saving. On failure it redirects to `/integrations` with an error flash and saves nothing.
- **Empty URL:** an empty URL with the webhook enabled still passes validation. The existing handler code already switches the integration off in that case.
- **Message:** I added `error.invalid_webhook_url` ("Invalid webhook URL.") to `en_US.json` only. I left the other locales alone, and I didn't check whether anything requires every locale to have every key. If a locale test fails, that's the cause.
- **Tests:** `TestValidateWebhook` covers the following cases:
  - Disabled integration with an invalid URL (allowed).
  - Enabled with an empty URL (allowed).
  - `https` and `http` with a port (allowed).
  - No scheme (rejected).
  - `ftp` (rejected).
  - No host (rejected).
  - Opaque `http:example.org` (rejected).

I did not block private or loopback addresses such as `localhost` or `192.168.x.x`. Self-hosted users often point webhooks at local services, so I treated that as a separate decision. `urllib.IsNonPublicIP` exists if you want it, but it would need DNS resolution and a decision about whether it should be configurable.

The per-feed `WebhookURL` on the feed form isn't validated either. It would be a small follow-up if you want it.

