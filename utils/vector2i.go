package utils

type Vector2i struct {
	X, Y int
}

func (v1 Vector2i) Add(v2 Vector2i) Vector2i {
	return Vector2i{
		X: v1.X + v2.X,
		Y: v1.Y + v2.Y,
	}
}
