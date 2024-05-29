package slicequeue

import (
	"testing"

	"github.com/nikolai-kramskoy/go-data-structures/queue"
	"github.com/stretchr/testify/assert"
)

func TestSliceQueue_New(t *testing.T) {
	intQueue := New[int]()
	assert.NotNil(t, intQueue)
	assertEmpty(t, intQueue)
}

func TestSliceQueue_Clear(t *testing.T) {
	intQueue := New[int]()

	intQueue.Push(50)
	assertSize(t, 1, intQueue)

	intQueue.Push(100)
	assertSize(t, 2, intQueue)

	intQueue.Clear()
	assertEmpty(t, intQueue)
}

func TestSliceQueue_1(t *testing.T) {
	intQueue := New[int]()

	intQueue.Push(50)
	assertSize(t, 1, intQueue)

	intQueue.Push(100)
	assertSize(t, 2, intQueue)

	assert.Equal(t, 50, intQueue.Pop())
	assertSize(t, 1, intQueue)

	assert.Equal(t, 100, intQueue.Pop())
	assertEmpty(t, intQueue)
}

func TestSliceQueue_2(t *testing.T) {
	intQueue := New[int]()

	intQueue.Push(50)
	assertSize(t, 1, intQueue)

	assert.Equal(t, 50, intQueue.Pop())
	assertEmpty(t, intQueue)

	intQueue.Push(100)
	assertSize(t, 1, intQueue)

	assert.Equal(t, 100, intQueue.Pop())
	assertEmpty(t, intQueue)
}

func TestSliceQueue_3(t *testing.T) {
	intQueue := New[int]()

	iterations := 1000

	for i := 0; i < iterations; i++ {
		intQueue.Push(i)
		assertSize(t, i+1, intQueue)
	}

	for i := 0; i < iterations; i++ {
		assertSize(t, iterations-i, intQueue)
		assert.Equal(t, i, intQueue.Pop())
	}

	assertEmpty(t, intQueue)
}

// this test checks enlargement when headIndex > tailIndex
// with initial len of slice == 4
func TestSliceQueue_4(t *testing.T) {
	intQueue := NewWithInitialCapacity[int](4)

	// work with consecutive natural numbers
	iterations := 4

	for i := 1; i <= iterations; i++ {
		intQueue.Push(i)
		assertSize(t, i, intQueue)
	}

	assert.Equal(t, 1, intQueue.Pop())
	assertSize(t, iterations-1, intQueue)

	// trigger headIndex > tailIndex
	intQueue.Push(5)
	assertSize(t, iterations, intQueue)

	// trigger allocation of new underlying slice
	intQueue.Push(6)
	assertSize(t, iterations+1, intQueue)

	// pop all elements
	for i := 2; i <= iterations+2; i++ {
		assert.Equal(t, i, intQueue.Pop())
		assertSize(t, iterations-i+2, intQueue)
	}

	// last checks

	intQueue.Push(50)
	assertSize(t, 1, intQueue)
	intQueue.Push(100)
	assertSize(t, 2, intQueue)

	assert.Equal(t, 50, intQueue.Pop())
	assertSize(t, 1, intQueue)
	assert.Equal(t, 100, intQueue.Pop())
	assertEmpty(t, intQueue)
}

func assertSize(t *testing.T, expectedSize int, intQueue queue.Queue[int]) {
	assert.Equal(t, expectedSize, intQueue.Size())

	if expectedSize == 0 {
		assert.True(t, intQueue.IsEmpty())
	} else {
		assert.False(t, intQueue.IsEmpty())
	}
}

func assertEmpty(t *testing.T, intQueue queue.Queue[int]) {
	assertSize(t, 0, intQueue)
	assert.Panics(t, func() { intQueue.Pop() })
	assert.Panics(t, func() { intQueue.Peek() })
}
