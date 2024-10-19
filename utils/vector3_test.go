package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestVector3n_Cross(t *testing.T) {
	type args[T Number] struct {
		v2 Vector3n[T]
	}
	type testCase[T Number] struct {
		name string
		v1   Vector3n[T]
		args args[T]
		want Vector3n[T]
	}
	tests := []testCase[int]{
		{"", Vector3i{X: 1}, args[int]{v2: Vector3i{Y: 1}}, Vector3i{Z: 1}},
		{"", Vector3i{Y: 1}, args[int]{v2: Vector3i{X: 1}}, Vector3i{Z: -1}},
		{"", Vector3i{Z: 1}, args[int]{v2: Vector3i{X: 1}}, Vector3i{Y: 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.v1.Cross(tt.args.v2), "Cross(%v)", tt.args.v2)
		})
	}
}

func TestVector3n_OrthogonalBase(t *testing.T) {
	type testCase[T Number] struct {
		name  string
		v1    Vector3n[T]
		want  Vector3n[T]
		want1 Vector3n[T]
	}
	tests := []testCase[int]{
		{"", Vector3i{Z: 1}, Vector3i{X: 1, Y: 1}, Vector3i{X: 1, Y: -1}},
		{"", Vector3i{X: -3, Y: 1, Z: 2}, Vector3i{X: -2, Y: 2, Z: -4}, Vector3i{X: 8, Y: 16, Z: 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.v1.OrthogonalBase()
			assert.Equalf(t, tt.want, got, "OrthogonalBase()")
			assert.Equalf(t, tt.want1, got1, "OrthogonalBase()")
		})
	}
}
