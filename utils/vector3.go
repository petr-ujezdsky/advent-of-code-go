package utils

import (
	"fmt"
)

type vector3[T any] struct {
	X, Y, Z T
}

type Vector3n[T Number] vector3[T]

type Vector3i = Vector3n[int]
type Vector3f = Vector3n[float64]

func NewVector3nRepeated[T Number](v T) Vector3n[T] {
	return Vector3n[T]{v, v, v}
}

func (v1 Vector3n[T]) Add(v2 Vector3n[T]) Vector3n[T] {
	return Vector3n[T]{
		X: v1.X + v2.X,
		Y: v1.Y + v2.Y,
		Z: v1.Z + v2.Z,
	}
}

func (v1 Vector3n[T]) Subtract(v2 Vector3n[T]) Vector3n[T] {
	return Vector3n[T]{
		X: v1.X - v2.X,
		Y: v1.Y - v2.Y,
		Z: v1.Z - v2.Z,
	}
}

func (v1 Vector3n[T]) MultiplyParts(v2 Vector3n[T]) Vector3n[T] {
	return Vector3n[T]{
		X: v1.X * v2.X,
		Y: v1.Y * v2.Y,
		Z: v1.Z * v2.Z,
	}
}

func (v1 Vector3n[T]) Divide(divisor T) Vector3n[T] {
	return Vector3n[T]{
		X: v1.X / divisor,
		Y: v1.Y / divisor,
		Z: v1.Z / divisor,
	}
}

func (v1 Vector3n[T]) Min(v2 Vector3n[T]) Vector3n[T] {
	return Vector3n[T]{
		X: Min(v1.X, v2.X),
		Y: Min(v1.Y, v2.Y),
		Z: Min(v1.Z, v2.Z),
	}
}

func (v1 Vector3n[T]) Max(v2 Vector3n[T]) Vector3n[T] {
	return Vector3n[T]{
		X: Max(v1.X, v2.X),
		Y: Max(v1.Y, v2.Y),
		Z: Max(v1.Z, v2.Z),
	}
}

func (v1 Vector3n[T]) ManhattanLength() T {
	return Abs(v1.X) + Abs(v1.Y) + Abs(v1.Z)
}

func (v1 Vector3n[T]) String() string {
	return fmt.Sprint([]T{v1.X, v1.Y, v1.Z})
}
