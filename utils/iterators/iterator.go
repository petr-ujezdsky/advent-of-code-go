package iterators

type Iterator[T any] interface {
	Next() T
	HasNext() bool
}

func ToSlice[T any](iterator Iterator[T]) []T {
	var slice []T

	for iterator.HasNext() {
		slice = append(slice, iterator.Next())
	}

	return slice
}

type joinedIterator[T any] struct {
	iterators       Iterator[Iterator[T]]
	currentIterator Iterator[T]
	currentValue    T
	valid           bool
}

func (i *joinedIterator[T]) Next() T {
	value := i.currentValue

	// find next in current iterator
	if i.currentIterator.HasNext() {
		i.currentValue = i.currentIterator.Next()
		return value
	}

	i.moveNextIterator()

	return value
}

func (i *joinedIterator[T]) moveNextIterator() {
	// go through remaining iterators
	for i.iterators.HasNext() {
		i.currentIterator = i.iterators.Next()

		if i.currentIterator.HasNext() {
			i.currentValue = i.currentIterator.Next()
			return
		}
	}

	// none has any value
	i.valid = false
}

func (i *joinedIterator[T]) HasNext() bool {
	return i.valid
}

func JoinIterators[T any](iterators ...Iterator[T]) Iterator[T] {
	i := joinedIterator[T]{
		iterators: NewSliceIterator(iterators),
		valid:     true,
	}

	i.moveNextIterator()

	return &i
}
