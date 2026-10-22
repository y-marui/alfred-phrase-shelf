package phraseshelf

// Categories returns the distinct Category values of snippets, in order of
// first appearance. Phrase Shelf's "list snippets" already returns results
// in its own category order (named categories first, "" uncategorized last
// if present), so first-appearance order preserves that.
func Categories(snippets []Snippet) []string {
	seen := make(map[string]bool)
	categories := make([]string, 0, len(snippets))
	for _, s := range snippets {
		if seen[s.Category] {
			continue
		}
		seen[s.Category] = true
		categories = append(categories, s.Category)
	}
	return categories
}

// CountByCategory returns how many snippets fall in each category.
func CountByCategory(snippets []Snippet) map[string]int {
	counts := make(map[string]int)
	for _, s := range snippets {
		counts[s.Category]++
	}
	return counts
}
