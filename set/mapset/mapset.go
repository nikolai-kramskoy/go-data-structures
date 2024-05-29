// Package mapset provides map-based Set.
package mapset

import "github.com/nikolai-kramskoy/go-data-structures/set"

type mapSet[T comparable] struct {
	tToEmptyStruct map[T]struct{}
}

var _ set.Set[struct{}] = (*mapSet[struct{}])(nil)

// New creates an empty Set.
func New[T comparable]() set.Set[T] {
	return &mapSet[T]{map[T]struct{}{}}
}

// NewFromElements creates a Set with non-duplicated entries from elements.
//
// Panics if elements == nil.
func NewFromElements[T comparable](elements ...T) set.Set[T] {
	if elements == nil {
		panic("mapset: elements == nil")
	}

	tToEmptyStruct := make(map[T]struct{}, len(elements))

	for _, elem := range elements {
		tToEmptyStruct[elem] = struct{}{}
	}

	return &mapSet[T]{tToEmptyStruct}
}

func (set *mapSet[T]) Add(element T) {
	set.tToEmptyStruct[element] = struct{}{}
}

func (set *mapSet[T]) Remove(element T) {
	delete(set.tToEmptyStruct, element)
}

func (set *mapSet[T]) Contains(element T) bool {
	_, isPresent := set.tToEmptyStruct[element]

	return isPresent
}

func (set *mapSet[T]) Elements() []T {
	elements := make([]T, len(set.tToEmptyStruct))

	i := 0
	for element := range set.tToEmptyStruct {
		elements[i] = element
		i++
	}

	return elements
}

func (set *mapSet[T]) Size() int {
	return len(set.tToEmptyStruct)
}

func (set *mapSet[T]) IsEmpty() bool {
	return len(set.tToEmptyStruct) == 0
}

func (set *mapSet[T]) Clear() {
	clear(set.tToEmptyStruct)
}
