package linkedlist

import (
	"errors"
)

type node struct {
	elem any
	next *node
}

type LinkedList struct {
	header *node
	count  int
}

func (l *LinkedList) Size() int {
	return l.count
}
func (l *LinkedList) Empty() bool {
	return l.count == 0
}
func (l *LinkedList) Clear() {
	l.count = 0
	l.header = nil
}
func (l *LinkedList) Push(e any) {
	l.header = &node{e, l.header}
	l.count++
}
func (l *LinkedList) Pop() (any, error) {
	if l.count == 0 {
		return nil, errors.New("list is empty")
	}
	result := l.header.elem
	l.header = l.header.next
	l.count--
	return result, nil
}

func (l *LinkedList) Top() (any, error) {
	if l.count == 0 {
		return nil, errors.New("list is empty")
	}
	return l.header.elem, nil
}
