package main

import (
	_ "embed"
	"fmt"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/parsers"
	"io"
	"strings"
)

type Cube struct {
	Name         string
	Index        int
	Box          utils.BoundingBox
	Stabilized   bool
	Above, Below []*Cube
}

func (cube Cube) String() string {
	return fmt.Sprintf("#%v %v %v", cube.Index+1, cube.Box.MinPoint(), cube.Box.MaxPoint())
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
		return false
	}

	// move 1 down
	movedBox := stepDown(cube.Box)

	// lookup collisions
	collided := false
	for _, otherCube := range cubes {
		if otherCube == cube {
			// skip self
			continue
		}

		if _, ok := movedBox.Intersection(otherCube.Box); ok {
			// can not move
			if otherCube.Stabilized {
				cube.Stabilized = true

				// link cubes
				cube.Below = append(cube.Below, otherCube)
				otherCube.Above = append(otherCube.Above, cube)
			}

			collided = true
		}
	}

	if collided {
		return false
	}

	// can move -> move
	cube.Box = movedBox

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
	for _, cubeAbove := range cube.Above {
		if len(cubeAbove.Below) <= 1 {
			// cubeAbove is sitting on *only* this cube -> can not disintegrate this cube
			return false
		}
	}

	return true
}

func DoWithInputPart02(world World) int {
	return 0
}

func ParseInput(r io.Reader) World {
	parseItem := func(str string, i int) *Cube {
		points := strings.Split(str, "~")

		pointA := parsePoint(points[0])
		pointB := parsePoint(points[1])

		return &Cube{
			Name:  string('A' + byte(i)),
			Index: i,
			Box:   utils.NewBoundingBoxPoints(pointA, pointB),
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
