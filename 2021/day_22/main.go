package day_22

import (
	"bufio"
	"fmt"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"io"
	"math"
	"regexp"
	"strings"
)

type Vector3i = utils.Vector3i

type IntersectionType int

const (
	None IntersectionType = iota
	Partial
	Inside
	Wraps
)

var metric = utils.Metric{}

type Cube struct {
	Low, High Vector3i
	Value     bool
}

func NewCubeSymmetric(halfSideLength int, value bool) Cube {
	return Cube{
		Low:   Vector3i{-halfSideLength, -halfSideLength, -halfSideLength},
		High:  Vector3i{halfSideLength, halfSideLength, halfSideLength},
		Value: value,
	}
}

func NewCubeOrigin(sideLength int, value bool) Cube {
	return Cube{
		Low:   Vector3i{0, 0, 0},
		High:  Vector3i{sideLength, sideLength, sideLength},
		Value: value,
	}
}

func (c Cube) Contains(p Vector3i) bool {
	return c.Low.X <= p.X && p.X <= c.High.X &&
		c.Low.Y <= p.Y && p.Y <= c.High.Y &&
		c.Low.Z <= p.Z && p.Z <= c.High.Z
}

func (c Cube) ContainsWholeCube(c2 Cube) bool {
	return c.Contains(c2.Low) && c.Contains(c2.High)
}

func intervalIntersection(lowA, highA, lowB, highB int) IntersectionType {
	low, high, ok := utils.IntervalIntersection(lowA, highA, lowB, highB)

	if !ok {
		return None
	}

	// common interval is whole B -> whole B is inside A
	if low == lowB && high == highB {
		return Inside
	}

	// common interval is whole A -> B wraps the whole A
	if low == lowA && high == highA {
		return Wraps
	}

	// partial otherwise
	return Partial
}

func (c Cube) Intersect(c2 Cube) IntersectionType {
	ix := intervalIntersection(c.Low.X, c.High.X, c2.Low.X, c2.High.X)
	iy := intervalIntersection(c.Low.Y, c.High.Y, c2.Low.Y, c2.High.Y)
	iz := intervalIntersection(c.Low.Z, c.High.Z, c2.Low.Z, c2.High.Z)

	// no intersection in at least 1 dimension -> no intersection at all
	if ix == None || iy == None || iz == None {
		return None
	}

	// c2 is inside in all dimensions -> c2 is inside c
	if ix == Inside && iy == Inside && iz == Inside {
		return Inside
	}

	// c2 wraps c in all dimensions -> c2 wraps c
	if ix == Wraps && iy == Wraps && iz == Wraps {
		return Wraps
	}

	// anything other is partial
	return Partial
}

func (c Cube) Size() Vector3i {
	return c.High.Subtract(c.Low).Add(Vector3i{1, 1, 1})
}

func (c Cube) Volume() int {
	size := c.Size()
	return size.X * size.Y * size.Z
}

//var splittingSteps = []Vector3i{
//	{0, 0, 0},  // #1 no move (origin)
//	{0, 1, 0},  // #2 backwards
//	{1, 0, 0},  // #3 right
//	{0, -1, 0}, // #4 onwards
//	//{-1, 0, 1}, // #5 left & up
//	//{0, 1, 0},  // #6 backwards
//	//{1, 0, 0},  // #7 right
//	//{0, -1, 0}, // #8 onwards
//}

func (c Cube) Split() []Cube {
	size := c.High.Subtract(c.Low) //.Add(Vector3i{1, 1, 1})

	if size == (Vector3i{0, 0, 0}) {
		// can not divide single cell
		panic("Can not divide single cell")
	}

	half := size.Divide(2)
	halfAndOne := half.Add(Vector3i{1, 1, 1})

	cubes := make([]Cube, 8)

	i := 0
	for kx := 0; kx < 2; kx++ {
		for ky := 0; ky < 2; ky++ {
			for kz := 0; kz < 2; kz++ {
				k := Vector3i{kx, ky, kz}
				low := c.Low.Add(halfAndOne.MultiplyParts(k))
				high := low.Add(half)

				cubes[i] = Cube{
					Low:   low,
					High:  high,
					Value: c.Value,
				}

				i++
			}
		}
	}

	return cubes
}

var regexCube = regexp.MustCompile("(on|off) x=(-?\\d+)\\.\\.(-?\\d+),y=(-?\\d+)\\.\\.(-?\\d+),z=(-?\\d+)\\.\\.(-?\\d+)")

func resolveOnOff(probe Cube, cubes []Cube) (bool, bool) {
	for _, cube := range cubes {
		intersectionType := cube.Intersect(probe)

		switch intersectionType {
		case None:
		// continue with other cubes
		case Inside:
			// probe is inside the cube
			return cube.Value, true
		case Wraps, Partial:
			return false, false
		}
	}

	panic("Should match on whole world")
	//return false, false
}

func countRecursive(probe Cube, cubes []Cube) int {
	on, ok := resolveOnOff(probe, cubes)
	if ok {
		if on {
			return probe.Volume()
		}

		return 0
	}

	// split
	subProbes := probe.Split()
	count := 0
	for _, subProbe := range subProbes {
		count += countRecursive(subProbe, cubes)
	}

	metric.TickSum(100_000, count)

	return count
}

