package twosum

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	testcases := []struct {
		input []int
		sum   int
		want  []int
	}{
		{
			input: []int{3, 1, 2, 5},
			sum:   8,
			want:  []int{3, 5},
		},
	}

	for _, tt := range testcases {
		got := twoSum(tt.input, tt.sum)
		if !reflect.DeepEqual(tt.want, got) {
			t.Fatalf("Expected %v got %v", tt.want, got)
		}
	}
}

func TestTwoSumOptimized(t *testing.T) {
	testcases := []struct {
		input []int
		sum   int
		want  []int
	}{
		{
			input: []int{3, 1, 2, 5},
			sum:   8,
			want:  []int{3, 5},
		},
		{
			input: []int{3, 1, 2, 4, 6, 10},
			sum:   10,
			want:  []int{4, 6},
		},
	}

	for _, tt := range testcases {
		got := twoSumOptimized(tt.input, tt.sum)
		if !reflect.DeepEqual(tt.want, got) {
			t.Fatalf("Expected %v got %v", tt.want, got)
		}
	}
}
