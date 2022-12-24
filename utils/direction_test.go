package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDirection8_Rotate(t *testing.T) {
	type args struct {
		steps int
	}
	tests := []struct {
		name string
		d    Direction8
		args args
		want Direction8
	}{
		{"", North, args{0}, North},
		{"", North, args{8}, North},
		{"", NorthWest, args{1}, North},
		{"", North, args{1}, NorthEast},

		{"", North, args{-8}, North},
		{"", North, args{-1}, NorthWest},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.d.Rotate(tt.args.steps), "Rotate(%v)", tt.args.steps)
		})
	}
}
