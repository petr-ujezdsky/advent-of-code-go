package utils

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewMatrix2_data_locality(t *testing.T) {
	matrix2 := NewMatrixInt(2, 3)

	i := 0
	for x := 0; x < 2; x++ {
		for y := 0; y < 3; y++ {
			matrix2.Columns[x][y] = i
			i++
		}
	}

	// ensured data locality
	assert.Equal(t, "[[0 1 2] [3 4 5]]", fmt.Sprint(matrix2.Columns))
}
