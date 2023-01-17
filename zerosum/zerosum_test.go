package zerosum

import "testing"

func TestZeroSumSubarray(t *testing.T) {
	testcases := []struct {
		input []int
		want  bool
	}{
		{
			input: []int{4, -6, 3, -1, 4, 2, 7},
			want:  true,
		},
		{
			input: []int{3, 4, -7, 3, 1, 3, 1, -4, -2, -2},
			want:  true,
		},
	}

	for _, tt := range testcases {
		got := zeroSumSubarray(tt.input)
		if got != tt.want {
			t.Fatalf("Expected %v got %v", tt.want, got)
		}
	}
}
