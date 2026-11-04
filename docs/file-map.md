# File Map

> File-level dependency map for alfred-phrase-shelf.
> Add to this as you explore the codebase during development.

## Entry Points

| File | Role |
|---|---|
| `cmd/phrase-shelf-alfred/main.go` | Alfred executes this binary — the sole entry point |

## Call Flow

```
cmd/phrase-shelf-alfred/main.go
  └─ dispatch(query)                              ← recover()-wrapped
       └─ internal/phraseshelfcmd.List(store, query)
            ├─ internal/phraseshelf.SplitQuery(query)
            ├─ store.ListSnippets() / ListSnippetsByCategory(category)
            ├─ internal/phraseshelf.Categories / CountByCategory
            └─ internal/phraseshelf.FuzzyMatch(filterText, ...)
       └─ scriptfilter.Response.Write(os.Stdout)
```

## Module Dependency Table

### `internal/scriptfilter/`

| File | Imports from | Notes |
|---|---|---|
| `scriptfilter.go` | stdlib only | Script Filter JSON types + `Response.Write` |

### `internal/phraseshelf/`

| File | Imports from | Notes |
|---|---|---|
| `snippet.go` | stdlib only | `Snippet` type, `Store` interface |
| `query.go` | stdlib only | `SplitQuery` — first-`#`-only split |
| `fuzzy.go` | stdlib only | `FuzzyMatch` — rune-based ASCII-case-folded subsequence match |
| `categories.go` | stdlib only | `Categories` / `CountByCategory` |
| `osascript.go` | `os/exec`, stdlib | `OsascriptStore` (real `Store`), `parseSnippetsJSON`, `escapeAppleScriptString` |

### `internal/phraseshelfcmd/`

| File | Imports from | Notes |
|---|---|---|
| `phraseshelfcmd.go` | `internal/phraseshelf`, `internal/scriptfilter` | Builds the Stage 1 / Stage 2 Script Filter response |

### `cmd/phrase-shelf-alfred/`

| File | Imports from | Notes |
|---|---|---|
| `main.go` | `internal/phraseshelf`, `internal/phraseshelfcmd`, `internal/scriptfilter` | Alfred boundary; argv dispatch only |

## Key Files for Customization

| File | What to change |
|---|---|
| `workflow/info.plist` | `bundleid`, keyword, UIDs, category, description |
| `internal/phraseshelf/` | The domain logic and the `Store` talking to Phrase Shelf |
| `internal/phraseshelfcmd/phraseshelfcmd.go` | Result formatting for both stages |
| `go.mod` | Module path |
| `workflow/icon.png` | Workflow icon (still the 1×1px template placeholder — replace before any Alfred Gallery submission) |
