package alg

// ArgZeroSecant finds function input x for which the f(x) = 0, using two starting points.
// see https://en.wikipedia.org/wiki/Secant_method
func ArgZeroSecant(x0, x1 int, f func(i int) int) (int, int) {
	for {
		x2 := x1 - f(x1)*(x1-x0)/(f(x1)-f(x0))
		//x2 := x1 - int(float64(f(x1)*(x1-x0))/float64(f(x1)-f(x0)))
		x0, x1 = x1, x2

		if x0-x1 == 0 {
			return x1, f(x1)
		}
	}
}
