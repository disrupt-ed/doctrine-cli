> Bash `{"command":"cd /private/tmp/claude-501/*[redacted] && grep -n \"rails \\|^RUBY\" -A0 Gemfile.lock | head -3; cat app/models/message.rb; cat app/controllers/messages_controller.rb; ls app/models app/models/message app/controllers/messages te…`

**Agent:** Bash is denied in this mode, so I'll explore with the file tools.

> Read `{"file_path":"[redacted]"}`

> Read `{"file_path":"[redacted]"}`

> Glob `{"pattern":"app/{models,controllers}/**/*.rb","path":"[redacted]"}`

> Read `{"file_path":"[redacted]"}`

> Read `{"file_path":"[redacted]"}`

> Read `{"file_path":"[redacted]"}`

> Glob `{"pattern":"{test/**/message*,test/**/boost*,app/views/messages/**,db/migrate/*}","path":"[redacted]"}`

> Read `{"file_path":"[redacted]"}`

> Read `{"file_path":"[redacted]"}`

> Read `{"file_path":"[redacted]"}`

> Read `{"file_path":"[redacted]"}`

> Read `{"file_path":"[redacted]"}`

> Read `{"file_path":"[redacted]"}`

> Read `{"file_path":"[redacted]"}`

> Read `{"file_path":"[redacted]"}`

> Read `{"file_path":"[redacted]"}`

> Read `{"file_path":"[redacted]","limit":20}`

> Read `{"file_path":"[redacted]"}`

> Write `{"file_path":"[redacted]","content":"class CreateMessageEdits < ActiveRecord::Migration[8.…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"  create_table \"messages\", force: :cascade do |t|","new_string":"…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"  add_foreign_key \"messages\", \"rooms\"","new_string":"  add_fore…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"2025_12_12_154340","new_string":"2025_12_15_090000"}`

> Write `{"file_path":"[redacted]","content":"class Message::Edit < ApplicationRecord\n  belongs_to :message\n\n  has_rich…`

> Write `{"file_path":"[redacted]","content":"module Message::Editable\n  extend ActiveSupport::Concern\n\n  included …`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"Broadcasts, Mentionee,","new_string":"Broadcasts, Editable…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"      resources :boosts\n    end\n  end\n\n  resources :searche…`

> Write `{"file_path":"[redacted]","content":"class Messages::EditsController < ApplicationController\n …`

> Write `{"file_path":"[redacted]","content":"<% @page_title = \"Edit history\" %>\n\n<div class=\"pad\">\n  …`

> Grep `{"pattern":"local_datetime_tag|def .*time.*tag|page_title","path":"[redacted]","output_mode":"content","head_limit":15}`

> Read `{"file_path":"[redacted]"}`

**Agent:** Checking which time styles the local-time controller supports, then adding the admin link and tests.

> Grep `{"pattern":"Target|Format","path":"[redacted]","output_mode":"content"}`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"        <% end %>\n      </div>\n    </div>…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"<div class=\"pad\">","new_string":"<div …`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"  private\n    def create_new","new_string":"  test …`

> Write `{"file_path":"[redacted]","content":"require \"test_helper\"\n\nclass Messages::EditsCont…`

**Agent:** Message edits are now recorded, and admins can see the original text. I couldn't run the migration or any tests, so none of this has been run. Bash was denied in this session.

**How it works**
- When a message's body changes, `Message::Editable` saves the previous body as a `Message::Edit` row in the new `message_edits` table. The body is stored as Action Text, so mentions and attachments are kept. Saves that don't change the body record nothing.
- This also covers edits made through the bot API, since it goes through the same model.
- `GET /messages/:message_id/edits` lists the edits newest first. It is admin-only, and non-admins get 403, including on their own messages.
- The message actions menu shows an "Edit history" link to admins, only on messages that have been edited. It reuses the `reply.svg` icon as a placeholder, so you'll want to swap in a proper one.
- Tests are in `test/models/message_test.rb` and the new `test/controllers/messages/edits_controller_test.rb`.

**Things to check**
- **Hand-edited schema:** I wrote the migration (`20251215090000_create_message_edits.rb`) and edited `db/schema.rb` by hand. Run `bin/rails db:migrate` and check the schema diff.
- **Tests:** run `bin/rails test`. One thing I couldn't confirm is that the old body is still readable in `before_update` when `update!` assigns a new body. That is the first place I'd look if the model test fails.
- **Deploy:** only edits made after deploy are tracked. Earlier edits can't be recovered, so there's no backfill.
- **No editor recorded:** I left out who made each edit, to avoid a foreign key to users. That's a one-column addition if you want it.

