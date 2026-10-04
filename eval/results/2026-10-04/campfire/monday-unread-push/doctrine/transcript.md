**Agent:** I'll start by looking at how the app handles push notifications, unread rooms, and recurring jobs.

> Bash `{"command":"grep -n \"^    rails \\|solid_queue\\|web-push\\|resque\\|sidekiq\" Gemfile.lock | head; cat config/recurring.yml; ls app/jobs app/models app/models/push app/models/room 2>/dev/null; grep -rn \"unread\" app/models | head -30"}`

> Bash `{"command":"cat app/models/push.rb app/models/push/subscription.rb app/models/room/message_pusher.rb app/models/membership.rb; ls app/jobs/* config; cat app/jobs/room/*.rb app/jobs/application_job.rb; grep -rn \"resque\\|Resque\\|schedule\" config lib Procfile* 2>/dev/null | head -20; ls test test/j…`

> Bash `{"command":"cat Procfile config/resque-pool.yml; ls lib lib/tasks config/deploy* 2>/dev/null; grep -rn \"web_push_pool\" config lib app | head; cat test/models/room/push_test.rb | head -50; ls test/lib test/fixtures; cat test/fixtures/memberships.yml | head -20; grep -rn \"cron\\|Cron\" --include=* …`

> Bash `{"command":"cat config/initializers/web_push.rb lib/tasks/*.rake; ls lib/tasks lib/web_push; ls -a . | head -40; grep -n \"unread_at\" db/schema.rb; cat test/test_helpers/*push* 2>/dev/null | head -30; grep -rn \"web_push_delivery\" test/test_helpers | head -3"}`

> Bash `{"command":"cat lib/web_push/notification.rb; grep -rn \"wait_for_web_push\\|stub_web_push\" test/test_helper.rb test/test_helpers; cat bin/start-app 2>/dev/null; ls docs; cat app/models/user.rb | head -30; ls app/models/user"}`

> Bash `{"command":"cat lib/web_push/pool.rb | head -40; grep -n \"root\\|resources :rooms\" config/routes.rb | head; grep -n \"unread\" test/fixtures/*.yml; sed -n 1,30p test/test_helpers/dns_test_helper.rb; grep -n \"Rake\\|cron\\|deploy\" -i docs/self-hosting.md | head"}`

**Agent:** I've read enough of the codebase to start writing. The app runs Resque with no scheduler, so I'll build the notification and job and then work out how to trigger it weekly.

> Bash `{"command":"mkdir -p app/jobs/push && cat > app/models/push/unread_reminder.rb <<'EOF'\nclass Push::UnreadReminder\n  def push\n    unread_counts.each do |user_id, count|\n      payload = {\n        title: \"Unread rooms\",\n        body: \"You have #{count} #{\"room\".pluralize(count)} with unread …`

**Agent:** Bash is denied in this mode, so I'll create the files with the Write tool instead.

> Write `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/campfire-monday-unread-push-doctrine-3312298145/app/models/push/unread_reminder.rb","content":"class Push::UnreadReminder\n  def push\n    unread_counts.each do |use…`

> Write `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/campfire-monday-unread-push-doctrine-3312298145/app/jobs/push/unread_reminder_job.rb","content":"class Push::UnreadReminderJob < ApplicationJob\n  def perform\n    P…`

> Write `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/campfire-monday-unread-push-doctrine-3312298145/lib/tasks/push.rake","content":"namespace :push do\n  desc \"Remind users with unread rooms (schedule weekly, Monday …`

> Write `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/campfire-monday-unread-push-doctrine-3312298145/test/models/push/unread_reminder_test.rb","content":"require \"test_helper\"\n\nclass Push::UnreadReminderTest < Acti…`

**Agent:** I've written the Monday push, but I couldn't run the tests, and nothing triggers it weekly yet. Bash was denied in this session, so the new test file hasn't run and the Ruby hasn't been executed.

**What I added**
- `app/models/push/unread_reminder.rb` finds users who have a push subscription and at least one unread room that isn't invisible. It sends each of them one push saying "You have N room(s) with unread messages", which opens the root path. It reuses the existing `web_push_pool` and the unread counts on `Membership`.
- `app/jobs/push/unread_reminder_job.rb` is a thin Active Job wrapper around that class.
- `lib/tasks/push.rake` adds `bin/rails push:unread_reminders`, which enqueues the job.
- `test/models/push/unread_reminder_test.rb` covers three cases: users with unread rooms are notified, users with none aren't, and invisible rooms are ignored.

**Scheduling is not set up.** The app runs Resque with no scheduler, and the Procfile only starts the web server, Redis and the workers. So nothing fires the task on Mondays yet.

Campfire is a single-container, self-hosted app, so I'd run the rake task from the host's cron or whatever scheduler runs the container. For example: `0 9 * * 1 bin/rails push:unread_reminders`.

The alternative is adding `resque-scheduler`. That is a new dependency and a new process, so I didn't add it without your say-so. Tell me if you want it.

**Things to know**
- "Monday morning" depends on the server's time zone. Users in other time zones will all get the push at the same moment.
- The `badge` count on the push is set by the existing `Push::Subscription#notification`, so it needs no change.

To check it, run `bin/rails test test/models/push test/models/room/push_test.rb`.

