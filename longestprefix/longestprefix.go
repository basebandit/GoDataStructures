package longestprefix

import (
	"strings"
)

func longestPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	prefix := strs[0]

	for i := 1; i <= len(strs)-1; i++ {
		for !strings.HasPrefix(strs[i], prefix) {
			prefix = prefix[0 : len(prefix)-1]
		}
	}
	return prefix
}
