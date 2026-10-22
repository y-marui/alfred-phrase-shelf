package phraseshelf

import "strings"

// SplitQuery splits the text after the "snip" keyword on the first "#".
//
// With no "#", the whole query is a category-filter text (hasCategory is
// false, filterText is query unchanged). With a "#", everything before the
// first one is the exact category (taken verbatim, never fuzzy-matched) and
// everything after — including any further literal "#" characters — is the
// snippet-filter text.
func SplitQuery(query string) (category string, filterText string, hasCategory bool) {
	idx := strings.IndexByte(query, '#')
	if idx < 0 {
		return "", query, false
	}
	return query[:idx], query[idx+1:], true
}
