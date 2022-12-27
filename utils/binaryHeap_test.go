package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBinaryHeap_Push(t *testing.T) {
	minHeap := NewBinaryHeap(func(i, j int) bool { return i < j })

	assert.True(t, minHeap.Empty())

	minHeap.Push(4)
	minHeap.Push(5)
	minHeap.Push(1)
	minHeap.Push(50)

	assert.False(t, minHeap.Empty())

	assert.Equal(t, 1, minHeap.Pop())
	assert.Equal(t, 4, minHeap.Pop())
	assert.Equal(t, 5, minHeap.Pop())
	assert.Equal(t, 50, minHeap.Pop())

	assert.True(t, minHeap.Empty())
}

func TestBinaryHeapInt_Push(t *testing.T) {
	minHeap := NewBinaryHeapInt[string]()

	assert.True(t, minHeap.Empty())

	minHeap.Push("aa", 4)
	minHeap.Push("bb", 5)
	minHeap.Push("cc", 1)
	minHeap.Push("dd", 50)

	assert.False(t, minHeap.Empty())

	assert.Equal(t, "cc", minHeap.Pop())
	assert.Equal(t, "aa", minHeap.Pop())
	assert.Equal(t, "bb", minHeap.Pop())
	assert.Equal(t, "dd", minHeap.Pop())

	assert.True(t, minHeap.Empty())
}
