# Alfred Gallery Readiness

Starting checklist for tracking this project's compliance with the
[Alfred Gallery submission requirements](https://alfred.app/submit/) and
[style guide](https://alfred.app/submit/styleguide/) — see dev-charter's
[`topics/alfred/ALFRED_GALLERY.md`](docs/dev-charter/topics/alfred/ALFRED_GALLERY.md)
for the generalized checklist this is based on. This is a checklist against
an external, occasionally-changing policy — re-read the linked pages before
acting on stale entries here.

Copy this file into new projects created from this template (it isn't itself
submitted to the Gallery, so nothing here applies to the template repo's own
demo Workflow) and update the Status/Notes columns as the project matures.

## Checklist

| Requirement | Status | Notes |
|---|---|---|
| Binaries signed and notarised | ⏳ Pending | `.github/workflows/release.yml` signs/notarises on a tagged release once the five Apple/GitHub secrets exist (see that workflow); unverified until a real release ships |
| No self-update | ✅ Done | Updates ship only as new `.alfredworkflow` releases — the architecture has no self-update code path |
| No self-installed external software | ✅ Done | `go.mod` carries no third-party dependencies; nothing is fetched at runtime (the `osascript` call talks to the user's own already-installed Phrase Shelf app, not a download) |
| Icon ≥ 256×256 px, square | ❌ Missing | `workflow/icon.png` is still the template's 1×1px placeholder — needs a real icon before Gallery submission; confirm size/squareness with `sips -g pixelWidth -g pixelHeight` |
| Keyword ≥ 3 characters | ✅ Done | `snip` (4 characters) |
| Category set | ✅ Done | `Productivity` in `workflow/info.plist`'s `category` key |
| User Configuration over environment variables | ✅ N/A | No user-facing settings yet; `userconfigurationconfig` is an empty array |
| English instructions in About/README | ✅ Done | `README.md` is the complete English/reference version |
| README follows Gallery style guide | ⏳ Verify | Current `## Usage` doesn't yet follow the Gallery's exact style (ending "via the `snip` keyword", one screenshot per entry point, modifier keys as a `* <kbd>...</kbd>` list) — revisit before submission |
| Screenshots (full Alfred window, shadow, no background) | ❌ Missing | Needs a real Alfred window capture — not something this repository's automation can produce; track as a GitHub issue when the project is ready for it |

## Out of scope here

- Posting to the Alfred Forum and the invitation-only Gallery submission itself — a per-project decision, not something this checklist mandates
- Exporting a Developer ID certificate, generating a notarization API key, and registering the GitHub Actions secrets `release.yml` expects — manual steps only the repository owner can perform
