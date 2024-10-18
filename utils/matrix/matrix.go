package matrix

import (
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/slices"
	slices2 "slices"
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
var NewMatrixFloat = NewMatrixNumber[float64]

var NewMatrixColumnNotationFloat = NewMatrixColumnNotationNumber[float64]

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

func NewMatrixColumnNotationNumber[T utils.Number](columns [][]T) MatrixNumber[T] {
	return MatrixNumber[T]{NewMatrixColumnNotation(columns)}
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

func (m MatrixNumber[T]) Determinant() T {
	if m.Rank() == 1 {
		return 1
	}

	if m.Rank() == 2 {
		return m.determinant2()
	}

	if m.Rank() == 3 {
		return m.determinant3()
	}

	m.assertSquare()

	return m.determinantN()
}

func (m MatrixNumber[T]) determinant2() T {
	return m.Columns[0][0]*m.Columns[1][1] - m.Columns[1][0]*m.Columns[0][1]
}

func (m MatrixNumber[T]) determinant3() T {
	return m.Columns[0][0]*m.Columns[1][1]*m.Columns[2][2] +
		m.Columns[1][0]*m.Columns[2][1]*m.Columns[0][2] +
		m.Columns[2][0]*m.Columns[0][1]*m.Columns[1][2] -
		m.Columns[0][0]*m.Columns[2][1]*m.Columns[1][2] -
		m.Columns[1][0]*m.Columns[0][1]*m.Columns[2][2] -
		m.Columns[2][0]*m.Columns[1][1]*m.Columns[0][2]
}

func (m MatrixNumber[T]) determinantN() T {
	m.assertSquare()

	// create clone
	m2 := NewMatrixColumnNotationNumber(m.Columns)

	// gauss elimination
	swapCount := m2.GaussElimination()

	sign := T(1)
	if swapCount%2 != 0 {
		sign = T(-1)
	}

	// determinant is product of diagonal
	return sign * m2.DiagonalProduct()
}

func (m MatrixNumber[T]) Rank() int {
	m.assertSquare()

	return m.Height
}

func (m Matrix[T]) assertSquare() {
	if m.Height != m.Width {
		panic("Matrix is not square")
	}
}

func (m MatrixNumber[T]) DiagonalProduct() T {
	product := T(1)

	for i := 0; i < m.Width; i++ {
		product *= m.Get(i, i)
	}

	return product
}

func (m MatrixNumber[T]) Inverse() bool {
	m2 := NewMatrixNumber[T](2*m.Width, m.Height)

	// copy m into m2
	for x, column := range m.Columns {
		for y, v := range column {
			m2.Set(x, y, v)
		}
	}

	// insert identity matrix to the right
	for x := m2.Width / 2; x < m2.Width; x++ {
		m2.Set(x, x-m2.Width/2, 1)
	}

	// do gauss elimination
	m2.GaussElimination()

	// do Jordan elimination
	ok := m2.JordanContinue()
	if !ok {
		return false
	}

	// now the identity matrix has become the inverse matrix - copy it back into m
	for x, column := range m.Columns {
		for y := range column {
			m.Set(x, y, m2.Get(x+m.Width, y))
		}
	}

	return true
}

func (m MatrixNumber[T]) MultiplyN(coef T) {
	for x, column := range m.Columns {
		for y, value := range column {
			m.Set(x, y, coef*value)
		}
	}
}

func (m MatrixNumber[T]) MultiplyV(a utils.VectorNn[T]) utils.VectorNn[T] {
	ma := utils.NewVectorNn[T](len(a.Items))

	for x, column := range m.Columns {
		for y, v := range column {
			ma.Items[y] = ma.Items[y] + v*a.Items[x]
		}
	}

	return ma
}

// GaussElimination
// see https://cs.wikipedia.org/wiki/Gaussova_elimina%C4%8Dn%C3%AD_metoda#Pseudok%C3%B3d
func (m MatrixNumber[T]) GaussElimination() int {
	// pivot row index
	h := 0

	// pivot column index
	k := 0

	rowSwapsCount := 0

	for k < m.Height && h < m.Width {
		// find k-th pivot
		iMax := -1
		v := T(0)
		for i := h; i < m.Height; i++ {
			candidate := utils.Abs(m.Get(k, i))
			if candidate >= v {
				iMax = i
				v = candidate
			}
		}

		if m.Get(k, iMax) == 0 {
			// no pivot in given column, move to next
			k++
			continue
		}

		// swap rows h and i_max
		if h != iMax {
			for i := 0; i < m.Width; i++ {
				m.Columns[i][iMax], m.Columns[i][h] = m.Columns[i][h], m.Columns[i][iMax]
			}
			rowSwapsCount++
		}
		// do for all rows under the pivot
		for i := h + 1; i < m.Height; i++ {
			f := m.Get(k, i) / m.Get(k, h)
			// fill the column part under the pivot with zeros
			m.Set(k, i, 0)

			// do for the rest of the remaining items in current row
			for j := k + 1; j < m.Width; j++ {
				m.Set(j, i, m.Get(j, i)-m.Get(j, h)*f)
			}
		}

		// move to the next pivot
		h++
		k++
	}

	return rowSwapsCount
}

// JordanContinue
// see https://cs.wikipedia.org/wiki/Gaussova_elimina%C4%8Dn%C3%AD_metoda#Pseudok%C3%B3d
func (m MatrixNumber[T]) JordanContinue() bool {
	// pivot row index
	h := utils.Min(m.Width, m.Height) - 1

	// pivot column index
	k := h

	for k >= 0 && h >= 0 {
		v := m.Get(k, h)
		if v == 0 {
			return false
		}

		if v != 1 {
			// divide whole row with pivot value
			for i := k; i < m.Width; i++ {
				m.Set(i, h, m.Get(i, h)/v)
			}
		}

		// update all items above
		for i := h - 1; i >= 0; i-- {
			p2 := m.Get(k, i)

			for j := k; j < m.Width; j++ {
				m.Set(j, i, m.Get(j, i)-p2*m.Get(j, h))
			}
		}
		h--
		k--
	}

	return true
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
	// no need to clone, new array is created in constructor during next calls
	return NewMatrixColumnNotation(m.Columns)
}

func EqualFunc[T any](m1, m2 Matrix[T], eq func(a, b T) bool) bool {
	return slices2.EqualFunc(m1.Columns, m2.Columns, func(col1, col2 []T) bool {
		return slices2.EqualFunc(col1, col2, eq)
	})
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
