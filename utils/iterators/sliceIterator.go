package iterators

type SliceIterator[T any] struct {
	slice    []T
	position int
	step     int
}

func NewSliceIterator[T any](s []T) *SliceIterator[T] {
	return &SliceIterator[T]{
		slice:    s,
		position: -1,
		step:     1,
	}
}

func NewReversedSliceIterator[T any](s []T) *SliceIterator[T] {
	return &SliceIterator[T]{
		slice:    s,
		position: len(s),
		step:     -1,
	}
}

func (i *SliceIterator[T]) HasNext() bool {
	nextPosition := i.position + i.step
	return nextPosition >= 0 && nextPosition < len(i.slice)
}

func (i *SliceIterator[T]) Next() T {
	i.position += i.step
	return i.slice[i.position]
}
