// Package set provides Set interface for collection of unique elements
// and specifies what common methods for this ADT must be implemented.
package set

// Set represents collection of unique elements.
type Set[T comparable] interface {
	// Add ads ths element if he's not already present in this Set.
	Add(element T)

	// Remove removes the element if he's present in this Set.
	Remove(element T)

	// Contains returns true if this Set contains element, false otherwise.
	Contains(element T) bool

	// Elements return a slice of all elements in this Set in
	// unspecified order.
	//
	// No operation on the returned slice may affect the state of this Set.
	Elements() []T

	// Size returns current size of this Set.
	Size() int

	// IsEmpty returns true if Size() == 0, false otherwise.
	IsEmpty() bool

	// Clear removes all elements in this Set.
	Clear()
}
