package uniquestring

import "testing"

func TestIsUnique(t *testing.T) {
	testcases := []struct {
		input string
		want  bool
	}{
		{input: "scskl", want: false},
		{input: "abcde", want: true},
	}

	for _, tt := range testcases {
		got := isUnique(tt.input)
		if tt.want != got {
			t.Fatalf("Expected %v got %v", tt.want, got)
		}
	}
}