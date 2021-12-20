package day_20

import (
	"bufio"
	"io"
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
	width := len(image.Pixels[0])
	height := len(image.Pixels)

	if x < 0 || x >= width || y < 0 || y >= height {
		return '.'
	}

	return rune(image.Pixels[y][x])
}
