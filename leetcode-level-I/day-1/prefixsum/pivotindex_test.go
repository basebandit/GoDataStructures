package main

import "testing"

func TestPivotIndex(t *testing.T) {
	tests := []struct {
		input []int
		want  int
	}{
		{input: []int{2, 1, -1}, want: 0},
		{input: []int{1, 2, 3}, want: -1},
		{input: []int{1, 7, 3, 6, 5, 6}, want: 3},
		{input: []int{0, 0, 0, 0}, want: 0},
	}
	for _, tt := range tests {
		got := pivotIndex(tt.input)
		if got != tt.want {
			t.Fatalf("expected %v got %v", tt.want, got)
		}
	}
}
