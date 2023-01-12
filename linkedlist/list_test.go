package linkedlist

import (
	"testing"
)

func TestPush(t *testing.T) {
	list := LinkedList{}

	type testCases struct {
		want  any
		input any
	}

	tests := []testCases{
		{1, 1},
		{2, 2},
		{3, 3},
	}

	for _, tt := range tests {
		list.Push(tt.input)
		got, err := list.Top()
		if err != nil {
			t.Fatal(err)
		}
		if got != tt.want {
			t.Fatalf("Expected %d Got %d", got, tt.want)
		}
	}

}
