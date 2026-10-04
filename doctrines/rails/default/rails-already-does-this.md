# Rails Already Does This

Check this list before writing your own version or adding a gem. The version is when the feature arrived. Check it against the app's Rails version.

## Models and data

| Need | Rails | Since |
|---|---|---|
| Clean up attribute values before saving | `normalizes` | 7.1 |
| Encrypt a column | `encrypts` (Active Record Encryption) | 7.0 |
| Secure random token on a record | `has_secure_token` | 5.0 |
| Expiring, purpose-specific tokens (password reset, email confirmation) | `generates_token_for` | 7.1 |
| Signed, tamper-proof IDs in URLs | `signed_id` / `find_signed` | 6.1 |
| Passwords | `has_secure_password`, `authenticate_by` | 3.1, 7.1 |
| A fixed set of values | `enum` | 4.1 |
| Several types sharing one table of common data | `delegated_type` | 6.1 |
| Store extra settings in a JSON column | `store`, `store_accessor` | 3.2 |
| Bulk insert or update | `insert_all`, `upsert_all` | 6.0 |
| Catch N+1 queries | `strict_loading` | 6.1 |
| Typed attributes on a non-database object | `ActiveModel::Attributes` | 5.2 |
| Order by a list of values | `in_order_of` | 7.0 |
| Run independent queries in parallel | `load_async` | 7.0 |

## Controllers and requests

| Need | Rails | Since |
|---|---|---|
| Rate limiting per action | `rate_limit` | 7.2 |
| Login, sessions, password reset | `bin/rails generate authentication` | 8.0 |
| HTTP caching (ETag, Last-Modified) | `fresh_when`, `stale?` | 3.x |
| Per-request global state (current user, account) | `ActiveSupport::CurrentAttributes` | 5.2 |
| Health check endpoint | `/up` (`Rails::HealthController`) | 7.1 |
| JSON responses | `render json:`, `as_json`, Jbuilder (in new apps by default) | |
| CSV, iCal or other formats | `respond_to` with a format and Ruby's standard libraries | |

## Background work and scheduling

| Need | Rails | Since |
|---|---|---|
| Background jobs without Redis | Solid Queue | 8.0 default |
| Recurring jobs | `config/recurring.yml` (Solid Queue) | 8.0 |
| Job dashboard | Mission Control – Jobs | |
| Long jobs that resume after a restart | Active Job Continuations | 8.1 |

## Caching and real-time

| Need | Rails | Since |
|---|---|---|
| Cache store without Redis | Solid Cache | 8.0 default |
| Fragment and low-level caching | `cache` helper, `Rails.cache.fetch` | |
| WebSockets without Redis | Solid Cable | 8.0 default |
| Live page updates | Turbo Streams (7.0 default), `broadcasts_refreshes` (turbo-rails 2.0) | 7.0 |

## Files, email and rich content

| Need | Rails | Since |
|---|---|---|
| File uploads, cloud storage, image variants | Active Storage | 5.2 |
| Rich text editing | Action Text | 6.0 |
| Receiving email | Action Mailbox | 6.0 |
| Sending email, previews | Action Mailer, mailer previews | |

## Operations

| Need | Rails | Since |
|---|---|---|
| Secrets | `bin/rails credentials:edit` | 5.2 |
| Per-environment YAML config | `config_for` | 4.2 |
| Error reporting interface | `Rails.error` | 7.0 |
| Instrumentation | `ActiveSupport::Notifications` | 3.0 |
| Structured events | `Rails.event` | 8.1 |
| Deployment | Kamal | 8.0 default |
| Local CI script | `bin/ci` | 8.1 |

## Not built in

Rails doesn't ship these. Check whether the app already solved them before adding a gem:

- pagination;
- authorization (policies, roles);
- full-text search beyond what the database offers;
- audit history of record changes;
- feature flags;
- OAuth login with external providers.
