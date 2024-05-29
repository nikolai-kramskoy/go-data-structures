// Package queue provides Queue interface for first-in-first-out (FIFO)
// collection of elements and specifies what common methods fot
// this ADT must be implemented.
package queue

// Queue represents first-in-first-out (FIFO) collection of elements.
type Queue[T any] interface {
	// Push adds the element to the tail of this Queue.
	Push(element T)

	// Pop removes the element from the head of this Queue.
	//
	// Panics if this Queue is empty.
	Pop() T

	// Pop returns the element from the head of the Queue without
	// removing it.
	//
	// Panics if this Queue is empty.
	Peek() T

	// Size returns current size of this Queue.
	Size() int

	// IsEmpty returns true if Size() == 0, false otherwise.
	IsEmpty() bool

	// Clear removes all elements in this Queue.
	Clear()
}
