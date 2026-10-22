// Package examplecmd builds the Alfred Script Filter response for the
// "example" keyword from internal/example.
package examplecmd

import (
	"github.com/y-marui/alfred-workflow-template/internal/example"
	"github.com/y-marui/alfred-workflow-template/internal/scriptfilter"
)

// List returns the Script Filter response for query: one row per matching
// shortcut (Enter opens its URL, via the workflow's native Open URL node),
// or a single non-actionable row when nothing matches.
func List(query string) scriptfilter.Response {
	shortcuts := example.Filter(query)

	if len(shortcuts) == 0 {
		return scriptfilter.Response{Items: []scriptfilter.Item{
			{
				Title:    "No matching shortcut",
				Subtitle: "Try a different query",
				Valid:    scriptfilter.BoolPtr(false),
			},
		}}
	}

	items := make([]scriptfilter.Item, len(shortcuts))
	for i, s := range shortcuts {
		items[i] = scriptfilter.Item{
			UID:          "example-" + s.Name,
			Title:        s.Name,
			Subtitle:     s.URL,
			Arg:          s.URL,
			Valid:        scriptfilter.BoolPtr(true),
			Autocomplete: s.Name,
		}
	}
	return scriptfilter.Response{Items: items}
}
