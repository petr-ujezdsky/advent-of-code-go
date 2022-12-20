package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"io"
	"math"
)

type JetDirection = int
type ShapePixels = utils.Matrix[bool]

type World struct {
	Bounds []IShape
}

var shapeLine = utils.NewMatrixRowNotation([][]bool{
	{true, true, true, true},
})

var shapePlus = utils.NewMatrixRowNotation([][]bool{
	{false, true, false},
	{true, true, true},
	{false, true, false},
})

var shapeL = utils.NewMatrixRowNotation([][]bool{
	{true, true, true},
	{false, false, true},
	{false, false, true},
})

var shapeI = utils.NewMatrixRowNotation([][]bool{
	{true},
	{true},
	{true},
	{true},
})

var shapeSquare = utils.NewMatrixRowNotation([][]bool{
	{true, true},
	{true, true},
})

var shapeTypes = []ShapePixels{
	shapeLine,
	shapePlus,
	shapeL,
	shapeI,
	shapeSquare,
}

type IShape interface {
	GetPixel(pos utils.Vector2i) bool
	BoundingBox() utils.BoundingBox
}

type PixelShape struct {
	//pixels    ShapePixels
	shapeTypeIndex int
	position       utils.Vector2i
}

func (s1 PixelShape) GetPixel(pos utils.Vector2i) bool {
	return shapeTypes[s1.shapeTypeIndex].GetV(pos.Subtract(s1.position))
}

func (s1 PixelShape) BoundingBox() utils.BoundingBox {
	return utils.BoundingBox{
		Horizontal: utils.IntervalI{Low: s1.position.X, High: s1.position.X + shapeTypes[s1.shapeTypeIndex].Width - 1},
		Vertical:   utils.IntervalI{Low: s1.position.Y, High: s1.position.Y + shapeTypes[s1.shapeTypeIndex].Height - 1},
	}
}

type BigShape struct {
	boundingBox utils.BoundingBox
}

func (s1 BigShape) GetPixel(pos utils.Vector2i) bool {
	return s1.boundingBox.Contains(pos)
}

func (s1 BigShape) BoundingBox() utils.BoundingBox {
	return s1.boundingBox
}

func Collides(s1, s2 IShape) bool {
	boundingBox, ok := s1.BoundingBox().Intersection(s2.BoundingBox())
	if !ok {
		return false
	}

	for x := boundingBox.Horizontal.Low; x <= boundingBox.Horizontal.High; x++ {
		for y := boundingBox.Vertical.Low; y <= boundingBox.Vertical.High; y++ {
			pos := utils.Vector2i{X: x, Y: y}

			s1pixel := s1.GetPixel(pos)
			s2pixel := s2.GetPixel(pos)

			if s1pixel && s2pixel {
				return true
			}
		}
	}

	return false
}

func CollidesMany(shape IShape, shapes []IShape) bool {
	// reverse order - the highest shapes are at the top
	for i := len(shapes) - 1; i >= 0; i-- {
		other := shapes[i]
		if Collides(shape, other) {
			return true
		}
	}

	return false
}

func MoveOrStay(shape PixelShape, step utils.Vector2i, shapes []IShape, bounds []IShape) (PixelShape, bool) {
	shapeMoved := shape
	shapeMoved.position = shapeMoved.position.Add(step)

	if CollidesMany(shapeMoved, bounds) || CollidesMany(shapeMoved, shapes) {
		return shape, false
	}

	return shapeMoved, true
}

func initWorld() World {
	left := BigShape{utils.BoundingBox{
		Horizontal: utils.IntervalI{Low: -1, High: -1},
		Vertical:   utils.IntervalI{Low: -1, High: math.MaxInt},
	}}

	right := BigShape{utils.BoundingBox{
		Horizontal: utils.IntervalI{Low: 7, High: 7},
		Vertical:   utils.IntervalI{Low: -1, High: math.MaxInt},
	}}

	floor := BigShape{utils.BoundingBox{
		Horizontal: utils.IntervalI{Low: -1, High: 7},
		Vertical:   utils.IntervalI{Low: -1, High: -1},
	}}

	return World{[]IShape{
		left,
		right,
		floor,
	}}
}

var metric = utils.NewMetric("Rocks count").Enable()

func InspectFallingRocks(jetDirections []JetDirection, rocksCount int) int {
	world := initWorld()

	iShapeType := 0
	iJetDirection := 0
	height := 0
	var shapes []IShape

	for iRock := 0; iRock < rocksCount; iRock++ {

		sameBeginning := iShapeType == 0 && iJetDirection == 0

		//shapeType := shapeTypes[iShapeType]

		shape := PixelShape{
			//pixels:         shapeType,
			shapeTypeIndex: iShapeType,
			position:       utils.Vector2i{X: 2, Y: height + 3},
		}
		iShapeType = (iShapeType + 1) % len(shapeTypes)

		for {
			jetDirection := jetDirections[iJetDirection%len(jetDirections)]
			iJetDirection = (iJetDirection + 1) % len(jetDirections)

			// move sideways using jet stream, if possible
			shape, _ = MoveOrStay(shape, utils.Vector2i{X: jetDirection, Y: 0}, shapes, world.Bounds)

			// move down
			var moved bool
			shape, moved = MoveOrStay(shape, utils.Vector2i{X: 0, Y: -1}, shapes, world.Bounds)

			// could not move -> rest
			if !moved {
				break
			}
		}

		if sameBeginning {
			fmt.Printf("Same beginning! Resting pos %v\n", shape.position)
		}

		// rest the shape
		shapes = append(shapes, shape)

		// keep only last 30 shapes
		if len(shapes) > 30 {
			shapes = shapes[1:]
		}

		metric.TickTime(1_000_000)

		// store new height if higher
		height = utils.Max(height, shape.BoundingBox().Vertical.High+1)
	}

	metric.Finished()

	return height
}

func ParseInput(r io.Reader) []JetDirection {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	var directions []JetDirection
	for scanner.Scan() {
		directions = make([]JetDirection, len(scanner.Text()))

		for i, char := range scanner.Text() {
			if char == '<' {
				directions[i] = -1
			} else {
				directions[i] = 1
			}
		}
	}

	return directions
}
