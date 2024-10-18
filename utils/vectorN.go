package utils

import "fmt"

type vectorN[T any] struct {
	Items []T
}

type VectorNn[T Number] vectorN[T]

type VectorNi = VectorNn[int]
type VectorNf = VectorNn[float64]

func NewVectorNn[T Number](length int) VectorNn[T] {
	return VectorNn[T]{Items: make([]T, length)}
}

func (v1 VectorNn[T]) Multiply(k T) VectorNn[T] {
	v2 := NewVectorNn[T](len(v1.Items))
	for i, v := range v1.Items {
		v2.Items[i] = v * k
	}

	return v2
}

func (v1 VectorNn[T]) Add(v2 VectorNn[T]) VectorNn[T] {
	v3 := NewVectorNn[T](len(v1.Items))

	for i, v := range v1.Items {
		v3.Items[i] = v + v2.Items[i]
	}

	return v3
}

func (v1 VectorNn[T]) String() string {
	return fmt.Sprint(v1.Items)
}
