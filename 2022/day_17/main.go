package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"io"
	"strings"
)

type JetDirection = int
type ShapeRowPixels = []uint16

var (
	walls        = utils.ParseBinary16("0001000000010000")
	fullPixelRow = utils.ParseBinary16("0001111111110000")
)

var (
	shapeLine = []uint16{
		utils.ParseBinary16("0000001111000000"),
	}
	shapePlus = []uint16{
		utils.ParseBinary16("0000000100000000"),
		utils.ParseBinary16("0000001110000000"),
		utils.ParseBinary16("0000000100000000"),
	}
	shapeL = []uint16{
		utils.ParseBinary16("0000001110000000"),
		utils.ParseBinary16("0000000010000000"),
		utils.ParseBinary16("0000000010000000"),
	}
	shapeI = []uint16{
		utils.ParseBinary16("0000001000000000"),
		utils.ParseBinary16("0000001000000000"),
		utils.ParseBinary16("0000001000000000"),
		utils.ParseBinary16("0000001000000000"),
	}
	shapeSquare = []uint16{
		utils.ParseBinary16("0000001100000000"),
		utils.ParseBinary16("0000001100000000"),
	}

	shapeTypes = []ShapeRowPixels{
		shapeLine,
		shapePlus,
		shapeL,
		shapeI,
		shapeSquare,
	}
)

type PixelShape struct {
	pixelRows ShapeRowPixels
	//shapeTypeIndex int
	yBottom int
}

func (s1 PixelShape) GetYTop() int {
	return s1.yBottom + len(s1.pixelRows) - 1
}

func (s1 PixelShape) Move(step utils.Vector2i) PixelShape {
	pixelRows := s1.pixelRows

	if step.X != 0 {
		pixelRows = utils.ShallowCopy(s1.pixelRows)

		for i := range pixelRows {
			if step.X > 0 {
				pixelRows[i] >>= step.X
			} else {
				pixelRows[i] <<= -step.X
			}
		}
	}

	return PixelShape{
		pixelRows: pixelRows,
		yBottom:   s1.yBottom + step.Y,
	}
}

func (s1 PixelShape) AddShape(shape PixelShape) PixelShape {
	for y := shape.yBottom; y < shape.yBottom+len(shape.pixelRows); y++ {
		pixelRowShape := shape.GetPixelRow(y)
		pixelRowWorld := s1.GetPixelRow(y)

		// add bits
		s1 = s1.SetPixelRow(y, pixelRowShape|pixelRowWorld)
		//s1.pixelRows[y-s1.yBottom] = pixelRowShape | pixelRowWorld
	}

	return s1
}

func (s1 PixelShape) Trunc(toSize int) PixelShape {
	from := utils.Max(0, len(s1.pixelRows)-toSize)

	s1.pixelRows = s1.pixelRows[from:]
	s1.yBottom += from

	return s1
}

func (s1 PixelShape) GetPixelRow(y int) uint16 {
	index := y - s1.yBottom

	if 0 <= index && index < len(s1.pixelRows) {
		return s1.pixelRows[index]
	}

	return walls
}

func (s1 PixelShape) SetPixelRow(y int, pixelRow uint16) PixelShape {
	index := y - s1.yBottom

	if index < 0 {
		panic("Can not set rows below shape")
	}

	// add extra rows if needed
	newRowsCount := index - len(s1.pixelRows) + 1
	for i := 0; i < newRowsCount; i++ {
		s1.pixelRows = append(s1.pixelRows, walls)
	}

	// set row
	s1.pixelRows[index] = pixelRow
	return s1
}

func (s1 PixelShape) String() string {
	sb := &strings.Builder{}

	for i := len(s1.pixelRows) - 1; i >= 0; i-- {
		// row height
		sb.WriteString(fmt.Sprintf("%10d ", s1.yBottom+i))

		pixelRow := s1.pixelRows[i]

		str := fmt.Sprintf("%.16b\n", pixelRow)
		str = strings.ReplaceAll(str, "0", " ")
		str = strings.ReplaceAll(str, "1", "#")
		sb.WriteString(str)
	}

	return sb.String()
}

func CollidesMany(shape PixelShape, world PixelShape) bool {
	for y := shape.yBottom; y < shape.yBottom+len(shape.pixelRows); y++ {
		pixelRowShape := shape.GetPixelRow(y)
		pixelRowWorld := world.GetPixelRow(y)

		// bits compare
		if pixelRowShape&pixelRowWorld == 0 {
			// no collision
			continue
		}

		return true
	}

	return false
}

func MoveOrStay(shape PixelShape, step utils.Vector2i, world PixelShape) (PixelShape, bool) {
	shapeMoved := shape.Move(step)

	if CollidesMany(shapeMoved, world) {
		return shape, false
	}

	return shapeMoved, true
}

var metric = utils.NewMetric("Rocks count").Enable()

func InspectFallingRocks(jetDirections []JetDirection, rocksCount int) int {
	iShapeType := 0
	iJetDirection := 0

	// start with floor
	world := PixelShape{
		pixelRows: ShapeRowPixels{fullPixelRow},
		yBottom:   0,
	}

	//fmt.Println(world.String())

	for iRock := 0; iRock < rocksCount; iRock++ {

		sameBeginning := iShapeType == 0 && iJetDirection == 0

		shape := PixelShape{
			pixelRows: shapeTypes[iShapeType],
			yBottom:   world.GetYTop() + 4,
		}
		iShapeType = (iShapeType + 1) % len(shapeTypes)

		for {
			jetDirection := jetDirections[iJetDirection%len(jetDirections)]
			iJetDirection = (iJetDirection + 1) % len(jetDirections)

			// move sideways using jet stream, if possible
			shape, _ = MoveOrStay(shape, utils.Vector2i{X: jetDirection, Y: 0}, world)

			// move down
			var moved bool
			shape, moved = MoveOrStay(shape, utils.Vector2i{X: 0, Y: -1}, world)

			// could not move -> rest
			if !moved {
				break
			}
		}

		if sameBeginning && iRock != 0 {
			fmt.Printf("Same beginning! Resting pos %v\n", shape.yBottom)
		}

		// rest the shape
		world = world.AddShape(shape)

		// keep only last 256 pixels
		world = world.Trunc(256)

		metric.TickTime(1_000_000)

		//fmt.Println(world.String())
	}

	metric.Finished()

	return world.GetYTop()
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