func FasterCount(world Cube, cubes []Cube) int {
	// start investigation with *last* added cube and so on
	cubes = utils.Reverse(cubes)

	// make world size divisible by 2 and symmetric
	size := world.Size()
	maxSize := utils.Max(utils.Max(size.X, size.Y), size.Z)

	nextPow2 := utils.NextPowOf2(maxSize + 1)
	biggerWorld := NewCubeOrigin(nextPow2-1, false)

	// shift world to original location
	biggerWorld.Low = biggerWorld.Low.Add(world.Low)
	biggerWorld.High = biggerWorld.High.Add(world.Low)

	// final cube is world itself
	cubes = append(cubes, biggerWorld)

	return countRecursive(biggerWorld, cubes)
}

func NaiveCount(world Cube, cubes []Cube) int {
	count := 0

	// start investigation with *last* added cube and so on
	cubes = utils.Reverse(cubes)

	// final cube is world itself
	cubes = append(cubes, world)

	return count
}

func ParseInput(r io.Reader) ([]Cube, Cube) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	world := Cube{
		Low:   utils.NewVector3nRepeated(math.MaxInt),
		High:  utils.NewVector3nRepeated(math.MinInt),
		Value: false,
	}

	var cubes []Cube
	for scanner.Scan() {
		parts := regexCube.FindStringSubmatch(scanner.Text())
		x1, x2 := utils.ParseInt(parts[2]), utils.ParseInt(parts[3])
		y1, y2 := utils.ParseInt(parts[4]), utils.ParseInt(parts[5])
		z1, z2 := utils.ParseInt(parts[6]), utils.ParseInt(parts[7])

		cube := Cube{
			Low:   Vector3i{x1, y1, z1},
			High:  Vector3i{x2, y2, z2},
			Value: parts[1] == "on",
		}

		world.Low = world.Low.Min(cube.Low)
		world.High = world.High.Max(cube.High)

		cubes = append(cubes, cube)
	}

	return cubes, world
}

type IntervalI = utils.IntervalI

type Cube2 struct {
	X, Y, Z IntervalI
	On      bool
}

func (c Cube2) Volume() int {
	return c.X.Size() * c.Y.Size() * c.Z.Size()
}

func (c Cube2) Subtract(c2 Cube2) []Cube2 {
	ix, ok := c.X.Intersection(c2.X)
	if !ok {
		return []Cube2{c}
	}

	iy, ok := c.Y.Intersection(c2.Y)
	if !ok {
		return []Cube2{c}
	}

	iz, ok := c.Z.Intersection(c2.Z)
	if !ok {
		return []Cube2{c}
	}

	if c.X == ix && c.Y == iy && c.Z == iz {
		// whole c is inside c2 -> result is empty
		// the algorithm below will work, this is just a shortcut
		return nil
	}

	xs := []IntervalI{{c.X.Low, ix.Low - 1}, ix, {ix.High + 1, c.X.High}}
	ys := []IntervalI{{c.Y.Low, iy.Low - 1}, iy, {iy.High + 1, c.Y.High}}
	zs := []IntervalI{{c.Z.Low, iz.Low - 1}, iz, {iz.High + 1, c.Z.High}}

	// generate all 27 sub-cubes
	subs := make([]Cube2, 0, 27)
	for xj, xi := range xs {
		if xi.IsInversed() {
			// invalid interval - intersection is at the beginning (or end) of the cube
			continue
		}
		for yj, yi := range ys {
			if yi.IsInversed() {
				// invalid interval - intersection is at the beginning (or end) of the cube
				continue
			}
			for zj, zi := range zs {
				// intersection cube is not part of the output
				if xj == 1 && yj == 1 && zj == 1 {
					continue
				}

				if zi.IsInversed() {
					// invalid interval - intersection is at the beginning (or end) of the cube
					continue
				}

				sub := Cube2{
					X:  xi,
					Y:  yi,
					Z:  zi,
					On: c.On,
				}

				subs = append(subs, sub)
			}
		}
	}

	return subs
}

func ExtraFastCount(cubes []Cube2) int {
	var onCubes []Cube2
	for i, cube := range cubes {
		if cube.On {
			subs := []Cube2{cube}

			// subtract all current on-cubes from cube
			for _, onCube := range onCubes {
				var newSubs []Cube2
				for _, sub := range subs {
					newSubs = append(newSubs, sub.Subtract(onCube)...)
				}
				subs = newSubs
			}
			onCubes = append(onCubes, subs...)
		} else {
			var newOnCubes []Cube2

			// subtract cube from all current on-cubes
			for _, onCube := range onCubes {
				newOnCubes = append(newOnCubes, onCube.Subtract(cube)...)
			}
			onCubes = newOnCubes
		}
		fmt.Printf("Cube #%v (%v), onCubes count: %v\n", i, cube.On, len(onCubes))
	}

	count := 0
	for _, onCube := range onCubes {
		count += onCube.Volume()
	}

	return count
}

func ParseInput2(r io.Reader) []Cube2 {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	var cubes []Cube2
	for scanner.Scan() {
		line := scanner.Text()
		ints := utils.ExtractInts(line, true)

		on := strings.HasPrefix(line, "on")

		cube := Cube2{
			X:  utils.NewInterval(ints[0], ints[1]),
			Y:  utils.NewInterval(ints[2], ints[3]),
			Z:  utils.NewInterval(ints[4], ints[5]),
			On: on,
		}

		cubes = append(cubes, cube)
	}

	return cubes
}
