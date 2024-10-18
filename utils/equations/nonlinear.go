package equations

import (
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/matrix"
)

type Fn = func(x utils.VectorNf) utils.VectorNf
type Jfn = func(x utils.VectorNf) matrix.MatrixFloat

// SolveNonLinearEquations solves nonlinear equations in form f(x) = 0
func SolveNonLinearEquations(F Fn, J Jfn, x0 utils.VectorNf) (utils.VectorNf, bool) {
	xi := x0
	i := 0
	iMax := 1000

	for {
		// fail on too much iterations
		if i > iMax {
			return xi, false
		}

		// compute F(xi)
		Fxi := F(xi)

		// convergence check
		if isNearZero(Fxi) {
			return xi, true
		}

		// compute J(xi)
		Jxi := J(xi)

		// compute J(xi)^-1
		ok := Jxi.Inverse()
		if !ok {
			return xi, false
		}

		delta := Jxi.MultiplyV(Fxi).Multiply(-1)

		xi = xi.Add(delta)
		i++
	}
}

func isNearZero(x utils.VectorNf) bool {
	for _, v := range x.Items {
		if utils.Abs(v) > 0.001 {
			return false
		}
	}

	return true
}
