package equations

import (
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/matrix"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestSolveNonLinearEquations_1(t *testing.T) {
	F := Xy(func(x, y float64) utils.VectorNf {
		return utils.VectorNf{Items: []float64{
			// x^2 + y^2 - 4
			x*x + y*y - 4,

			// ln(2-x) - y
			math.Log(2-x) - y,
		}}
	})

	J := Xy(func(x, y float64) matrix.MatrixFloat {
		return matrix.NewMatrixNumberRowNotation([][]float64{
			{2 * x, 2 * y},

			{-1 / (2 - x), -1},
		})
	})

	x0 := utils.VectorNf{Items: []float64{1, 2}}

	solution, ok := SolveNonLinearEquations(F, J, x0)
	assert.True(t, ok)

	expected := utils.VectorNf{Items: []float64{-1.5479975789882978, 1.2663842449509168}}
	assert.Equal(t, expected, solution)
}

func TestSolveNonLinearEquations_2(t *testing.T) {
	F := Xy(func(x, y float64) utils.VectorNf {
		return utils.VectorNf{Items: []float64{
			x*x - 4*x + y*y,

			y - math.Exp(-x) - 1,
		}}
	})

	J := Xy(func(x, y float64) matrix.MatrixFloat {
		return matrix.NewMatrixNumberRowNotation([][]float64{
			{2*x - 4, 2 * y},

			{math.Exp(-x), 1},
		})
	})

	x0 := utils.VectorNf{Items: []float64{0, -2}}

	solution, ok := SolveNonLinearEquations(F, J, x0)
	assert.False(t, ok)

	x0 = utils.VectorNf{Items: []float64{4, 1}}

	solution, ok = SolveNonLinearEquations(F, J, x0)
	assert.True(t, ok)

	expected := utils.VectorNf{Items: []float64{3.717927396241469, 1.024279204292909}}
	assert.Equal(t, expected, solution)
}
