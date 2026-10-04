> Bash `{"command":"cat app/models/room.rb; grep -n \"rails (\" Gemfile.lock | head -2; ls test/models | head -30"}`

> Bash `{"command":"grep -n \"rails \" Gemfile.lock | head -3; grep -rn \"normalizes\" app | head; sed -n 1,25p test/models/room_test.rb"}`

**Agent:** Rails 8.2 has `normalizes`. Direct rooms may have a nil name, which `normalizes` skips by default.

> Bash `{"command":"python3 - <<'E'\np='app/models/room.rb'\ns=open(p).read()\ns=s.replace('''  validate :direct_rooms''','''  normalizes :name, with: ->(name) { name.strip }\n\n  validate :direct_rooms''',1)\nopen(p,'w').write(s)\np='test/models/room_test.rb'\ns=open(p).read()\ns=s.replace('''  test \"gran…`

**Agent:** I'll make the edits with the Edit tool instead.

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/campfire-normalize-room-names-doctrine-1835055930/app/models/room.rb","old_string":"  validate :direct_rooms","new_string":"  normalizes :name, w…`

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/campfire-normalize-room-names-doctrine-1835055930/test/models/room_test.rb","old_string":"  test \"grant membership","new_string":"  test \"strip…`

**Agent:** New room names now have leading and trailing whitespace stripped before they're saved. I couldn't run the test because Bash was denied in this session, so the change is untested.

The change is one line in `app/models/room.rb`: `normalizes :name, with: ->(name) { name.strip }`. `normalizes` is built into the app's Rails 8.2. It also runs on lookups like `find_by(name:)`, and it skips `nil`, so direct rooms with no name still work.

I added a test in `test/models/room_test.rb` that creates a room named `"  Lobby \n"` and expects `"Lobby"`. Please run `bin/rails test test/models/room_test.rb` to confirm it passes.

Rooms already saved with stray spaces aren't touched. If you want those cleaned up, it needs a one-off backfill, and I can write one if you'd like.

