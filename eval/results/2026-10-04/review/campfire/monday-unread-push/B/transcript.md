**Agent:** I'll start by looking at how the app handles push notifications and unread state.

> Grep `{"pattern":"unread|Push|WebPush","glob":"*.{rb,yml,rake}","output_mode":"files_with_matches"}`

> Bash `{"command":"ls; ls app/jobs config lib/tasks 2>/dev/null; cat config/recurring.yml 2>/dev/null"}`

> Bash `{"command":"cat app/models/room/message_pusher.rb app/jobs/room/push_message_job.rb app/models/push.rb app/models/push/subscription.rb lib/web_push/notification.rb config/resque-pool.yml Procfile lib/tasks/resque.rake; grep -n \"unread\" -B2 -A8 app/models/membership.rb app/models/user.rb; grep -rn …`

> Bash `{"command":"cat lib/web_push/pool.rb config/initializers/web_push.rb; cat app/models/membership.rb | head -14; cat test/models/room/push_test.rb | head -40; grep -n \"queue\\|active_job\\|time_zone\" config/application.rb config/environments/production.rb; ls bin script; grep -rn \"Procfile\\|resque…`

> Bash `{"command":"grep -n \"push_subscriptions\\|memberships\\|scope\\|active\" app/models/user.rb; grep -n \"root\" config/routes.rb; sed -n 1,40p test/test_helper.rb; cat test/fixtures/memberships.yml | head -20; ls test/jobs test/models | head -30"}`

> Bash `{"command":"cat > app/models/push/unread_reminder.rb <<'EOF'\nclass Push::UnreadReminder\n  def self.deliver_all\n    User.active.joins(:push_subscriptions).distinct.find_each do |user|\n      new(user).deliver\n    end\n  end\n\n  attr_reader :user\n\n  def initialize(user)\n    @user = user\n  end…`

**Agent:** Bash is denied in this mode, so I'll write the files with the Write tool instead.

> Write `{"file_path":"[redacted]","content":"class Push::UnreadReminder\n  MAX_NAMED_ROOMS = 3\n\n  def self.delive…`

> Write `{"file_path":"[redacted]","content":"class Push::UnreadRemindersJob < ApplicationJob\n  def perform\n   …`

> Write `{"file_path":"[redacted]","content":"namespace :push do\n  desc \"Enqueue the unread rooms reminder push (schedule for Mon…`

> Write `{"file_path":"[redacted]","content":"require \"test_helper\"\n\nclass Push::UnreadReminderTest < Acti…`

**Agent:** I've written the reminder code, but none of it has been run, and nothing triggers it on Mondays yet. The shell was denied in this session, so I couldn't run the tests or the rake task.

**What I added**
- `app/models/push/unread_reminder.rb`: the logic. It covers active users who have a push subscription and at least one unread room that isn't hidden (hidden rooms are the `invisible` involvement). Each user gets one push, not one per room.
  - The title is "You have unread rooms". The body names up to three rooms, then "and N more".
  - A single unread room links straight to it. Several unread rooms link to the root page.
  - It goes out through the existing web push pool, so the badge count and the cleanup of expired subscriptions work as they do for message pushes.
- `app/jobs/push/unread_reminders_job.rb`: a job that calls `Push::UnreadReminder.deliver_all`.
- `lib/tasks/push.rake`: `bin/rails push:unread_reminders` enqueues that job.
- `test/models/push/unread_reminder_test.rb`: three tests. They cover one push per subscription when several rooms are unread, no push when nothing is unread, and ignoring hidden rooms.

**Scheduling**
The app has no scheduler: it runs on Resque plus resque-pool, with no cron or recurring-job setup. So you'll need to pick one:
- **Cron on the host:** `0 9 * * 1 cd /path/to/app && bin/rails push:unread_reminders`. This is the least work, but it runs on server time, not each user's time zone.
- **resque-scheduler:** a new gem and an extra process in the `Procfile`, so the schedule lives in the app.

I'd go with cron unless you want the schedule in the repo. Say which you prefer and I'll set it up.

To check it, run `bin/rails test test/models/push` and `bin/rails push:unread_reminders` once yourself. I also haven't committed anything.

