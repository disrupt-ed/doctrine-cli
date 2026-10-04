# Project Manifest

A project's Doctrine setup lives in `.doctrine/doctrine.yml`. The developer owns this file. The CLI writes it once on `doctrine init` or `doctrine install`, and only edits `extends` and `doctrines` afterwards (`add`, `remove`, `update`), keeping comments.

```yaml
version: 1         # manifest format
doctrines: 0.1     # official doctrine release

extends:
  - rails/default
  - go/default

engineering:       # optional
  kiss: strong
  yagni: strong
  dependencies: conservative
  abstractions: late
```

| Key | Required | Meaning |
|---|---|---|
| `version` | yes | Manifest format. Always `1` for now. |
| `doctrines` | yes | The official doctrine release the project is pinned to. |
| `extends` | yes | Top-level doctrines. Parents come along automatically. |
| `engineering` | no | Preferences that adjust the defaults. |

## Engineering preferences

| Key | Values | Default |
|---|---|---|
| `kiss` | `strong`, `normal` | `strong` |
| `yagni` | `strong`, `normal` | `strong` |
| `dependencies` | `conservative`, `balanced`, `open` | `conservative` |
| `abstractions` | `late`, `balanced` | `late` |

Keys that are present are written into the generated output as project preferences. Unknown keys or values are an error.

## Local overrides

Markdown files in `.doctrine/local/` are the project's own rules. See [doctrine.md](doctrine.md#overrides).
