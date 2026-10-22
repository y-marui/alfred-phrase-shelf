package phraseshelf

import "testing"

func TestSplitQueryNoHashIsCategoryFilterText(t *testing.T) {
	category, filterText, hasCategory := SplitQuery("wor")
	if hasCategory {
		t.Fatalf("hasCategory = true, want false")
	}
	if category != "" {
		t.Errorf("category = %q, want empty", category)
	}
	if filterText != "wor" {
		t.Errorf("filterText = %q, want %q", filterText, "wor")
	}
}

func TestSplitQueryEmptyIsCategoryFilterText(t *testing.T) {
	category, filterText, hasCategory := SplitQuery("")
	if hasCategory {
		t.Fatalf("hasCategory = true, want false")
	}
	if category != "" || filterText != "" {
		t.Errorf("got (%q, %q), want (\"\", \"\")", category, filterText)
	}
}

func TestSplitQuerySplitsOnFirstHashOnly(t *testing.T) {
	category, filterText, hasCategory := SplitQuery("work#greeting#formal")
	if !hasCategory {
		t.Fatalf("hasCategory = false, want true")
	}
	if category != "work" {
		t.Errorf("category = %q, want %q", category, "work")
	}
	if filterText != "greeting#formal" {
		t.Errorf("filterText = %q, want %q", filterText, "greeting#formal")
	}
}

func TestSplitQueryHashWithEmptyCategoryIsUncategorized(t *testing.T) {
	category, filterText, hasCategory := SplitQuery("#hello")
	if !hasCategory {
		t.Fatalf("hasCategory = false, want true")
	}
	if category != "" {
		t.Errorf("category = %q, want empty (uncategorized)", category)
	}
	if filterText != "hello" {
		t.Errorf("filterText = %q, want %q", filterText, "hello")
	}
}

func TestSplitQueryTrailingHashIsEmptyFilterText(t *testing.T) {
	category, filterText, hasCategory := SplitQuery("work#")
	if !hasCategory {
		t.Fatalf("hasCategory = false, want true")
	}
	if category != "work" {
		t.Errorf("category = %q, want %q", category, "work")
	}
	if filterText != "" {
		t.Errorf("filterText = %q, want empty", filterText)
	}
}
