package string

// Given two strings s and t, return true if s is a subsequence of t, or false otherwise.
//
// A subsequence of a string is a new string that is formed from the original string by deleting some (can be none) of the characters without disturbing the relative positions of the remaining characters. (i.e., "ace" is a subsequence of "abcde" while "aec" is not).
//
// Example 1:
//
// Input: s = "abc", t = "ahbgdc"
// Output: true
// Example 2:
//
// Input: s = "axc", t = "ahbgdc"
// Output: false
// Solution: We use the 2 pointers approach, the first pointer say i points to the start of the first string
// and the second pointer points to the start of the second string. We increment the first pointer (i) only if the
// characters at the current index in both first and second string match. We increment the second pointer as we move along.
// If i reaches end of s1,that mean we found all characters of s1 in s2, so s1 is subsequence of s2
func isSubsequence(s1 string, s2 string) bool {
	i, j := 0, 0
	for i < len(s1) && j < len(s2) {
		if s1[i] == s2[j] {
			i++
		}
		j++
	}

	//If i reaches end of s1,that mean we found all characters of s1 in s2,
	return i == len(s1)
}
