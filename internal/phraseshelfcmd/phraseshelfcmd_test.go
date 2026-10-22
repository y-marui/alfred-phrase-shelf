package phraseshelfcmd

import (
	"errors"
	"testing"

	"github.com/y-marui/alfred-phrase-shelf/internal/phraseshelf"
)

// fakeStore is a Store test double — real osascript/the real app is never
// called from tests.
type fakeStore struct {
	all        []phraseshelf.Snippet
	byCategory map[string][]phraseshelf.Snippet
	err        error
}

func (f fakeStore) ListSnippets() ([]phraseshelf.Snippet, error) {
	if f.err != nil {
		return nil, f.err
	}
	return f.all, nil
}

func (f fakeStore) ListSnippetsByCategory(category string) ([]phraseshelf.Snippet, error) {
	if f.err != nil {
		return nil, f.err
	}
	return f.byCategory[category], nil
}

func newFixtureStore() fakeStore {
	work := []phraseshelf.Snippet{
		{ID: "1", Name: "greeting", Text: "Hello there!\nSecond line.", Category: "work"},
		{ID: "2", Name: "bye", Text: "Bye!", Category: "work"},
	}
	loose := []phraseshelf.Snippet{
		{ID: "3", Name: "memo", Text: "Memo.", Category: ""},
	}
	return fakeStore{
		all: append(append([]phraseshelf.Snippet{}, work...), loose...),
		byCategory: map[string][]phraseshelf.Snippet{
			"work": work,
			"":     loose,
		},
	}
}

func TestListNoHashListsCategories(t *testing.T) {
	resp := List(newFixtureStore(), "")
	if len(resp.Items) != 2 {
		t.Fatalf("len(Items) = %d, want 2", len(resp.Items))
	}
	if resp.Items[0].Title != "work" {
		t.Errorf("Items[0].Title = %q, want %q", resp.Items[0].Title, "work")
	}
	if resp.Items[0].Autocomplete != "work#" {
		t.Errorf("Items[0].Autocomplete = %q, want %q", resp.Items[0].Autocomplete, "work#")
	}
	if resp.Items[0].Valid == nil || *resp.Items[0].Valid {
		t.Error("Items[0].Valid should be false (autocomplete-only)")
	}
}

func TestListNoHashUncategorizedLabel(t *testing.T) {
	resp := List(newFixtureStore(), "")
	last := resp.Items[len(resp.Items)-1]
	if last.Title != "(Uncategorized)" {
		t.Errorf("last.Title = %q, want %q", last.Title, "(Uncategorized)")
	}
	if last.Autocomplete != "#" {
		t.Errorf("last.Autocomplete = %q, want %q (empty category + '#')", last.Autocomplete, "#")
	}
}

func TestListNoHashFuzzyFiltersCategories(t *testing.T) {
	resp := List(newFixtureStore(), "wk")
	if len(resp.Items) != 1 || resp.Items[0].Title != "work" {
		t.Fatalf("Items = %+v, want a single \"work\" category", resp.Items)
	}
}

func TestListNoHashNoMatchReturnsNonActionableRow(t *testing.T) {
	resp := List(newFixtureStore(), "nonexistent-category")
	if len(resp.Items) != 1 {
		t.Fatalf("len(Items) = %d, want 1", len(resp.Items))
	}
	if resp.Items[0].Valid == nil || *resp.Items[0].Valid {
		t.Error("Valid should be false for the no-match row")
	}
}

func TestListWithHashListsSnippetsInCategory(t *testing.T) {
	resp := List(newFixtureStore(), "work#")
	if len(resp.Items) != 2 {
		t.Fatalf("len(Items) = %d, want 2", len(resp.Items))
	}
	got := resp.Items[0]
	if got.Title != "greeting" {
		t.Errorf("Title = %q, want %q", got.Title, "greeting")
	}
	if got.Arg != "Hello there!\nSecond line." {
		t.Errorf("Arg = %q, want the full snippet text", got.Arg)
	}
	if got.Subtitle != "Hello there!" {
		t.Errorf("Subtitle = %q, want the first line only", got.Subtitle)
	}
	if got.Valid == nil || !*got.Valid {
		t.Error("Valid should be true — this is the final, actionable item")
	}
}

func TestListWithHashVerbatimCategoryNotFuzzyMatched(t *testing.T) {
	// "wor" must NOT match "work" — the category before '#' is exact.
	resp := List(newFixtureStore(), "wor#")
	if len(resp.Items) != 1 {
		t.Fatalf("len(Items) = %d, want 1", len(resp.Items))
	}
	if resp.Items[0].Valid == nil || *resp.Items[0].Valid {
		t.Fatal("expected the no-snippets-in-category row, got an actionable item")
	}
}

func TestListWithHashFuzzyFiltersSnippetNames(t *testing.T) {
	resp := List(newFixtureStore(), "work#by")
	if len(resp.Items) != 1 || resp.Items[0].Title != "bye" {
		t.Fatalf("Items = %+v, want a single \"bye\" snippet", resp.Items)
	}
}

func TestListWithHashUncategorizedCategory(t *testing.T) {
	resp := List(newFixtureStore(), "#memo")
	if len(resp.Items) != 1 || resp.Items[0].Title != "memo" {
		t.Fatalf("Items = %+v, want a single \"memo\" snippet", resp.Items)
	}
}

func TestListWithHashEmptyCategoryShowsInfoItem(t *testing.T) {
	resp := List(newFixtureStore(), "nonexistent-category#")
	if len(resp.Items) != 1 {
		t.Fatalf("len(Items) = %d, want 1", len(resp.Items))
	}
	if resp.Items[0].Valid == nil || *resp.Items[0].Valid {
		t.Error("Valid should be false for the empty-category info row")
	}
}

func TestListWithHashLiteralFurtherHashesStayInFilterText(t *testing.T) {
	resp := List(newFixtureStore(), "work#a#b")
	// "a#b" as a fuzzy query against "bye"/"greeting" names: neither
	// contains the literal '#', so nothing should match.
	if len(resp.Items) != 1 || resp.Items[0].Valid == nil || *resp.Items[0].Valid {
		t.Fatalf("Items = %+v, want the no-match row (filter text keeps literal '#')", resp.Items)
	}
}

func TestListStoreErrorReturnsSingleNonValidItem(t *testing.T) {
	store := fakeStore{err: errors.New("Phrase Shelf is unreachable: boom")}

	for _, query := range []string{"", "work#"} {
		resp := List(store, query)
		if len(resp.Items) != 1 {
			t.Fatalf("List(%q): len(Items) = %d, want 1", query, len(resp.Items))
		}
		if resp.Items[0].Valid == nil || *resp.Items[0].Valid {
			t.Errorf("List(%q): Valid should be false on a Store error", query)
		}
	}
}

func TestListLeadingSpaceListsAllCategories(t *testing.T) {
	want := len(List(newFixtureStore(), "").Items)
	got := len(List(newFixtureStore(), " ").Items)
	if got != want {
		t.Fatalf("List(\" \"): len(Items) = %d, want %d", got, want)
	}
}
