package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLineOrthogonal2i_Intersection(t *testing.T) {
	type fields struct {
		A Vector2i
		B Vector2i
	}
	type args struct {
		line2 LineOrthogonal2i
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   LineOrthogonal2i
		want1  bool
	}{
		{
			"Horizontal vs Vertical",
			fields{
				A: Vector2i{0, 3},
				B: Vector2i{5, 3}},
			args{NewLineOrthogonal2i(Vector2i{4, 0}, Vector2i{4, 5})},
			NewLineOrthogonal2i(Vector2i{4, 3}, Vector2i{4, 3}),
			true,
		},
		{
			"Horizontal vs Vertical",
			fields{
				A: Vector2i{8, 5},
				B: Vector2i{3, 5}},
			args{NewLineOrthogonal2i(Vector2i{6, 7}, Vector2i{6, 3})},
			NewLineOrthogonal2i(Vector2i{6, 5}, Vector2i{6, 5}),
			true,
		},
		{
			"Horizontal vs Horizontal",
			fields{
				A: Vector2i{0, 3},
				B: Vector2i{5, 3}},
			args{
				NewLineOrthogonal2i(Vector2i{2, 3}, Vector2i{7, 3})},
			NewLineOrthogonal2i(Vector2i{2, 3}, Vector2i{5, 3}),
			true,
		},
		{
			"None",
			fields{
				A: Vector2i{0, 3},
				B: Vector2i{5, 3}},
			args{
				NewLineOrthogonal2i(Vector2i{0, 4}, Vector2i{5, 4})},
			LineOrthogonal2i{},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			line := NewLineOrthogonal2i(tt.fields.A, tt.fields.B)
			got, got1 := line.Intersection(tt.args.line2)
			assert.Equalf(t, tt.want, got, "Intersection(%v)", tt.args.line2)
			assert.Equalf(t, tt.want1, got1, "Intersection(%v)", tt.args.line2)
		})
	}
}
