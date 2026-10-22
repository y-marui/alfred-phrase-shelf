package examplecmd

import "testing"

func TestListEmptyQueryReturnsAllShortcuts(t *testing.T) {
	resp := List("")
	if len(resp.Items) != 3 {
		t.Fatalf("len(Items) = %d, want 3", len(resp.Items))
	}
}

func TestListMatchingQueryReturnsFilteredItem(t *testing.T) {
	resp := List("doc")
	if len(resp.Items) != 1 {
		t.Fatalf("len(Items) = %d, want 1", len(resp.Items))
	}
	got := resp.Items[0]
	if got.Title != "docs" {
		t.Errorf("Title = %q, want %q", got.Title, "docs")
	}
	if got.Arg == "" {
		t.Error("Arg is empty, want the shortcut's URL")
	}
	if got.Valid == nil || !*got.Valid {
		t.Error("Valid should be true for a matching shortcut")
	}
}

func TestListNoMatchReturnsNonActionableRow(t *testing.T) {
	resp := List("nonexistent-shortcut-name")
	if len(resp.Items) != 1 {
		t.Fatalf("len(Items) = %d, want 1", len(resp.Items))
	}
	got := resp.Items[0]
	if got.Valid == nil || *got.Valid {
		t.Error("Valid should be false for the no-match row")
	}
	if got.Arg != "" {
		t.Errorf("Arg = %q, want empty for the no-match row", got.Arg)
	}
}
