> Grep `{"pattern":"class Message|validates|maxlength","path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/campfire-message-length-limit-doctrine-1084005371","glob":"app/{models/message*,views/messages/**,views/rooms/show/**,javasc…`

> Glob `{"pattern":"app/**/*composer*","path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/campfire-message-length-limit-doctrine-1084005371"}`

> Read `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/campfire-message-length-limit-doctrine-1084005371/app/models/message.rb"}`

> Read `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/campfire-message-length-limit-doctrine-1084005371/app/views/rooms/show/_composer.html.erb"}`

> Read `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/campfire-message-length-limit-doctrine-1084005371/app/javascript/controllers/composer_controller.js"}`

> Bash `{"command":"cd /private/tmp/claude-501/*/scratchpad/evalwork/campfire-message-length-limit-doctrine-1084005371 && cat app/controllers/messages_controller.rb && cat app/assets/stylesheets/composer.css | head -60 && ls test/models | head -30 && grep -n \"limit\\|create\" test/controllers/messages_cont…`

> Read `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/campfire-message-length-limit-doctrine-1084005371/app/controllers/messages_controller.rb"}`

> Grep `{"pattern":"composer__input-hint|typing-indicator \\{","path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/campfire-message-length-limit-doctrine-1084005371/app/assets/stylesheets","output_mode":"content","-n":true}`

> Glob `{"pattern":"test/**/message*_test.rb","path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/campfire-message-length-limit-doctrine-1084005371"}`

> Grep `{"pattern":"isBlank|toPlainText|get (value|editorElement|editorContentElement)","path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/campfire-message-length-limit-doctrine-1084005371","glob":"!{log,tmp,node_modules}/**","out…`

> Read `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/campfire-message-length-limit-doctrine-1084005371/test/models/message_test.rb"}`

**Agent:** I'm adding a length validation to the model, then a counter in the composer that blocks sending when the message is over the limit.

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/campfire-message-length-limit-doctrine-1084005371/app/models/message.rb","old_string":"  has_rich_text :body\n\n","new_string":"  has_rich_text :…`

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/campfire-message-length-limit-doctrine-1084005371/app/models/message.rb","old_string":"      Sound.find_by_name match[:name]\n    end\n  end\n","…`

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/campfire-message-length-limit-doctrine-1084005371/test/models/message_test.rb","old_string":"  test \"mentionees\" do","new_string":"  test \"bod…`

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/campfire-message-length-limit-doctrine-1084005371/app/views/rooms/show/_composer.html.erb","old_string":"                </div>\n\n              …`

> Grep `{"pattern":"def composer_form_tag","path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/campfire-message-length-limit-doctrine-1084005371/app/helpers","output_mode":"content","-A":12}`

> Grep `{"pattern":"def composer_data_options","path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/campfire-message-length-limit-doctrine-1084005371/app/helpers/rooms_helper.rb","output_mode":"content","-A":16}`

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/campfire-message-length-limit-doctrine-1084005371/app/helpers/rooms_helper.rb","old_string":"composer_room_id_value: room.id\n","new_string":"com…`

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/campfire-message-length-limit-doctrine-1084005371/app/views/rooms/show/_composer.html.erb","old_string":"\n                      data-composer-ta…`

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/campfire-message-length-limit-doctrine-1084005371/app/views/rooms/show/_composer.html.erb","old_string":"lexxy:change->composer#saveDraft ","new_…`

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/campfire-message-length-limit-doctrine-1084005371/app/javascript/controllers/composer_controller.js","old_string":"  static targets = [ \"clienti…`

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/campfire-message-length-limit-doctrine-1084005371/app/javascript/controllers/composer_controller.js","old_string":"  submit(event) {","new_string…`

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/campfire-message-length-limit-doctrine-1084005371/app/javascript/controllers/composer_controller.js","old_string":"    return !this.textTarget.is…`

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/campfire-message-length-limit-doctrine-1084005371/app/javascript/controllers/composer_controller.js","old_string":"    return html.body.innerText…`

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/campfire-message-length-limit-doctrine-1084005371/app/javascript/controllers/composer_controller.js","old_string":"    localStorage.removeItem(th…`

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/campfire-message-length-limit-doctrine-1084005371/app/javascript/controllers/composer_controller.js","old_string":"    this.#restoreDraft()\n","n…`

> Read `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/campfire-message-length-limit-doctrine-1084005371/app/assets/stylesheets/composer.css","offset":150,"limit":15}`

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/campfire-message-length-limit-doctrine-1084005371/app/assets/stylesheets/composer.css","old_string":"/* Typing indicator */","new_string":".compo…`

**Agent:** Messages are now limited to 5,000 characters, and the composer shows a counter. I haven't run anything: Bash was denied in this session, so the new model test and the UI are untested.

- **Server limit:** `Message` validates the plain-text length of the body against `MAX_BODY_LENGTH = 5_000` (`app/models/message.rb`). This covers the API and bots as well as the web composer. There's a boundary test in `test/models/message_test.rb`: 5,000 characters is valid and 5,001 is not.
- **Counter:** the composer shows `n / 5000`, updated on each editor change, on load, and after a send. It turns red when you're over. The limit is passed in from the model through `composer_data_options`, so the number lives in one place.
- **Blocking:** while a message is over the limit, the send button and Enter do nothing. The message stays in the editor so you can trim it.

- **Server rejection:** if a too-long message reaches the server, `create_with_attachment!` raises `RecordInvalid`. I didn't add a custom rescue. I'm assuming Rails answers with a 422 and the composer's existing failure handling (`submitEnd`) marks the pending message as failed. I haven't confirmed that.
- **Count mismatch:** the browser counts the editor HTML's text content, while the server counts `to_plain_text`. Mentions and line breaks can differ by a few characters near the limit.
- **Red colour:** the `.composer__counter--over` rule uses `--color-negative` with a plain `red` fallback. I didn't check that variable exists in the stylesheets.
- **Existing messages:** messages already over 5,000 characters aren't touched. They will fail validation if someone edits them.

Run `bin/rails test test/models/message_test.rb` and try typing past the limit in a room.

