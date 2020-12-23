# The Stack Abstract Data Type

Stacks are containers, and as such they hold values of some type T, where T is the type of elements held in the stack.

## Essential operations supported by the Stack ADT
- push(t T) - Add element t at the top of the stack. This is done implicitly.
- pop()     - Remove and return the top element of the stack. The precondition is that the stack is not empty.
- empty()   - Return the Boolean value true just in case the stack is empty.
- top()     - Return the top element of the stack without removing it. The precondition here is that the stack is not empty.

## Axioms of the Stack ADT
> An axiom is a statement about the operations of an abstract data type that must always be true.      

- For any stack `s` and element `e`, `s.push(e).pop() = s.`
- For any stack `s` and element `e`, `s.push(e).top() = e.`
- For any stack `s` and element `e`, `s.push(e).empty() = false.`

## The Stack Interface

The Stack interface need only describe behaviour for pushing elements, poping elements and peeking at the top element of the stack.
```Go
type Stack interface{
  Push(e interface{})
  Pop() (interface{},error)
  Top() (interface{},error)
  Empty() bool
  Size() int
}
```

## The Stack ADT Implementation

There are two approaches to implementing the carrier set for the stack ADT
- A contiguous implementation using *arrays*.
- A linked implementation using *singly linked lists*.

### Contiguous Implementation Of The Stack ADT

Requires using an array to hold the contents of the stack and a marker to keep track of the top of the stack. The marker can record the location of the top element, or the location where the top element would go on the next Push() operation. 

> If a static (fixed-size) array is used, then the stack can become full. We wouldn't want that to happen now, would we? If a
> dynamic (resizable) array is used, then the stack is essentially unbounded. Sounds like something we would want to use, right?

We will use a dnamic array (Slice) as the underlying store to hold the stack elements. The size of the store will be the size of the stack and we will resize the dynamic array when elements are pushed of popped from the stack. We choose this design because resizing slices is inexpensive in Go (unless the underlying array needs to be expanded).





