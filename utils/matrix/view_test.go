package matrix

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewMatrixView_directly(t *testing.T) {
	var view View[int] = NewMatrixNumberRowNotation([][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
		{10, 11, 12},
	}).Matrix

	assert.Equal(t, 1, view.Get(0, 0))
	assert.Equal(t, 4, view.Get(0, 1))
	assert.Equal(t, 7, view.Get(0, 2))
	assert.Equal(t, 10, view.Get(0, 3))
}

func TestNewMatrixViewFlippedUpDown(t *testing.T) {
	m := NewMatrixNumberRowNotation([][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
		{10, 11, 12},
	}).Matrix

	view := NewMatrixViewFlippedUpDown[int](m)

	assert.Equal(t, 10, view.Get(0, 0))
	assert.Equal(t, 7, view.Get(0, 1))
	assert.Equal(t, 4, view.Get(0, 2))
	assert.Equal(t, 1, view.Get(0, 3))
}

func TestNewMatrixViewFlippedLeftRight(t *testing.T) {
	m := NewMatrixNumberRowNotation([][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
		{10, 11, 12},
	})

	view := NewMatrixViewFlippedLeftRight[int](m)

	assert.Equal(t, 3, view.Get(0, 0))
	assert.Equal(t, 6, view.Get(0, 1))
	assert.Equal(t, 9, view.Get(0, 2))
	assert.Equal(t, 12, view.Get(0, 3))
}
