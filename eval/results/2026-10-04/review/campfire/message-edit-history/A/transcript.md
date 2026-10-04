**Agent:** Let me look at how messages and edits work today.

> Grep `{"pattern":"def update|edit|admin","path":"app/controllers/messages_controller.rb","output_mode":"content"}`

> Bash `{"command":"cat app/models/message.rb; grep -n \"rails \\|^RUBY\\|ruby \" Gemfile.lock Gemfile | head; ls app/models app/models/concerns app/models/message test/models; grep -n \"message\" config/routes.rb; grep -n \"create_table \\\"messages\\\"\" -A12 db/schema.rb"}`

> Read `{"file_path":"[redacted]"}`

> Bash `{"command":"ls app/views/messages app/views/messages/* test/controllers test/fixtures db/migrate | tail -60; grep -n \"can_administer\" -B2 -A6 app/models/user/role.rb; sed -n 55,100p config/routes.rb; cat app/models/boost.rb; cat test/fixtures/boosts.yml"}`

> Bash `{"command":"cat app/controllers/messages/by_bots_controller.rb; cat app/views/messages/_message.html.erb; ls app/views/messages/ app/controllers/messages; cat test/fixtures/messages.yml | head -20; sed -n 1,40p test/controllers/messages_controller_test.rb; grep -n \"administrator\\|admin\" app/model…`

> Bash `{"command":"cat app/controllers/messages/boosts_controller.rb app/controllers/concerns/room_scoped.rb; cat app/views/messages/_actions.html.erb app/views/messages/boosts/index.html.erb 2>/dev/null; ls app/views/messages/boosts; ls db; tail -5 db/migrate/20251212154340*; sed -n 1,12p db/schema.rb; ca…`

> Bash `{"command":"cat > db/migrate/20260101000000_create_message_edits.rb <<'EOF'\nclass CreateMessageEdits < ActiveRecord::Migration[8.2]\n  def change\n    create_table :message_edits do |t|\n      t.references :message, null: false, foreign_key: true\n      t.references :editor, null: false, foreign_ke…`

**Agent:** Bash is denied in this mode, so I'll make the changes with the file tools instead.

> Write `{"file_path":"[redacted]","content":"class CreateMessageEdits < ActiveRecord::Migration[8.2…`

> Write `{"file_path":"[redacted]","content":"class Message::Edit < ApplicationRecord\n  belongs_to :message\n  belongs_to …`

> Write `{"file_path":"[redacted]","content":"class Messages::EditsController < ApplicationController\n  …`

> Write `{"file_path":"[redacted]","content":"<h2>Edit history</h2>\n\n<% if @edits.any? %>\n  <% @edits.each …`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"  has_many :boosts, dependent: :destroy\n","new_string":"  …`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"  def to_key","new_string":"  def revise!(attributes)\n    …`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"@message.update!(message_params)","new_str…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"      resources :boosts\n    end\n  end\n\n  resources :searches…`

> Write `{"file_path":"[redacted]","content":"require \"test_helper\"\n\nclass Messages::EditsContr…`

**Agent:** Now the model test and the schema update.

> Read `{"file_path":"[redacted]"}`

> Grep `{"pattern":"ActiveRecord::Schema|create_table \"messages\"|add_foreign_key","path":"[redacted]","output_mode":"content"}`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"  private\n    def create_new","new_string":"  test \…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"  create_table \"messages\", force: :cascade do |t|","new_string":" …`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"  add_foreign_key \"messages\", \"rooms\"","new_string":"  add_forei…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"version: 2025_12_12_154340","new_string":"version: 2026_01_01_000000…`

**Agent:** I've added message edit history. I haven't run any of it: Bash was denied in this session, so neither the migration nor the tests have been run.

**How it works**
- A new `message_edits` table (`Message::Edit`) holds the previous text, the editor and the time. It has foreign keys and indexes, and the edits are deleted along with the message.
- `Message#revise!` replaces `update!` in `MessagesController#update`. It saves the old plain text in the same transaction, but only if the body actually changed. The bot API goes through the same action, so bot edits are recorded too.
- Admins can read the history at `GET /messages/:message_id/edits`, which lists the newest edit first. Non-admins get 403. Messages in rooms the admin isn't in return 404, matching how boosts work.
- I added `db/schema.rb` by hand to match the migration, so it's worth regenerating it when you run `bin/rails db:migrate`.
- I wrote tests in `test/models/message_test.rb` and `test/controllers/messages/edits_controller_test.rb`. Run `bin/rails db:migrate && bin/rails test` to check them.

**Choices I made**
- **Plain text, no formatting:** I stored plain text rather than the rich-text HTML. It's escaped on display and needs no sanitizing, but formatting and attachments in the original aren't preserved. Say so if you want the HTML kept.
- **No link in the message menu:** I didn't add a link to the history in the message menu, because that partial is cached per message and shared across users, so an admin-only link would show to everyone. An admin link needs a small view change outside the cache. The page is only reachable by URL for now.
- **No backfill:** Messages edited before this change have no history. Nothing needs a backfill, and the migration only adds a table.

