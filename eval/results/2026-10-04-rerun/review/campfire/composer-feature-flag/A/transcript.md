**Agent:** I'll start by checking whether the app already has a feature-flag mechanism and where the composer lives.

> Grep `{"pattern":"flipper|feature_flag|feature flag|rollout|Feature\\.|flag","-i":true,"glob":"!{node_modules,log,tmp,vendor}/**","output_mode":"files_with_matches","head_limit":40}`

> Grep `{"pattern":"composer","-i":true,"path":"app","output_mode":"files_with_matches"}`

> Bash `{"command":"grep -E \"^    (rails|flipper|rollout) \" Gemfile.lock; cat .ruby-version 2>/dev/null; ls"}`

**Agent:** No existing flag mechanism. Next I'll read the composer view, the user model and the related tests.

> Bash `{"command":"cat app/views/rooms/show/_composer.html.erb; grep -n -i composer app/views/rooms/show.html.erb app/views/messages/edit.html.erb app/helpers/*.rb; cat app/models/user.rb; ls app/models app/models/concerns test test/models; grep -n \"users\" -A25 db/schema.rb | head -50; git log --oneline …`

> Bash `{"command":"git branch -a; git log --oneline -15; ls app/views/rooms/show app/views/messages app/javascript/controllers | head -80; grep -rniE \"new.composer|composer.*(v2|new)\" app docs CONTRIBUTING.md | head"}`

**Agent:** I haven't changed anything yet, because the repo has no "new composer" to roll out.

- **One composer only.** `app/views/rooms/show/_composer.html.erb` is the only message composer, and it's used in `rooms/show.html.erb:24`. There is no second version, branch, or leftover variant in the code or git history.
- **Lexxy and drafts are already live.** The Lexxy editor (#224) and composer drafts (#283) are both merged into `main`, so every user already has them.
- **No flag mechanism.** There is no feature-flag gem or helper, and nothing on `User` or in the schema to hang a percentage on. Rails doesn't ship flags either.

**What I'd build once I know what to gate.** I'd go with the simplest option: a plain `Feature` model in `app/models`, with no gem. It would decide by hashing the user's id against a percentage, so a user stays in or out and raising 10% to 50% only adds people. The percentage would come from an env var, or from `config_for` if you want per-environment settings. The partial would pick the new or old composer with `<%= render Feature.enabled?(:new_composer, Current.user) ? ... %>`. That needs no migration and no new dependency, and I'd add model tests for the bucketing. Flipper would add a UI and per-user overrides, but it needs tables and a gem for what is one flag today.

**Which of these do you mean?**
1. The new composer is on another branch or in your head and not in this repo yet. Where is it, or what should it do differently?
2. You want the current composer, with Lexxy and drafts, gated to 10% of users. That means shipping the flag with the old Trix-style composer as the fallback, which no longer exists in the codebase.

