**Agent:** I'll start by checking whether the app already has feature flags or a composer.

> Grep `{"pattern":"feature.?flag|flipper|rollout|Feature\\.|enabled_for|beta","-i":true,"glob":"!{log,node_modules,vendor,tmp}/**","output_mode":"content","head_limit":40}`

> Grep `{"pattern":"composer","-i":true,"path":"app","output_mode":"files_with_matches"}`

> Bash `{"command":"grep -E \"^    (rails|flipper|redis|solid_cache|solid_queue) \" Gemfile.lock; cat .ruby-version 2>/dev/null; git log --oneline -8"}`

> Bash `{"command":"cat app/views/rooms/show/_composer.html.erb; git show --stat HEAD~2 | head -20; ls app/views/rooms/show app/views/messages; grep -n \"composer\" app/views/rooms/show.html.erb app/views/messages/edit.html.erb app/helpers/*.rb; ls app/models/concerns app/models/user; grep -n \"cache_store\…`

**Agent:** I haven't changed anything yet, because there's nothing in the repo for a rollout to switch on.

- **No new composer exists.** The only composer is `app/views/rooms/show/_composer.html.erb`, with `composer_controller.js` behind it. Recent commits (#283 drafts, #284 mention menu) only touch that one. There's no second composer partial, branch or stub.
- **No feature-flag mechanism exists.** Rails doesn't ship one, and the app has no Flipper or similar. The app runs Redis and Resque.

**Question:** where is the new composer? Is it on another branch or PR, or do you want me to build it?

Once there's a new composer to gate, my plan for the rollout is a small one:

- **Bucketing:** a `users.composer_rollout` check in a model concern. It would hash the user ID into a stable bucket from 0 to 99 and enable the new composer for buckets below 10. A user always sees the same composer, and raising the percentage later only adds users.
- **Rendering:** `rooms/show.html.erb` would pick the old or new composer partial based on that check.
- **Percentage:** it would come from an environment variable or credential, so you can change it without a deploy.
- **Overrides:** an explicit per-user override wouldn't be in v1. I'd only mention it as a likely next need.
- **No gem:** I wouldn't add Flipper for one flag. It's worth revisiting if you expect more flags.

Is that approach okay? And do you want the percentage to apply per user (my assumption) or per account?

