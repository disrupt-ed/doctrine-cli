# Gems

Many popular gems were added to Rails apps to fill gaps that Rails has since closed. Generic advice and older tutorials still recommend them.

## Gems Rails may have replaced

In a new app, or when adding the capability for the first time, prefer the Rails option. In an app that already uses the gem, keep using the gem unless the developer asks to migrate.

| Gem | Rails option | Notes |
|---|---|---|
| Sidekiq, Resque, Delayed Job | Solid Queue | Sidekiq still makes sense at very high job volume or if the app already runs Redis. |
| whenever, sidekiq-cron | `config/recurring.yml` | Only with Solid Queue. |
| redis-store for caching | Solid Cache | |
| Devise (for basic email and password login) | Authentication generator | Devise still helps for OAuth, confirmable, lockable and other extras. |
| rack-attack (per-action limits) | `rate_limit` | rack-attack still fits app-wide rules, blocklists and IP-based throttling. |
| dotenv | Credentials | dotenv is still common for local development. |
| Webpacker, Sprockets | Propshaft, importmap | |
| Active Model Serializers | `as_json`, Jbuilder | Check its health before using it in new code. |
| attr_encrypted | `encrypts` | |

## Gems that are often not needed

- **HTTP clients** (HTTParty, Faraday, RestClient): `Net::HTTP` is enough for simple calls. Check whether the app already has an HTTP client or wrapper first.
- **Decorators and presenters** (Draper): helpers and model methods usually do the job.
- **Form objects and service frameworks** (dry-rb, Trailblazer, Interactor): plain Ruby classes do the job without a new vocabulary.
- **Settings gems**: `config_for`, credentials or a `store` column.
- **Small utility gems** for a handful of lines of code.

## When a gem is the right call

Some needs are worth a well-maintained gem: payments, OAuth providers, PDF generation, complex authorization. Check the facts first, as in the engineering doctrine. Also check:

- it supports the app's Rails and Ruby versions;
- it doesn't monkey-patch Rails in surprising ways;
- the app doesn't already have another gem for the same job.
