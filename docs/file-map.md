# File Map

> File-level dependency map for the alfred-workflow-template.
> Add to this as you explore the codebase during development.

## Entry Points

| File | Role |
|---|---|
| `cmd/example-alfred/main.go` | Alfred executes this binary — the sole entry point |

## Call Flow

```
cmd/example-alfred/main.go
  └─ dispatch(query)                          ← recover()-wrapped
       └─ internal/examplecmd.List(query)
            └─ internal/example.Filter(query)
       └─ scriptfilter.Response.Write(os.Stdout)
```

## Module Dependency Table

### `internal/scriptfilter/`

| File | Imports from | Notes |
|---|---|---|
| `scriptfilter.go` | stdlib only | Script Filter JSON types + `Response.Write` |

### `internal/`

| File | Imports from | Notes |
|---|---|---|
| `example/example.go` | stdlib only | Static shortcut list + substring filter — replace with your own |
| `examplecmd/examplecmd.go` | `internal/example`, `internal/scriptfilter` | Builds the Script Filter response |

### `cmd/example-alfred/`

| File | Imports from | Notes |
|---|---|---|
| `main.go` | `internal/examplecmd`, `internal/scriptfilter` | Alfred boundary; argv dispatch only |

## Key Files for Customization

When building a new workflow from this template, replace or update these files:

| File | What to change |
|---|---|
| `workflow/info.plist` | `bundleid`, keyword, UIDs, category, description |
| `cmd/example-alfred/` | Rename to `cmd/<workflow-name>-alfred/` |
| `internal/example/example.go` | Replace with your own domain logic |
| `internal/examplecmd/examplecmd.go` | Adjust result formatting for your domain |
| `go.mod` | Module path |
| `workflow/icon.png` | Workflow icon |
