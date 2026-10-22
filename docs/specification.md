# Specification

> Functional specification, behavior definition, and data flow for the
> alfred-phrase-shelf workflow.

## Overview

This workflow is an Alfred 5 Script Filter that searches and inserts
snippets saved in the "Phrase Shelf" macOS app. It has no snippet storage
of its own — it is a client that calls Phrase Shelf's AppleScript interface
(`PhraseShelf.sdef`) via `osascript`.

## Commands

### `snip` (the only command)

**Trigger:** `snip`, `snip <category>`, `snip <category>#`,
`snip <category>#<query>`

The query (the text after the keyword) is split on `#`:

- **No `#`** → the whole query (possibly empty) is a category-filter text
  → **Stage 1**.
- **Has `#`** → split at the *first* `#` only. Everything before it is the
  exact `category` (never fuzzy-matched); everything after it — including
  any further literal `#` characters — is the snippet-filter text →
  **Stage 2**.

Implemented by `internal/phraseshelf.SplitQuery`.

#### Stage 1 — list/filter categories

1. `osascript ... to list snippets` → every snippet, sortOrder ascending.
2. Derive the category list: distinct `category` values in order of first
   appearance (`internal/phraseshelf.Categories`) — already the app's own
   category order (named categories, then `""`/uncategorized last if
   present).
3. Fuzzy-filter category names against the filter text
   (`internal/phraseshelf.FuzzyMatch`). Empty filter text → show all.
4. No categories / Phrase Shelf unreachable → a single non-valid
   explanatory item, never an empty list.

Each result item is **autocomplete-only**, not directly actionable
(`valid: false`): its `autocomplete` is `"<category>#"`, so accepting it
(Tab, or Return since it's invalid) rewrites the Alfred query to
`snip <category>#`, entering Stage 2.

| Field | Source |
|---|---|
| `title` | Category name (`(Uncategorized)` for `""`) |
| `subtitle` | Snippet count in that category |
| `valid` | `false` |
| `autocomplete` | `"<category>#"` (the real `""` category value, not the display label) |

#### Stage 2 — list/filter snippets in a category

1. `osascript ... to list snippets category "<category>"` with `category`
   taken verbatim (exact lookup — an unknown category simply returns an
   empty array, handled as a dedicated info row, not an error).
2. Fuzzy-filter the snippets' `name` against the filter text
   (after the first `#`). Empty filter text → show all, in sortOrder.
3. Empty category / no match / Phrase Shelf unreachable → a single
   non-valid explanatory item.

Each result item **is** the final, actionable item:

| Field | Source |
|---|---|
| `title` | Snippet name |
| `subtitle` | First line of `text`, truncated to 60 runes |
| `arg` | Snippet's full `text` |
| `uid` | `"phrase-shelf-" + Snippet.ID` |
| `valid` | `true` |

## Data Flow

```
Alfred input (keyword "snip" + query string)
  │
  ▼
cmd/phrase-shelf-alfred/main.go   reads os.Args[1]
  │
  ▼
dispatch(query)                   recover()-wraps the call below → error item on panic
  │
  ▼
internal/phraseshelfcmd.List(store, query)
  │
  ├─ phraseshelf.SplitQuery(query)         → category / filterText / hasCategory
  ├─ store.ListSnippets() /
  │    store.ListSnippetsByCategory(category)   → osascript → Phrase Shelf app
  ├─ phraseshelf.Categories / CountByCategory    (Stage 1)
  └─ phraseshelf.FuzzyMatch(filterText, ...)     (both stages)
  │
  ▼
scriptfilter.Response.Write()     prints JSON to stdout → Alfred renders result list
  │
  ▼ (user presses Enter on a Stage-2 item)
alfred.workflow.output.clipboard  native node, autopaste enabled — unchanged by this binary
```

## Error Handling

- Any panic during `internal/phraseshelfcmd.List` is recovered by
  `dispatch` in `main.go`, which emits a single error result item
  containing the panic value.
- A `Store` error (Phrase Shelf not launchable, AppleScript error, bad
  JSON) is turned into a single non-valid "Phrase Shelf is unreachable"
  item by `phraseshelfcmd`'s `errorResponse`, never a crash or empty
  output.

## Configuration Variables

None — `workflow/info.plist`'s `userconfigurationconfig` is an empty
array. See `docs/configuration-builder.md` for the mechanism to use if a
future setting needs one.

## Constraints

- Script Filter response time target: **< 100 ms** (a compiled binary,
  usually well within budget, aside from the `osascript` round-trip to
  Phrase Shelf itself).
- All output must go through `scriptfilter.Response.Write()` — never
  `fmt.Print*` directly.
- `cmd/phrase-shelf-alfred/main.go` contains no business logic; it only
  dispatches argv and writes the response.
