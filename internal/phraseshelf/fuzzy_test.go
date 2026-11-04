package phraseshelf

import "testing"

func TestFuzzyMatchEmptyQueryMatchesEverything(t *testing.T) {
	if !FuzzyMatch("", "anything") {
		t.Fatal("FuzzyMatch(\"\", ...) = false, want true")
	}
}

func TestFuzzyMatchSubsequenceNotNecessarilyContiguous(t *testing.T) {
	if !FuzzyMatch("gtg", "greeting") {
		t.Fatal(`FuzzyMatch("gtg", "greeting") = false, want true`)
	}
}

func TestFuzzyMatchIsCaseInsensitiveForASCII(t *testing.T) {
	if !FuzzyMatch("GRE", "greeting") {
		t.Fatal(`FuzzyMatch("GRE", "greeting") = false, want true`)
	}
	if !FuzzyMatch("gre", "GREETING") {
		t.Fatal(`FuzzyMatch("gre", "GREETING") = false, want true`)
	}
}

func TestFuzzyMatchRejectsOutOfOrderCharacters(t *testing.T) {
	if FuzzyMatch("fa", "abcdef") {
		t.Fatal(`FuzzyMatch("fa", "abcdef") = true, want false ('f' comes after 'a' in "abcdef", not before)`)
	}
}

func TestFuzzyMatchRejectsMissingCharacters(t *testing.T) {
	if FuzzyMatch("xyz", "greeting") {
		t.Fatal(`FuzzyMatch("xyz", "greeting") = true, want false`)
	}
}

func TestFuzzyMatchIsNotPrefixOnly(t *testing.T) {
	if !FuzzyMatch("eeting", "greeting") {
		t.Fatal(`FuzzyMatch("eeting", "greeting") = false, want true (not anchored to the start)`)
	}
}

func TestFuzzyMatchJapaneseSubsequence(t *testing.T) {
	// "あいさつ" (greeting) matched by the non-contiguous subsequence "あつ".
	if !FuzzyMatch("あつ", "あいさつ") {
		t.Fatal(`FuzzyMatch("あつ", "あいさつ") = false, want true`)
	}
}

func TestFuzzyMatchJapaneseNoFalsePositive(t *testing.T) {
	if FuzzyMatch("つあ", "あいさつ") {
		t.Fatal(`FuzzyMatch("つあ", "あいさつ") = true, want false (wrong order)`)
	}
}
