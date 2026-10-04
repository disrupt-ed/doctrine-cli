**Agent:** I'll start by looking for the composer and any existing feature-flag mechanism.

> Grep `{"pattern":"feature.?flag|flipper|rollout|Feature\\.|feature_enabled","-i":true,"output_mode":"files_with_matches","head_limit":30}`

> Grep `{"pattern":"composer","-i":true,"output_mode":"files_with_matches","head_limit":30}`

> Grep `{"pattern":"feature.?flag|flipper|rollout","-i":true,"output_mode":"content","glob":"!*.lock"}`

> Read `{"file_path":"/private/tmp/claude-501/-Users-luc-Developer-luuuc-oss-doctrine/4644ffcf-8822-4777-a6cc-1b1a1cf0efe9/scratchpad/evalwork/campfire-composer-feature-flag-baseline-947558210/app/views/rooms/show/_composer.html.erb"}`

> Bash `{"command":"git log --oneline -8 && ls app/models config; grep -n \"gem \" Gemfile"}`

**Agent:** I haven't changed anything, because I can't tell what the 10% should gate.

- **No flag system:** The repo has no feature-flag or percentage-rollout mechanism. There's no Flipper or similar gem, and the only "rollout" hit is a sound effect. A 10% rollout needs one to be built.
- **No old and new composer:** There's one composer, in `app/views/rooms/show/_composer.html.erb` with `composer_controller.js`. Recent commits (Lexxy replacing Trix, composer drafts) changed it for everyone. I found no second implementation to switch between.

Two questions:

1. **What counts as the "new composer"?** Is it a change that isn't in the repo yet, or something already merged that you want to hide behind a flag? If it's already merged, I need to know which parts, and what users in the other 90% should see.
2. **How should the 10% be chosen?**
   - **Deterministic bucket on user id:** `user.id % 100 < 10`, stored as a constant or a setting. It's simple, needs no new dependency, and keeps each user in the same bucket. Changing the percentage means a deploy.
   - **Flipper (`percentage_of_actors`):** The percentage can change at runtime without a deploy. It adds a gem and a migration.
   - **Per-account flag:** This is the better fit if you want whole rooms or accounts to see the same composer.

My default would be the deterministic bucket on user id, a `User#new_composer?` helper, and a branch in `_composer.html.erb`. Tell me what the new composer is and whether that default works, and I'll implement it.

