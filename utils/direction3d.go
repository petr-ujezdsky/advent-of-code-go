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

// Direction3D26Arr contains all directions in 3D space
// +Y axis is right
// +Z axis is up
// +X axis is toward you
var Direction3D26Arr = [26]Vector3i{
	{X: 0, Y: 0, Z: 1},
	{X: 0, Y: 0, Z: -1},
	{X: 0, Y: 1, Z: 0},
	{X: 0, Y: 1, Z: 1},
	{X: 0, Y: 1, Z: -1},
	{X: 0, Y: -1, Z: 0},
	{X: 0, Y: -1, Z: 1},
	{X: 0, Y: -1, Z: -1},
	{X: 1, Y: 0, Z: 0},
	{X: 1, Y: 0, Z: 1},
	{X: 1, Y: 0, Z: -1},
	{X: 1, Y: 1, Z: 0},
	{X: 1, Y: 1, Z: 1},
	{X: 1, Y: 1, Z: -1},
	{X: 1, Y: -1, Z: 0},
	{X: 1, Y: -1, Z: 1},
	{X: 1, Y: -1, Z: -1},
	{X: -1, Y: 0, Z: 0},
	{X: -1, Y: 0, Z: 1},
	{X: -1, Y: 0, Z: -1},
	{X: -1, Y: 1, Z: 0},
	{X: -1, Y: 1, Z: 1},
	{X: -1, Y: 1, Z: -1},
	{X: -1, Y: -1, Z: 0},
	{X: -1, Y: -1, Z: 1},
	{X: -1, Y: -1, Z: -1},
}

//func Generate() {
//	steps := []int{0, 1, -1}
//	for _, x := range steps {
//		for _, y := range steps {
//			for _, z := range steps {
//				if x == 0 && y == 0 && z == 0 {
//					continue
//				}
//				fmt.Printf("{X: %v, Y: %v, Z: %v},\n", x, y, z)
//			}
//		}
//	}
//}
