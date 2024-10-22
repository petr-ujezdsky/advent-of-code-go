package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func Test_01_parse(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	world := ParseInput(reader)

	assert.Equal(t, 15, len(world.Components))
	jqt := world.Components["jqt"]
	assert.Equal(t, "jqt", jqt.Name)
	assert.Equal(t, 4, len(jqt.Neighbours))
	assert.Equal(t, "rhn", jqt.Neighbours[0].Name)
	assert.Equal(t, "xhk", jqt.Neighbours[1].Name)
	assert.Equal(t, "nvd", jqt.Neighbours[2].Name)
	assert.Equal(t, "ntq", jqt.Neighbours[3].Name)
}

func Test_01_print(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	world := ParseInput(reader)

	for _, component := range world.Components {
		fmt.Println(component.String())
	}
}

// see https://csacademy.com/app/graph_editor/
func Test_01_print_csacademy(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	//reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	world := ParseInput(reader)

	for _, component := range world.Components {
		for _, neighbour := range component.Neighbours {
			fmt.Printf("%s %s\n", component.Name, neighbour.Name)
		}
	}
}

func Test_01_example(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	world := ParseInput(reader)

	result := DoWithInputPart01(world)
	assert.Equal(t, 0, result)
}

func Test_01(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	world := ParseInput(reader)

	result := DoWithInputPart01(world)
	assert.Equal(t, 0, result)
}

func Test_02_example(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	world := ParseInput(reader)

	result := DoWithInputPart02(world)
	assert.Equal(t, 0, result)
}

func Test_02(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	world := ParseInput(reader)

	result := DoWithInputPart02(world)
	assert.Equal(t, 0, result)
}
