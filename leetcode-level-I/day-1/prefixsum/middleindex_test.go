package main

import "testing"

func TestMiddleIndex(t *testing.T) {
	tests := []struct {
		input []int
		want  int
	}{
		{input: []int{2, 3, -1, 8, 4}, want: 3},
		{input: []int{1, -1, 4}, want: 2},
		{input: []int{2, 5}, want: -1},
		{input: []int{0, 0, 0, 0}, want: 0},
	}
	for _, tt := range tests {
		got := findMiddleIndex(tt.input)
		if got != tt.want {
			t.Fatalf("expected %v got %v", tt.want, got)
		}
	}
}
