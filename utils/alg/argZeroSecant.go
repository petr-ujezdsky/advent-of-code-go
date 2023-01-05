package alg

import "github.com/petr-ujezdsky/advent-of-code-go/utils"

// ArgZeroSecant finds function input x for which the f(x) = 0, using two starting points.
// see https://en.wikipedia.org/wiki/Secant_method
func ArgZeroSecant(x0, x1, yEpsilon float64, f func(float64) float64) (float64, float64) {
	for {
		x2 := x1 - f(x1)*(x1-x0)/(f(x1)-f(x0))
		x0, x1 = x1, x2

		y := f(x1)
		if utils.Abs(y) < yEpsilon {
			return x1, y
		}
	}
}
