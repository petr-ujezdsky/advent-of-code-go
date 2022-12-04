package utils

type Interval[T Number] struct {
	Low, High T
}

type IntervalI = Interval[int]
type IntervalF = Interval[float64]

func NewInterval[T Number](a, b T) Interval[T] {
	if a > b {
		return Interval[T]{b, a}
	}

	return Interval[T]{a, b}
}

func (i Interval[T]) Intersection(i2 Interval[T]) (Interval[T], bool) {
	low, high, ok := IntervalIntersection[T](i.Low, i.High, i2.Low, i2.High)
	return NewInterval(low, high), ok
}

func (i Interval[T]) IntersectionDetail(i2 Interval[T]) (IntersectionType, Interval[T]) {
	intersectionType, low, high := IntervalIntersectionDetail[T](i.Low, i.High, i2.Low, i2.High)
	return intersectionType, NewInterval(low, high)
}
