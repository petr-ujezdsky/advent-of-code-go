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

func (v1 Vector3n[T]) Multiply(k T) Vector3n[T] {
	return Vector3n[T]{
		X: v1.X * k,
		Y: v1.Y * k,
		Z: v1.Z * k,
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

func (v1 Vector3n[T]) Cross(v2 Vector3n[T]) Vector3n[T] {
	return Vector3n[T]{
		X: v1.Y*v2.Z - v1.Z*v2.Y,
		Y: v1.Z*v2.X - v1.X*v2.Z,
		Z: v1.X*v2.Y - v1.Y*v2.X,
	}
}

func (v1 Vector3n[T]) Dot(v2 Vector3n[T]) T {
	return v1.X*v2.X + v1.Y*v2.Y + v1.Z*v2.Z
}

// OrthogonalBase creates 2 more perpendicular vectors to create a base vectors in 3D space
// see https://math.stackexchange.com/a/4112622
func (v1 Vector3n[T]) OrthogonalBase() (Vector3n[T], Vector3n[T]) {
	if v1.X == 0 && v1.Y == 0 && v1.Z == 0 {
		panic("Can not build orthogonal base from zero length vector")
	}

	v2 := Vector3n[T]{
		X: Copysign(v1.Z, v1.X),
		Y: Copysign(v1.Z, v1.Y),
		Z: -Copysign(Abs(v1.X)+Abs(v1.Y), v1.Z),
		//Z: -Copysign(v1.X, v1.Z) - Copysign(v1.Y, v1.Z),
	}

	v3 := v2.Cross(v1)

	return v2, v3
}

func (v1 Vector3n[T]) ManhattanLength() T {
	return Abs(v1.X) + Abs(v1.Y) + Abs(v1.Z)
}

func (v1 Vector3n[T]) String() string {
	return fmt.Sprint([]T{v1.X, v1.Y, v1.Z})
}
