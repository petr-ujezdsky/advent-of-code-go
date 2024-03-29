package day_05

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_01_example_parse(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	lines, err := ParseInput(reader)
	assert.Nil(t, err)

	assert.Equal(t, 10, len(lines))

	line := lines[0]
	assert.Equal(t, NewLine(0, 9, 5, 9), line)

	line = lines[9]
	assert.Equal(t, NewLine(5, 5, 8, 2), line)
}

func Test_01_example_intersections(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	lines, err := ParseInput(reader)
	assert.Nil(t, err)

	intersectionsCount, area := CountIntersections(lines, false)

	area.Print()

	assert.Equal(t, 5, intersectionsCount)
}

func Test_01_intersections(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	lines, err := ParseInput(reader)
	assert.Nil(t, err)

	intersectionsCount, _ := CountIntersections(lines, false)

	assert.Equal(t, 4745, intersectionsCount)
}

func Test_02_example_intersections(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	lines, err := ParseInput(reader)
	assert.Nil(t, err)

	intersectionsCount, area := CountIntersections(lines, true)

	area.Print()

	assert.Equal(t, 12, intersectionsCount)
}

func Test_02_intersections(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	lines, err := ParseInput(reader)
	assert.Nil(t, err)

	intersectionsCount, _ := CountIntersections(lines, true)

	assert.Equal(t, 18442, intersectionsCount)
}
