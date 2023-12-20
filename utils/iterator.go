package utils

type Iterator[T any] interface {
	Next() T
	HasNext() bool
}
