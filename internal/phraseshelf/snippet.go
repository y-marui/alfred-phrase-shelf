// Package phraseshelf is the domain logic for talking to the "Phrase Shelf"
// macOS app and filtering its snippets. It never imports
// internal/scriptfilter: this package has no Alfred JSON concerns.
package phraseshelf

// Snippet is one saved phrase, as returned by Phrase Shelf's AppleScript
// "list snippets" command. Category is "" for the uncategorized group.
type Snippet struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Text     string `json:"text"`
	Category string `json:"category"`
}

// Store fetches snippets from Phrase Shelf. The real implementation
// (OsascriptStore) shells out to osascript; tests use a fake instead.
type Store interface {
	// ListSnippets returns every snippet, sortOrder ascending.
	ListSnippets() ([]Snippet, error)
	// ListSnippetsByCategory returns only the snippets in category.
	// category == "" returns the uncategorized group.
	ListSnippetsByCategory(category string) ([]Snippet, error)
}
