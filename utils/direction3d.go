package utils

type Direction3D6 int

//const (
//	Up Direction3D6 = iota
//	Right
//	Down
//	Left
//)

// Direction3D6Arr contains all major directions in 3D space
// +Y axis is right
// +Z axis is up
// +X axis is toward you
var Direction3D6Arr = [6]Vector3i{
	{X: 0, Y: 0, Z: 1},  //up
	{X: 0, Y: 1, Z: 0},  //right
	{X: 0, Y: 0, Z: -1}, //down
	{X: 0, Y: -1, Z: 0}, //left
	{X: -1, Y: 0, Z: 0}, //behind
	{X: 1, Y: 0, Z: 0},  //upfront
}
