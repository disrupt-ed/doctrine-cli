# The Rails Way

Rails is opinionated on purpose. Working with its conventions gives less code and code any Rails developer can read. Working against them costs more than it seems.

## Check the version first

Read the Rails version in `Gemfile.lock` before using a feature. Many built-in features are recent. Don't use one the app's version doesn't have, and don't add a gem for something the app's version already does.

## Follow the app before the default

This doctrine describes plain Rails. Many apps have their own conventions: service objects, RSpec, a component library, a custom pattern for jobs. Follow the app. Mention a Rails-native alternative only if it would clearly be simpler for the task at hand.

## Models

- Business logic lives in models. Use concerns to share behavior between models or to split a large model by topic.
- Use callbacks for things that must always happen with the record. Don't use them for side effects that only some callers want, like sending an email.
- Don't add a service object, interactor or command layer to an app that doesn't have one. A model method or a plain Ruby class in `app/models` usually does the job.
- Use validations, scopes, enums and associations before writing the same thing by hand.

## Controllers and routes

- Stick to the seven REST actions. When an action doesn't fit, it's often a new resource: `resource :archive` beats `post :archive`.
- Keep controllers thin. Load, authorize, call the model, render.
- Use `params.expect` (Rails 8) or `params.require(...).permit(...)` for strong parameters.
- Use `before_action` for shared setup, not for business logic.

## Views and the front end

- Server-rendered HTML first. Use Turbo Frames and Turbo Streams for partial updates, and Stimulus for small bits of behavior.
- Reach for a JavaScript framework only if the app already uses one or the feature truly needs heavy client-side state.
- Use partials and helpers. Add a component library only if the app already has one.
- Use the asset setup the app already has (Propshaft or Sprockets, importmap or a bundler). Don't switch it as a side effect.

## Jobs

- Use Active Job with the app's backend. Pass records, not their attributes, so jobs load fresh data.
- Make jobs safe to run twice. Retries happen.
- Schedule recurring work with the backend's scheduler (Solid Queue's `config/recurring.yml`), not system cron, if the app uses Solid Queue.

## Tests

- Use the app's framework (Minitest or RSpec) and its setup (fixtures or factories). Don't add a second one.
- Prefer model and request tests. Add system tests for flows that need a browser.
- Don't mock Active Record. Use the test database.

## Security

Rails protects you by default. Don't turn the protection off.

- Keep CSRF protection on. Escape output; use `raw` and `html_safe` only on content you control.
- Use parameterized queries. Never put user input into SQL strings.
- Store secrets in credentials or environment variables, never in the code.
- Check authorization on every action that loads a record by ID.
