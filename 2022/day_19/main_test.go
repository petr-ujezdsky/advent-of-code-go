package main

import (
	"github.com/stretchr/testify/assert"
	"os"
	"strings"
	"testing"
)

func Test_01_parse(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	world := ParseInput(reader)

	assert.Equal(t, 2, len(world.Blueprints))
	assert.Equal(t, 8, world.Blueprints[1].RobotsCosts[Obsidian][Clay])
}

func Test_01_example_blueprint_1(t *testing.T) {
	//reader := strings.NewReader("Blueprint 1: Each ore robot costs 5000 ore. Each clay robot costs 5000 ore. Each obsidian robot costs 5000 ore and 5000 clay. Each geode robot costs 0 ore and 0 obsidian.")
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	world := ParseInput(reader)

	// 3s
	geodes, state := maxGeodeCountInTime(world.Blueprints[0], 24)
	printState(&state)
	assert.Equal(t, 9, geodes)
}

func Test_01_example_blueprint_2(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	world := ParseInput(reader)

	// 2s
	geodes, state := maxGeodeCountInTime(world.Blueprints[1], 24)
	printState(&state)
	assert.Equal(t, 12, geodes)
}

func Test_01_example(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	world := ParseInput(reader)

	result := DoWithInputParallel(world)
	assert.Equal(t, 33, result)
}

func Test_01(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	world := ParseInput(reader)

	// very slow, about 4h
	result := DoWithInputParallel(world)
	assert.Equal(t, 1192, result)
}

func Test_02_example_blueprint_1(t *testing.T) {
	//reader := strings.NewReader("Blueprint 1: " +
	//	"Each ore robot costs 5000 ore. " +
	//	"Each clay robot costs 5000 ore. " +
	//	"Each obsidian robot costs 5000 ore and 5000 clay. " +
	//	"Each geode robot costs 0 ore and 0 obsidian.")
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	world := ParseInput(reader)

	// 32s
	geodes, state := maxGeodeCountInTime(world.Blueprints[0], 32)
	printState(&state)
	assert.Equal(t, 56, geodes)
}

func Test_02_example_blueprint_2(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	world := ParseInput(reader)

	// 18s
	geodes, state := maxGeodeCountInTime(world.Blueprints[1], 32)
	printState(&state)
	assert.Equal(t, 62, geodes)
}

func Test_02_example(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	world := ParseInput(reader)

	result := DoWithInputParallelFirstN(world, 2)
	assert.Equal(t, 3472, result)
}

func Test_02(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	world := ParseInput(reader)

	result := DoWithInputParallelFirstN(world, 3)
	assert.Equal(t, 0, result)
}
