# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Changed

- **Rewrite in Go**, matching the `cmd/`+`internal/` architecture every
  workflow generated from this template has already independently converged
  on (`alfred-quick-txt-save`, `alfred-markdown-ref`,
  `alfred-note-md-template`, and others). Removed `search`/`open`/
  `config`/`help` and the `ExampleService`/`ApiClient`/persistent config
  store/TTL cache/file logger they depended on — none of those were kept
  past the template defaults in any real migration; two were explicitly
  removed after being ported (`alfred-paste-formatted-date`'s config store:
  "nothing ever wrote to the config store"). The new single example command
  (`cmd/example-alfred` + `internal/example` + `internal/examplecmd` +
  `internal/scriptfilter`) lists static name→URL shortcuts, filtered by
  query, opened via the workflow's existing native Open URL node — the same
  minimal shape as the old `open_cmd.py`, with nothing else. Single
  universal (amd64+arm64) binary via `lipo`; no vendored dependencies, no
  runtime interpreter selection (`use_uv` is gone).
- `docs/configuration-builder.md` now covers only this project's own
  Configuration Builder settings; the general mechanism reference moved
  to `docs/alfred-workflow-notes/configuration-builder.md`

### Added

- `docs/alfred-workflow-notes/` — canonical, subtree-pullable reference for
  cross-project Alfred workflow development knowledge (Configuration
  Builder mechanism, `info.plist` object schema)

### Fixed

- `docs/alfred-workflow-notes/README.md`'s "Installing in a consuming
  project" sample commands: a plain `git subtree add`/`pull` against
  `alfred-workflow-template` pulled in that repo's entire root instead of
  just `docs/alfred-workflow-notes/`, since (unlike `docs/dev-charter/`)
  the shared content here lives in a subdirectory, not at the upstream
  repo's root. Both the first-time install and the `update-workflow-notes`
  Makefile target now capture `git subtree split`'s output SHA directly
  (no named branch, so nothing that can collide with a branch the
  consuming repo already has, or needs cleanup afterwards) and run under
  `set -e` (so a failed split aborts before `merge` runs against a stale
  or empty SHA) — verified end-to-end in a sandbox repo (first-time
  install, idempotent re-run, a real update, and a forced split failure
  leaving the stash untouched rather than popped onto a broken state)
- `scripts/check-charter-subtree-edit.sh` (and its mirrored sample in
  `docs/alfred-workflow-notes/README.md`) rejected the merge commit that
  finalizes a `git subtree add`/`pull`/`merge` whenever that merge needed
  manual conflict resolution, since a plain `git commit` runs the
  pre-commit hook and the old check couldn't tell that apart from a direct
  edit under the subtree — getting the working tree stuck mid-merge. Now
  skipped only when `MERGE_HEAD` carries the `git-subtree-dir: $PREFIX`
  trailer `git subtree` itself writes, so an unrelated merge that happens
  to touch the same prefix is still blocked — verified in a sandbox repo
  (a real conflicting subtree pull now finalizes; a standalone direct edit
  and an unrelated merge touching the prefix are still rejected)

## [0.1.0] - 2024-01-01

### Added

- Initial release of the Alfred Workflow Template
- Alfred SDK: `response`, `cache`, `config`, `logger`, `router`, `safe_run`
- Command-based UX: `search`, `open`, `config`, `help`
- Vendor packaging via `scripts/vendor.sh`
- Build pipeline via `scripts/build.sh`
- GitHub Actions CI (lint, test, build)
- GitHub Actions Release (tag → `.alfredworkflow` → GitHub Release)
- Full pytest test suite

[Unreleased]: https://github.com/yourname/alfred-workflow-template/compare/v0.1.0...HEAD
[0.1.0]: https://github.com/yourname/alfred-workflow-template/releases/tag/v0.1.0
