package utils_test

import (
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Permute_2(t *testing.T) {
	values := []int{1, 2}
	output := utils.Permute(values)

	assert.Equal(t, []int{1, 2}, <-output)
	assert.Equal(t, []int{2, 1}, <-output)
}

func Test_Permute_3(t *testing.T) {
	values := []int{1, 2, 3}
	output := utils.Permute(values)

	assert.Equal(t, []int{1, 2, 3}, <-output)
	assert.Equal(t, []int{1, 3, 2}, <-output)
	assert.Equal(t, []int{2, 1, 3}, <-output)
	assert.Equal(t, []int{2, 3, 1}, <-output)
	assert.Equal(t, []int{3, 2, 1}, <-output)
	assert.Equal(t, []int{3, 1, 2}, <-output)

	_, ok := <-output
	assert.False(t, ok)
}
