package utils

import "fmt"

type vector4[T any] struct {
	X, Y, Z, W T
}

type Vector4n[T Number] vector4[T]

type Vector4i = Vector4n[int]
type Vector4f = Vector4n[float64]

func (v1 Vector4n[T]) Add(v2 Vector4n[T]) Vector4n[T] {
	return Vector4n[T]{
		X: v1.X + v2.X,
		Y: v1.Y + v2.Y,
		Z: v1.Z + v2.Z,
		W: v1.W + v2.W,
	}
}

func (v1 Vector4n[T]) String() string {
	return fmt.Sprint([]T{v1.X, v1.Y, v1.Z, v1.W})
}
