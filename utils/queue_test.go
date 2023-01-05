package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQueue_All(t *testing.T) {
	queue := NewQueue[int]()
	assert.Equal(t, 0, queue.Length())

	queue.Push(1)
	queue.Push(2)
	queue.Push(3)
	assert.Equal(t, 3, queue.Length())

	assert.Equal(t, "[front <- 1 <- 2 <- 3 <- rear]", queue.String())
	assert.Equal(t, []int{1, 2, 3}, queue.PeekAll())

	assert.Equal(t, 1, queue.Pop())
	assert.Equal(t, 2, queue.Length())
	assert.Equal(t, 2, queue.Peek())
	assert.Equal(t, "[front <- 2 <- 3 <- rear]", queue.String())
}
