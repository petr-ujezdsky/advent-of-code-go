package equations

import (
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/matrix"
	"reflect"
	"testing"
)

func TestSolveLinearEquations2(t *testing.T) {
	type args struct {
		a matrix.MatrixFloat
		b utils.VectorNf
	}
	tests := []struct {
		name  string
		args  args
		want  utils.VectorNf
		want1 bool
	}{
		{"", args{matrix.NewMatrixColumnNotationFloat([][]float64{{1, 4}, {2, 5}}), utils.VectorNf{Items: []float64{3, 6}}}, utils.VectorNf{Items: []float64{-1, 2}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := SolveLinearEquations(tt.args.a, tt.args.b)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SolveLinearEquations() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("SolveLinearEquations() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
