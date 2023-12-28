package utils

import (
	"fmt"
	"math"
)

type MaxFinder[T any] struct {
	value int
	item  T
}

func NewMaxFinder[T any]() *MaxFinder[T] {
	return &MaxFinder[T]{value: math.MinInt}
}

func (mf *MaxFinder[T]) Inspect(value int, item T) {
	if value > mf.value {
		mf.value = value
		mf.item = item
	}
}

func (mf *MaxFinder[T]) Result() (int, T) {
	return mf.value, mf.item
}

func (mf *MaxFinder[T]) String() string {
	return fmt.Sprintf("Max %d for %v", mf.value, mf.item)
}
