# Alfred Workflow Template

> **これは日本語版（正本）です。**
> 英語版（参照）は [README.md](README.md) を参照してください。

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)
[![CI](https://github.com/y-marui/alfred-workflow-template/actions/workflows/ci.yml/badge.svg)](https://github.com/y-marui/alfred-workflow-template/actions/workflows/ci.yml)
[![Charter Check](https://github.com/y-marui/alfred-workflow-template/actions/workflows/dev-charter-check.yml/badge.svg)](https://github.com/y-marui/alfred-workflow-template/actions/workflows/dev-charter-check.yml)
[![GitHub Sponsors](https://img.shields.io/github/sponsors/y-marui)](https://github.com/sponsors/y-marui)
[![Buy Me a Coffee](https://img.shields.io/badge/Buy%20Me%20a%20Coffee-ffdd00?logo=buy-me-a-coffee&logoColor=black)](https://www.buymeacoffee.com/y.marui)

| 項目 | 内容 |
|---|---|
| 開発対象 | Alfred 5 Script Filter ワークフロー |
| 開発環境 | 個人〜小規模チーム（1〜3人） |
| 主言語 | 英語（OSS） |
| ライセンス | MIT |
| 動作環境 | Go（`go.mod` 参照）, Alfred 5 |
| AI ツール | Claude Code / GitHub Copilot / Gemini CLI |

> Alfred 5 Script Filter ワークフローのプロダクションレディな Go テンプレート。
> 10分で開発を開始できます。

## Features

- ✅ **レイヤードアーキテクチャ** — Alfred 境界（`cmd/`）とドメインロジック（`internal/`）を分離
- ✅ **再利用可能な Script Filter JSON レイヤー** — `internal/scriptfilter`（Alfred 非依存）
- ✅ **単一ユニバーサルバイナリ** — `darwin/amd64` + `darwin/arm64` を `lipo` でマージ、ランタイムインタプリタ不要
- ✅ **フルテストスイート** — `go test` で Alfred なしにテスト実行可能
- ✅ **CI/CD** — GitHub Actions でリント・テスト・ビルド・リリースを自動化
- ✅ **依存ゼロ** — `go.mod` に `require` ブロックなし
- ✅ **AI 対応** — `AI_CONTEXT.md` + `CLAUDE.md` で AI アシスタントのコンテキストを管理

## Requirements

- Alfred 5（Script Filter には Powerpack が必要）
- Go（バージョンは `go.mod` を参照）
- [pre-commit](https://pre-commit.com/)（セキュリティフック用）

## Quick Start

### Using this template

1. GitHub で **"Use this template"** → **"Create a new repository"** をクリック
2. 新しいリポジトリをクローンして AI ツールで開く
3. AI に「AI_CONTEXT.md の初期セットアップを実行して」と伝える — AI が以下を実行します:
   - GitHub リポジトリ設定を適用
   - `README_TEMPLATE.md` → `README.md`、`README_TEMPLATE-jp.md` → `README-jp.md` にリネーム
   - `{user}`、`{repo}`、`{keyword}` プレースホルダを置換
   - dev-charter の cron スケジュールをランダム化
4. ワークフローをカスタマイズする（`DEVELOPING.md` 参照）

> **Note:** README を書き換える際、Charter Check バッジ（`dev-charter-check.yml` の badge）が
> 消えやすい。`README.md` / `README-jp.md` の両方に残っているか確認する。

### Development (this template)

```bash
git clone https://github.com/y-marui/alfred-workflow-template
cd alfred-workflow-template

# Alfred をローカルでシミュレート
go run ./cmd/example-alfred
go run ./cmd/example-alfred "doc"

# テストを実行
make test

# ワークフローパッケージをビルド
make build-workflow
# → dist/workflow-template-1.0.0.alfredworkflow
```

`dist/*.alfredworkflow` をダブルクリックして Alfred にインストールします。

## Usage

`tpl` キーワードでサンプルショートカット（repo, docs, issues）を検索します。

```
tpl          -> サンプルショートカット一覧（repo, docs, issues）
tpl <query>  -> 名前でショートカットを絞り込み
```

* <kbd>↩︎</kbd> 選択したショートカットの URL を開く

### Troubleshooting

**結果が表示されない場合**
- Alfred のデバッガーを確認: Alfred を開いて `⌘D`

## Project Structure

```
alfred-workflow-template/
├── cmd/
│   └── example-alfred/     # Alfred が実行するバイナリ（argv ディスパッチのみ）
├── internal/
│   ├── example/            # ドメインロジック — 自分のものに置き換える
│   ├── examplecmd/         # Script Filter レスポンスを組み立てる
│   └── scriptfilter/       # Script Filter JSON 型（Alfred 非依存、そのまま再利用可）
├── workflow/                 # Alfred パッケージ (info.plist, icon.png)
├── scripts/                  # build-workflow.sh, extract-changelog.sh
└── docs/                     # アーキテクチャ・リファレンスドキュメント
```

## Documentation

| ドキュメント | 内容 |
|---|---|
| [docs/architecture.md](docs/architecture.md) | モジュール・レイヤー構造 |
| [docs/file-map.md](docs/file-map.md) | ファイルレベルの依存関係マップ |
| [docs/specification.md](docs/specification.md) | 機能仕様・データフロー |
| [docs/ui-design.md](docs/ui-design.md) | Alfred 結果アイテムの UI 設計指針 |
| [docs/configuration-builder.md](docs/configuration-builder.md) | このプロジェクトの Configuration Builder 設定 |
| [docs/alfred-workflow-notes/](docs/alfred-workflow-notes/README.md) | Alfred workflow開発の汎用リファレンス(他プロジェクトへの正本。`git subtree`で配布) |

## AI-Assisted Development

このテンプレートは AI 支援開発に対応しています。

| ツール | 役割 |
|---|---|
| Claude Code | アーキテクチャ設計・大規模変更・リファクタリング |
| GitHub Copilot | バグ修正・細かな実装・単体テスト作成 |
| Gemini CLI | ドキュメント管理 |

セッションコンテキスト: [`AI_CONTEXT.md`](AI_CONTEXT.md)、[`CLAUDE.md`](CLAUDE.md)

## Customizing This Template

上記 Quick Start の初期セットアップが完了したら、ワークフローをカスタマイズします:

1. `workflow/info.plist` を編集:
   - `bundleid` を自分のバンドル ID に変更（例: `com.yourname.workflowname`）
   - キーワード（`wf`）を自分のトリガーキーワードに変更
   - `uuidgen` で生成した UUID に置き換え
2. `cmd/example-alfred` を `cmd/<workflow名>-alfred` にリネーム
3. `internal/example` + `internal/examplecmd` を自分のドメインロジックに置き換え
   （`internal/scriptfilter` はそのまま使う）
4. `go.mod` のモジュールパスを更新
5. `workflow/icon.png` を追加

クリップボード・キーストローク・通知など macOS 固有の操作をGoで書く前に、
[docs/alfred-workflow-notes/workflow-object-schema.md](docs/alfred-workflow-notes/workflow-object-schema.md)
を確認する — Alfred ネイティブのオブジェクトで、コード無しに実現できることがある。

## Release

```bash
# 1. workflow/info.plist のバージョンを更新
# 2. タグを付けてプッシュ
git tag v1.2.3
git push --tags
# GitHub Actions が .alfredworkflow をビルドして GitHub Release を作成
```

## License

MIT — [LICENSE](LICENSE) を参照

---

*この文書には英語版（参照版）[README.md](README.md) があります。編集時は同一コミットで更新してください。*
