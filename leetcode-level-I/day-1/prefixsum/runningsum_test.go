package main

import (
	"reflect"
	"testing"
)

func TestRunningSum(t *testing.T) {
	tests := []struct {
		nums []int
		want []int
	}{
		{
			nums: []int{1, 2, 3, 4},
			want: []int{1, 3, 6, 10},
		},
		{
			nums: []int{1, 1, 1, 1, 1},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			nums: []int{3, 1, 2, 10, 1},
			want: []int{3, 4, 6, 16, 17},
		},
	}

	for _, tt := range tests {
		got := runningSum(tt.nums)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("expected %v got %v", tt.want, got)
		}
	}

	//runningSum2
	for _, tt := range tests {
		got := runningSum2(tt.nums)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("expected %v got %v", tt.want, got)
		}
	}
}
