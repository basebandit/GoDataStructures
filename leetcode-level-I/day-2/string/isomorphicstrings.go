package string

// Given two strings s and t, determine if they are isomorphic.
// Two strings s and t are isomorphic if the characters in s can be replaced to get t.
// All occurrences of a character must be replaced with another character while preserving the order of characters. No two characters may map to
// the same character, but a character may map to itself.

// Example 1:
//
// Input: s = "egg", t = "add"
// Output: true
// Example 2:
//
// Input: s = "foo", t = "bar"
// Output: false
// Example 3:
//
// Input: s = "paper", t = "title"
// Output: true
func isIsomorphic(s string, t string) bool {
	sT := make(map[uint8]int, len(s))
	tS := make(map[uint8]int, len(t))

	for i := 0; i < len(s); i++ {
		if sT[s[i]] != tS[t[i]] {
			return false
		}
		// example "egg" and "add", we map e -> a, g -> d. That means e and g cannot map to any other character. Therefore they should have equal values
		// i.e sT['e'] = 1 && tS['a'] = 1
		sT[s[i]] = i + 1
		tS[t[i]] = i + 1
	}
	return true
}
