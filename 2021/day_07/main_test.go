package day_07

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_01_example_parse(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	positions, err := ParseInput(reader)
	assert.Nil(t, err)

	assert.Equal(t, []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}, positions)
}

func Test_01_example_costs(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	positions, err := ParseInput(reader)
	assert.Nil(t, err)

	costs := CalculateFuelCosts(positions, CostSimple)

	fmt.Println(costs)

	assert.Equal(t, 37, costs[1])
	assert.Equal(t, 39, costs[2])
	assert.Equal(t, 71, costs[9])
}

func Test_01_example_costs_min(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	positions, err := ParseInput(reader)
	assert.Nil(t, err)

	position, cost := LowestAlignment(positions, CostSimple)

	assert.Equal(t, 2, position)
	assert.Equal(t, 37, cost)
}

func Test_01(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	positions, err := ParseInput(reader)
	assert.Nil(t, err)

	position, cost := LowestAlignment(positions, CostSimple)

	assert.Equal(t, 343, position)
	assert.Equal(t, 353800, cost)
}

func Test_02_example_costs_min(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	positions, err := ParseInput(reader)
	assert.Nil(t, err)

	position, cost := LowestAlignment(positions, CostSteppingUp)

	assert.Equal(t, 5, position)
	assert.Equal(t, 168, cost)
}

func Test_02(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	positions, err := ParseInput(reader)
	assert.Nil(t, err)

	position, cost := LowestAlignment(positions, CostSteppingUp)

	assert.Equal(t, 480, position)
	assert.Equal(t, 98119739, cost)
}
