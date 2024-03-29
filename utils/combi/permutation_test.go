package combi_test

import (
	"github.com/petr-ujezdsky/advent-of-code-go/utils/combi"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Permute_2(t *testing.T) {
	quit := make(chan interface{})
	values := []int{1, 2}
	output := combi.Permute(quit, values)

	assert.Equal(t, []int{1, 2}, <-output)
	assert.Equal(t, []int{2, 1}, <-output)
}

func Test_Permute_3(t *testing.T) {
	quit := make(chan interface{})
	values := []int{1, 2, 3}
	output := combi.Permute(quit, values)

	assert.Equal(t, []int{1, 2, 3}, <-output)
	assert.Equal(t, []int{1, 3, 2}, <-output)
	assert.Equal(t, []int{2, 1, 3}, <-output)
	assert.Equal(t, []int{2, 3, 1}, <-output)
	assert.Equal(t, []int{3, 2, 1}, <-output)
	assert.Equal(t, []int{3, 1, 2}, <-output)

	_, ok := <-output
	assert.False(t, ok)
}

func Test_Permute_quit(t *testing.T) {
	quit := make(chan interface{})
	values := []int{1, 2, 3}
	output := combi.Permute(quit, values)

	assert.Equal(t, []int{1, 2, 3}, <-output)
	assert.Equal(t, []int{1, 3, 2}, <-output)
	assert.Equal(t, []int{2, 1, 3}, <-output)

	// stop the permutation
	close(quit)

	// next read is nil
	_, ok := <-output
	assert.False(t, ok)
}
