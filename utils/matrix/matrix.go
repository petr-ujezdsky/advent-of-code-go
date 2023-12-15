package matrix

import (
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/slices"
)

// Matrix is array of columns of values T
// This matrix values are also indexes of given values:
// 0 3 6
// 1 4 7
// 2 5 8
//
// The multi-array of these values is following
// [[0, 1, 2], [3, 4, 5], [6, 7, 8]]
type Matrix[T any] struct {
	Columns       [][]T
	Width, Height int
}

type MatrixNumber[T utils.Number] struct {
	Matrix[T]
}

type MatrixInt = MatrixNumber[int]
type MatrixFloat = MatrixNumber[float64]

var NewMatrixInt = NewMatrixNumber[int]

func NewMatrixNumber[T utils.Number](width, height int) MatrixNumber[T] {
	return MatrixNumber[T]{NewMatrix[T](width, height)}
}

func NewMatrix[T any](width, height int) Matrix[T] {
	matrixCols := make([][]T, width)
	cells := make([]T, width*height)

	// ensure data locality
	for col := range matrixCols {
		matrixCols[col], cells = cells[:height], cells[height:]
	}

	return Matrix[T]{matrixCols, width, height}
}
func NewMatrixColumnNotation[T any](columns [][]T) Matrix[T] {
	width := len(columns)
	height := len(columns[0])

	matrix := NewMatrix[T](width, height)

	for x, col := range columns {
		for y, value := range col {
			matrix.Columns[x][y] = value
		}
	}

	return matrix
}

// NewMatrixNumberRowNotation converts matrix from row-first notation to column-first notation
func NewMatrixNumberRowNotation[T utils.Number](rows [][]T) MatrixNumber[T] {
	return MatrixNumber[T]{NewMatrixRowNotation(rows)}
}

// NewMatrixRowNotation converts matrix from row-first notation to column-first notation
func NewMatrixRowNotation[T any](rows [][]T) Matrix[T] {
	width := len(rows[0])
	height := len(rows)

	matrix := NewMatrix[T](width, height)

	for y, row := range rows {
		for x, value := range row {
			matrix.Columns[x][y] = value
		}
	}

	return matrix
}

func (m Matrix[T]) SetAll(value T) Matrix[T] {
	for x := 0; x < m.Width; x++ {
		for y := 0; y < m.Height; y++ {
			m.Columns[x][y] = value
		}
	}

	return m
}

func (m MatrixNumber[T]) SetAll(value T) MatrixNumber[T] {
	return MatrixNumber[T]{m.Matrix.SetAll(value)}
}

func (m Matrix[T]) Get(x, y int) T {
	return m.Columns[x][y]
}

func (m Matrix[T]) SetSafe(x, y int, value T) bool {
	if x < 0 || x >= m.Width || y < 0 || y >= m.Height {
		return false
	}

	m.Columns[x][y] = value
	return true
}

func (m Matrix[T]) Set(x, y int, value T) {
	m.Columns[x][y] = value
}

func (m Matrix[T]) SetV(pos utils.Vector2i, value T) {
	m.Columns[pos.X][pos.Y] = value
}

func (m Matrix[T]) SetVSafe(pos utils.Vector2i, value T) bool {
	return m.SetSafe(pos.X, pos.Y, value)
}

func (m Matrix[T]) GetSafe(x, y int) (T, bool) {
	if x < 0 || x >= m.Width || y < 0 || y >= m.Height {
		var nothing T
		return nothing, false
	}

	return m.Columns[x][y], true
}

func (m Matrix[T]) GetV(pos utils.Vector2i) T {
	return m.Columns[pos.X][pos.Y]
}

func (m Matrix[T]) GetVSafe(pos utils.Vector2i) (T, bool) {
	return m.GetSafe(pos.X, pos.Y)
}

func (m Matrix[T]) GetWidth() int {
	return m.Width
}

func (m Matrix[T]) GetHeight() int {
	return m.Height
}

func (m Matrix[T]) Transpose() Matrix[T] {
	transposed := NewMatrix[T](m.Height, m.Width)

	for x, col := range m.Columns {
		for y, val := range col {
			transposed.Columns[y][x] = val
		}
	}

	return transposed
}

func (m MatrixNumber[T]) Transpose() MatrixNumber[T] {
	return MatrixNumber[T]{m.Matrix.Transpose()}
}

func (m Matrix[T]) Rotate90CounterClockwise(steps int) Matrix[T] {
	steps = utils.ModFloor(steps, 4)

	switch steps {
	case 0:
		return m
	case 1:
		return NewMatrixRowNotation(slices.Reverse(m.Columns))
	case 2:
		cols := slices.Reverse(m.Columns)
		for i, col := range cols {
			cols[i] = slices.Reverse(col)
		}
		return NewMatrixColumnNotation(cols)
	case 3:
		rows := slices.Clone(m.Columns)
		for i, row := range rows {
			rows[i] = slices.Reverse(row)
		}
		return NewMatrixRowNotation(rows)
	}

	panic("Can not happen")
}

func (m Matrix[T]) FlipHorizontal() Matrix[T] {
	return NewMatrixColumnNotation(slices.Reverse(m.Columns))
}

func (m Matrix[T]) FlipVertical() Matrix[T] {
	cols := slices.Clone(m.Columns)
	for i, col := range cols {
		cols[i] = slices.Reverse(col)
	}
	return NewMatrixColumnNotation(cols)
}

func (m MatrixNumber[T]) ArgMax() (utils.Vector2i, T) {
	max, xmax, ymax := m.Columns[0][0], 0, 0

	for x, col := range m.Columns {
		for y, val := range col {
			if val > max {
				max, xmax, ymax = val, x, y
			}
		}
	}

	return utils.Vector2i{xmax, ymax}, max
}

func (m Matrix[T]) Bounds() utils.BoundingRectangle {
	return utils.BoundingRectangle{
		Horizontal: utils.IntervalI{
			Low:  0,
			High: m.Width,
		},
		Vertical: utils.IntervalI{
			Low:  0,
			High: m.Height,
		},
	}
}

func (m Matrix[T]) Clone() Matrix[T] {
	cols := slices.Clone(m.Columns)
	for i, col := range cols {
		cols[i] = slices.Clone(col)
	}
	return NewMatrixColumnNotation(cols)
}

func (m Matrix[T]) String() string {
	return StringFmt[T](m, FmtNative[T])
}

func (m Matrix[T]) StringFmt(formatter ValueFormatter[T]) string {
	return StringFmtSeparator[T](m, " ", formatter)
}

func (m Matrix[T]) StringFmtSeparator(separator string, formatter ValueFormatter[T]) string {
	return StringFmtSeparator[T](m, separator, formatter)
}
