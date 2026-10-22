# Configuration Builder — This Project

For the general Alfred Configuration Builder mechanism (widget types,
config keys, how variables reach scripts), see the cross-project reference:
[`docs/alfred-workflow-notes/configuration-builder.md`](alfred-workflow-notes/configuration-builder.md).

This file covers only what's specific to this project.

Variable names in this project use **lowercase with underscores**.

## This project's configuration

None currently — `workflow/info.plist`'s `userconfigurationconfig` is an
empty array, since the example command (a static shortcut list) has nothing
to configure.

When you add a setting a future command needs, register it here in this
same table (name, `select`/`checkbox`/`textfield`/`file`/`password` type,
default, and what it does), following the general reference's widget
patterns. Read it via `os.Getenv("your_variable_name")` in Go.
