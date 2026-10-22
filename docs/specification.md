# Specification

> Functional specification, behavior definition, and data flow for the
> alfred-workflow-template. Adapt this document when building a workflow from
> this template.

## Overview

This workflow is an Alfred 5 Script Filter that accepts a keyword + query,
lists matching results, and lets Alfred's native Open URL action node open
whichever one the user selects.

## Commands

### `example` (the only command)

**Trigger:** `tpl` or `tpl <query>`

**Behavior:**
1. Filter `internal/example.Shortcuts` by whether `query` is a
   case-insensitive substring of the shortcut's name. An empty query
   matches every shortcut.
2. If no matches → display "No matching shortcut" item (valid: false).
3. Otherwise → display one result item per matching shortcut; pressing
   Enter passes its URL as `arg`, which the workflow's native Open URL
   action node opens.

**Result item fields:**

| Field | Source | Notes |
|---|---|---|
| `title` | `Shortcut.Name` | Primary display text |
| `subtitle` | `Shortcut.URL` | Secondary display text |
| `arg` | `Shortcut.URL` | Passed to the native Open URL node on Enter |
| `uid` | `"example-" + Shortcut.Name` | Used by Alfred for learned ordering |
| `autocomplete` | `Shortcut.Name` | Filled into the query box on Tab |

**Shortcuts (template defaults — replace with project-specific URLs):**

| Name | URL |
|---|---|
| `repo` | `https://github.com/y-marui/alfred-workflow-template` |
| `docs` | `https://github.com/y-marui/alfred-workflow-template/tree/main/docs` |
| `issues` | `https://github.com/y-marui/alfred-workflow-template/issues` |

## Data Flow

```
Alfred input (keyword "tpl" + query string)
  │
  ▼
cmd/example-alfred/main.go        reads os.Args[1]
  │
  ▼
dispatch(query)                   recover()-wraps the call below → error item on panic
  │
  ▼
internal/examplecmd.List(query)
  │
  ▼
internal/example.Filter(query)    substring-match against Shortcuts
  │
  ▼
scriptfilter.Response.Write()     prints JSON to stdout → Alfred renders result list
  │
  ▼ (user presses Enter)
alfred.workflow.action.openurl    native node, unchanged by this binary
```

## Error Handling

- Any panic during `internal/examplecmd.List` is recovered by `dispatch` in
  `main.go`, which emits a single error result item containing the panic
  value.
- Callers never crash Alfred silently (empty output = Alfred spinner
  forever).

## Configuration Variables

None currently — `workflow/info.plist`'s `userconfigurationconfig` is an
empty array. See `docs/configuration-builder.md` for the mechanism to use
when a future command needs one.

## Constraints

- Script Filter response time target: **< 100 ms** (a compiled binary,
  usually well within budget).
- All output must go through `scriptfilter.Response.Write()` — never
  `fmt.Print*` directly.
- `cmd/example-alfred/main.go` contains no business logic; it only
  dispatches argv and writes the response.
