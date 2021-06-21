package container

import "errors"

var (
	//ErrTopEmpty reports an empty stack while attempting a Top() operation
	ErrTopEmpty = errors.New("Top: stack cannot be empty")
	//ErrPopEmpty reports an empty stack while attempting a Pop() operation
	ErrPopEmpty = errors.New("Pop: stack cannot be empty")
)

//ArrayStack is an implementation of a stack using a dynamic array(slice).
type ArrayStack struct {
	store []interface{}
}

//A Stack supports:
// pop() returns the element removed from the stack
// push(element interface{}) adds a new element to the top of the stack
// top() returns the top of the stack len(stack)-1
//empty() convenience method to check if stack is empty. returns true if empty false otherwise.

//Pop removes the top most element from the stack and returns it.
func (s *ArrayStack) Pop() (interface{}, error) {
	if len(s.store) == 0 {
		return nil, ErrPopEmpty
	}

	top := s.store[len(s.store)-1]
	s.store = s.store[:len(s.store)-1]
	return top, nil
}

//Push adds a new element to the top of the stack.
func (s *ArrayStack) Push(e interface{}) {
	s.store = append(s.store, e)
}

//Top returns the top of the stack len(stack)-1
func (s *ArrayStack) Top() (interface{}, error) {
	if len(s.store) == 0 {
		return nil, ErrTopEmpty
	}

	return s.store[len(s.store)-1], nil
}

//Empty returns state of the stack. returns true if stack is empty otherwise false
func (s *ArrayStack) Empty() bool {
	return len(s.store) > 0
}

//Size returns the length of the stack.
func (s *ArrayStack) Size() int {
	return len(s.store)
}

//Clear empties the stack
func (s *ArrayStack) Clear() {
	s.store = make([]interface{}, 0, 10)
}

//node an aggregate variable with data and link(reference) fields. This forms our
//backing linked data structure for the stack ADT linked implementation.
type node struct {
	item interface{}
	next *node
}

//LinkedStack is an implementation of the stack ADT using a linked data structure (a node).
type LinkedStack struct {
	topPtr *node //singly-linked list of values
	count  int   //how many elements are present
}

//Pop removes the top most element from the stack and returns it.
func (s *LinkedStack) Pop() (interface{}, error) {
	if s.count == 0 {
		return nil, ErrPopEmpty
	}

	result := s.topPtr.item
	s.topPtr = s.topPtr.next
	s.count--

	return result, nil
}

//Push adds a new element on to the top of the stack.
func (s *LinkedStack) Push(e interface{}) {
	s.topPtr = &node{item: e, next: s.topPtr}
	s.count++
}

//Top returns the top of the stack
func (s *LinkedStack) Top() (interface{}, error) {
	if s.count == 0 {
		return nil, ErrTopEmpty
	}

	return s.topPtr.item, nil
}

//Empty returns the state of the stack, true if is empty false otherwise.
func (s *LinkedStack) Empty() bool {
	return s.count > 0
}

//Size returns the number of elements held by the stack.
func (s *LinkedStack) Size() int {
	return s.count
}

//Clear empties the stack
func (s *LinkedStack) Clear() {
	s.count = 0
	s.topPtr = nil
}
