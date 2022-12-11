package utils

import (
	"fmt"
	"strings"
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

type MatrixNumber[T Number] struct {
	Matrix[T]
}

// number aliases

type MatrixInt = MatrixNumber[int]
type MatrixFloat = MatrixNumber[float64]

// constructor aliases

var NewMatrixInt = NewMatrixNumber[int]

func NewMatrixNumber[T Number](width, height int) MatrixNumber[T] {
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

// NewMatrixNumberRowNotation converts matrix from row-first notation to column-first notation
func NewMatrixNumberRowNotation[T Number](rows [][]T) MatrixNumber[T] {
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

func (m Matrix[T]) SetSafe(x, y int, value T) bool {
	if x < 0 || x >= m.Width || y < 0 || y >= m.Height {
		return false
	}

	m.Columns[x][y] = value
	return true
}

func (m Matrix[T]) SetV(pos Vector2i, value T) {
	m.Columns[pos.X][pos.Y] = value
}

func (m Matrix[T]) SetVSafe(pos Vector2i, value T) bool {
	return m.SetSafe(pos.X, pos.Y, value)
}

func (m Matrix[T]) GetSafe(x, y int) (T, bool) {
	if x < 0 || x >= m.Width || y < 0 || y >= m.Height {
		var nothing T
		return nothing, false
	}

	return m.Columns[x][y], true
}

func (m Matrix[T]) GetV(pos Vector2i) T {
	return m.Columns[pos.X][pos.Y]
}

func (m Matrix[T]) GetVSafe(pos Vector2i) (T, bool) {
	return m.GetSafe(pos.X, pos.Y)
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

func (m MatrixNumber[T]) ArgMax() (Vector2i, T) {
	max, xmax, ymax := m.Columns[0][0], 0, 0

	for x, col := range m.Columns {
		for y, val := range col {
			if val > max {
				max, xmax, ymax = val, x, y
			}
		}
	}

	return Vector2i{xmax, ymax}, max
}

func (m Matrix[T]) String() string {
	return m.StringFmt(FmtNative[T])
}

type ValueFormatter[T any] func(value T) string

func (m Matrix[T]) StringFmt(formatter ValueFormatter[T]) string {
	var sb strings.Builder

	for y := 0; y < m.Height; y++ {
		for x := 0; x < m.Width; x++ {
			val := m.Columns[x][y]

			if x > 0 {
				sb.WriteString(" ")
			}
			sb.WriteString(formatter(val))
		}
		if y < m.Height-1 {
			sb.WriteString("\n")
		}
	}

	return sb.String()
}

func FmtNative[T any](value T) string {
	return fmt.Sprint(value)
}

func FmtFmt[T any](format string) func(v T) string {
	return func(val T) string {
		return fmt.Sprintf(format, val)
	}
}
func FmtConstant[T any](value string) func(v T) string {
	return func(val T) string {
		return value
	}
}

func FmtBoolean[T comparable](val T) string {
	return FmtBooleanConst[T](".", "#")(val)
}

func FmtBooleanConst[T comparable](falseVal, trueVal string) ValueFormatter[T] {
	return FmtBooleanCustom[T](FmtConstant[T](falseVal), FmtConstant[T](trueVal))
}

func FmtBooleanCustom[T comparable](formatterFalse, formatterTrue ValueFormatter[T]) func(v T) string {
	return func(val T) string {
		var empty T
		if val == empty {
			return formatterFalse(empty)
		} else {
			return formatterTrue(val)
		}
	}
}
