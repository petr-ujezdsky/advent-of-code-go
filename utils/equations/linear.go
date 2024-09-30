package equations

import (
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/matrix"
	"math"
)

// SolveLinearEquations solves linear equations in form Ax = b
func SolveLinearEquations(a matrix.MatrixFloat, b utils.VectorNf) (utils.VectorNf, bool) {
	detA := a.Determinant()

	if math.Abs(detA) < 0.001 {
		return utils.VectorNf{}, false
	}

	x := utils.VectorNf{Items: make([]float64, len(b.Items))}

	for i, _ := range x.Items {
		Ai := swapColumn(i, a, b)
		detAi := Ai.Determinant()

		x.Items[i] = detAi / detA
	}

	return x, true
}

func swapColumn(index int, a matrix.MatrixFloat, b utils.VectorNf) matrix.MatrixFloat {
	swapped := matrix.NewMatrixFloat(a.Width, a.Height)

	for x, column := range swapped.Columns {
		for y := range column {
			if x == index {
				swapped.Columns[x][y] = b.Items[y]
			} else {
				swapped.Columns[x][y] = a.Columns[x][y]
			}
		}
	}

	return swapped
}
