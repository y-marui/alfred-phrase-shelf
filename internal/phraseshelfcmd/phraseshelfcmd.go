// Package phraseshelfcmd builds the Alfred Script Filter response for the
// "snip" keyword from internal/phraseshelf.
package phraseshelfcmd

import (
	"fmt"
	"strings"

	"github.com/y-marui/alfred-phrase-shelf/internal/phraseshelf"
	"github.com/y-marui/alfred-phrase-shelf/internal/scriptfilter"
)

// uncategorizedLabel is the title shown for the "" category in Stage 1.
// The underlying category value stays "" wherever it flows into Stage 2.
const uncategorizedLabel = "(Uncategorized)"

// List dispatches query to Stage 1 (category list) or Stage 2 (snippet
// list within a category), per internal/phraseshelf.SplitQuery.
func List(store phraseshelf.Store, query string) scriptfilter.Response {
	// A lone space after the keyword (e.g. "snip ") can reach us as " "; left
	// in, it fuzzy-matches only categories whose name contains a space.
	query = strings.TrimLeft(query, " \t")
	category, filterText, hasCategory := phraseshelf.SplitQuery(query)
	if !hasCategory {
		return stage1(store, filterText)
	}
	return stage2(store, category, filterText)
}

// stage1 lists categories matching filterText. Each item is autocomplete-
// only (not directly actionable): accepting it rewrites the Alfred query to
// "<category>#", dropping the user into Stage 2.
func stage1(store phraseshelf.Store, filterText string) scriptfilter.Response {
	all, err := store.ListSnippets()
	if err != nil {
		return errorResponse(err)
	}

	counts := phraseshelf.CountByCategory(all)
	var matched []string
	for _, category := range phraseshelf.Categories(all) {
		if phraseshelf.FuzzyMatch(filterText, category) {
			matched = append(matched, category)
		}
	}

	if len(matched) == 0 {
		return scriptfilter.Response{Items: []scriptfilter.Item{
			{
				Title:    "No matching category",
				Subtitle: "Try a different query",
				Valid:    scriptfilter.BoolPtr(false),
			},
		}}
	}

	items := make([]scriptfilter.Item, len(matched))
	for i, category := range matched {
		items[i] = scriptfilter.Item{
			Title:        displayCategory(category),
			Subtitle:     fmt.Sprintf("%d snippet(s)", counts[category]),
			Valid:        scriptfilter.BoolPtr(false),
			Autocomplete: category + "#",
		}
	}
	return scriptfilter.Response{Items: items}
}

// stage2 lists snippets in category matching filterText. Each item is the
// final, actionable result: its Arg carries the snippet's full text, for
// the workflow's Copy to Clipboard (auto-paste) node downstream.
func stage2(store phraseshelf.Store, category, filterText string) scriptfilter.Response {
	snippets, err := store.ListSnippetsByCategory(category)
	if err != nil {
		return errorResponse(err)
	}

	if len(snippets) == 0 {
		return scriptfilter.Response{Items: []scriptfilter.Item{
			{
				Title:    "No snippets in " + displayCategory(category),
				Subtitle: "Try a different category",
				Valid:    scriptfilter.BoolPtr(false),
			},
		}}
	}

	var matched []phraseshelf.Snippet
	for _, s := range snippets {
		if phraseshelf.FuzzyMatch(filterText, s.Name) {
			matched = append(matched, s)
		}
	}

	if len(matched) == 0 {
		return scriptfilter.Response{Items: []scriptfilter.Item{
			{
				Title:    "No matching snippet",
				Subtitle: "Try a different query",
				Valid:    scriptfilter.BoolPtr(false),
			},
		}}
	}

	items := make([]scriptfilter.Item, len(matched))
	for i, s := range matched {
		items[i] = scriptfilter.Item{
			UID:      "phrase-shelf-" + s.ID,
			Title:    s.Name,
			Subtitle: preview(s.Text),
			Arg:      s.Text,
			Valid:    scriptfilter.BoolPtr(true),
		}
	}
	return scriptfilter.Response{Items: items}
}

func displayCategory(category string) string {
	if category == "" {
		return uncategorizedLabel
	}
	return category
}

// preview returns the first line of text, truncated for the subtitle.
const previewMaxLen = 60

func preview(text string) string {
	for i, r := range text {
		if r == '\n' {
			text = text[:i]
			break
		}
	}
	runes := []rune(text)
	if len(runes) > previewMaxLen {
		return string(runes[:previewMaxLen]) + "…"
	}
	return text
}

// errorResponse shows the Store failure (app not launchable, AppleScript
// error) as a single non-valid item rather than crashing or returning
// nothing.
func errorResponse(err error) scriptfilter.Response {
	return scriptfilter.Response{Items: []scriptfilter.Item{
		{
			Title:    "Phrase Shelf is unreachable",
			Subtitle: err.Error(),
			Valid:    scriptfilter.BoolPtr(false),
		},
	}}
}
