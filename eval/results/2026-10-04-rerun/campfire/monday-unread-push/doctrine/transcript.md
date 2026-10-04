**Agent:** I'll start by looking at how the app handles unread rooms, push notifications, and recurring jobs.

> Grep `{"pattern":"unread|push|Push","glob":"*.{rb,yml}","output_mode":"files_with_matches"}`

> Bash `{"command":"cat config/recurring.yml; ls app/jobs app/models app/models/concerns; grep -E \"^    (rails|solid_queue|web-push) \" Gemfile.lock; cat .ruby-version"}`

> Bash `{"command":"cat app/models/room/message_pusher.rb app/jobs/room/push_message_job.rb app/models/push.rb app/models/push/subscription.rb app/channels/unread_rooms_channel.rb; grep -n \"unread\\|push\\|def \\|scope\" app/models/user.rb app/models/membership.rb; grep -n -B2 -A12 'create_table \"membersh…`

> Bash `{"command":"cat config/resque-pool.yml; grep -rniE \"cron|schedul\" --include=* -l . 2>/dev/null | grep -v -E \"^./(node_modules|log|tmp|\\.git|vendor)/\" | head; ls; cat lib/web_push/pool.rb | head -50; cat app/jobs/application_job.rb; grep -n \"badge\\|title\\|body\" lib/web_push/notification.rb; …`

> Bash `{"command":"cat Procfile; ls lib/tasks script docs 2>/dev/null; grep -rniE \"cron|schedul\" Dockerfile script docs README.md bin Procfile 2>/dev/null | head; grep -n \"unread\" -r app/models | head -20"}`

> Bash `{"command":"sed -n 1,50p app/models/user.rb; cat lib/tasks/resque.rake; grep -n \"root\" config/routes.rb; grep -n \"unread\" -B3 -A10 test/models/user_test.rb | head -30; grep -n \"kevin\\|david\" test/fixtures/memberships.yml | head"}`

> Bash `{"command":"cat > app/models/push/unread_digest.rb <<'EOF'\nclass Push::UnreadDigest\n  def push\n    users_with_unread_rooms.find_each do |user|\n      count = user.memberships.visible.unread.count\n      next if count.zero?\n\n      Rails.configuration.x.web_push_pool.queue(payload_for(count), use…`

**Agent:** I'm writing the digest, job, rake task and test with the file tools, because Bash was denied.

> Write `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/campfire-monday-unread-push-doctrine-1592305355/app/models/push/unread_digest.rb","content":"class Push::UnreadDigest\n  def push\n    users_with_unread_rooms.find_e…`

> Write `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/campfire-monday-unread-push-doctrine-1592305355/app/jobs/push/unread_digest_job.rb","content":"class Push::UnreadDigestJob < ApplicationJob\n  def perform\n    Push:…`

> Write `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/campfire-monday-unread-push-doctrine-1592305355/lib/tasks/push.rake","content":"namespace :push do\n  desc \"Enqueue the unread rooms push notification (schedule for…`

> Write `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/campfire-monday-unread-push-doctrine-1592305355/test/models/push/unread_digest_test.rb","content":"require \"test_helper\"\n\nclass Push::UnreadDigestTest < ActiveSu…`

**Agent:** The unread-rooms push is written, but nothing runs it on Mondays yet, and I haven't run any of the code. Bash was denied partway through, so the new tests are unrun.

**What I added**
- `app/models/push/unread_digest.rb` finds active users who have a push subscription and at least one unread, visible room. Each gets one push saying how many rooms are unread (for example "You have 3 rooms with unread messages"), linking to the root path. It sends through the same web push pool that message notifications use.
- `app/jobs/push/unread_digest_job.rb` is a thin job around it.
- `lib/tasks/push.rake` adds `bin/rails push:unread_digest`, which enqueues that job.
- `test/models/push/unread_digest_test.rb` covers the main case (a user with an unread room gets pushed) and the skips (invisible rooms, deactivated users). I never ran it, so run `bin/rails test test/models/push/unread_digest_test.rb` first.

**Scheduling is the open question.** The app uses Resque and has no scheduler: no `recurring.yml`, no resque-scheduler, no cron in the Procfile or Dockerfile. I didn't want to add a new gem or extra process without asking, so the rake task is the hook. You can run it from a Monday cron entry, for example `0 9 * * 1 cd /app && bin/rails push:unread_digest`. The alternative is adding `resque-scheduler` and a schedule file, which means a new dependency and a new process in the Procfile. Tell me if you'd rather go that way and I'll run the dependency review first.

**Assumptions**
- "Monday morning" is whenever the cron fires, in the server's timezone. Users aren't targeted by their own timezone.
- There's no per-user opt-out beyond not having a push subscription.

