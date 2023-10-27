package utils

type BoundingRectangle struct {
	Horizontal, Vertical IntervalI
}

func NewBoundingRectangleFromPoints(point1, point2 Vector2i) BoundingRectangle {
	return BoundingRectangle{
		Horizontal: NewInterval(Min(point1.X, point2.X), Max(point1.X, point2.X)),
		Vertical:   NewInterval(Min(point1.Y, point2.Y), Max(point1.Y, point2.Y)),
	}
}

func (b1 BoundingRectangle) Contains(pos Vector2i) bool {
	return b1.Horizontal.Contains(pos.X) && b1.Vertical.Contains(pos.Y)
}

func (b1 BoundingRectangle) Intersection(b2 BoundingRectangle) (BoundingRectangle, bool) {
	horizontal, ok := b1.Horizontal.Intersection(b2.Horizontal)
	if !ok {
		return BoundingRectangle{}, false
	}

	vertical, ok := b1.Vertical.Intersection(b2.Vertical)
	if !ok {
		return BoundingRectangle{}, false
	}

	return BoundingRectangle{horizontal, vertical}, true
}

// Enlarge grows bounding rectangle to contain given point
func (b1 BoundingRectangle) Enlarge(point Vector2i) BoundingRectangle {
	return BoundingRectangle{
		Horizontal: b1.Horizontal.Enlarge(point.X),
		Vertical:   b1.Vertical.Enlarge(point.Y),
	}
}

func (b1 BoundingRectangle) Width() int {
	return b1.Horizontal.Size()
}

func (b1 BoundingRectangle) Height() int {
	return b1.Vertical.Size()
}
