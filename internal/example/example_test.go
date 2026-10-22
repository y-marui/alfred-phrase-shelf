package example

import "testing"

func TestFilterEmptyQueryReturnsAllShortcuts(t *testing.T) {
	got := Filter("")
	if len(got) != len(Shortcuts) {
		t.Fatalf("len(Filter(\"\")) = %d, want %d", len(got), len(Shortcuts))
	}
}

func TestFilterWhitespaceOnlyQueryReturnsAllShortcuts(t *testing.T) {
	got := Filter("   ")
	if len(got) != len(Shortcuts) {
		t.Fatalf("len(Filter(\"   \")) = %d, want %d", len(got), len(Shortcuts))
	}
}

func TestFilterMatchesSubstringCaseInsensitively(t *testing.T) {
	got := Filter("DOC")
	if len(got) != 1 || got[0].Name != "docs" {
		t.Fatalf("Filter(\"DOC\") = %+v, want a single \"docs\" shortcut", got)
	}
}

func TestFilterNoMatchReturnsEmpty(t *testing.T) {
	got := Filter("nonexistent-shortcut-name")
	if len(got) != 0 {
		t.Fatalf("Filter(\"nonexistent-shortcut-name\") = %+v, want empty", got)
	}
}

func TestFilterPreservesShortcutsOrder(t *testing.T) {
	got := Filter("")
	for i, s := range got {
		if s != Shortcuts[i] {
			t.Fatalf("Filter(\"\")[%d] = %+v, want %+v (order must match Shortcuts)", i, s, Shortcuts[i])
		}
	}
}
