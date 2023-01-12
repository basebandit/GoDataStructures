package selectionsort

import (
	"reflect"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	arr := []int{3, 5, 4, 1, 2}
	want := []int{1, 2, 3, 4, 5}
	selectionSort(arr)
	if !reflect.DeepEqual(arr, want) {
		t.Fatalf("Expected %v got %v", want, arr)
	}
}
