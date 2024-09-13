package utils

type BoundingBox struct {
	XInterval, YInterval, ZInterval IntervalI
}

func NewBoundingBox(point Vector3i) BoundingBox {
	return BoundingBox{
		XInterval: IntervalI{Low: point.X, High: point.X},
		YInterval: IntervalI{Low: point.Y, High: point.Y},
		ZInterval: IntervalI{Low: point.Z, High: point.Z},
	}
}

func NewBoundingBoxPoints(pointA, pointB Vector3i) BoundingBox {
	return BoundingBox{
		XInterval: NewInterval(pointA.X, pointB.X),
		YInterval: NewInterval(pointA.Y, pointB.Y),
		ZInterval: NewInterval(pointA.Z, pointB.Z),
	}
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

// Volume returns total box volume
func (b1 BoundingBox) Volume() int {
	return b1.XInterval.Size() * b1.YInterval.Size() * b1.ZInterval.Size()
}

func (b1 BoundingBox) MinPoint() Vector3i {
	return Vector3i{
		X: b1.XInterval.Low,
		Y: b1.YInterval.Low,
		Z: b1.ZInterval.Low,
	}
}

func (b1 BoundingBox) MaxPoint() Vector3i {
	return Vector3i{
		X: b1.XInterval.High,
		Y: b1.YInterval.High,
		Z: b1.ZInterval.High,
	}
}
