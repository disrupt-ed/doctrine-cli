> Grep `{"pattern":"class Message\\b|body|composer","path":"app","glob":"{models/message.rb,views/messages/_composer*,views/rooms/show*}","output_mode":"content"}`

> Glob `{"pattern":"app/{views/messages/*,javascript/controllers/*composer*,models/message*}"}`

> Read `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/campfire-message-length-limit-doctrine-1726128333/app/models/message.rb"}`

> Read `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/campfire-message-length-limit-doctrine-1726128333/app/javascript/controllers/composer_controller.js"}`

> Read `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/campfire-message-length-limit-doctrine-1726128333/app/views/rooms/show/_composer.html.erb"}`

> Bash `{"command":"cd /private/tmp/claude-501/*/scratchpad/evalwork/campfire-message-length-limit-doctrine-1726128333 && cat app/controllers/messages_controller.rb; grep -rn \"composer__input-hint\\|\\.composer__attachment-btn\" app/assets/stylesheets | head; ls test/models | head -30; grep -n \"rescue_fro…`

> Read `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/campfire-message-length-limit-doctrine-1726128333/app/controllers/messages_controller.rb"}`

> Grep `{"pattern":"composer__attachment-btn|composer__rich-text-btn","path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/campfire-message-length-limit-doctrine-1726128333/app/assets/stylesheets","output_mode":"content"}`

> Glob `{"pattern":"test/{models/message_test.rb,controllers/messages_controller_test.rb}","path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/campfire-message-length-limit-doctrine-1726128333"}`

> Read `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/campfire-message-length-limit-doctrine-1726128333/test/models/message_test.rb"}`

> Read `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/campfire-message-length-limit-doctrine-1726128333/app/assets/stylesheets/composer.css","limit":110}`

**Agent:** Adding the model validation first, then the composer counter. I'm checking how the editor element exposes its text.

> Grep `{"pattern":"isBlank|toPlainText|get value","path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/campfire-message-length-limit-doctrine-1726128333/app/javascript","output_mode":"content"}`

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/campfire-message-length-limit-doctrine-1726128333/app/models/message.rb","old_string":"  has_rich_text :body\n\n","new_string":"  has_rich_text :…`

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/campfire-message-length-limit-doctrine-1726128333/app/models/message.rb","old_string":"      Sound.find_by_name match[:name]\n    end\n  end\n","…`

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/campfire-message-length-limit-doctrine-1726128333/test/models/message_test.rb","old_string":"  private\n    def create","new_string":"  test \"bo…`

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/campfire-message-length-limit-doctrine-1726128333/app/views/rooms/show/_composer.html.erb","old_string":"                <label class=\"btn btn--…`

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/campfire-message-length-limit-doctrine-1726128333/app/assets/stylesheets/composer.css","old_string":"/* Attachments */","new_string":"/* Characte…`

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/campfire-message-length-limit-doctrine-1726128333/app/javascript/controllers/composer_controller.js","old_string":"  saveDraft() {\n    if","new_…`

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/campfire-message-length-limit-doctrine-1726128333/app/javascript/controllers/composer_controller.js","old_string":"    this.textTarget.value = co…`

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/campfire-message-length-limit-doctrine-1726128333/app/javascript/controllers/composer_controller.js","old_string":"      this.textTarget.value = …`

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/campfire-message-length-limit-doctrine-1726128333/app/javascript/controllers/composer_controller.js","old_string":"    return !this.textTarget.is…`

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/campfire-message-length-limit-doctrine-1726128333/app/javascript/controllers/composer_controller.js","old_string":"    this.textTarget.value = \"…`

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/campfire-message-length-limit-doctrine-1726128333/app/javascript/controllers/composer_controller.js","old_string":"\"clientid\", \"fields\",","ne…`

**Agent:** Messages are now limited to 5,000 characters, and the composer shows a live `n / 5000` counter. I haven't run any of it: the tests, the browser behaviour and the counter styling are all unchecked, because Bash was denied here.

- **Server:** `Message` now validates the plain-text length of the body (limit in `Message::MAX_BODY_LENGTH`). That also covers the API and bots. The existing `create!` raises on a too-long message, which I expect Rails to return as a 422, but I didn't check. I added a model test for exactly 5,000 characters (valid) and 5,001 (invalid) in `test/models/message_test.rb`.
- **Composer:**
  - `composer_controller.js` updates the counter whenever the draft changes, is restored, is replaced or is reset.
  - The counter is hidden when the composer is empty. It turns red and bold when the text is over the limit.
  - Sending is blocked while the text is over the limit. The server check is the real enforcement.
- **Markup and style:** the counter span is in `_composer.html.erb`, next to the attach button. Its style is in `composer.css`. The colour variables there are my guesses (`--color-text-lighter`, `--color-negative`) with fallbacks, so check the look.

The browser counts the characters of the editor's HTML text, while the server uses Rails' `to_plain_text`. The server adds newlines between paragraphs, so a message with many line breaks can be a few characters over on the server even when the counter says it fits.

Messages already over 5,000 characters stay as they are, but editing one will fail validation until it is shortened.

