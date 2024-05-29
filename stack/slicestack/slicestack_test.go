package slicestack

import (
	"testing"

	"github.com/nikolai-kramskoy/go-data-structures/stack"
	"github.com/stretchr/testify/assert"
)

func TestSliceStack_New(t *testing.T) {
	intStack := New[int]()
	assert.NotNil(t, intStack)
	assertEmpty(t, intStack)
}

func TestSliceStack_Clear(t *testing.T) {
	intStack := New[int]()

	intStack.Push(50)
	assertSize(t, 1, intStack)

	intStack.Push(2)
	assertSize(t, 2, intStack)

	intStack.Clear()
	assertEmpty(t, intStack)
}

func TestSliceStack_1(t *testing.T) {
	intStack := New[int]()

	intStack.Push(50)
	assertSize(t, 1, intStack)

	intStack.Push(100)
	assertSize(t, 2, intStack)

	assert.Equal(t, 100, intStack.Pop())
	assertSize(t, 1, intStack)

	assert.Equal(t, 50, intStack.Pop())
	assertEmpty(t, intStack)
}

func TestSliceStack_2(t *testing.T) {
	intStack := New[int]()

	intStack.Push(50)
	assertSize(t, 1, intStack)

	assert.Equal(t, 50, intStack.Pop())
	assertEmpty(t, intStack)

	intStack.Push(100)
	assertSize(t, 1, intStack)

	assert.Equal(t, 100, intStack.Pop())
	assertEmpty(t, intStack)
}

func TestSliceStack_3(t *testing.T) {
	intStack := New[int]()

	iterations := 1000

	for i := 0; i < iterations; i++ {
		intStack.Push(i)
		assertSize(t, i+1, intStack)
	}

	for i := iterations - 1; i >= 0; i-- {
		assertSize(t, i+1, intStack)
		assert.Equal(t, i, intStack.Pop())
	}

	assertEmpty(t, intStack)
}

func assertSize(t *testing.T, expectedSize int, intStack stack.Stack[int]) {
	assert.Equal(t, expectedSize, intStack.Size())

	if expectedSize == 0 {
		assert.True(t, intStack.IsEmpty())
	} else {
		assert.False(t, intStack.IsEmpty())
	}
}

func assertEmpty(t *testing.T, intStack stack.Stack[int]) {
	assertSize(t, 0, intStack)
	assert.Panics(t, func() { intStack.Pop() })
	assert.Panics(t, func() { intStack.Peek() })
}
