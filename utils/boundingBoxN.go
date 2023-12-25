package utils

type BoundingBoxN struct {
	Intervals []IntervalI
}

func (b1 BoundingBoxN) Contains(coordinates []int) bool {
	if len(b1.Intervals) != len(coordinates) {
		panic("Dimension mismatch")
	}

	for i, interval := range b1.Intervals {
		if !interval.Contains(coordinates[i]) {
			return false
		}
	}

	return true
}

func (b1 BoundingBoxN) Intersection(b2 BoundingBoxN) (BoundingBoxN, bool) {
	if len(b1.Intervals) != len(b2.Intervals) {
		panic("Dimension mismatch")
	}

	intersections := make([]IntervalI, len(b1.Intervals))

	for i, interval1 := range b1.Intervals {
		interval2 := b2.Intervals[i]

		intersection, ok := interval1.Intersection(interval2)
		if !ok {
			return BoundingBoxN{}, false
		}
		intersections[i] = intersection
	}

	return BoundingBoxN{intersections}, true
}

// Enlarge grows bounding box to contain given point
func (b1 BoundingBoxN) Enlarge(coordinates []int) BoundingBoxN {
	if len(b1.Intervals) != len(coordinates) {
		panic("Dimension mismatch")
	}

	enlarged := make([]IntervalI, len(b1.Intervals))

	for i, interval := range b1.Intervals {
		enlarged[i] = interval.Enlarge(coordinates[i])
	}

	return BoundingBoxN{enlarged}
}

// Volume returns total box volume
func (b1 BoundingBoxN) Volume() int {
	volume := 1

	for _, interval := range b1.Intervals {
		volume *= interval.Size()
	}

	return volume
}
