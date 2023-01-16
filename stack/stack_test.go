package stack_test

import (
	"DesignPatterns/algorithms/stack"
	"fmt"
	"testing"
)

func TestStack(t *testing.T) {
	mystack := stack.New(3)
	mystack.Push(2)
	mystack.Push(4)
	mystack.Push(5)
	got, err := mystack.Top()
	checkErr(t, err)
	want := 5
	if got != want {
		t.Fatalf("Expected %d got %d", want, got)
	}

	mystack.Push(6)
	want = 6
	got, err = mystack.Pop()
	checkErr(t, err)
	if got != want {
		t.Fatalf("Expected %d got %d", want, got)
	}

	want = 4
	got = mystack.Size()
	fmt.Println(mystack.Pop())
	if got != want {
		t.Fatalf("Expected %d got %d", want, got)
	}

	//fmt.Println(, mystack.Empty())
	//fmt.Println(mystack.Pop())
	//fmt.Println(mystack.Size(), mystack.Empty())
	//fmt.Println(mystack.Pop())
	//fmt.Println(mystack.Size(), mystack.Empty())
	//fmt.Println(mystack.Top())
	//fmt.Println(mystack.Size(), mystack.Empty())
}

func checkErr(t *testing.T, err error) {
	if err != nil {
		t.Fatalf("oops, error occurred: %v", err)
	}
}
