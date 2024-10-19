package equations

import (
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/matrix"
)

type Fn = func(x utils.VectorNf) utils.VectorNf
type Jfn = func(x utils.VectorNf) matrix.MatrixFloat

func Xy[T utils.Number, R any](f func(x, y T) R) func(nn utils.VectorNn[T]) R {
	return func(v utils.VectorNn[T]) R {
		return f(v.Items[0], v.Items[1])
	}
}

// SolveNonLinearEquations solves nonlinear equations in form f(x) = 0, stops when f(x) is below 0.001
func SolveNonLinearEquations(F Fn, J Jfn, x0 utils.VectorNf) (utils.VectorNf, bool) {
	return SolveNonLinearEquationsThreshold(F, J, x0, 0.001, 1000)
}

// SolveNonLinearEquationsThreshold solves nonlinear equations in form f(x) = 0, stops when f(x) is below threshold
func SolveNonLinearEquationsThreshold(F Fn, J Jfn, x0 utils.VectorNf, threshold float64, iterationsMax int) (utils.VectorNf, bool) {
	xi := x0
	i := 0

	for {
		// fail on too much iterations
		if i > iterationsMax {
			return xi, false
		}

		// compute F(xi)
		Fxi := F(xi)

		// convergence check
		if isBelowThreshold(Fxi, threshold) {
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

func isBelowThreshold(x utils.VectorNf, threshold float64) bool {
	for _, v := range x.Items {
		if utils.Abs(v) > threshold {
			return false
		}
	}

	return true
}
