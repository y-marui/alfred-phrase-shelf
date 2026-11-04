# Phrase Shelf (Alfred Workflow)

> **This is the English (reference) version.**
> For the Japanese canonical version, see [README-jp.md](README-jp.md).

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)
[![CI](https://github.com/y-marui/alfred-phrase-shelf/actions/workflows/ci.yml/badge.svg)](https://github.com/y-marui/alfred-phrase-shelf/actions/workflows/ci.yml)
[![Charter Check](https://github.com/y-marui/alfred-phrase-shelf/actions/workflows/dev-charter-check.yml/badge.svg)](https://github.com/y-marui/alfred-phrase-shelf/actions/workflows/dev-charter-check.yml)
[![GitHub Sponsors](https://img.shields.io/github/sponsors/y-marui?style=social)](https://github.com/sponsors/y-marui)
[![Buy Me a Coffee](https://img.shields.io/badge/Buy%20Me%20a%20Coffee-donate-yellow.svg)](https://www.buymeacoffee.com/y.marui)

Search and insert snippets saved in the "Phrase Shelf" macOS app, without
leaving Alfred — fuzzy-filter by category and name (including Japanese
text), then auto-paste the result into whatever app you were typing in.

## Requirements

- Alfred 5 (Powerpack required for Script Filter)
- The "Phrase Shelf" macOS app, installed and able to run — this workflow
  is a client for it and has no snippet storage of its own. On first use,
  macOS prompts for Automation permission so this workflow can talk to it;
  grant it, or snippets won't load.

## Setup

```bash
# Download the latest release:
# https://github.com/y-marui/alfred-phrase-shelf/releases/latest
```

Double-click `*.alfredworkflow` to install in Alfred.

Or build from source (requires Go — see `go.mod` for the version):

```bash
git clone https://github.com/y-marui/alfred-phrase-shelf
cd alfred-phrase-shelf
make build-workflow
# → dist/*.alfredworkflow
```

## Usage

Open Alfred and type `snip`.

```
snip                    -> list all categories
snip <category>         -> fuzzy-filter categories by name
snip <category>#        -> list every snippet in that category
snip <category>#<query> -> fuzzy-filter that category's snippets by name
```

Examples:

```
snip work               -> categories whose name fuzzy-matches "work"
snip work#              -> every snippet in the "work" category
snip work#gre           -> "work" snippets whose name fuzzy-matches "gre"
snip #memo              -> uncategorized snippets whose name fuzzy-matches "memo"
```

On a category row, press Tab (or Enter) to drill into it — it rewrites the
query to `snip <category>#` for you. On a snippet row, press Enter: its
full text is copied to the clipboard and automatically pasted into
whichever app was frontmost.

Filtering is a genuine fuzzy (subsequence) match, not a prefix match, and
works on Japanese names since it compares characters directly rather than
relying on word boundaries.

The category typed (or autocompleted) before the first `#` is looked up
exactly, not fuzzy-matched — only the text after it is fuzzy-filtered.
Only the *first* `#` in the query is a separator; any further `#`
characters are treated as literal text in the filter.

## License

MIT — see [LICENSE](LICENSE)

---

*This is the reference (English) version. The canonical Japanese version is [README-jp.md](README-jp.md). Update both files in the same commit.*
