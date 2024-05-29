// Package stack provides Stack interface for last-in-first-out (LIFO)
// collection of elements and specifies what common methods fot
// this ADT must be implemented.
package stack

// Stack represents last-in-first-out (LIFO) collection of elements.
type Stack[T any] interface {
	// Push adds the element to the top of this Stack.
	Push(element T)

	// Pop removes the element from the top of this Stack.
	//
	// Panics if this Stack is empty.
	Pop() T

	// Pop returns the element from the top of this Stack without
	// removing it.
	//
	// Panics if this Stack is empty.
	Peek() T

	// Size returns current size of this Stack.
	Size() int

	// IsEmpty returns true if Size() == 0, false otherwise.
	IsEmpty() bool

	// Clear removes all elements in this Stack.
	Clear()
}
