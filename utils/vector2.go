package utils

import "fmt"

type vector2[T any] struct {
	X, Y T
}

type Number interface {
	int | float64
}

type vector2n[T Number] vector2[T]

type Vector2i = vector2n[int]
type Vector2f = vector2n[float64]

func (v1 vector2n[T]) Add(v2 vector2n[T]) vector2n[T] {
	return vector2n[T]{
		X: v1.X + v2.X,
		Y: v1.Y + v2.Y,
	}
}

func (v1 vector2n[T]) Subtract(v2 vector2n[T]) vector2n[T] {
	return vector2n[T]{
		X: v1.X - v2.X,
		Y: v1.Y - v2.Y,
	}
}

func (v1 vector2n[T]) Abs() vector2n[T] {
	return vector2n[T]{
		X: Abs(v1.X),
		Y: Abs(v1.Y),
	}
}

func (v1 vector2n[T]) Signum() vector2n[T] {
	return vector2n[T]{
		X: Signum(v1.X),
		Y: Signum(v1.Y),
	}
}

func (v1 vector2n[T]) ArgMin() (int, T) {
	return ArgMin(v1.X, v1.Y)
}

func (v1 vector2n[T]) ArgMax() (int, T) {
	return ArgMax(v1.X, v1.Y)
}

func (v1 vector2n[T]) Min(v2 vector2n[T]) vector2n[T] {
	return vector2n[T]{
		X: Min(v1.X, v2.X),
		Y: Min(v1.Y, v2.Y),
	}
}

func (v1 vector2n[T]) Max(v2 vector2n[T]) vector2n[T] {
	return vector2n[T]{
		X: Max(v1.X, v2.X),
		Y: Max(v1.Y, v2.Y),
	}
}

func (v1 vector2n[T]) Change(i int, v T) vector2n[T] {
	if i == 0 {
		return vector2n[T]{
			X: v,
			Y: v1.Y,
		}
	}
	return vector2n[T]{
		X: v1.X,
		Y: v,
	}
}

func (v1 vector2n[T]) Reverse() vector2n[T] {
	return vector2n[T]{v1.Y, v1.X}
}

func (v1 vector2n[T]) InvY() vector2n[T] {
	return vector2n[T]{v1.X, -v1.Y}
}

func (v1 vector2n[T]) LengthSq() T {
	return v1.X*v1.X + v1.Y*v1.Y
}

func (v1 vector2n[T]) String() string {
	return fmt.Sprint([]T{v1.X, v1.Y})
}
