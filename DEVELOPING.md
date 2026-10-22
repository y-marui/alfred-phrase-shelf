# Developing

This document covers the development workflow, conventions, and guidelines for contributors to this project.

## Prerequisites

- macOS (required for Alfred)
- Go (see `go.mod` for the toolchain version)
- Alfred 5 with Powerpack
- `jq` (optional, for pretty-printed dev output): `brew install jq`
- `gh` CLI (required for releases): `brew install gh`

## Development Setup

```bash
git clone https://github.com/y-marui/alfred-workflow-template
cd alfred-workflow-template
go build ./...
```

## Daily Workflow

### Simulate Alfred locally

```bash
go run ./cmd/example-alfred            # all shortcuts
go run ./cmd/example-alfred "doc"      # filtered
```

Pipe through `jq` for pretty-printed JSON:

```bash
go run ./cmd/example-alfred | jq .
```

### Run tests

```bash
make test          # go test ./...
```

### Lint and format

```bash
make lint          # gofmt -l + go vet
make fmt           # gofmt -w (auto-fix)
```

## Adding a New Command

This template ships a single command (`example`). To add another Script Filter
command:

1. Add the domain logic to a new `internal/<domain>/` package (stdlib only, Alfred-independent,
   unit-testable) — see `internal/example/example.go` for the shape to follow.
2. Add an `internal/<domain>cmd/` package with a function returning `scriptfilter.Response`,
   following `internal/examplecmd/examplecmd.go`'s shape.
3. If the workflow needs a second Alfred-invoked binary (e.g. a Script Filter plus a
   separate Run Script action with a side effect), add another `cmd/<name>-alfred/`
   directory — see `alfred-note-md-template`'s `cmd/note-md-template-alfred` +
   `cmd/note-md-template-paste-alfred` for that pattern. For a single command with an
   optional subcommand (e.g. a `help` case), dispatch with a plain `switch` inside
   `main.go` instead — see `alfred-password-generator`'s `passgencmd.handleHelp` for
   the pattern; don't introduce a generic router abstraction for this.
4. Add tests for both new packages.
5. Add a Script Filter (and, if it has a side effect, a Run Script) node to
   `workflow/info.plist`, wired to the new subcommand. Check
   `docs/alfred-workflow-notes/workflow-object-schema.md` before deciding a feature
   needs a Go binary at all — an Alfred-native object may already cover it.
6. Update `docs/specification.md`, `README.md`/`README-jp.md`, and `CHANGELOG.md`.

## Building the Package

```bash
make build-workflow
```

Output: `dist/<name>-<version>.alfredworkflow`

Install during development: double-click the `.alfredworkflow` file,
or drag it into Alfred Preferences → Workflows.

## Testing in Alfred

1. Build: `make build-workflow`
2. Install: open `dist/*.alfredworkflow`
3. Open Alfred, type `tpl`

During rapid iteration you can symlink `workflow/` to Alfred's workflow directory,
but `go run ./cmd/example-alfred "query"` is usually faster for logic changes.

## Naming Conventions

| Scope | Convention | Example |
|---|---|---|
| Go packages | short, lowercase, no underscores | `example`, `examplecmd`, `scriptfilter` |
| Exported functions / types | `PascalCase` | `Filter`, `Response`, `Item` |
| Unexported functions / variables | `camelCase` | `writeResponse`, `defaultTTL` |
| Alfred command names | lowercase | `"example"` |
| Alfred variable names | `lowercase_with_underscores` | (none currently — add via Config Builder) |
| Commit messages | Conventional Commits | `feat:`, `fix:`, `docs:`, `chore:` |
| Branch names | `feat/`, `fix/`, `docs/`, `chore/`, `work/` | `feat/add-open-browser` |

## Code Style

- **Formatter:** `gofmt`. CI enforces this (`make lint`).
- **Linter:** `go vet`.
- **Comments:** Write *why*, not *what*. Do not comment self-evident code.
- **No debug prints:** Remove all stray `fmt.Print*` statements before committing;
  the only writer to stdout is `scriptfilter.Response.Write`.
- **No third-party dependencies** unless clearly justified — keep `go.mod` dependency-free.

## Releasing

```bash
# 1. Update version in workflow/info.plist
# 2. Update CHANGELOG.md
git add workflow/info.plist CHANGELOG.md
git commit -m "chore: release v1.2.3"

# 3. Tag and push
git tag v1.2.3
git push origin main --tags
# GitHub Actions builds .alfredworkflow and creates a GitHub Release
```

## Commit Guidelines

- Commit per **feature unit**, after confirming it works.
- **No WIP commits** — do not commit code that does not run.
- **No `--no-verify`** — never skip pre-commit hooks.

### Commit Message Format

Follow [Conventional Commits](https://www.conventionalcommits.org/):

```
feat: add clipboard copy action
fix: substring filter matches on stale query
chore: update Go toolchain to 1.28
docs: update README usage section
refactor: simplify examplecmd dispatch logic
```

## Pull Request Checklist

- [ ] `make lint` passes
- [ ] `make test` passes
- [ ] `make build-workflow` succeeds
- [ ] New commands have tests
- [ ] `README.md`/`README-jp.md` updated if user-facing changes
- [ ] `CHANGELOG.md` entry added under `[Unreleased]`

## Code Review Guidelines

**Reviewers check for:**
- Architectural constraints respected (no business logic in `cmd/example-alfred`, no layer skipping)
- No hardcoded absolute paths (use `$HOME` / env vars)
- No debug prints in production code
- No Unicode emoji in Alfred result item `title` / `subtitle`
- Tests cover the new or changed behavior
- Alfred env variables managed via Config Builder, not `variables` key

**Security-sensitive changes** (auth, encryption, data access) require explicit
security review before merge.

**Self-review:** Individual contributors open a PR and self-review before merging
to `main`.
