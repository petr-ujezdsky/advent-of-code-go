package utils

import "fmt"

type vectorN[T any] struct {
	Items []T
}

type VectorNn[T Number] vectorN[T]

type VectorNi = VectorNn[int]
type VectorNf = VectorNn[float64]

func (v1 VectorNn[T]) String() string {
	return fmt.Sprint(v1.Items)
}
