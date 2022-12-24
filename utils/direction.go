package utils

type Direction4 int

const (
	Up Direction4 = iota
	Right
	Down
	Left
)

var direction4ToStep = [4]Vector2i{
	{X: 0, Y: 1},
	{X: 1, Y: 0},
	{X: 0, Y: -1},
	{X: -1, Y: 0},
}

func (d Direction4) ToStep() Vector2i {
	return direction4ToStep[d]
}

func (d Direction4) Rotate(steps int) Direction4 {
	return Direction4(ModFloor(int(d)+steps, 4))
}

type Direction8 int

const (
	North Direction8 = iota
	NorthEast
	East
	SouthEast
	South
	SouthWest
	West
	NorthWest
)

var direction8ToStep = [8]Vector2i{
	{X: 0, Y: 1},
	{X: 1, Y: 1},
	{X: 1, Y: 0},
	{X: 1, Y: -1},
	{X: 0, Y: -1},
	{X: -1, Y: -1},
	{X: -1, Y: 0},
	{X: -1, Y: 1},
}

func (d Direction8) ToStep() Vector2i {
	return direction8ToStep[d]
}

func (d Direction8) Rotate(steps int) Direction8 {
	return Direction8(ModFloor(int(d)+steps, 8))
}
