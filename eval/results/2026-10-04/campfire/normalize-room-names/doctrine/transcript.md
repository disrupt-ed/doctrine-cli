> Bash `{"command":"cat app/models/room.rb; grep -n \"rails (\" Gemfile.lock | head -3; grep -rn \"normalizes\" app/models | head"}`

> Bash `{"command":"cat test/models/room_test.rb | head -30"}`

> Bash `{"command":"python3 - <<'E'\np='app/models/room.rb'\ns=open(p).read()\ns=s.replace('  validate :direct_rooms','  normalizes :name, with: ->(name) { name.strip }\\n\\n  validate :direct_rooms',1)\nopen(p,'w').write(s)\np='test/models/room_test.rb'\ns=open(p).read()\ns=s.replace('  test \"type\" do','…`

**Agent:** I'll use the edit tools instead.

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/campfire-normalize-room-names-doctrine-1776824604/app/models/room.rb","old_string":"  validate :direct_rooms_keep_their_type, on: :update","new_s…`

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/campfire-normalize-room-names-doctrine-1776824604/test/models/room_test.rb","old_string":"  test \"type\" do","new_string":"  test \"strips surro…`

**Agent:** Room names are now stripped of leading and trailing whitespace before they're saved. I haven't run the tests: Bash was denied in this session, so the new test is unverified.

- **Fix:** I added `normalizes :name, with: ->(name) { name.strip }` to `app/models/room.rb`. The app is on Rails 8.2, so the built-in `normalizes` covers this without a gem or callback. Nil names pass through untouched, because `normalizes` skips nil by default.
- **Test:** I added a test in `test/models/room_test.rb` that creates a room named `"  Lobby \n"` and expects `"Lobby"`.
- **Existing data:** Rooms already saved with stray spaces keep them until they're next saved. If you want those cleaned up, it needs a one-off backfill, which I haven't done.

To check it, run `bin/rails test test/models/room_test.rb`.

