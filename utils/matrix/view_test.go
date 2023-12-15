package matrix

import (
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
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

	str := StringFmt(view, FmtNative[int])
	expected := utils.Msg(`
1 2 3
4 5 6
7 8 9
10 11 12`)

	assert.Equal(t, expected, str)
}

func TestNewMatrixViewFlippedUpDown(t *testing.T) {
	m := NewMatrixNumberRowNotation([][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
		{10, 11, 12},
	}).Matrix

	view := NewMatrixViewFlippedUpDown[int](m)

	str := StringFmt(view, FmtNative[int])
	expected := utils.Msg(`
10 11 12
7 8 9
4 5 6
1 2 3`)

	assert.Equal(t, expected, str)
}

func TestNewMatrixViewFlippedLeftRight(t *testing.T) {
	m := NewMatrixNumberRowNotation([][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
		{10, 11, 12},
	})

	view := NewMatrixViewFlippedLeftRight[int](m)

	str := StringFmt(view, FmtNative[int])
	expected := utils.Msg(`
3 2 1
6 5 4
9 8 7
12 11 10`)

	assert.Equal(t, expected, str)
}
