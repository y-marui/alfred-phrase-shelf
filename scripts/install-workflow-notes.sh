#!/usr/bin/env bash
# alfred-workflow-notes installer/updater
#
# docs/alfred-workflow-notes/ is shared documentation that lives inside this
# repo (alfred-workflow-template) at that subdirectory, not at the repo
# root (unlike dev-charter, whose repo root IS the shared content). A plain
# `git subtree add/pull` would therefore pull this whole template repo in;
# instead we `git subtree split` that subdirectory's history into a
# synthetic commit first, then add/merge just that commit.
#
# Usage (from an adopting project's root):
#   bash <(curl -fsSL https://raw.githubusercontent.com/y-marui/alfred-workflow-template/main/scripts/install-workflow-notes.sh)
#
# Environment variables (all optional):
#   WORKFLOW_NOTES_REMOTE  git remote name    (default: alfred-workflow-notes)
#   WORKFLOW_NOTES_URL     repository URL     (default: https://github.com/y-marui/alfred-workflow-template)
#   WORKFLOW_NOTES_BRANCH  source branch      (default: main)
#   WORKFLOW_NOTES_PREFIX  install directory  (default: docs/alfred-workflow-notes)

set -euo pipefail

REMOTE_NAME="${WORKFLOW_NOTES_REMOTE:-alfred-workflow-notes}"
REMOTE_URL="${WORKFLOW_NOTES_URL:-https://github.com/y-marui/alfred-workflow-template}"
BRANCH="${WORKFLOW_NOTES_BRANCH:-main}"
PREFIX="${WORKFLOW_NOTES_PREFIX:-docs/alfred-workflow-notes}"

if ! git rev-parse --git-dir > /dev/null 2>&1; then
    echo "Error: not in a git repository. Run this script from your project root." >&2
    exit 1
fi

if ! git remote get-url "$REMOTE_NAME" > /dev/null 2>&1; then
    echo "Adding remote '$REMOTE_NAME'..."
    git remote add "$REMOTE_NAME" "$REMOTE_URL"
fi

echo "Fetching $REMOTE_NAME..."
git fetch "$REMOTE_NAME"

# git subtree fails on a dirty working tree, so stash first and restore
# afterward (like the dev-charter installer).
STASHED=0
if ! git diff --quiet || ! git diff --cached --quiet || [ -n "$(git ls-files --others --exclude-standard)" ]; then
    echo "Stashing uncommitted changes before updating..."
    git stash push -u -m "install-workflow-notes.sh update"
    STASHED=1
fi

echo "Splitting $PREFIX out of ${REMOTE_NAME}/${BRANCH}..."
SPLIT_SHA=$(git subtree split --prefix="$PREFIX" "${REMOTE_NAME}/${BRANCH}")

if [ -d "$PREFIX" ]; then
    echo "$PREFIX already exists. Merging..."
    git subtree merge --prefix="$PREFIX" "$SPLIT_SHA" --squash
else
    echo "Installing $PREFIX for the first time..."
    git subtree add --prefix="$PREFIX" "$SPLIT_SHA" --squash
fi

if [ "$STASHED" = "1" ]; then
    echo "Restoring stashed changes..."
    git stash pop
fi

echo "Done."
