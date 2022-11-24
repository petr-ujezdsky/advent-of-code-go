package utils

type Matrix3[T any] struct {
	Cells                 [][][]T
	Width, Height, Length int
}

type Matrix3n[T Number] Matrix3[T]

type Matrix3i = Matrix3n[int]
type Matrix3f = Matrix3n[float64]

func NewMatrix3[T Number](width, height, length int) Matrix3n[T] {
	matrixCols := make([][][]T, width)
	cells := make([]T, width*height*length)

	// ensure data locality
	for col := range matrixCols {
		matrixCols[col] = make([][]T, height)
		for row := range matrixCols[col] {
			matrixCols[col][row], cells = cells[:length], cells[length:]
		}
	}

	return Matrix3n[T]{matrixCols, width, height, length}
}

func (m Matrix3n[T]) Get(x, y, z int) T {
	return m.Cells[x][y][z]
}

func (m Matrix3n[T]) GetSafe(x, y, z int) (T, bool) {
	if x < 0 || x >= m.Width || y < 0 || y >= m.Height || z < 0 || z >= m.Length {
		var nothing T
		return nothing, false
	}

	return m.Get(x, y, z), true
}

func (m Matrix3n[T]) Set(x, y, z int, value T) {
	m.Cells[x][y][z] = value
}
