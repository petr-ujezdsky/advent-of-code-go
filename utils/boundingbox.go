package utils

type BoundingBox struct {
	Horizontal, Vertical IntervalI
}

func (b1 BoundingBox) Contains(pos Vector2i) bool {
	return b1.Horizontal.Contains(pos.X) && b1.Vertical.Contains(pos.Y)
}

func (b1 BoundingBox) Intersection(b2 BoundingBox) (BoundingBox, bool) {
	horizontal, ok := b1.Horizontal.Intersection(b2.Horizontal)
	if !ok {
		return BoundingBox{}, false
	}

	vertical, ok := b1.Vertical.Intersection(b2.Vertical)
	if !ok {
		return BoundingBox{}, false
	}

	return BoundingBox{horizontal, vertical}, true
}
