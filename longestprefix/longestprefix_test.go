package longestprefix

import "testing"

func TestLongestPrefix(t *testing.T) {
	testcases := []struct {
		input []string
		want  string
	}{
		{
			input: []string{"flower", "flow", "flight"},
			want:  "fl",
		},
		{
			input: []string{"dog", "racecar", "flower"},
			want:  "",
		},
	}
	for _, tt := range testcases {
		got := longestPrefix(tt.input)
		if got != tt.want {
			t.Fatalf("expected %v got %v", tt.want, got)
		}
	}
}
