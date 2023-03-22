package string

import "testing"

func TestIsIsomorphic(t *testing.T) {
	tests := []struct {
		s1   string
		s2   string
		want bool
	}{
		{s1: "egg", s2: "add", want: true},
		{s1: "foo", s2: "bar", want: false},
		{s1: "bbbaaaba", s2: "aaabbbba", want: false},
		{s1: "paper", s2: "title", want: true},
	}

	for _, tt := range tests {
		got := isIsomorphic(tt.s1, tt.s2)
		if got != tt.want {
			t.Fatalf("for %v and %v, expected %v got %v", tt.s1, tt.s2, tt.want, got)
		}
	}
}
