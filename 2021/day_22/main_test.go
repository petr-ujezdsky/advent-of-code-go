package day_22

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_regexCube(t *testing.T) {
	matches := regexCube.FindStringSubmatch("on x=10..12,y=10..12,z=10..12")
	assert.Equal(t, []string{"on x=10..12,y=10..12,z=10..12", "on", "10", "12", "10", "12", "10", "12"}, matches)
}

func Test_01_example_1(t *testing.T) {
	reader, err := os.Open("data-00-example-1.txt")
	assert.Nil(t, err)

	cubes := ParseInput(reader)
	assert.Equal(t, 4, len(cubes))
}
