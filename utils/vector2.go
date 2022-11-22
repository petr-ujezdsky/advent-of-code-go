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

func (v1 vector2n[T]) String() string {
	return fmt.Sprint([]T{v1.X, v1.Y})
}
