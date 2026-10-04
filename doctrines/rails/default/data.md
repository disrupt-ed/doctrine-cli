# Data and Migrations

## Migrations

- One purpose per migration. Generate them with `bin/rails generate migration`.
- Add database constraints alongside validations: `null: false`, foreign keys, unique indexes. Validations alone allow bad data from races and other code paths.
- Index foreign keys and columns you filter or sort by.
- Write reversible migrations. Use `change` when Rails can reverse it, `up` and `down` otherwise.
- Don't use application models inside migrations. They change later; the migration doesn't.

## Changing live data

On a table that's already big in production:

- Backfill data in a separate step or job, in batches (`in_batches`, `find_each`), not in the schema migration.
- Add a column, deploy, backfill, then add the constraint.
- Removing a column takes two deploys: first `ignored_columns`, then the migration.
- Say in your summary when a change needs a backfill or a specific deploy order.

## Queries

- Use `includes` or `preload` when you loop over associations.
- Use `find_each` to go through many records.
- Use `exists?`, `pluck` and `count` instead of loading records you don't need.
- Use the database's features (JSON columns, full-text search, constraints) through the database the app already has. Don't add a second datastore for one feature.
