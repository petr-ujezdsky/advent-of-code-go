package utils

// matrix2 is array of columns of values T
// This matrix values are also indexes of given values:
// 0 3 6
// 1 4 7
// 2 5 8
//
// The multi-array of these values is following
// [[0, 1, 2], [3, 4, 5], [6, 7, 8]]
type matrix2[T any] struct {
	Columns       [][]T
	Width, Height int
}

type Matrix2n[T Number] matrix2[T]

type Matrix2i = Matrix2n[int]
type Matrix2f = Matrix2n[float64]

func NewMatrix2[T Number](width, height int) Matrix2n[T] {
	matrixCols := make([][]T, width)
	cells := make([]T, width*height)

	for col := range matrixCols {
		matrixCols[col], cells = cells[:height], cells[height:]
	}

	return Matrix2n[T]{matrixCols, width, height}
}

// NewMatrix2RowNotation converts matrix from row-first notation to column-first notation
func NewMatrix2RowNotation[T Number](rows [][]T) Matrix2n[T] {
	width := len(rows[0])
	height := len(rows)

	matrix := NewMatrix2[T](width, height)

	for y, row := range rows {
		for x, value := range row {
			matrix.Set(x, y, value)
		}
	}

	return matrix
}

func (m Matrix2n[T]) Get(x, y int) T {
	return m.Columns[x][y]
}

func (m Matrix2n[T]) GetSafe(x, y int) (T, bool) {
	if x < 0 || x >= m.Width || y < 0 || y >= m.Height {
		var nothing T
		return nothing, false
	}

	return m.Get(x, y), true
}

func (m Matrix2n[T]) Set(x, y int, value T) {
	m.Columns[x][y] = value
}
