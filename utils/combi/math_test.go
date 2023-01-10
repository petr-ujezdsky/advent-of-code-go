package combi

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_printFactorial(t *testing.T) {
	for i := 0; i < 100; i++ {
		//fmt.Printf("%3v: %v\n", i, factorial(i))
		fmt.Printf("%v,\n", computeFactorial(i))
	}
}

func Test_factorial(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{0}, 1},
		{"", args{1}, 1},
		{"", args{2}, 2},
		{"", args{3}, 6},
		{"", args{4}, 24},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, computeFactorial(tt.args.n), "factorial(%v)", tt.args.n)
		})
	}
}

func TestCombinationsWithoutRepetition(t *testing.T) {
	type args struct {
		n int
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{3, 1}, 3},
		{"", args{3, 2}, 3},
		{"", args{3, 3}, 1},
		{"", args{20, 20}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, CombinationsWithoutRepetition(tt.args.n, tt.args.k), "CombinationsWithoutRepetition(%v, %v)", tt.args.n, tt.args.k)
		})
	}
}

func TestCombinationsWithRepetition(t *testing.T) {
	type args struct {
		n int
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{2, 1}, 2},
		{"", args{2, 2}, 3},
		{"", args{2, 3}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, CombinationsWithRepetition(tt.args.n, tt.args.k), "CombinationsWithRepetition(%v, %v)", tt.args.n, tt.args.k)
		})
	}
}
