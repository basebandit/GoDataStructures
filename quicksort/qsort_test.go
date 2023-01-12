package quicksort

import (
	"reflect"
	"testing"
)

func TestQsortWithRandomPivot(t *testing.T) {
	a := []int{9, 12, 9, 2, 17, 1, 6}
	want := []int{1, 2, 6, 9, 9, 12, 17}
	got := qsortWithRandomPivot(a)
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Expected %v got %v", want, got)
	}
}

func TestQsort(t *testing.T) {
	a := []int{9, 12, 9, 2, 17, 1, 6}
	want := []int{1, 2, 6, 9, 9, 12, 17}
	got := qsort(a)
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Expected %v got %v", want, got)
	}
}

func BenchmarkQsortWithRandomPivot(b *testing.B) {
	a := []int{9, 12, 9, 2, 17, 1, 6}
	want := []int{1, 2, 6, 9, 9, 12, 17}
	got := qsortWithRandomPivot(a)
	if !reflect.DeepEqual(got, want) {
		b.Fatalf("Expected %v got %v", want, got)
	}
}

func BenchmarkQsort(b *testing.B) {
	a := []int{9, 12, 9, 2, 17, 1, 6}
	want := []int{1, 2, 6, 9, 9, 12, 17}
	got := qsort(a)
	if !reflect.DeepEqual(got, want) {
		b.Fatalf("Expected %v got %v", want, got)
	}
}
