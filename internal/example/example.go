// Package example is the placeholder domain logic this template ships —
// replace it with your own. It never imports internal/scriptfilter: this
// package has no Alfred JSON concerns, only pure logic that's unit
// testable without Alfred running.
package example

import "strings"

// Shortcut is a named link this example workflow can open.
type Shortcut struct {
	Name string
	URL  string
}

// Shortcuts is the static list this example ships. Replace with your own
// domain data, or with a function that computes results at request time.
var Shortcuts = []Shortcut{
	{Name: "repo", URL: "https://github.com/y-marui/alfred-workflow-template"},
	{Name: "docs", URL: "https://github.com/y-marui/alfred-workflow-template/tree/main/docs"},
	{Name: "issues", URL: "https://github.com/y-marui/alfred-workflow-template/issues"},
}

// Filter returns the shortcuts whose Name contains query, case-insensitively,
// preserving Shortcuts' order. An empty (or whitespace-only) query returns
// every shortcut.
func Filter(query string) []Shortcut {
	query = strings.ToLower(strings.TrimSpace(query))
	if query == "" {
		return Shortcuts
	}

	matched := make([]Shortcut, 0, len(Shortcuts))
	for _, s := range Shortcuts {
		if strings.Contains(strings.ToLower(s.Name), query) {
			matched = append(matched, s)
		}
	}
	return matched
}
