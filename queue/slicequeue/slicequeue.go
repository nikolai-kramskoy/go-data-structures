// Package slicequeue provides non-shrinking and auto-enlarging
// slice-based Queue.
package slicequeue

import (
	"math"

	"github.com/nikolai-kramskoy/go-data-structures/queue"
)

type sliceQueue[T any] struct {
	slice []T

	// We take an element from queue's head
	headIndex int

	// We add an element to queue's tail
	tailIndex int
	size      int

	initialCapacity   int
	enlargementFactor float64
}

var _ queue.Queue[struct{}] = (*sliceQueue[struct{}])(nil)

// New creates an empty Queue with default initialCapacity == 8
// and enlargementFactor == 2.0.
func New[T any]() queue.Queue[T] {
	// length == capacity == 8 for underlying slice is a reasonable
	// default enlargementFactor == 2.0 is a reasonable default
	return NewWithInitialCapacityAndEnlargementFactor[T](8, 2.0)
}

// NewWithInitialCapacity creates an empty Queue with given
// initialCapacity and default enlargementFactor == 2.0.
//
// Panics if initialCapacity < 1.
func NewWithInitialCapacity[T any](initialCapacity int) queue.Queue[T] {
	return NewWithInitialCapacityAndEnlargementFactor[T](initialCapacity, 2.0)
}

// NewWithEnlargementFactor creates an empty Queue with given
// enlargementFactor and default initialCapacity == 8.
//
// Panics if enlargementFactor <= 1.0.
func NewWithEnlargementFactor[T any](enlargementFactor float64) queue.Queue[T] {
	return NewWithInitialCapacityAndEnlargementFactor[T](8, enlargementFactor)
}

// NewWithInitialCapacityAndEnlargementFactor creates an empty
// Queue with given initialCapacity and enlargementFactor.
//
// Panics if initialCapacity < 1 or enlargementFactor <= 1.0.
func NewWithInitialCapacityAndEnlargementFactor[T any](
	initialCapacity int,
	enlargementFactor float64) queue.Queue[T] {
	if initialCapacity < 1 {
		panic("slicequeue: initialCapacity < 1")
	}
	if enlargementFactor <= 1.0 {
		panic("slicequeue: enlargementFactor <= 1.0")
	}

	return createEmptySliceQueue[T](initialCapacity, enlargementFactor)
}

func (queue *sliceQueue[T]) Push(element T) {
	// len(queue.slice) is always >= 1

	if queue.size == 0 {
		queue.slice[queue.tailIndex] = element
	} else {
		if sliceSize := len(queue.slice); queue.size < sliceSize {
			queue.tailIndex = (queue.tailIndex + 1) % sliceSize
			queue.slice[queue.tailIndex] = element
		} else {
			newSliceSize := int(math.Round(
				float64(sliceSize) * queue.enlargementFactor))

			if newSliceSize <= sliceSize {
				// maybe there is a better solution...
				panic("slicequeue: can't allocate bigger underlying " +
					"slice, maybe enlargementFactor is too big")
			}

			newSlice := make([]T, newSliceSize)

			// i for newSlice, j for slice
			for i, j := 0, queue.headIndex; i < queue.size; i, j =
				i+1, (j+1)%sliceSize {
				newSlice[i] = queue.slice[j]
			}

			newSlice[queue.size] = element

			// change current queue internals to new internals

			queue.slice = newSlice
			queue.headIndex = 0
			queue.tailIndex = queue.size
		}
	}

	queue.size++
}

func (queue *sliceQueue[T]) Pop() T {
	if queue.size == 0 {
		panic("slicequeue: pop from an empty queue")
	}

	element := queue.slice[queue.headIndex]

	if queue.size > 1 {
		queue.headIndex = (queue.headIndex + 1) % len(queue.slice)
	}

	queue.size--

	return element
}

func (queue *sliceQueue[T]) Peek() T {
	if queue.size == 0 {
		panic("slicequeue: peek on an empty queue")
	}

	return queue.slice[queue.headIndex]
}

func (queue *sliceQueue[T]) Size() int {
	return queue.size
}

func (queue *sliceQueue[T]) IsEmpty() bool {
	return queue.size == 0
}

func (queue *sliceQueue[T]) Clear() {
	*queue = *createEmptySliceQueue[T](
		queue.initialCapacity,
		queue.enlargementFactor)
}

func createEmptySliceQueue[T any](
	initialCapacity int,
	enlargementFactor float64) *sliceQueue[T] {
	return &sliceQueue[T]{
		make([]T, initialCapacity),
		0,
		0,
		0,
		initialCapacity,
		enlargementFactor,
	}
}
