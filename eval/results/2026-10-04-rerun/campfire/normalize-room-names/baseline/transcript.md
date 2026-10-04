> Grep `{"pattern":"normalizes|before_validation|class Room","path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/campfire-normalize-room-names-baseline-3356337760/app/models","output_mode":"content"}`

> Read `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/campfire-normalize-room-names-baseline-3356337760/app/models/room.rb"}`

**Agent:** Direct rooms may have nil names, so I'll guard for that.

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/campfire-normalize-room-names-baseline-3356337760/app/models/room.rb","old_string":"  validate :direct_rooms_keep_their_type, on: :update\n","new…`

> Edit `{"replace_all":false,"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/campfire-normalize-room-names-baseline-3356337760/app/models/room.rb","old_string":"  private\n","new_string":"  private\n    def strip_name\n   …`

**Agent:** Room names now get stripped of leading and trailing whitespace before validation, so new and updated rooms save clean. I haven't run the tests or added a test for it.

The change is in `app/models/room.rb`: a `before_validation :strip_name` callback. It skips nil names, which direct rooms can have.

Names already saved with stray spaces stay as they are until the room is next saved. If you'd like, I can add a migration to clean up the existing ones.

