package linkedlist

import (
	"testing"
)

func TestSinglyLinkedList(t *testing.T) {
	list := SinglyLinkedList{}
	list.Prepend(1)
	want := 1
	got := list.Get(0)
	if got != want {
		t.Fatalf("Expected %d got %d", want, got)
	}
	list.Prepend(2)
	want = 2
	got = list.Get(0)
	if got != want {
		t.Fatalf("Expected %d got %d", want, got)
	}
	want = 1
	got = list.Get(1)
	if got != want {
		t.Fatalf("Expected %d got %d", want, got)
	}

	err := list.AddAt(2, 3)
	if err != nil {
		t.Fatalf("error occured: %v", err)
	}
	want = 3
	got = list.Get(2)
	if got != want {
		t.Fatalf("Expected %d got %d", want, got)
	}
	err = list.AddAt(3, 4)
	if err != nil {
		t.Fatalf("error occured: %v", err)
	}
	want = 4
	got = list.Get(3)
	if got != want {
		t.Fatalf("Expected %d got %d", want, got)
	}
	err = list.DeleteAt(3)
	if err != nil {
		t.Fatalf("error occured: %v", err)
	}

	got = list.Get(3)
	if got != nil {
		t.Fatalf("Expected nil got %v", got)
	}

	want = 3
	got = list.Get(2)
	if got != want {
		t.Fatalf("Expected %v got %v", want, got)
	}
}
