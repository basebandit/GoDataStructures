package binarysearch

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	arr := []int{6, 7, 8, 9, 10, 1, 2, 3, 4, 5}
	want := 0
	got := binarySearch(arr, 6)
	if got != want {
		t.Fatalf("Expected %d got %d", want, got)
	}
}
