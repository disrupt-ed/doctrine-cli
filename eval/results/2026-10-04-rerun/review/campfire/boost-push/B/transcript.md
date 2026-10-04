> Grep `{"pattern":"boost","-i":true,"glob":"app/{models,jobs}/**/*.rb","output_mode":"files_with_matches"}`

> Glob `{"pattern":"app/**/*push*"}`

> Read `{"file_path":"[redacted]"}`

> Read `{"file_path":"[redacted]"}`

> Read `{"file_path":"[redacted]"}`

> Read `{"file_path":"[redacted]"}`

> Grep `{"pattern":"PushMessageJob|boosts","output_mode":"content","path":"[redacted]"}`

> Bash `{"command":"cd /private/tmp/claude-501/*/*/scratchpad/evalwork/campfire-boost-push-baseline-2032778083 && cat app/controllers/boosts_controller.rb; ls test/models test/models/room test/jobs 2>/dev/null; cat test/models/room/message_pusher_test.rb | head -60"}`

> Read `{"file_path":"[redacted]","offset":70,"limit":25}`

> Read `{"file_path":"[redacted]"}`

> Glob `{"pattern":"test/**/*{pusher,boost}*","path":"[redacted]"}`

> Bash `{"command":"cd /private/tmp/claude-501/*/*/scratchpad/evalwork/campfire-boost-push-baseline-2032778083 && cat app/models/push/subscription.rb; grep -n \"scope\\|disconnected\" app/models/membership.rb app/models/membership/*.rb; cat test/fixtures/boosts.yml; ls test/models test/models/room"}`

**Agent:** Bash isn't available here, so I'm using the file tools to read the membership scopes and the subscription model before I write the boost push.

> Grep `{"pattern":"scope|disconnected|has_many|belongs_to","path":"[redacted]","glob":"{membership.rb,membership/*.rb,push/*.rb,user.rb}","output_…`

> Write `{"file_path":"[redacted]","content":"class Boost::Pusher\n  attr_reader :boost\n\n  delegate :message, :booster, to: :boost…`

> Write `{"file_path":"[redacted]","content":"class Boost::PushJob < ApplicationJob\n  def perform(boost)\n    Boost::Pusher.new(boo…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"  scope :ordered","new_string":"  after_create_commit -> { Boost::Push…`

**Agent:** Boosting a message now queues a web push to the message's author. I haven't run it or added tests, because Bash was denied in this session. Nothing is committed.

- **`app/models/boost.rb`:** an `after_create_commit` hook enqueues `Boost::PushJob`.
- **`app/jobs/boost/push_job.rb`:** a new job that calls `Boost::Pusher`.
- **`app/models/boost/pusher.rb`:** a new class modelled on `Room::MessagePusher`, with these behaviours:
  - **Who gets it:** only the author's push subscriptions. Nothing is sent when you boost your own message.
  - **When it's sent:** only if the author has a visible membership in the room and isn't currently connected to it, as with message pushes.
  - **Payload:** the title is the room name, or the booster's name in a direct room. The body is "Name boosted: <message text, cut to 100 characters>". Tapping it opens the room.

Boosts made through the bot endpoint go through the same model, so they trigger the same push.

I skipped the "involved in everything / mentions" notification settings. A boost is aimed at one person, so I didn't gate it on those. If you'd rather it respect them, that's a small change.

I'd add a test for `Boost::Pusher` and run the suite before merging. Do you want me to write that?

