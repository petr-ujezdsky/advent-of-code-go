package utils

import "fmt"

type vector3[T any] struct {
	X, Y, Z T
}

type Vector3n[T Number] vector3[T]

type Vector3i = Vector3n[int]
type Vector3f = Vector3n[float64]

func (v1 Vector3n[T]) Add(v2 Vector3n[T]) Vector3n[T] {
	return Vector3n[T]{
		X: v1.X + v2.X,
		Y: v1.Y + v2.Y,
		Z: v1.Z + v2.Z,
	}
}

func (v1 Vector3n[T]) String() string {
	return fmt.Sprint([]T{v1.X, v1.Y, v1.Z})
}
