package combi

import (
	"fmt"
)

var factorial = []int{
	1,
	1,
	2,
	6,
	24,
	120,
	720,
	5040,
	40320,
	362880,
	3628800,
	39916800,
	479001600,
	6227020800,
	87178291200,
	1307674368000,
	20922789888000,
	355687428096000,
	6402373705728000,
	121645100408832000,
	2432902008176640000, //20!
}

func computeFactorial(n int) int {
	if n < 0 {
		panic(fmt.Sprintf("Wrong input n=%v", n))
	}

	if n > 20 {
		panic("Out of 64bit range")
	}

	for i := len(factorial) - 1; i <= n; i++ {
		factorial = append(factorial, factorial[i]*(i+1))
	}

	return factorial[n]
}

func CombinationsWithoutRepetition(n, k int) int {
	if n < 1 || k < 1 {
		panic(fmt.Sprintf("Wrong input n=%v, k=%v", n, k))
	}

	return factorial[n] / factorial[k] / factorial[n-k]
}

func CombinationsWithRepetition(n, k int) int {
	if n < 1 || k < 1 {
		panic(fmt.Sprintf("Wrong input n=%v, k=%v", n, k))
	}

	return factorial[n+k-1] / factorial[k] / factorial[n-1]
}
