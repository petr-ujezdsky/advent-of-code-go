package utils

type BoundingBox struct {
	XInterval, YInterval, ZInterval IntervalI
}

func (b1 BoundingBox) Contains(pos Vector3i) bool {
	return b1.XInterval.Contains(pos.X) && b1.YInterval.Contains(pos.Y) && b1.ZInterval.Contains(pos.Z)
}

func (b1 BoundingBox) Intersection(b2 BoundingBox) (BoundingBox, bool) {
	xInt, ok := b1.XInterval.Intersection(b2.XInterval)
	if !ok {
		return BoundingBox{}, false
	}

	yInt, ok := b1.YInterval.Intersection(b2.YInterval)
	if !ok {
		return BoundingBox{}, false
	}

	zInt, ok := b1.ZInterval.Intersection(b2.ZInterval)
	if !ok {
		return BoundingBox{}, false
	}

	return BoundingBox{xInt, yInt, zInt}, true
}

// Enlarge grows bounding box to contain given point
func (b1 BoundingBox) Enlarge(point Vector3i) BoundingBox {
	return BoundingBox{
		XInterval: b1.XInterval.Enlarge(point.X),
		YInterval: b1.YInterval.Enlarge(point.Y),
		ZInterval: b1.ZInterval.Enlarge(point.Z),
	}
}
