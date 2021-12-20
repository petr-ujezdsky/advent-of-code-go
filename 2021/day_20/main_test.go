package day_20

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_01_example(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	image, err := NewImage(reader)
	assert.Nil(t, err)

	var expected string
	expected = `..#.#..#####.#.#.#.###.##.....###.##.#..###.####..#####..#....#..#..##..###..######.###...####..#..#####..##..#.#####...##.#.#..#.##..#.#......#.###.######.###.####...#.##.##..#..#..#####.....#.#....###..#.##......#.....#..#..#..##..#...##.######.####.####.#.#...#.......#..#.#.#...####.##.#......#..#...##.#.##..#...##.#.##..###.#......#.#.......#.#.#.####.###.##...#.....####.#..#..#.##.#....##..#.####....##...##..#...#......#.#.......#.......##..####..#...#.#.#...##..#.#..###..#####........#..####......#..#

#..#.
#....
##..#
..#..
..###
`

	assert.Equal(t, expected, image.String())

	var enhanced *Image
	enhanced = image.Enhance()

	fmt.Println(image.String())
	fmt.Println("------------------------")
	fmt.Println(enhanced.String())

	expected = `..#.#..#####.#.#.#.###.##.....###.##.#..###.####..#####..#....#..#..##..###..######.###...####..#..#####..##..#.#####...##.#.#..#.##..#.#......#.###.######.###.####...#.##.##..#..#..#####.....#.#....###..#.##......#.....#..#..#..##..#...##.######.####.####.#.#...#.......#..#.#.#...####.##.#......#..#...##.#.##..#...##.#.##..###.#......#.#.......#.#.#.####.###.##...#.....####.#..#..#.##.#....##..#.####....##...##..#...#......#.#.......#.......##..####..#...#.#.#...##..#.#..###..#####........#..####......#..#

..........
..##.##...
.#..#.#...
.##.#..#..
.####..#..
..#..##...
...##..#..
....#.#...
..........
..........
`

	assert.Equal(t, expected, enhanced.String())

	enhanced = image.Enhance().Enhance()

	fmt.Println(image.String())
	fmt.Println("------------------------")
	fmt.Println(enhanced.String())

	expected = `..#.#..#####.#.#.#.###.##.....###.##.#..###.####..#####..#....#..#..##..###..######.###...####..#..#####..##..#.#####...##.#.#..#.##..#.#......#.###.######.###.####...#.##.##..#..#..#####.....#.#....###..#.##......#.....#..#..#..##..#...##.######.####.####.#.#...#.......#..#.#.#...####.##.#......#..#...##.#.##..#...##.#.##..###.#......#.#.......#.#.#.####.###.##...#.....####.#..#..#.##.#....##..#.####....##...##..#...#......#.#.......#.......##..####..#...#.#.#...##..#.#..###..#####........#..####......#..#

...............
...............
.........#.....
...#..#.#......
..#.#...###....
..#...##.#.....
..#.....#.#....
...#.#####.....
....#.#####....
.....##.##.....
......###......
...............
...............
...............
...............
`

	assert.Equal(t, expected, enhanced.String())

	assert.Equal(t, 35, enhanced.LightPixelsCount())
}

func Test_01(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	image, err := NewImage(reader)
	assert.Nil(t, err)

	fmt.Println(image.String())
	fmt.Println("------------------------")

	var enhanced = image.Enhance()
	fmt.Println(enhanced.String())
	fmt.Println("------------------------")

	enhanced = enhanced.Enhance()
	fmt.Println(enhanced.String())

	assert.Equal(t, 5379, enhanced.LightPixelsCount())
}

func Test_02_example(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	image, err := NewImage(reader)
	assert.Nil(t, err)

	fmt.Println(image.String())
	fmt.Println("------------------------")

	var enhanced = image
	for i := 0; i < 50; i++ {
		enhanced = *enhanced.Enhance()
	}
	fmt.Println(enhanced.String())

	assert.Equal(t, 3351, enhanced.LightPixelsCount())
}

func Test_02(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	image, err := NewImage(reader)
	assert.Nil(t, err)

	fmt.Println(image.String())
	fmt.Println("------------------------")

	var enhanced = image
	for i := 0; i < 50; i++ {
		enhanced = *enhanced.Enhance()
	}
	fmt.Println(enhanced.String())

	assert.Equal(t, 17917, enhanced.LightPixelsCount())
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
