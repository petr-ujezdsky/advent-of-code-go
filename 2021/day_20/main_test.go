package day_20

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_01_example(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	image, err := NewImage(reader)
	assert.Nil(t, err)

	expected := `..#.#..#####.#.#.#.###.##.....###.##.#..###.####..#####..#....#..#..##..###..######.###...####..#..#####..##..#.#####...##.#.#..#.##..#.#......#.###.######.###.####...#.##.##..#..#..#####.....#.#....###..#.##......#.....#..#..#..##..#...##.######.####.####.#.#...#.......#..#.#.#...####.##.#......#..#...##.#.##..#...##.#.##..###.#......#.#.......#.#.#.####.###.##...#.....####.#..#..#.##.#....##..#.####....##...##..#...#......#.#.......#.......##..####..#...#.#.#...##..#.#..###..#####........#..####......#..#

#..#.
#....
##..#
..#..
..###
`

	assert.Equal(t, expected, image.String())
}

func TestGetPixel(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	image, err := NewImage(reader)
	assert.Nil(t, err)

	assert.Equal(t, '#', image.GetPixel(0, 0))
	assert.Equal(t, '.', image.GetPixel(1, 0))
	assert.Equal(t, '.', image.GetPixel(0, 3))

	assert.Equal(t, '.', image.GetPixel(-5, -9))
	assert.Equal(t, '.', image.GetPixel(200, 50))
}
