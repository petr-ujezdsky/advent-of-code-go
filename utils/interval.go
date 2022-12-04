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

func (i Interval[T]) Intersection(i2 Interval[T]) (T, T, bool) {
	return IntervalIntersection[T](i.Low, i.High, i2.Low, i2.High)
}

func (i Interval[T]) IntersectionDetail(i2 Interval[T]) (IntersectionType, T, T) {
	return IntervalIntersectionDetail[T](i.Low, i.High, i2.Low, i2.High)
}
