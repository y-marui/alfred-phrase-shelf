# Architecture

## Overview

An Alfred Workflow (Go): `cmd/example-alfred` is the single universal
(amd64+arm64) binary `workflow/info.plist` invokes. The "tpl" Script Filter
node runs it with the query as `$1`, and Alfred's existing native Open URL
action node opens whichever shortcut's URL the user picks — the binary never
opens URLs itself.

```
Alfred
  │  keyword "tpl" + query
  ▼
cmd/example-alfred/main.go     ← Alfred boundary; argv dispatch + recover() only
  │
  ▼
internal/examplecmd.List()     ← builds the Script Filter response
  │
  ▼
internal/example.Filter()      ← domain logic: substring-filter the shortcut list
  │
  ▼
internal/scriptfilter.Response ← Script Filter JSON types + writer

Alfred (user presses Enter)
  │  selected shortcut's URL, as {query}
  ▼
alfred.workflow.action.openurl (native, unchanged by this binary)
```

## Layers

### `cmd/example-alfred/`

The only binary Alfred executes. Dispatches argv to `internal/examplecmd`,
recovers from any panic into a visible Script Filter error item, and writes
the response — no business logic beyond that.

### `internal/scriptfilter/`

Script Filter JSON types (`Item`, `Response`, `Icon`, `Mod`) and the writer
that encodes them to stdout. Alfred-independent; ported as-is from
`alfred-note-table-converter`'s version. Keep this if you replace the rest
of `internal/` with your own domain logic.

### `internal/example/`

Pure, Alfred-independent domain logic — the placeholder this template ships
for you to replace. A static list of name→URL shortcuts and a
case-insensitive substring filter. Unit tested without Alfred running, and
never imports `internal/scriptfilter`.

### `internal/examplecmd/`

Builds the `scriptfilter.Response` from `internal/example`'s filtered
results — the only layer that imports both `internal/example` and
`internal/scriptfilter`.

## Dependency Flow

```
cmd/example-alfred → internal/examplecmd → internal/example
                                          → internal/scriptfilter
```

`internal/example` never imports `internal/scriptfilter` — it has no Alfred
JSON concerns.

## Packaging

At build time (`make build-workflow` / `scripts/build-workflow.sh`):

```
.build/                     ← temporary build directory
├── info.plist              ← copied from workflow/
├── icon.png
└── example-alfred          ← universal (amd64+arm64) binary
```

`cmd/example-alfred` is built for `darwin/amd64` and `darwin/arm64` and
merged into a single universal binary with `lipo`, so the packaged workflow
runs natively on both Intel and Apple Silicon Macs without needing a runtime
interpreter or vendored dependencies. The entire `.build/` directory is
zipped to `dist/<name>-<version>.alfredworkflow`.

## Alfred Configuration Builder (`userconfigurationconfig`)

Alfred 5 の Configuration Builder は `info.plist` の `userconfigurationconfig` キーで定義する。
利用可能な全型・各キーの詳細は [`docs/alfred-workflow-notes/configuration-builder.md`](alfred-workflow-notes/configuration-builder.md)、
このプロジェクトの設定項目は [`docs/configuration-builder.md`](configuration-builder.md) を参照。
このテンプレートの例示コマンドは現在 Config Builder 変数を使っていない
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
