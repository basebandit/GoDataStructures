package palindrome

import "testing"

func TestIsPalindrome(t *testing.T) {
	testcases := []struct {
		input int
		want  bool
	}{
		{input: 123, want: false},
		{input: 1221, want: true},
		{input: 234, want: false},
		{input: 3456543, want: true},
	}

	for _, tt := range testcases {
		got := isPalindrome(tt.input)
		if got != tt.want {
			t.Fatalf("Expected %v got %v", tt.want, got)
		}
	}
}
