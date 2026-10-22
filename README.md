# Alfred Workflow Template

> **This is the English (reference) version.**
> For the Japanese canonical version, see [README-jp.md](README-jp.md).

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)
[![CI](https://github.com/y-marui/alfred-workflow-template/actions/workflows/ci.yml/badge.svg)](https://github.com/y-marui/alfred-workflow-template/actions/workflows/ci.yml)
[![Charter Check](https://github.com/y-marui/alfred-workflow-template/actions/workflows/dev-charter-check.yml/badge.svg)](https://github.com/y-marui/alfred-workflow-template/actions/workflows/dev-charter-check.yml)
[![GitHub Sponsors](https://img.shields.io/github/sponsors/y-marui)](https://github.com/sponsors/y-marui)
[![Buy Me a Coffee](https://img.shields.io/badge/Buy%20Me%20a%20Coffee-ffdd00?logo=buy-me-a-coffee&logoColor=black)](https://www.buymeacoffee.com/y.marui)

| Field | Value |
|---|---|
| Target | Alfred 5 Script Filter workflow |
| Team size | Individual to small team (1–3 people) |
| Language | English (OSS) |
| License | MIT |
| Runtime | Go (see `go.mod`), Alfred 5 |
| AI tools | Claude Code / GitHub Copilot / Gemini CLI |

> Production-ready template for building Alfred 5 Script Filter workflows in Go.
> Start shipping in 10 minutes.

## Features

- ✅ **Layered architecture** — Alfred boundary (`cmd/`) isolated from domain logic (`internal/`)
- ✅ **Reusable Script Filter JSON layer** — `internal/scriptfilter`, Alfred-independent
- ✅ **Single universal binary** — `darwin/amd64` + `darwin/arm64` merged via `lipo`, no runtime interpreter
- ✅ **Full test suite** — `go test`, no Alfred required to run tests
- ✅ **CI/CD** — lint, test, build, and release via GitHub Actions
- ✅ **Zero dependencies** — `go.mod` has no `require` block
- ✅ **AI-ready** — `AI_CONTEXT.md` + `CLAUDE.md` for AI assistant context

## Requirements

- Alfred 5 (Powerpack required for Script Filter)
- Go (see `go.mod` for the toolchain version)
- [pre-commit](https://pre-commit.com/) (for security hooks)

## Quick Start

### Using this template

1. Click **"Use this template"** → **"Create a new repository"** on GitHub
2. Clone your new repository and open it in your AI tool
3. Tell the AI: "Run the initial setup from AI_CONTEXT.md" — it will:
   - Apply GitHub repository settings
   - Rename `README_TEMPLATE.md` → `README.md` and `README_TEMPLATE-jp.md` → `README-jp.md`
   - Replace `{user}`, `{repo}`, `{keyword}` placeholders
   - Randomize the dev-charter cron schedule
4. Customize the workflow (see `DEVELOPING.md`)

> **Note:** When rewriting the README, the Charter Check badge (for `dev-charter-check.yml`)
> is easy to lose. Verify it's still present in both `README.md` and `README-jp.md`.

### Development (this template)

```bash
git clone https://github.com/y-marui/alfred-workflow-template
cd alfred-workflow-template

# Simulate Alfred locally
go run ./cmd/example-alfred
go run ./cmd/example-alfred "doc"

# Run tests
make test

# Build workflow package
make build-workflow
# → dist/workflow-template-1.0.0.alfredworkflow
```

Double-click `dist/*.alfredworkflow` to install in Alfred.

## Usage

Search example shortcuts (repo, docs, issues) via the `tpl` keyword.

```
tpl          -> list example shortcuts (repo, docs, issues)
tpl <query>  -> filter shortcuts by name
```

* <kbd>↩︎</kbd> Open the selected shortcut's URL

### Troubleshooting

**No results appear**
- Check Alfred's debugger: open Alfred → `⌘D`

## Project Structure

```
alfred-workflow-template/
├── cmd/
│   └── example-alfred/     # The binary Alfred invokes (argv dispatch only)
├── internal/
│   ├── example/            # Domain logic — replace with your own
│   ├── examplecmd/         # Builds the Script Filter response
│   └── scriptfilter/       # Script Filter JSON types (Alfred-independent, reusable)
├── workflow/                # Alfred package (info.plist, icon.png)
├── scripts/                 # build-workflow.sh, extract-changelog.sh
└── docs/                    # Architecture and reference documentation
```

## Documentation

| Document | Description |
|---|---|
| [docs/architecture.md](docs/architecture.md) | Module and layer design |
| [docs/file-map.md](docs/file-map.md) | File-level dependency map |
| [docs/specification.md](docs/specification.md) | Feature specification and data flow |
| [docs/ui-design.md](docs/ui-design.md) | Alfred result item UI conventions |
| [docs/configuration-builder.md](docs/configuration-builder.md) | This project's Configuration Builder settings |
| [docs/alfred-workflow-notes/](docs/alfred-workflow-notes/README.md) | Cross-project Alfred workflow dev reference (canonical source for other projects, via `git subtree`) |

## AI-Assisted Development

This template is configured for AI-assisted development.

| Tool | Role |
|---|---|
| Claude Code | Architecture, large-scale changes, refactoring |
| GitHub Copilot | Bug fixes, small implementation, unit tests |
| Gemini CLI | Documentation management |

See [`AI_CONTEXT.md`](AI_CONTEXT.md) and [`CLAUDE.md`](CLAUDE.md) for session context.

## Customizing This Template

After running the initial setup (see Quick Start above), customize the workflow:

1. Edit `workflow/info.plist`:
   - Replace `bundleid` with your bundle ID (`com.yourname.workflowname`)
   - Replace the `keyword` (`wf`) with your trigger keyword
   - Run `uuidgen` and replace the placeholder UIDs
2. Rename `cmd/example-alfred` to `cmd/<your-workflow-name>-alfred`
3. Replace `internal/example` + `internal/examplecmd` with your own domain logic
   (keep `internal/scriptfilter` as-is)
4. Update the module path in `go.mod`
5. Add your `workflow/icon.png`

Before writing any Go code to talk to macOS (clipboard, keystrokes, notifications),
check [docs/alfred-workflow-notes/workflow-object-schema.md](docs/alfred-workflow-notes/workflow-object-schema.md) —
an Alfred-native object may already cover it without any code at all.

## Release

```bash
# 1. Bump version in workflow/info.plist
# 2. Tag and push
git tag v1.2.3
git push --tags
# GitHub Actions builds .alfredworkflow and creates a GitHub Release
```

## License

MIT — see [LICENSE](LICENSE)

---

*This is the reference (English) version. The canonical Japanese version is [README-jp.md](README-jp.md). Update both files in the same commit.*
