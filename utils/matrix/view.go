package matrix

import "github.com/petr-ujezdsky/advent-of-code-go/utils"

type coordinatesTransformer func(x, y int) (int, int)

type View[T any] interface {
	Get(x, y int) T
	GetV(pos utils.Vector2i) T
	Set(x, y int, value T)
	SetV(pos utils.Vector2i, value T)
	GetWidth() int
	GetHeight() int
}

type transformingView[T any] struct {
	view                   View[T]
	coordinatesTransformer coordinatesTransformer
}

func NewMatrixView[T any](view View[T]) View[T] {
	transformer := func(x, y int) (int, int) {
		return x, y
	}

	return transformingView[T]{
		view:                   view,
		coordinatesTransformer: transformer,
	}
}

func NewMatrixViewFlippedUpDown[T any](view View[T]) View[T] {
	transformer := func(x, y int) (int, int) {
		return x, view.GetHeight() - y - 1
	}

	return transformingView[T]{
		view:                   view,
		coordinatesTransformer: transformer,
	}
}

func NewMatrixViewFlippedLeftRight[T any](view View[T]) View[T] {
	transformer := func(x, y int) (int, int) {
		return view.GetWidth() - x - 1, y
	}

	return transformingView[T]{
		view:                   view,
		coordinatesTransformer: transformer,
	}
}

func (v transformingView[T]) Get(x, y int) T {
	x, y = v.coordinatesTransformer(x, y)
	return v.view.Get(x, y)
}

func (v transformingView[T]) GetV(pos utils.Vector2i) T {
	return v.Get(pos.X, pos.Y)
}

func (v transformingView[T]) Set(x, y int, value T) {
	x, y = v.coordinatesTransformer(x, y)
	v.view.Set(x, y, value)
}

func (v transformingView[T]) SetV(pos utils.Vector2i, value T) {
	v.Set(pos.X, pos.Y, value)
}

func (v transformingView[T]) GetWidth() int {
	return v.GetWidth()
}

func (v transformingView[T]) GetHeight() int {
	return v.GetHeight()
}
