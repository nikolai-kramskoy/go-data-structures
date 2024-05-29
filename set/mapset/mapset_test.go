package mapset

import (
	"testing"

	"github.com/nikolai-kramskoy/go-data-structures/set"
	"github.com/stretchr/testify/assert"
)

func TestMapSet_New(t *testing.T) {
	intSet := New[int]()
	assert.NotNil(t, intSet)
	assertSize(t, 0, intSet)
	assert.Equal(t, []int{}, intSet.Elements())
}

func TestMapSet_NewFromValues(t *testing.T) {
	intSet := NewFromElements(*newTestSlice()...)
	assert.NotNil(t, intSet)

	assertMapSetElements(t, intSet)
}

func TestMapSet_Contains(t *testing.T) {
	intSet := New[int]()

	assert.False(t, intSet.Contains(2))

	intSet.Add(2)
	assertSize(t, 1, intSet)
	assert.True(t, intSet.Contains(2))
}

func TestSliceStack_Clear(t *testing.T) {
	intSet := New[int]()

	intSet.Add(50)
	assertSize(t, 1, intSet)

	intSet.Add(2)
	assertSize(t, 2, intSet)

	intSet.Clear()
	assertSize(t, 0, intSet)
}

func TestMapSet_1(t *testing.T) {
	intSet := New[int]()

	intSet.Add(5)
	assertSize(t, 1, intSet)

	intSet.Add(5)
	assertSize(t, 1, intSet)

	intSet.Add(1)
	assertSize(t, 2, intSet)
}

func TestMapSet_2(t *testing.T) {
	intSet := New[int]()

	intSet.Add(1)
	intSet.Add(2)
	intSet.Add(2)
	intSet.Add(2)
	intSet.Add(3)
	intSet.Add(-100)
	intSet.Add(0)

	assertMapSetElements(t, intSet)
}

func assertSize(t *testing.T, expectedSize int, intSet set.Set[int]) {
	assert.Equal(t, expectedSize, intSet.Size())

	if expectedSize == 0 {
		assert.True(t, intSet.IsEmpty())
	} else {
		assert.False(t, intSet.IsEmpty())
	}
}

func newTestSlice() *[]int {
	return &[]int{1, 2, 2, 2, 3, -100, 0}
}

func assertMapSetElements(t *testing.T, intSet set.Set[int]) {
	elements := intSet.Elements()
	assert.Equal(t, 5, len(elements))

	testSlice := *newTestSlice()

	for _, element := range testSlice {
		assert.Contains(t, elements, element)
	}
}
