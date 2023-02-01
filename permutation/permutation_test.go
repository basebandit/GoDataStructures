package permutation

import "testing"

func TestIsPermutation(t *testing.T) {
	testcases := []struct {
		input []string
		want  bool
	}{
		{input: []string{"abcd", "dacb"}, want: true},
		{input: []string{"decds", "dcds"}, want: false},
	}

	for _, tt := range testcases {
		got := isPermutation(tt.input[0], tt.input[1])
		if tt.want != got {
			t.Fatalf("expected %v got %v", tt.want, got)
		}

		got = isPermutation2(tt.input[0], tt.input[1])
		if tt.want != got {
			t.Fatalf("expected %v got %v", tt.want, got)
		}
	}
}
