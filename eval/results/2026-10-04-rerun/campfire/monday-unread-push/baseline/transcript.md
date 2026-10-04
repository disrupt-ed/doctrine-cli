**Agent:** I'll start by looking at how the app handles unread rooms and push notifications.

> Grep `{"pattern":"unread|WebPush|web_push","glob":"!{node_modules,vendor,log}/**","output_mode":"files_with_matches","-i":true}`

> Bash `{"command":"ls; ls app/jobs app/models config; cat config/recurring.yml 2>/dev/null"}`

> Bash `{"command":"cat app/models/room/message_pusher.rb app/models/push/subscription.rb lib/web_push/notification.rb config/resque-pool.yml Procfile; grep -n -i \"unread\" -B2 -A8 app/models/membership.rb app/models/room.rb; ls app/jobs/room lib config/initializers; cat app/models/push.rb; grep -rn -i \"r…`

> Bash `{"command":"cat app/jobs/room/push_message_job.rb lib/web_push/pool.rb config/initializers/web_push.rb; ls lib/tasks; cat test/models/room/push_test.rb | head -40; grep -n \"cron\\|schedul\" -ri README.md docs Dockerfile script bin | head; grep -n \"memberships\\|push_sub\" app/models/user.rb"}`

> Bash `{"command":"cat lib/tasks/resque.rake; grep -n \"wait_for_web_push\\|stub_web_push\" -A8 test/test_helper.rb | head -40; grep -n \"unread\" test/fixtures/memberships.yml; cat app/models/user.rb | sed -n 1,30p; grep -rn \"web_push\\|Procfile\\|cron\" Dockerfile bin/* | head"}`

> Bash `{"command":"grep -rn \"web_push_pool\\|wait_for_web_push\" test/test_helper.rb test/test_helpers | head; grep -n \"unread\" -B3 test/fixtures/*.yml | head; cat bin/boot | sed -n 35,70p; grep -n \"scope\\|def \" app/models/user/*.rb | grep -i \"active\\|bann\""}`

> Bash `{"command":"sed -n 15,60p test/test_helper.rb; cat test/fixtures/memberships.yml | head -30; grep -n \"root\" config/routes.rb"}`

**Agent:** I've finished reading the push code, and I'm now writing the reminder model, job, rake task and a test. The app has no scheduler, so Monday's trigger will be a rake task that cron runs.

> Bash `{"command":"grep -rn \"def wait_for_web_push\" -A8 test | head -15; grep -n \"active\\b\" app/models/user.rb | head -3"}`

> Bash `{"command":"sed -n 60,90p test/models/room/push_test.rb\nmkdir -p app/models/push\ncat > app/models/push/unread_reminder.rb <<'EOF'\nclass Push::UnreadReminder\n  def deliver\n    unread_counts.each do |user_id, count|\n      payload = build_payload(count)\n      Rails.configuration.x.web_push_pool.…`

**Agent:** Bash was denied, so I'll use the file tools instead.

> Write `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/campfire-monday-unread-push-baseline-125691061/app/models/push/unread_reminder.rb","content":"class Push::UnreadReminder\n  def deliver\n    unread_counts.each do |u…`

> Write `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/campfire-monday-unread-push-baseline-125691061/app/jobs/push/unread_reminder_job.rb","content":"class Push::UnreadReminderJob < ApplicationJob\n  def perform\n    Pu…`

> Write `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/campfire-monday-unread-push-baseline-125691061/lib/tasks/push.rake","content":"namespace :push do\n  desc \"Enqueue the unread rooms reminder. Schedule for Monday mo…`

> Read `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/campfire-monday-unread-push-baseline-125691061/test/models/room/push_test.rb","offset":60}`

> Write `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/campfire-monday-unread-push-baseline-125691061/test/models/push/unread_reminder_test.rb","content":"require \"test_helper\"\n\nclass Push::UnreadReminderTest < Activ…`

**Agent:** I wrote the Monday unread reminder, but I haven't run any of it. Bash was denied in this session, so the new tests are unrun and I haven't confirmed the fixtures (e.g. that `kevin` has a push subscription) match what they assume.

**What it does**
- **`Push::UnreadReminder`** (`app/models/push/unread_reminder.rb`) finds active users who have a push subscription and at least one unread, non-invisible room. It sends each of them one push saying "You have N rooms with unread messages", which opens `root_path`. It goes to all of that user's devices, through the existing `WebPush::Pool`, and the app-icon badge still gets the unread count.
- **`Push::UnreadReminderJob`** (`app/jobs/push/unread_reminder_job.rb`) runs it in the background.
- **`rake push:unread_reminders`** (`lib/tasks/push.rake`) enqueues that job.
- **`test/models/push/unread_reminder_test.rb`** covers four cases:
  - a user with an unread room gets a push;
  - a user with no unread rooms gets none;
  - a user with unread rooms but no subscription gets none;
  - an invisible membership doesn't count.

**Scheduling**
The app has no scheduler: resque is there but resque-scheduler isn't, and there's no cron setup. So nothing fires on Monday until you schedule the rake task, for example `0 9 * * 1 cd /app && bin/rails push:unread_reminders` in the host's crontab.

**Time zone**
That runs at 9:00 server time for everyone. Users have no time zone stored, so "Monday morning" can't be per-user without adding one.

**Decision for you**
If you'd rather have it built into the app, I can add `resque-scheduler` and a schedule entry, or a recurring process in the `Procfile`. That adds a dependency and a process to run, so I held off. Tell me which you prefer.

To check it, run `bin/rails test test/models/push/unread_reminder_test.rb`.

