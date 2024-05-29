package sets

import (
	"github.com/nikolai-kramskoy/go-data-structures/set"
	"github.com/nikolai-kramskoy/go-data-structures/set/mapset"
)

// Intersection returns a Set that contains all elements that are
// present in both a and b.
//
// Panics if a == nil or b == nil.
func Intersection[T comparable](a set.Set[T], b set.Set[T]) set.Set[T] {
	if a == nil {
		panic("sets: a == nil")
	}
	if b == nil {
		panic("sets: b == nil")
	}

	intersection := mapset.New[T]()

	// It is sufficient to only iterate over one set, not both
	for _, aElement := range a.Elements() {
		if b.Contains(aElement) {
			intersection.Add(aElement)
		}
	}

	return intersection
}
