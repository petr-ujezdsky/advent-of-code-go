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

func TestNewMatrixViewRotated90CounterClockwise(t *testing.T) {
	m := NewMatrixNumberRowNotation([][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
		{10, 11, 12},
	})

	view := NewMatrixViewRotated90CounterClockwise[int](m, 0)

	str := StringFmt(view, FmtNative[int])
	expected := utils.Msg(`
1 2 3
4 5 6
7 8 9
10 11 12`)

	assert.Equal(t, expected, str)

	view = NewMatrixViewRotated90CounterClockwise[int](m, 1)

	str = StringFmt(view, FmtNative[int])
	expected = utils.Msg(`
3 6 9 12
2 5 8 11
1 4 7 10`)

	assert.Equal(t, expected, str)

	view = NewMatrixViewRotated90CounterClockwise[int](m, 2)

	str = StringFmt(view, FmtNative[int])
	expected = utils.Msg(`
12 11 10
9 8 7
6 5 4
3 2 1`)

	assert.Equal(t, expected, str)

	view = NewMatrixViewRotated90CounterClockwise[int](m, 3)

	str = StringFmt(view, FmtNative[int])
	expected = utils.Msg(`
10 7 4 1
11 8 5 2
12 9 6 3`)

	assert.Equal(t, expected, str)

	view = NewMatrixViewRotated90CounterClockwise[int](m, 4)

	str = StringFmt(view, FmtNative[int])
	expected = utils.Msg(`
1 2 3
4 5 6
7 8 9
10 11 12`)

	assert.Equal(t, expected, str)

	view = NewMatrixViewRotated90CounterClockwise[int](m, -1)

	str = StringFmt(view, FmtNative[int])
	expected = utils.Msg(`
10 7 4 1
11 8 5 2
12 9 6 3`)

	assert.Equal(t, expected, str)
}

func TestNewMatrixViewTransposed(t *testing.T) {
	m := NewMatrixNumberRowNotation([][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
		{10, 11, 12},
	})

	view := NewMatrixViewTransposed[int](m)

	str := StringFmt(view, FmtNative[int])
	expected := utils.Msg(`
1 4 7 10
2 5 8 11
3 6 9 12`)

	assert.Equal(t, expected, str)
}
