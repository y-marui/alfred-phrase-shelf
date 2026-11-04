# Architecture

## Overview

An Alfred Workflow (Go): `cmd/phrase-shelf-alfred` is the single universal
(amd64+arm64) binary `workflow/info.plist` invokes. The "snip" Script
Filter node runs it with the query as `$1`; the workflow's native Copy to
Clipboard action node (auto-paste enabled) copies and pastes whichever
Stage-2 item's `arg` the user picks — the binary never touches the
clipboard or simulates keystrokes itself.

```
Alfred
  │  keyword "snip" + query
  ▼
cmd/phrase-shelf-alfred/main.go  ← Alfred boundary; argv dispatch + recover() only
  │
  ▼
internal/phraseshelfcmd.List()  ← builds the Script Filter response (Stage 1 / Stage 2)
  │
  ▼
internal/phraseshelf            ← domain logic: query splitting, fuzzy matching,
  │                                category extraction, and the Store that calls
  │                                osascript
  ▼
internal/scriptfilter.Response  ← Script Filter JSON types + writer

Alfred (user presses Enter on a Stage-2 item)
  │  selected snippet's text, as {query}
  ▼
alfred.workflow.output.clipboard (native, autopaste enabled — unchanged by this binary)
```

## Layers

### `cmd/phrase-shelf-alfred/`

The only binary Alfred executes. Dispatches argv to
`internal/phraseshelfcmd`, recovers from any panic into a visible Script
Filter error item, and writes the response — no business logic beyond
that.

### `internal/scriptfilter/`

Script Filter JSON types (`Item`, `Response`, `Icon`, `Mod`) and the writer
that encodes them to stdout. Alfred-independent, reused unchanged from the
template.

### `internal/phraseshelf/`

Domain logic — Alfred-independent, no `internal/scriptfilter` import.

- `Snippet` — the data shape `list snippets` returns (`id`/`name`/`text`/`category`).
- `Store` — the interface fetching snippets; tests use a fake, never the
  real app.
- `OsascriptStore` — the real `Store`, shelling out to `osascript` to run
  `PhraseShelf.sdef`'s `list snippets [category "..."]` command. Escapes
  the category into the AppleScript string literal
  (`escapeAppleScriptString`) and separates JSON decoding
  (`parseSnippetsJSON`) so it's unit-testable without calling `osascript`.
- `SplitQuery` — splits a query on the first `#` only (see
  `docs/specification.md`).
- `FuzzyMatch` — a rune-based, ASCII-case-folded subsequence matcher (not
  Alfred's own built-in filter, which doesn't fuzzy-match CJK text well —
  hence `alfredfiltersresults: false` on the Script Filter node).
- `Categories` / `CountByCategory` — distinct categories in first-appearance
  order, and per-category snippet counts.

### `internal/phraseshelfcmd/`

Builds the `scriptfilter.Response` for both stages from
`internal/phraseshelf`, including every non-valid informational row (no
match, empty category, Store error) — the only layer that imports both
`internal/phraseshelf` and `internal/scriptfilter`.

## Dependency Flow

```
cmd/phrase-shelf-alfred → internal/phraseshelfcmd → internal/phraseshelf
                                                   → internal/scriptfilter
```

`internal/phraseshelf` never imports `internal/scriptfilter` — it has no
Alfred JSON concerns.

## Packaging

At build time (`make build-workflow` / `scripts/build-workflow.sh`):

```
.build/                     ← temporary build directory
├── info.plist              ← copied from workflow/
├── icon.png
└── phrase-shelf-alfred     ← universal (amd64+arm64) binary
```

`cmd/phrase-shelf-alfred` is built for `darwin/amd64` and `darwin/arm64`
and merged into a single universal binary with `lipo`, so the packaged
workflow runs natively on both Intel and Apple Silicon Macs without
needing a runtime interpreter or vendored dependencies. The entire
`.build/` directory is zipped to `dist/<name>-<version>.alfredworkflow`.

## Alfred Configuration Builder (`userconfigurationconfig`)

Alfred 5 の Configuration Builder は `info.plist` の `userconfigurationconfig` キーで定義する。
利用可能な全型・各キーの詳細は [`docs/alfred-workflow-notes/configuration-builder.md`](alfred-workflow-notes/configuration-builder.md)、
このプロジェクトの設定項目は [`docs/configuration-builder.md`](configuration-builder.md) を参照。
`snip` コマンドは現在 Config Builder 変数を使っていない
（`userconfigurationconfig` は空配列） — 設定が必要な機能を追加する際は
上記ドキュメントのパターンに従う。

### Passing Variables

Alfred はスクリプト実行時に各 `variable` を環境変数として渡す。
インストール直後は `prefs.plist` が存在しないため変数は未セットになる場合がある。
スクリプト側で常にデフォルト値を持たせること。

~~~go
// Go
value := os.Getenv("my_variable")
if value == "" {
    value = "fallback"
}
~~~

**注意:** `checkbox` 型の unchecked 値は `"0"` ではなく空文字 `""` になる。
`value == "1"` で判定し、`"0"` との比較は避けること。

### Relationship Between `variables` / `prefs.plist` / `default`

| 場所 | 役割 |
|---|---|
| `userconfigurationconfig[].config.default` | Configuration Builder UI の初期表示のみ。変数への書き込みは行わない。 |
| `prefs.plist`（同ディレクトリ） | ユーザーが Configuration Builder で保存した値。Alfred が自動生成・更新する。 |
| `info.plist` の `variables` | スクリプトに常に渡したい固定の環境変数。Configuration Builder で管理する変数はここに入れない。 |
