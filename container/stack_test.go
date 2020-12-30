package container

import (
	"testing"
)

func TestArrayStack(t *testing.T) {

	var (
		tPush  = "Push Test"
		tEmpty = "Empty Test"
		pTop   = "Positive Top Test"
		nTop   = "Negative Top Test"
		pPop   = "Positive Pop Test"
		nPop   = "Negative Pop Test"
	)

	in := []struct {
		name string
		want interface{}
		errw error
	}{
		{name: tPush,
			want: 1,
			errw: nil,
		},
		{
			name: pTop,
			want: 4,
			errw: nil,
		},
		{
			name: nTop,
			want: nil,
			errw: ErrTopEmpty,
		},
		{
			name: pPop,
			want: 0,
			errw: nil,
		},
		{
			name: nPop,
			want: nil,
			errw: ErrPopEmpty,
		},
		{
			name: tEmpty,
			want: true,
		},
	}

	s := ArrayStack{}

	for _, tt := range in {
		switch tt.name {
		case tPush: //Test for the Push() happy path
			t.Log(tt.name)
			s.Push(4)
			got := s.Size()
			if got != tt.want {
				t.Fatalf("want %d got %d", tt.want, got)
			}
			s.Clear()
		case tEmpty:
			t.Log(tt.name)
			s.Push(4)
			got := s.Empty()
			if got != tt.want {
				t.Fatalf("want %t got %t", tt.want, got)
			}
		case pTop: //Test for the Top() happy path
			t.Log(tt.name)
			s.Push(4)
			got, err := s.Top()
			if err != nil {
				t.Fatal(err)
			}

			if got != tt.want {
				t.Fatalf("want %d got %d", tt.want, got)
			}
			s.Clear()

		case pPop: //Test for the Pop() Happy path
			t.Log(tt.name)
			s.Push(4)
			s.Pop()

			got := s.Size()
			if got != tt.want {
				t.Fatalf("want %d got %d", tt.want, got)
			}

		case nTop: //Test for the Top() Unhappy path
			t.Log(tt.name)
			_, err := s.Top()
			if err != nil {
				if err != tt.errw {
					t.Fatalf("want %v got %v", tt.errw, err)
				}
			}

		case nPop: //Test for the Pop() Unhappy path
			t.Log(tt.name)
			_, err := s.Pop()
			if err != nil {
				if err != tt.errw {
					t.Fatalf("want %v got %v", tt.errw, err)
				}
			}

		}
	}
}

func TestLinkedStack(t *testing.T) {
	var (
		tPush  = "Push Test"
		tEmpty = "Empty Test"
		pTop   = "Positive Top Test"
		nTop   = "Negative Top Test"
		pPop   = "Positive Pop Test"
		nPop   = "Negative Pop Test"
	)

	in := []struct {
		name string
		want interface{}
		errw error
	}{
		{name: tPush,
			want: 1,
			errw: nil,
		},
		{
			name: pTop,
			want: 4,
			errw: nil,
		},
		{
			name: nTop,
			want: nil,
			errw: ErrTopEmpty,
		},
		{
			name: pPop,
			want: 0,
			errw: nil,
		},
		{
			name: nPop,
			want: nil,
			errw: ErrPopEmpty,
		},
		{
			name: tEmpty,
			want: true,
		},
	}

	s := LinkedStack{}

	for _, tt := range in {
		switch tt.name {
		case tPush: //Test for the Push() happy path
			t.Log(tt.name)
			s.Push(4)
			got := s.Size()
			if got != tt.want {
				t.Fatalf("want %d got %d", tt.want, got)
			}
			s.Clear()
		case tEmpty:
			t.Log(tt.name)
			s.Push(4)
			got := s.Empty()
			if got != tt.want {
				t.Fatalf("want %t got %t", tt.want, got)
			}
		case pTop: //Test for the Top() happy path
			t.Log(tt.name)
			s.Push(4)
			got, err := s.Top()
			if err != nil {
				t.Fatal(err)
			}

			if got != tt.want {
				t.Fatalf("want %d got %d", tt.want, got)
			}
			s.Clear()

		case pPop: //Test for the Pop() Happy path
			t.Log(tt.name)
			s.Push(4)
			s.Pop()

			got := s.Size()
			if got != tt.want {
				t.Fatalf("want %d got %d", tt.want, got)
			}

		case nTop: //Test for the Top() Unhappy path
			t.Log(tt.name)
			_, err := s.Top()
			if err != nil {
				if err != tt.errw {
					t.Fatalf("want %v got %v", tt.errw, err)
				}
			}

		case nPop: //Test for the Pop() Unhappy path
			t.Log(tt.name)
			_, err := s.Pop()
			if err != nil {
				if err != tt.errw {
					t.Fatalf("want %v got %v", tt.errw, err)
				}
			}

		}
	}
}
