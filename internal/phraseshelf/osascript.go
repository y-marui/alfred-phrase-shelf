package phraseshelf

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"
)

// OsascriptStore is the real Store: it shells out to osascript to drive the
// "Phrase Shelf" app's AppleScript interface (PhraseShelf.sdef). `tell
// application "Phrase Shelf"` auto-launches the app if it isn't running,
// and the first call triggers a one-time macOS Automation permission
// prompt for the calling process.
type OsascriptStore struct{}

// NewOsascriptStore returns a Store backed by the real "Phrase Shelf" app.
func NewOsascriptStore() OsascriptStore {
	return OsascriptStore{}
}

func (OsascriptStore) ListSnippets() ([]Snippet, error) {
	return runListSnippets(`tell application "Phrase Shelf" to list snippets`)
}

func (OsascriptStore) ListSnippetsByCategory(category string) ([]Snippet, error) {
	script := fmt.Sprintf(`tell application "Phrase Shelf" to list snippets category "%s"`,
		escapeAppleScriptString(category))
	return runListSnippets(script)
}

func runListSnippets(script string) ([]Snippet, error) {
	cmd := exec.Command("osascript", "-e", script)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		msg := strings.TrimSpace(stderr.String())
		if msg == "" {
			msg = err.Error()
		}
		return nil, fmt.Errorf("Phrase Shelf is unreachable: %s", msg)
	}

	snippets, err := parseSnippetsJSON(stdout.Bytes())
	if err != nil {
		return nil, fmt.Errorf("parsing Phrase Shelf's response: %w", err)
	}
	return snippets, nil
}

// parseSnippetsJSON decodes the JSON array "list snippets" returns. Split
// out from runListSnippets so it's testable without calling osascript.
func parseSnippetsJSON(data []byte) ([]Snippet, error) {
	var snippets []Snippet
	if err := json.Unmarshal(data, &snippets); err != nil {
		return nil, err
	}
	return snippets, nil
}

// escapeAppleScriptString escapes s for interpolation into a double-quoted
// AppleScript string literal.
func escapeAppleScriptString(s string) string {
	var b strings.Builder
	for _, r := range s {
		switch r {
		case '\\':
			b.WriteString(`\\`)
		case '"':
			b.WriteString(`\"`)
		default:
			b.WriteRune(r)
		}
	}
	return b.String()
}
