> Grep `{"pattern":"class Room|before_validation|normalizes","path":"[redacted]","output_mode":"content"}`

> Read `{"file_path":"[redacted]"}`

> Bash `{"command":"grep -E \"^    rails \\(\" Gemfile.lock; grep -rn \"normalizes\" app | head -3"}`

**Agent:** `normalizes` skips nil values by default, so direct rooms with no name are unaffected.

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"  validate :direct_rooms_keep_their_type, on: :update\n","new_…`

> Bash `{"command":"ls test/models | grep -i room; grep -n \"^  test\" test/models/room_test.rb | head -5"}`

> Read `{"file_path":"[redacted]","limit":30}`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"  test \"type\" do","new_string":"  test \"strips surrou…`

> Bash `{"command":"bin/rails test test/models test/controllers/rooms 2>&1 | tail -8"}`

> Bash `{"command":"bin/rails test test/models/room_test.rb 2>&1 | head -5"}`

**Agent:** Room names now have leading and trailing spaces stripped before they're saved. I couldn't run the tests, so the change is untested.

The change is one line in `app/models/room.rb`: `normalizes :name, with: ->(name) { name.strip }`. It applies to open, closed and direct rooms. A room with no name, like a direct room, is left as nil.

I added a test in `test/models/room_test.rb` that creates a room named `"  Hello!  "` and checks the name comes back as `"Hello!"`.

The test run failed before reaching any tests. The Gemfile pins Rails to a git checkout that isn't installed here, and Bundler asks for `bundle install`. Run that, then `bin/rails test test/models/room_test.rb`.

Existing rooms that already have stray spaces in their names aren't fixed by this. Rails normalizes the name when you assign it, so those rooms stay as they are until someone renames them. A one-off data migration would clean them up. I haven't written one, so say if you want it.

