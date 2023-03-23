package string

import "testing"

func TestSubsequence(t *testing.T) {
	tests := []struct {
		s1   string
		s2   string
		want bool
	}{
		{s1: "abc", s2: "ahbgdc", want: true},
		{s1: "axc", s2: "ahbgdc", want: false},
	}

	for _, tt := range tests {
		got := isSubsequence(tt.s1, tt.s2)
		if got != tt.want {
			t.Fatalf("expected %v got %v", tt.want, got)
		}
	}
}
