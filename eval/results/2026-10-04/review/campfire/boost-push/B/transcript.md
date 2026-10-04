**Agent:** I'll start by looking at how boosts and push notifications work in this codebase.

> Grep `{"pattern":"Boost|push|Push","path":"[redacted]","glob":"*.rb","output_mode":"files_with_matches"}`

> Read `{"file_path":"[redacted]"}`

> Read `{"file_path":"[redacted]"}`

> Read `{"file_path":"[redacted]"}`

> Read `{"file_path":"[redacted]"}`

**Agent:** Let me see how the message push is triggered, plus the by_bots controller, the message model, and existing tests.

> Grep `{"pattern":"PushMessageJob|MessagePusher|boosts","path":"[redacted]","glob":"*.{rb,yml}","output_mode":"content","head_limit":0}`

> Read `{"file_path":"[redacted]"}`

> Read `{"file_path":"[redacted]","offset":60}`

> Read `{"file_path":"[redacted]"}`

> Read `{"file_path":"[redacted]"}`

> Bash `{"command":"cat app/models/membership.rb app/models/push/subscription.rb; grep -n \"boost\" -A8 test/fixtures/boosts.yml | head -30; grep -n \"push_subscriptions\" test/fixtures/*.yml | head; cat test/fixtures/push/subscriptions.yml | head -30"}`

> Bash `{"command":"mkdir -p app/models/boost app/jobs/boost test/models/boost\ncat > app/models/boost/pusher.rb <<'EOF'\nclass Boost::Pusher\n  attr_reader :boost\n\n  delegate :message, :booster, to: :boost\n  delegate :room, to: :message\n\n  def initialize(boost:)\n    @boost = boost\n  end\n\n  def pus…`

**Agent:** Bash is denied, so I'll create the files with the file tools. I won't be able to run the tests.

> Write `{"file_path":"[redacted]","content":"class Boost::Pusher\n  attr_reader :boost\n\n  delegate :message, :booster, to: :boost\…`

> Write `{"file_path":"[redacted]","content":"class Boost::PushJob < ApplicationJob\n  def perform(boost)\n    Boost::Pusher.new(boos…`

> Edit `{"replace_all":false,"file_path":"[redacted]","old_string":"  scope :ordered","new_string":"  after_create_commit -> { Boost::PushJ…`

> Write `{"file_path":"[redacted]","content":"require \"test_helper\"\n\nclass Boost::PushTest < ActiveSupport::TestCase\n  inclu…`

**Agent:** I added a push notification that goes to a message's author when someone boosts it. I couldn't run anything, because Bash was denied here. The code and tests are untested.

It follows the existing `Room::MessagePusher` / `Room::PushMessageJob` pattern:

- **`app/models/boost.rb`:** an `after_create_commit` enqueues `Boost::PushJob`. It skips the job when you boost your own message.
- **`app/jobs/boost/push_job.rb`:** a thin job that calls the pusher.
- **`app/models/boost/pusher.rb`:** builds the payload and queues it on the existing `web_push_pool`.
  - The title is the room name, or the booster's name in a direct room.
  - The body is "Name boosted your message: 👍".
  - The path is the room.
  - Only the message author's subscriptions are used, and only if their membership is `mentions` or `everything` and they aren't currently connected to the room. That is how new-message pushes already behave.
- **`test/models/boost/push_test.rb`:** covers four cases. The author is notified, a self-boost sends nothing, an author connected to the room gets nothing, and an author with `nothing` involvement gets nothing.

Decision for you: authors with `mentions` involvement also get boost notifications, because I treated a boost as directed at them. If you want boosts only for `everything`, change `%w[ mentions everything ]` in `Boost::Pusher`.

Bot boosts go through `Messages::Boosts::ByBotsController`, which also calls `create!`, so they trigger the push too. No migration or deploy step is needed.

Please run `bin/rails test test/models/boost test/models/room/push_test.rb test/controllers/messages`. The test setup is copied from `room/push_test.rb`.

