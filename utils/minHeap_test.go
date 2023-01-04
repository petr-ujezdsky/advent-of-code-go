package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMinHeapInt_Push(t *testing.T) {
	minHeap := NewMinHeapInt[string]()

	assert.Equal(t, 0, minHeap.Len())
	assert.True(t, minHeap.Empty())

	minHeap.Push("aa", 4)
	minHeap.Push("bb", 5)
	minHeap.Push("cc", 1)
	minHeap.Push("dd", 50)

	assert.Equal(t, 4, minHeap.Len())
	assert.False(t, minHeap.Empty())

	minHeap.Fix("dd", 0)

	assert.Equal(t, "dd", minHeap.Pop())
	assert.Equal(t, "cc", minHeap.Pop())
	assert.Equal(t, "aa", minHeap.Pop())
	assert.Equal(t, "bb", minHeap.Pop())

	assert.Equal(t, 0, minHeap.Len())
	assert.True(t, minHeap.Empty())
}
