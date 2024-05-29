// Package slicestack provides non-shrinking and auto-enlarging
// slice-based Stack.
package slicestack

import (
	"github.com/nikolai-kramskoy/go-data-structures/stack"
)

type sliceStack[T any] struct {
	slice []T
}

var _ stack.Stack[struct{}] = (*sliceStack[struct{}])(nil)

// New creates an empty Stack.
func New[T any]() stack.Stack[T] {
	return createEmptySliceStack[T]()
}

func (stack *sliceStack[T]) Push(element T) {
	stack.slice = append(stack.slice, element)
}

func (stack *sliceStack[T]) Pop() T {
	size := len(stack.slice)

	if size == 0 {
		panic("slicestack: pop from an empty stack")
	}

	element := stack.slice[size-1]

	stack.slice = stack.slice[:size-1]

	return element
}

func (stack *sliceStack[T]) Peek() T {
	size := len(stack.slice)

	if size == 0 {
		panic("slicestack: peek on an empty stack")
	}

	return stack.slice[size-1]
}

func (stack *sliceStack[T]) Size() int {
	return len(stack.slice)
}

func (stack *sliceStack[T]) IsEmpty() bool {
	return len(stack.slice) == 0
}

func (stack *sliceStack[T]) Clear() {
	*stack = *createEmptySliceStack[T]()
}

func createEmptySliceStack[T any]() *sliceStack[T] {
	return &sliceStack[T]{
		// capacity == 8 is a reasonable default
		make([]T, 0, 8),
	}
}
