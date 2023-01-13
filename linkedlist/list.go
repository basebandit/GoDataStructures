package linkedlist

import (
	"fmt"
)

type node struct {
	elem any
	next *node
}

type SinglyLinkedList struct {
	head *node
	tail *node
	size int
}

// Prepend inserts a new node at the beginning
func (s *SinglyLinkedList) Prepend(data any) {
	if s.head == nil {
		s.head = &node{elem: data}
		s.tail = s.head
	} else {
		n := &node{elem: data}
		n.next = s.head
		s.head = n
	}
	s.size++
}

// Append inserts a new node at the end
func (s *SinglyLinkedList) Append(data any) {
	if s.head == nil {
		s.head = &node{elem: data}
		s.tail = s.head
	} else {
		n := &node{elem: data}
		s.tail.next = n
		s.tail = n
	}
	s.size++
}

// Get retrieves the value of the index-th node in the linked list. Returns nil if node not found
// takes O(n) time.
func (s *SinglyLinkedList) Get(index int) any {
	if index < 0 || s.head == nil || index >= s.size {
		return nil
	}

	current := s.head
	for i := 0; i < index; i++ {
		current = current.next
	}

	if current == nil {
		return nil
	}
	
	return current.elem
}

// AddAt inserts a new node at the index-th of the linked list
// Takes O(n) time
func (s *SinglyLinkedList) AddAt(index int, data any) error {
	if index < 0 || index > s.size {
		return fmt.Errorf("index %d is out of bounds", index)
	}

	if index == s.size {
		s.Append(data)
	} else {
		current := s.head
		for i := 0; i < index-1; i++ {
			current = current.next
		}
		n := &node{elem: data}
		n.next = current.next
		current.next = n
	}
	s.size++
	return nil
}

// DeleteAt deletes the index-th node of the linkedlist
// takesO(n) time
func (s *SinglyLinkedList) DeleteAt(index int) error {
	if index < 0 || index >= s.size {
		return fmt.Errorf("index %d is out of bounds", index)
	}
	if index == 0 {
		s.head = s.head.next
		if s.head == nil {
			s.tail = nil
		}
	} else {
		current := s.head
		for i := 0; i < index-1; i++ {
			current = current.next
		}
		if current.next.next == nil {
			current.next = nil
			s.tail = current
		} else {
			current.next = current.next.next
		}
	}
	s.size--
	return nil
}
