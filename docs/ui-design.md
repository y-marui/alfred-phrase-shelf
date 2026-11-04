# UI Design

Alfred Script Filter workflows present results as a list of items in the Alfred
launcher. This document defines the UI conventions for result items in this
workflow.

## Result Item Structure

Alfred result items are JSON objects with the following fields used in this workflow:

| Field | Type | Required | Description |
|---|---|---|---|
| `title` | string | yes | Primary text (large, always visible) |
| `subtitle` | string | no | Secondary text (small, below title) |
| `arg` | string | no | Value passed to Alfred's action on Enter |
| `uid` | string | no | Unique ID for Alfred's learned ordering |
| `valid` | bool | yes | If false, Enter does not trigger an action |
| `autocomplete` | string | no | Text inserted into Alfred's input on Tab |
| `icon` | object | no | Custom icon (`{ "path": "icon.png" }`) |

## Text Guidelines

### No Unicode Emoji in `title` / `subtitle`

- **Prohibited:** `🔍 Search`, `✅ Done`, `📄 Document`
- **Allowed:** ASCII symbols — `>`, `*`, `[x]`, `(!)`, `--`
- **Reason:** Emoji rendering is inconsistent across Alfred versions and macOS
  updates. ASCII symbols are universally stable.

### Capitalization

- `title`: Sentence case for action labels; preserve original casing for data values
  (category/snippet names come straight from the user's own Phrase Shelf data).
- `subtitle`: Sentence case. Use short phrases, not full sentences.

### Empty / Error States

- No matching category/snippet → an informative item (e.g. `No matching category`)
  with `valid: false`.
- Empty category (`snip <category>#` where `<category>` doesn't exist) →
  `No snippets in <category>` with `valid: false` — not an error, since Phrase
  Shelf's `list snippets category "..."` simply returns an empty array for an
  unknown category.
- Phrase Shelf unreachable (app not launchable, AppleScript error) →
  `Phrase Shelf is unreachable` with `valid: false` and the underlying message
  as `subtitle`.
- Any other error → `main.go`'s `recover()`-wrapped `dispatch` automatically
  shows an error item; do not hide errors silently.

## Icon

- Workflow icon: `workflow/icon.png` (PNG, any size — Alfred scales it). Still
  the 1×1px template placeholder; replace with a real icon before any Alfred
  Gallery submission (see `docs/alfred-gallery-readiness.md`).
- Alfred controls light/dark mode; do not ship separate light/dark icons.
- Per-item icons are optional. If omitted, the workflow icon is used.

## Keyboard Shortcuts

These are standard Alfred behaviors — do not override them in the workflow:

| Key | Action |
|---|---|
| ↩ Enter | Stage 1: autocomplete into `<category>#` (item is `valid: false`). Stage 2: copy `arg` to clipboard and auto-paste it |
| ⌘↩ | Alfred action picker |
| ⇥ Tab | Stage 1: insert `autocomplete` text (`<category>#`) into Alfred's input |
| ⌘C | Copy `arg` to clipboard |
| ⌘L | Show `title` in Large Type |

## Layout Conventions by Command

### Stage 1 — category row

```
title:        <category name> (or "(Uncategorized)" for "")
subtitle:      <N> snippet(s)
valid:        false
autocomplete: <category>#
```

### Stage 2 — snippet row (final, actionable)

```
title:    <snippet name>
subtitle: <first line of snippet text, truncated>
arg:      <full snippet text>
uid:      phrase-shelf-<snippet id>
valid:    true
```

### No-match / empty-category / error rows

```
title:    No matching category | No matching snippet | No snippets in <category> | Phrase Shelf is unreachable
subtitle: Try a different query | Try a different category | <error message>
valid:    false
```
