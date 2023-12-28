package utils

import "math"

type MinFinder[T any] struct {
	value int
	item  T
}

func NewMinFinder[T any]() *MinFinder[T] {
	return &MinFinder[T]{value: math.MaxInt}
}

func (mf *MinFinder[T]) Inspect(value int, item T) {
	if value < mf.value {
		mf.value = value
		mf.item = item
	}
}

func (mf *MinFinder[T]) Result() (int, T) {
	return mf.value, mf.item
}
