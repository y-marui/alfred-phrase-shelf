# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added

- `snip` Alfred Script Filter workflow for the "Phrase Shelf" macOS app:
  `snip` lists categories, `snip <category>` fuzzy-filters them,
  `snip <category>#` lists that category's snippets, and
  `snip <category>#<query>` fuzzy-filters them by name. Selecting a
  snippet copies its text and auto-pastes it into the frontmost app via
  the workflow's native Copy to Clipboard node.
- `internal/phraseshelf`: the domain package — the `Store` interface and
  its `osascript`-backed implementation talking to Phrase Shelf's
  AppleScript interface, the first-`#`-only query splitter, a Unicode-aware
  (Japanese-capable) subsequence fuzzy matcher, and category
  extraction/ordering from the snippet list.
- `internal/phraseshelfcmd`: builds the Script Filter response for both
  stages, including the non-valid informational rows (no matching
  category/snippet, empty category, Phrase Shelf unreachable).
- The native Copy to Clipboard node sets `ignoredynamicplaceholders: true`,
  so a snippet whose text happens to contain an Alfred placeholder like
  `{date}` or `{clipboard}` is pasted verbatim instead of being expanded.
