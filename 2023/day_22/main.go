package main

import (
	_ "embed"
	"fmt"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/parsers"
	"io"
	"strconv"
	"strings"
)

type Cube struct {
	Name         string
	Id           int
	Box          utils.BoundingBox
	Stabilized   bool
	Above, Below map[*Cube]struct{}
}

func (cube Cube) String() string {
	return fmt.Sprintf("%v %v %v", cube.Name, cube.Box.MinPoint(), cube.Box.MaxPoint())
}

type World struct {
	Cubes []*Cube
}

func DoWithInputPart01(world World) int {
	fallDown(world.Cubes)

	//for _, cube := range world.Cubes {
	//	fmt.Println(cube)
	//}

	return countDisintegratable(world.Cubes)
}

func fallDown(cubes []*Cube) {
	movableCubes := cubes

	for {
		var nextMovableCubes []*Cube
		//fmt.Printf("Movable cubes: %d\n", len(movableCubes))

		for _, cube := range movableCubes {
			for tryStepDown(cube, cubes) {
			}

			if !cube.Stabilized {
				nextMovableCubes = append(nextMovableCubes, cube)
			}
		}

		movableCubes = nextMovableCubes

		if len(movableCubes) == 0 {
			// all stabilized
			break
		}
	}
}

func tryStepDown(cube *Cube, cubes []*Cube) bool {
	if cube.Stabilized {
		return false
	}

	if cube.Box.ZInterval.Low == 1 {
		// at floor -> stabilize
		cube.Stabilized = true
		//fmt.Printf("Cube %v stabilized by floor\n", cube.Name)
		return false
	}

	// move 1 down
	movedBox := stepDown(cube.Box)

	// lookup collisions
	var below map[*Cube]struct{}

	for _, otherCube := range cubes {
		if otherCube == cube {
			// skip self
			continue
		}

		if _, ok := movedBox.Intersection(otherCube.Box); ok {
			// can not move
			if otherCube.Stabilized {
				cube.Stabilized = true
				//fmt.Printf("Cube %v stabilized by %v\n", cube.Name, otherCube.Name)
			}

			// link cubes
			if below == nil {
				below = make(map[*Cube]struct{})
			}
			below[otherCube] = struct{}{}

			if otherCube.Above == nil {
				otherCube.Above = make(map[*Cube]struct{})
			}
			otherCube.Above[cube] = struct{}{}
		}
	}

	cube.Below = below

	if len(below) > 0 {
		return false
	}

	// can move -> move
	//fmt.Printf("Cube %v moved down\n", cube.Name)
	cube.Box = movedBox

	// disconnect from above
	for otherCube := range cube.Above {
		delete(otherCube.Below, cube)
	}
	cube.Above = nil

	return true
}

func stepDown(box utils.BoundingBox) utils.BoundingBox {
	return utils.BoundingBox{
		XInterval: box.XInterval,
		YInterval: box.YInterval,
		ZInterval: utils.IntervalI{
			Low:  box.ZInterval.Low - 1,
			High: box.ZInterval.High - 1,
		},
	}
}

func countDisintegratable(cubes []*Cube) int {
	count := 0

	for _, cube := range cubes {
		if isDisintegratable(cube) {
			count++
			//fmt.Printf("%v can be disintegrated\n", cube)
		} else {
			//fmt.Printf("%v can NOT be disintegrated\n", cube)
		}
	}

	return count
}

func isDisintegratable(cube *Cube) bool {
	for cubeAbove := range cube.Above {
		if len(cubeAbove.Below) <= 1 {
			// cubeAbove is sitting on *only* this cube -> can not disintegrate this cube
			return false
		}
	}

	return true
}

func fallingCount(cube *Cube, removed map[*Cube]struct{}) {
	toBeRemoved := make(map[*Cube]struct{})

	for cubeAbove := range cube.Above {
		if _, ok := removed[cubeAbove]; ok {
			// already removed
			continue
		}

		// count below cubes that are not yet removed
		hasCubesBelow := false
		for cubeBelow := range cubeAbove.Below {
			if _, ok := removed[cubeBelow]; !ok {
				hasCubesBelow = true
				break
			}
		}

		if hasCubesBelow {
			// can not remove - will not fall down
			continue
		}

		// cubeAbove has no cubes below -> store to remove later
		toBeRemoved[cubeAbove] = struct{}{}
	}

	for cubeToRemove := range toBeRemoved {
		removed[cubeToRemove] = struct{}{}
	}

	for cubeRemoved := range toBeRemoved {
		fallingCount(cubeRemoved, removed)
	}
}

func DoWithInputPart02(world World) int {
	fallDown(world.Cubes)

	fallSum := 0

	for _, cube := range world.Cubes {
		removed := make(map[*Cube]struct{})
		removed[cube] = struct{}{}

		fallingCount(cube, removed)

		fallSum += len(removed) - 1
	}

	return fallSum
}

func ParseInput(r io.Reader) World {
	parseItem := func(str string, i int) *Cube {
		points := strings.Split(str, "~")

		pointA := parsePoint(points[0])
		pointB := parsePoint(points[1])

		id := i + 1

		return &Cube{
			Name: "#" + strconv.Itoa(id),
			Id:   id,
			Box:  utils.NewBoundingBoxPoints(pointA, pointB),
		}
	}

	items := parsers.ParseToObjectsIndexed(r, parseItem)
	return World{Cubes: items}
}

func parsePoint(str string) utils.Vector3i {
	coordinates := strings.Split(str, ",")

	return utils.Vector3i{
		X: utils.ParseInt(coordinates[0]),
		Y: utils.ParseInt(coordinates[1]),
		Z: utils.ParseInt(coordinates[2]),
	}
}
