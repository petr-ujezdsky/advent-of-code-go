package day_12

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_01_example_parse(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	world, err := ParseInput(reader)
	assert.Nil(t, err)

	assert.Equal(t, "NNCB", world.template)
	assert.Equal(t, 16, len(world.rules))
	assert.Equal(t, "B", world.rules["CH"])
	assert.Equal(t, "C", world.rules["CN"])
}
