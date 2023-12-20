package slices

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSliceIterator(t *testing.T) {
	s := []int{1, 2, 3, 4}
	iterator := NewSliceIterator(s)

	assert.True(t, iterator.HasNext())
	assert.Equal(t, 1, iterator.Next())

	assert.True(t, iterator.HasNext())
	assert.Equal(t, 2, iterator.Next())

	assert.True(t, iterator.HasNext())
	assert.Equal(t, 3, iterator.Next())

	assert.True(t, iterator.HasNext())
	assert.Equal(t, 4, iterator.Next())

	assert.False(t, iterator.HasNext())
}

func TestReversedSliceIterator(t *testing.T) {
	s := []int{1, 2, 3, 4}
	iterator := NewReversedSliceIterator(s)

	assert.True(t, iterator.HasNext())
	assert.Equal(t, 4, iterator.Next())

	assert.True(t, iterator.HasNext())
	assert.Equal(t, 3, iterator.Next())

	assert.True(t, iterator.HasNext())
	assert.Equal(t, 2, iterator.Next())

	assert.True(t, iterator.HasNext())
	assert.Equal(t, 1, iterator.Next())

	assert.False(t, iterator.HasNext())
}

func TestSliceIterator_Empty(t *testing.T) {
	var s []int
	iterator := NewSliceIterator(s)

	assert.False(t, iterator.HasNext())
}
