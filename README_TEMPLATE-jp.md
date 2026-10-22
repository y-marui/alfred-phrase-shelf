# {WorkflowName}

> **これは日本語版（正本）です。**
> 英語版（参照）は [README.md](README.md) を参照してください。

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)
[![CI](https://github.com/{user}/{repo}/actions/workflows/ci.yml/badge.svg)](https://github.com/{user}/{repo}/actions/workflows/ci.yml)
[![Charter Check](https://github.com/{user}/{repo}/actions/workflows/dev-charter-check.yml/badge.svg)](https://github.com/{user}/{repo}/actions/workflows/dev-charter-check.yml)
[![GitHub Sponsors](https://img.shields.io/github/sponsors/{user}?style=social)](https://github.com/sponsors/{user})
[![Buy Me a Coffee](https://img.shields.io/badge/Buy%20Me%20a%20Coffee-donate-yellow.svg)](https://www.buymeacoffee.com/{bmc_username})

{一行概要: 何を・誰のために・どう解決するか}

## Requirements

- Alfred 5（Script Filter には Powerpack が必要）

## Setup

```bash
# 最新リリースをダウンロード:
# https://github.com/{user}/{repo}/releases/latest
```

`*.alfredworkflow` をダブルクリックして Alfred にインストールします。

またはソースからビルド（Go が必要 — バージョンは `go.mod` を参照）:

```bash
git clone https://github.com/{user}/{repo}
cd {repo}
make build-workflow
# → dist/*.alfredworkflow
```

## Usage

Alfred を開いて `{keyword}` と入力します。

```
{keyword}          -> 全てのサンプルショートカットを一覧表示
{keyword} <query>  -> 名前でショートカットを絞り込み
```

Enter で選択したショートカットの URL を開きます。

## License

MIT — [LICENSE](LICENSE) を参照

---

*この文書には英語版（参照版）[README.md](README.md) があります。編集時は同一コミットで更新してください。*
