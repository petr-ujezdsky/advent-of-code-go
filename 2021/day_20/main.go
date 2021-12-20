package day_20

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

type Image struct {
	Pixels   []string
	Enhancor string
}

func NewImage(r io.Reader) (Image, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	image := Image{}

	for scanner.Scan() {
		row := scanner.Text()

		if row == "" {
			continue
		}

		if image.Enhancor == "" {
			image.Enhancor = row
			continue
		}

		image.Pixels = append(image.Pixels, row)
	}

	return image, scanner.Err()
}

func (image *Image) String() string {
	var sb strings.Builder

	sb.WriteString(image.Enhancor)
	sb.WriteRune('\n')
	sb.WriteRune('\n')

	for _, row := range image.Pixels {
		sb.WriteString(row)
		sb.WriteRune('\n')
	}

	return sb.String()
}

func (image *Image) GetPixel(x, y int) rune {
	if x < 0 || x >= image.Width() || y < 0 || y >= image.Height() {
		return '.'
	}

	return rune(image.Pixels[y][x])
}

func (image *Image) Width() int {
	return len(image.Pixels[0])
}

func (image *Image) Height() int {
	return len(image.Pixels)
}

func (image *Image) Enhance() *Image {
	enhanced := &Image{Enhancor: image.Enhancor}

	for y := -2; y < image.Height()+3; y++ {
		row := ""
		for x := -2; x < image.Width()+3; x++ {
			index := readPixelsToIndex(image, x, y)

			pixel := rune(image.Enhancor[index])

			row += string(pixel)
		}
		enhanced.Pixels = append(enhanced.Pixels, row)
	}

	return enhanced
}

func (image *Image) LightPixelsCount() int {
	count := 0

	for _, row := range image.Pixels {
		for _, pixel := range row {
			if pixel == '#' {
				count++
			}
		}
	}

	return count
}

func readPixelsToIndex(image *Image, x, y int) int {
	result := ""

	for offsetY := -1; offsetY <= 1; offsetY++ {
		for offsetX := -1; offsetX <= 1; offsetX++ {
			pixel := image.GetPixel(x+offsetX, y+offsetY)
			if pixel == '.' {
				result += "0"
			} else {
				result += "1"
			}
		}
	}

	index, _ := strconv.ParseInt(result, 2, 0)

	return int(index)
}
