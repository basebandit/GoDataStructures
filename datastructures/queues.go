package container

import (
	"errors"
	"fmt"
)

//Queue interface defines the method set (behaviour) for the Queue ADT
type Queue interface {
	Enter(e interface{})
	Leave(e interface{}) (interface{}, error)
	Front() (interface{}, error)
	Empty() bool
}

//ArrayQueue is a contiguous implementation of the Queue interface.
//A slice is used to store the data, and it expands as necessary if the
//queue becomes full. The front location is recorded, and the next open slot
//at the rear is at store[(frontIndex+count)%len(store)]
//count is the number of elements currently present in the queue and length is the
//length of the store.
type ArrayQueue struct {
	store      []interface{} //the queue's storage, holds all the elemens of the queue. circular buffer/ ring buffer for queue elements
	frontIndex int           //holds the index of the element at the front of the queue
	count      int           //holds the number of elements in the queue
}

//Leave removes and returns the front element on the queue.
//Precondition: the queue is not empty.
//Precondition violation: return nil and an error indication.
//Normal return: the front element and nil.
func (aq *ArrayQueue) Leave(e interface{}) (interface{}, error) {
	if aq.count == 0 {
		return nil, errors.New("Leave: the  queue cannot be empty")
	}
	result := aq.store[aq.frontIndex]
	aq.frontIndex++
	aq.store = aq.store[aq.frontIndex:]
	aq.count--
	return result, nil
}

//Enter adds a new element to the rear of the queue.
func (aq *ArrayQueue) Enter(e interface{}) {
	//if the queue cannot accomodate we reallocate a new store/slice(underlying array)
	if len(aq.store) == 0 {
		aq.store = make([]interface{}, 4)
	}

	//if the queue is full i.e the elements in the queue equal the length of the queue.
	if aq.count == len(aq.store) {
		aq.store = append(aq.store, aq.store...)
	}

	aq.store[(aq.frontIndex+aq.count)%len(aq.store)] = e
	aq.count++
}

//Empty returns true if the queue is empty.
func (aq *ArrayQueue) Empty() bool {
	return aq.count == 0
}

//Front returns the front element in the queue without removing it.
//Precondition: the queue is not empty.
//Precondition violation: return nil and an error indication.
//Normal return: the front element and nil.
func (aq *ArrayQueue) Front() (interface{}, error) {
	if aq.count == 0 {
		return nil, errors.New("Front: the queue cannot be empty")
	}
	return aq.store[aq.frontIndex], nil
}

//String makes a report on the queue container.
func (aq *ArrayQueue) String() string {
	return fmt.Sprintf("ArrayQueue instance:\nsize: %d\nfrontIndex: %d\nstore len: %d\nstore cap: %d\n"+"store: %v\n", aq.count, aq.frontIndex, len(aq.store), cap(aq.store), aq.store)
}
