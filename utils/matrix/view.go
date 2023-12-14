package matrix

import "github.com/petr-ujezdsky/advent-of-code-go/utils"

type coordinatesTransformer func(x, y int) (int, int)

type View[T any] struct {
	Width, Height          int
	matrix                 Matrix[T]
	coordinatesTransformer coordinatesTransformer
}

func NewMatrixViewFlippedUpDown[T any](matrix Matrix[T]) View[T] {
	transformer := func(x, y int) (int, int) {
		return x, matrix.Height - y - 1
	}

	return View[T]{
		Width:                  matrix.Width,
		Height:                 matrix.Height,
		matrix:                 matrix,
		coordinatesTransformer: transformer,
	}
}

func NewMatrixViewFlippedLeftRight[T any](matrix Matrix[T]) View[T] {
	transformer := func(x, y int) (int, int) {
		return matrix.Width - x - 1, y
	}

	return View[T]{
		Width:                  matrix.Width,
		Height:                 matrix.Height,
		matrix:                 matrix,
		coordinatesTransformer: transformer,
	}
}

func (v View[T]) Get(x, y int) T {
	x, y = v.coordinatesTransformer(x, y)
	return v.matrix.Columns[x][y]
}

func (v View[T]) GetV(pos utils.Vector2i) T {
	return v.Get(pos.X, pos.Y)
}

func (v View[T]) Set(x, y int, value T) {
	x, y = v.coordinatesTransformer(x, y)
	v.matrix.Columns[x][y] = value
}

func (v View[T]) SetV(pos utils.Vector2i, value T) {
	v.Set(pos.X, pos.Y, value)
}
