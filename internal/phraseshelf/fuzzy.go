package phraseshelf

// FuzzyMatch reports whether every rune of query appears in target, in
// order but not necessarily contiguously (a subsequence match), case-folded
// for ASCII letters only. Non-ASCII runes (e.g. Japanese) are compared as-is,
// since they have no case. An empty query matches everything.
func FuzzyMatch(query, target string) bool {
	q := []rune(query)
	if len(q) == 0 {
		return true
	}

	qi := 0
	for _, r := range target {
		if foldASCII(r) == foldASCII(q[qi]) {
			qi++
			if qi == len(q) {
				return true
			}
		}
	}
	return false
}

func foldASCII(r rune) rune {
	if r >= 'A' && r <= 'Z' {
		return r + ('a' - 'A')
	}
	return r
}
