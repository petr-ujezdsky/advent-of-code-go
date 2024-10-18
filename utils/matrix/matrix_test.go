package matrix

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewMatrix2_data_locality(t *testing.T) {
	matrix2 := NewMatrixInt(2, 3)

	i := 0
	for x := 0; x < 2; x++ {
		for y := 0; y < 3; y++ {
			matrix2.Columns[x][y] = i
			i++
		}
	}

	// ensured data locality
	assert.Equal(t, "[[0 1 2] [3 4 5]]", fmt.Sprint(matrix2.Columns))
}

func TestMatrix_Rotate90CounterClockwise(t *testing.T) {
	m0 := NewMatrixRowNotation[int]([][]int{
		{0, 2},
		{1, 3},
	})

	m1 := NewMatrixRowNotation[int]([][]int{
		{2, 3},
		{0, 1},
	})

	m2 := NewMatrixRowNotation[int]([][]int{
		{3, 1},
		{2, 0},
	})

	m3 := NewMatrixRowNotation[int]([][]int{
		{1, 0},
		{3, 2},
	})

	type args struct {
		steps int
	}
	type testCase[T any] struct {
		name  string
		m     Matrix[T]
		args  args
		wantM Matrix[T]
	}
	tests := []testCase[int]{
		{"", m0, args{0}, m0},
		{"", m0, args{1}, m1},
		{"", m0, args{2}, m2},
		{"", m0, args{3}, m3},

		{"", m0, args{4}, m0},

		{"", m0, args{-1}, m3},
		{"", m0, args{-2}, m2},
		{"", m0, args{-3}, m1},
		{"", m0, args{-4}, m0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.wantM, tt.m.Rotate90CounterClockwise(tt.args.steps), "Rotate90CounterClockwise(%v)", tt.args.steps)
		})
	}
}

func TestMatrix_FlipHorizontal(t *testing.T) {
	m := NewMatrixRowNotation[int]([][]int{
		{0, 2},
		{1, 3},
	})

	expected := NewMatrixRowNotation[int]([][]int{
		{2, 0},
		{3, 1},
	})

	assert.Equal(t, expected, m.FlipHorizontal())
}

func TestMatrix_FlipVertical(t *testing.T) {
	m := NewMatrixRowNotation[int]([][]int{
		{0, 2},
		{1, 3},
	})

	expected := NewMatrixRowNotation[int]([][]int{
		{1, 3},
		{0, 2},
	})

	assert.Equal(t, expected, m.FlipVertical())
}

func TestMatrixNumber_Determinant(t *testing.T) {
	type testCase[T float64] struct {
		name string
		m    MatrixNumber[T]
		want T
	}
	tests := []testCase[float64]{
		{"Rank 1", NewMatrixColumnNotationFloat([][]float64{{3}}), 1},
		{"Rank 2", NewMatrixColumnNotationFloat([][]float64{{3, 1}, {7, -4}}), -19},
		{"Rank 3", NewMatrixColumnNotationFloat([][]float64{{0, 3, 1}, {1, 2, 1}, {2, 1, 0}}), 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.m.Determinant(), "Determinant()")
		})
	}
}

func TestMatrixNumber_GaussElimination(t *testing.T) {
	actual := NewMatrixNumberRowNotation([][]float64{
		{2, 1, -1, 8},
		{-3, -1, 2, -11},
		{-2, 1, 2, -3},
	})

	actual.GaussElimination()

	expected := NewMatrixNumberRowNotation([][]float64{
		{-3, -1, 2, -11},
		{0, 1.6666666666666665, 0.6666666666666667, 4.333333333333333},
		{0, 0, 0.19999999999999987, -0.1999999999999997},
	})

	fmt.Println(actual.StringFmt(FmtFmt[float64]("%9.6f")))

	assert.Equal(t, expected, actual)
}

func TestMatrixNumber_JordanContinue(t *testing.T) {
	actual := NewMatrixNumberRowNotation([][]float64{
		{2, 1, -1, 8},
		{0, 0.5, 0.5, 1},
		{0, 0, -1, 1},
	})

	actual.JordanContinue()

	expected := NewMatrixNumberRowNotation([][]float64{
		{1, 0, 0, 2},
		{0, 1, 0, 3},
		{0, 0, 1, -1},
	})

	fmt.Println(actual.StringFmt(FmtFmt[float64]("%9.6f")))

	assert.Equal(t, expected, actual)
}

func TestMatrixNumber_Inverse_1(t *testing.T) {
	actual := NewMatrixNumberRowNotation([][]float64{
		{1, 2},
		{2, 3},
	})

	ok := actual.Inverse()
	assert.True(t, ok)

	expected := NewMatrixNumberRowNotation([][]float64{
		{-3, 2},
		{2, -1},
	})

	fmt.Println(actual.StringFmt(FmtFmt[float64]("%9.6f")))

	assert.Equal(t, expected, actual)
}

func TestMatrixNumber_Inverse_2(t *testing.T) {
	actual := NewMatrixNumberRowNotation([][]float64{
		{1, 2, 0},
		{2, 4, 1},
		{2, 1, 0},
	})

	ok := actual.Inverse()
	assert.True(t, ok)

	expected := NewMatrixNumberRowNotation([][]float64{
		{-0.3333333333333333, 0, 0.6666666666666666},
		{0.6666666666666666, 0, -0.3333333333333333},
		{-2, 1, 0},
	})

	fmt.Println(actual.StringFmt(FmtFmt[float64]("%9.6f")))

	assert.Equal(t, expected, actual)
}
