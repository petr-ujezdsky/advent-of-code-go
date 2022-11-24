package utils

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewMatrix3_data_locality(t *testing.T) {
	matrix3 := NewMatrix3[int](1, 3, 2)

	i := 0
	for x := 0; x < 1; x++ {
		for y := 0; y < 3; y++ {
			for z := 0; z < 2; z++ {
				matrix3.Cells[x][y][z] = i
				i++
			}
		}
	}

	// ensured data locality
	assert.Equal(t, "[[[0 1] [2 3] [4 5]]]", fmt.Sprint(matrix3.Cells))
}
