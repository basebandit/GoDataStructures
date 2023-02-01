package permutation

// Given two strings, write a method to decide if one is a permutation of the
// other.
// time complexity O(n) space complexity O(1)
func isPermutation(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	
	var totalAsciiVal1 int

	for i := 0; i < len(s1); i++ {
		totalAsciiVal1 += int(s1[i])
	}
	var totalAsciiVal2 int
	for i := 0; i < len(s2); i++ {
		totalAsciiVal2 += int(s2[i])
	}

	if totalAsciiVal1 != totalAsciiVal2 {
		return false
	}

	return true
}

// time complexity O(n) space complexity O(n)
func isPermutation2(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	count := map[rune]int{}

	for _, runeValue := range s1 {
		count[runeValue]++
	}

	for _, runeValue := range s2 {
		count[runeValue]--
	}

	for _, val := range count {
		if val != 0 {
			return false
		}
	}

	return true
}
