package phraseshelf

import (
	"encoding/json"
	"reflect"
	"testing"
)

// fixture mirrors a real "list snippets" response: named categories in the
// app's configured order, then the uncategorized ("") group last.
const fixtureJSON = `[
	{"id": "1", "name": "hello", "text": "Hello!", "category": "work"},
	{"id": "2", "name": "bye", "text": "Bye!", "category": "work"},
	{"id": "3", "name": "memo", "text": "Memo.", "category": "personal"},
	{"id": "4", "name": "loose", "text": "Loose.", "category": ""}
]`

func fixtureSnippets(t *testing.T) []Snippet {
	t.Helper()
	snippets, err := parseSnippetsJSON([]byte(fixtureJSON))
	if err != nil {
		t.Fatalf("parseSnippetsJSON: %v", err)
	}
	return snippets
}

func TestCategoriesOrderedByFirstAppearance(t *testing.T) {
	got := Categories(fixtureSnippets(t))
	want := []string{"work", "personal", ""}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Categories() = %v, want %v", got, want)
	}
}

func TestCategoriesDeduplicates(t *testing.T) {
	got := Categories(fixtureSnippets(t))
	if len(got) != 3 {
		t.Fatalf("len(Categories()) = %d, want 3 (deduplicated)", len(got))
	}
}

func TestCategoriesEmptyInput(t *testing.T) {
	got := Categories(nil)
	if len(got) != 0 {
		t.Fatalf("Categories(nil) = %v, want empty", got)
	}
}

func TestCountByCategory(t *testing.T) {
	got := CountByCategory(fixtureSnippets(t))
	want := map[string]int{"work": 2, "personal": 1, "": 1}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("CountByCategory() = %v, want %v", got, want)
	}
}

func TestParseSnippetsJSONDecodesFields(t *testing.T) {
	var snippets []Snippet
	if err := json.Unmarshal([]byte(fixtureJSON), &snippets); err != nil {
		t.Fatalf("Unmarshal: %v", err)
	}
	if len(snippets) != 4 {
		t.Fatalf("len(snippets) = %d, want 4", len(snippets))
	}
	first := snippets[0]
	if first.ID != "1" || first.Name != "hello" || first.Text != "Hello!" || first.Category != "work" {
		t.Errorf("snippets[0] = %+v, want {ID:1 Name:hello Text:Hello! Category:work}", first)
	}
}
