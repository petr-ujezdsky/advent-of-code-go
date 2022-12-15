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

func (i Interval[T]) IsInversed() bool {
	return i.Low > i.High
}

func (i Interval[T]) Size() T {
	return i.High - i.Low + 1
}

func (i Interval[T]) Contains(v T) bool {
	return i.Low <= v && v <= i.High
}

func (i Interval[T]) Intersection(i2 Interval[T]) (Interval[T], bool) {
	low, high, ok := IntervalIntersection[T](i.Low, i.High, i2.Low, i2.High)
	return NewInterval(low, high), ok
}

func (i Interval[T]) IntersectionDetail(i2 Interval[T]) (IntersectionType, Interval[T]) {
	intersectionType, low, high := IntervalIntersectionDetail[T](i.Low, i.High, i2.Low, i2.High)
	return intersectionType, NewInterval(low, high)
}

func (i Interval[T]) Subtract(i2 Interval[T]) []Interval[T] {
	intersection, ok := i.Intersection(i2)

	if ok {
		// A ∩ B = A  ->  A - B = ∅
		if i == intersection {
			return nil
		}

		subs := make([]Interval[T], 0, 2)
		if intersection.Low > i.Low {
			subs = append(subs, NewInterval(i.Low, intersection.Low-1))
		}

		if intersection.High < i.High {
			subs = append(subs, NewInterval(intersection.High+1, i.High))
		}

		return subs
	}

	return []Interval[T]{i}
}

// Enlarge grows interval to contain given value
func (i Interval[T]) Enlarge(value T) Interval[T] {
	low := Min(i.Low, value)
	high := Max(i.High, value)

	return NewInterval(low, high)
}
