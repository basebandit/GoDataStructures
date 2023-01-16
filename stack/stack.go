package stack

import "errors"

type Stack struct {
	store []any
	top   int
}

// New creates a new stack with size as its initial capacity and length.
func New(size int) *Stack {
	return &Stack{
		store: make([]any, size),
	}
}
func (s *Stack) Push(e any) {
	s.store = append(s.store, e)
	s.top = len(s.store) - 1
}

func (s *Stack) Pop() (any, error) {
	if len(s.store) == 0 {
		return nil, errors.New("stack is empty")
	}
	elem := s.store[s.top]
	s.store = s.store[:s.top]
	s.top--
	return elem, nil
}

func (s *Stack) Size() int {
	return len(s.store)
}

func (s *Stack) Empty() bool {
	return len(s.store) == 0
}
func (s *Stack) Top() (any, error) {
	if len(s.store) == 0 {
		return nil, errors.New("stack is empty")
	}
	return s.store[s.top], nil
}
