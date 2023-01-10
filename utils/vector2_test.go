package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_vector2n_Rotate90CounterClockwisePivot(t *testing.T) {
	type args[T Number] struct {
		steps int
		pivot vector2n[T]
	}
	type testCase[T Number] struct {
		name string
		v1   vector2n[T]
		args args[T]
		want vector2n[T]
	}
	tests := []testCase[int]{
		{"", vector2n[int]{X: 0, Y: 0}, args[int]{0, vector2n[int]{X: 10, Y: 20}}, vector2n[int]{X: 0, Y: 0}},

		{"", vector2n[int]{X: 0, Y: 0}, args[int]{1, vector2n[int]{X: 10, Y: 20}}, vector2n[int]{X: 30, Y: 10}},
		{"", vector2n[int]{X: 0, Y: 0}, args[int]{2, vector2n[int]{X: 10, Y: 20}}, vector2n[int]{X: 20, Y: 40}},
		{"", vector2n[int]{X: 0, Y: 0}, args[int]{3, vector2n[int]{X: 10, Y: 20}}, vector2n[int]{X: -10, Y: 30}},
		{"", vector2n[int]{X: 0, Y: 0}, args[int]{4, vector2n[int]{X: 10, Y: 20}}, vector2n[int]{X: 0, Y: 0}},

		{"", vector2n[int]{X: 0, Y: 0}, args[int]{-1, vector2n[int]{X: 10, Y: 20}}, vector2n[int]{X: -10, Y: 30}},
		{"", vector2n[int]{X: 0, Y: 0}, args[int]{-2, vector2n[int]{X: 10, Y: 20}}, vector2n[int]{X: 20, Y: 40}},
		{"", vector2n[int]{X: 0, Y: 0}, args[int]{-3, vector2n[int]{X: 10, Y: 20}}, vector2n[int]{X: 30, Y: 10}},
		{"", vector2n[int]{X: 0, Y: 0}, args[int]{-4, vector2n[int]{X: 10, Y: 20}}, vector2n[int]{X: 0, Y: 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.v1.Rotate90CounterClockwisePivot(tt.args.steps, tt.args.pivot), "Rotate90CounterClockwisePivot(%v, %v)", tt.args.steps, tt.args.pivot)
		})
	}
}
