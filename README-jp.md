# Phrase Shelf (Alfred Workflow)

> **これは日本語版（正本）です。**
> 英語版（参照）は [README.md](README.md) を参照してください。

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)
[![CI](https://github.com/y-marui/alfred-phrase-shelf/actions/workflows/ci.yml/badge.svg)](https://github.com/y-marui/alfred-phrase-shelf/actions/workflows/ci.yml)
[![Charter Check](https://github.com/y-marui/alfred-phrase-shelf/actions/workflows/dev-charter-check.yml/badge.svg)](https://github.com/y-marui/alfred-phrase-shelf/actions/workflows/dev-charter-check.yml)
[![GitHub Sponsors](https://img.shields.io/github/sponsors/y-marui?style=social)](https://github.com/sponsors/y-marui)
[![Buy Me a Coffee](https://img.shields.io/badge/Buy%20Me%20a%20Coffee-donate-yellow.svg)](https://www.buymeacoffee.com/y.marui)

macOS アプリ「Phrase Shelf」に保存したスニペットを、Alfred を離れずに検索して
挿入するワークフロー。カテゴリ名・スニペット名を日本語を含めてファジー検索し、
選んだスニペットをそのままフロントのアプリへ自動ペーストする。

## Requirements

- Alfred 5（Script Filter には Powerpack が必要）
- macOS アプリ「Phrase Shelf」— インストール済みで起動可能であること。
  このワークフローは同アプリのクライアントであり、スニペット自体は保持しない。
  初回利用時に macOS のオートメーション許可プロンプトが表示されるので許可する
  こと（許可しないとスニペットを取得できない）。

## Setup

```bash
# 最新リリースをダウンロード:
# https://github.com/y-marui/alfred-phrase-shelf/releases/latest
```

`*.alfredworkflow` をダブルクリックして Alfred にインストールします。

またはソースからビルド（Go が必要 — バージョンは `go.mod` を参照）:

```bash
git clone https://github.com/y-marui/alfred-phrase-shelf
cd alfred-phrase-shelf
make build-workflow
# → dist/*.alfredworkflow
```

## Usage

Alfred を開いて `snip` と入力します。

```
snip                    -> 全カテゴリを一覧表示
snip <category>         -> カテゴリ名をファジー検索
snip <category>#        -> そのカテゴリの全スニペットを一覧表示
snip <category>#<query> -> そのカテゴリ内でスニペット名をファジー検索
```

例:

```
snip work               -> 名前が "work" にファジーマッチするカテゴリ
snip work#              -> "work" カテゴリの全スニペット
snip work#gre           -> "work" カテゴリ内で "gre" にファジーマッチするスニペット
snip #memo              -> 未分類スニペットのうち "memo" にファジーマッチするもの
```

カテゴリの行では Tab（または Enter）でそのカテゴリに入ります — クエリが
`snip <category>#` に書き換わります。スニペットの行で Enter を押すと、
その全文がクリップボードにコピーされ、フロントのアプリへ自動的にペースト
されます。

検索は前方一致ではなく、文字が順番通りに（連続していなくても）含まれて
いるかを見る本来の意味でのファジー（部分列）マッチで、単語境界に頼らない
ため日本語の名前でも機能します。

最初の `#` より前（直接入力または Stage 1 の補完で入る）のカテゴリは
ファジー検索せず完全一致で扱う。ファジー検索の対象は `#` の後ろの部分だけ。
クエリ中で区切りとして扱うのは最初の `#` のみで、それ以降の `#` は検索
対象の文字列としてそのまま扱う。

## License

MIT — [LICENSE](LICENSE) を参照

---

*この文書には英語版（参照版）[README.md](README.md) があります。編集時は同一コミットで更新してください。*
