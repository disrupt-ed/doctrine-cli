# Ruby

Ruby is built for programmer happiness: code that reads like what it does. Write the Ruby an experienced Rubyist would expect.

## Check the version first

Read `.ruby-version` or the `ruby` line in the `Gemfile` before using a newer feature.

## Standard library first

Ruby ships a lot. Check it before writing your own or adding a gem:

| Need | Ruby |
|---|---|
| CSV | `csv` |
| JSON | `json` |
| HTTP requests | `net/http`, `open-uri` for simple reads |
| URLs | `URI` |
| Dates and times | `Time`, `Date` |
| Hashing, HMAC, random tokens | `digest`, `openssl`, `securerandom` |
| Base64 | `base64` |
| Sets | `Set` |
| Immutable value objects | `Data.define` (3.2) |
| Simple structs | `Struct` |
| Temp files and directories | `tempfile`, `tmpdir` |
| Running commands | `open3` |
| Benchmarks | `benchmark` |
| Logging | `logger` |

Since Ruby 3.4, some of these (for example `csv` and `base64`) are bundled gems: they need a line in the `Gemfile`, but they are still the standard library, maintained with Ruby. That isn't a new dependency.

## Idioms

- Use `Enumerable` (`map`, `select`, `each_with_object`, `group_by`, `sum`, `tally`) instead of manual loops and accumulators.
- Use guard clauses and early returns instead of nested conditionals.
- Use keyword arguments when a method takes more than two or three arguments or a boolean.
- Use `case`/`in` pattern matching to take apart hashes and arrays when it reads better than chained lookups.
- Use `fetch` when a missing key is a bug.
- Use `&.` sparingly. If you need it everywhere, the nil is the problem.
- Raise specific errors. Rescue only what you can handle, never a bare `rescue Exception`.

## Design

- Plain Ruby objects and modules before frameworks and DSLs.
- Duck typing over type checks. Don't check `is_a?` to decide behavior.
- Avoid metaprogramming (`method_missing`, `define_method`, `instance_variable_get`) unless it removes a lot of repetition and stays readable.
- Don't monkey-patch core classes. Use refinements or a plain method if you must.
- Keep classes small, but don't split a class just to make it smaller.

## Style

Follow the project's formatter and linter (RuboCop, Standard). Don't reformat code you're not changing.
