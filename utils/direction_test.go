package utils

import (
	"fmt"
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

func TestGenerateSteps3D(t *testing.T) {
	steps := []int{0, 1, -1}
	for _, x := range steps {
		for _, y := range steps {
			for _, z := range steps {
				if x == 0 && y == 0 && z == 0 {
					continue
				}
				fmt.Printf("{X: %v, Y: %v, Z: %v},\n", x, y, z)
			}
		}
	}
}

func TestGenerateSteps4D(t *testing.T) {
	steps := []int{0, 1, -1}
	for _, x := range steps {
		for _, y := range steps {
			for _, z := range steps {
				for _, w := range steps {
					if x == 0 && y == 0 && z == 0 && w == 0 {
						continue
					}
					fmt.Printf("{X: %v, Y: %v, Z: %v, W: %v},\n", x, y, z, w)
				}
			}
		}
	}
}
