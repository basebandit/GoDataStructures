package uniquestring

// Implement an algorithm to determine if a string has all unique characters. What if you
// cannot use additional data structures?
func isUnique(s string) bool {
	seen := map[rune]struct{}{}
	for _, runeValue := range s {
		if _, ok := seen[runeValue]; ok {
			return false
		}
		seen[runeValue] = struct{}{}
	}
	return true
}
