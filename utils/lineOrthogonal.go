package utils

type LineOrthogonal2i struct {
	A, B Vector2i
}

func NewLineOrthogonal2i(a, b Vector2i) LineOrthogonal2i {
	if a.X != b.X && a.Y != b.Y {
		panic("Points must be orthogonal")
	}

	return LineOrthogonal2i{A: a, B: b}
}

func (line LineOrthogonal2i) Intersection(line2 LineOrthogonal2i) (LineOrthogonal2i, bool) {
	br1 := NewBoundingRectangleFromPoints(line.A, line.B)
	br2 := NewBoundingRectangleFromPoints(line2.A, line2.B)

	intersection, ok := br1.Intersection(br2)
	if !ok {
		return LineOrthogonal2i{}, false
	}

	intersectingLine := NewLineOrthogonal2i(
		Vector2i{X: intersection.Horizontal.Low, Y: intersection.Vertical.Low},
		Vector2i{X: intersection.Horizontal.High, Y: intersection.Vertical.High})

	return intersectingLine, true
}

func (line LineOrthogonal2i) Length() int {
	return line.B.Subtract(line.A).LengthManhattan() + 1
}
