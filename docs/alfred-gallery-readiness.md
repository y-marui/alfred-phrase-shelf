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
| No self-installed external software | ✅ Done | `go.mod` carries no third-party dependencies; nothing is fetched at runtime |
| Icon ≥ 256×256 px, square | ⏳ Verify | Replace `workflow/icon.png` with a real icon for your project; confirm size/squareness with `sips -g pixelWidth -g pixelHeight` |
| Keyword ≥ 3 characters | ⏳ Verify | Confirm your chosen keyword (replacing the template's `tpl`) is 3+ characters and doesn't collide with common workflows |
| Category set | ⏳ Verify | Replace the placeholder `Tools & Utilities` in `workflow/info.plist`'s `category` key with the value that actually fits your workflow — see the valid values in [`docs/alfred-workflow-notes/workflow-object-schema.md`](docs/alfred-workflow-notes/workflow-object-schema.md#category) |
| User Configuration over environment variables | ⏳ Verify | Add to Alfred's Configuration Builder (`workflow/info.plist`'s `userconfigurationconfig`) if your project needs user-facing settings; N/A otherwise |
| English instructions in About/README | ⏳ Verify | Confirm `README.md` (the reference/English version per `LANGUAGE_POLICY.md`) is complete once renamed from `README_TEMPLATE.md` |
| README follows Gallery style guide | ⏳ Verify | `## Usage` should open with 1-2 sentences ending "via the `<keyword>` keyword" (or "via the Universal Action"), followed by one screenshot per entry point; modifier keys as a `* <kbd>...</kbd> action` bullet list, not a table |
| Screenshots (full Alfred window, shadow, no background) | ❌ Missing | Needs a real Alfred window capture — not something this repository's automation can produce; track as a GitHub issue when the project is ready for it |

## Out of scope here

- Posting to the Alfred Forum and the invitation-only Gallery submission itself — a per-project decision, not something this checklist mandates
- Exporting a Developer ID certificate, generating a notarization API key, and registering the GitHub Actions secrets `release.yml` expects — manual steps only the repository owner can perform
