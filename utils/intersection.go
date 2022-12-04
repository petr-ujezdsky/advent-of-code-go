package utils

type IntersectionType int

const (
	None IntersectionType = iota
	Partial
	Inside
	Identical
	Wraps
)

// IntervalIntersection finds common intersection of two intervals (lowA, highA) and (lowB, highB)
// see https://scicomp.stackexchange.com/a/26260
func IntervalIntersection(lowA, highA, lowB, highB int) (int, int, bool) {
	if highA < lowB || highB < lowA {
		// no intersection
		return 0, 0, false
	}

	low := Max(lowA, lowB)
	high := Min(highA, highB)

	return low, high, true
}

func IntervalIntersectionDetail(lowA, highA, lowB, highB int) (IntersectionType, int, int) {
	// same
	if lowA == lowB && highA == highB {
		return Identical, lowA, highA
	}

	low, high, ok := IntervalIntersection(lowA, highA, lowB, highB)

	if !ok {
		return None, 0, 0
	}

	// common interval is whole B -> whole B is inside A
	if low == lowB && high == highB {
		return Inside, low, high
	}

	// common interval is whole A -> B wraps the whole A
	if low == lowA && high == highA {
		return Wraps, low, high
	}

	// partial otherwise
	return Partial, low, high
}
