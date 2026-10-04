**Agent:** I'll start by looking at how the composer and any feature flags are set up in this repo.

> Grep `{"pattern":"feature.?flag|flipper|rollout|Feature\\.|feature_enabled|enabled_for","-i":true,"output_mode":"files_with_matches","head_limit":30}`

> Grep `{"pattern":"composer","-i":true,"output_mode":"files_with_matches","head_limit":40}`

> Read `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/campfire-composer-feature-flag-baseline-1355062977/app/views/rooms/show/_composer.html.erb"}`

> Grep `{"pattern":"feature.?flag|flipper|rollout","-i":true,"output_mode":"content","glob":"{CONTRIBUTING.md,app/models/sound.rb,app/javascript/models/client_message.js}"}`

> Grep `{"pattern":"composer","-i":true,"output_mode":"content","glob":"{app/views/rooms/show.html.erb,app/helpers/*.rb,app/views/messages/edit.html.erb}"}`

> Bash `{"command":"grep -n -i \"gem \" Gemfile | head -50; git log --oneline -8 -- app/views/rooms/show/_composer.html.erb"}`

**Agent:** I haven't changed anything, because the repo has nothing I can roll out to 10% of users yet.

- **No feature flag system.** There's no Flipper-style gem in the `Gemfile`, and no flag or percentage-rollout code in `app/`. The only hits for "rollout" were a sound name and a CONTRIBUTING mention.
- **No "new" composer to gate.** There is one composer, `app/views/rooms/show/_composer.html.erb`, which uses Lexxy. Trix was replaced with Lexxy in 47bc5f5 and draft saving landed in #283. No old and new versions exist side by side. So I don't know which change you mean by "the new composer".

Two questions before I build anything:

1. **Which change is the new composer?** Is it an unmerged branch or PR, or something you want me to build? If it's existing work, point me at it. If it's the Lexxy composer that's already in `main`, say so. Gating that would mean bringing back a Trix fallback, which is a much bigger job.
2. **Is a small homegrown flag OK, or do you want a gem like Flipper?** My default would be a small `Feature` model with no new dependency. It would hash `user.id` with the flag name into a stable bucket from 0 to 99 and enable the flag when the bucket is below the configured percentage. The percentage would come from an env var or setting so you can raise it without a deploy. `rooms/show.html.erb` would then render the new or old composer per user.

Once you answer, I'll implement it and add tests for the bucketing and for both composer paths.

