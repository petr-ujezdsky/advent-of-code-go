package day_24

import (
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_01_example_1(t *testing.T) {
	input := utils.Msg(`
inp x
mul x -1
`)

	instructions := ParseInput(strings.NewReader(input))

	assert.Equal(t, [...]int{0, -5, 0, 0}, Run(instructions, "5"))
	assert.Equal(t, [...]int{0, -1, 0, 0}, Run(instructions, "1"))
}

func Test_01_example_2(t *testing.T) {
	input := utils.Msg(`
inp z
inp x
mul z 3
eql z x
`)

	instructions := ParseInput(strings.NewReader(input))

	assert.Equal(t, [...]int{0, 9, 0, 1}, Run(instructions, "39"))
	assert.Equal(t, [...]int{0, 9, 0, 0}, Run(instructions, "49"))
}

func Test_01_example_3(t *testing.T) {
	input := utils.Msg(`
inp w
add z w
mod z 2
div w 2
add y w
mod y 2
div w 2
add x w
mod x 2
div w 2
mod w 2
`)

	instructions := ParseInput(strings.NewReader(input))

	assert.Equal(t, [...]int{0, 0, 0, 1}, Run(instructions, "1"))
	assert.Equal(t, [...]int{0, 0, 1, 0}, Run(instructions, "2"))
	assert.Equal(t, [...]int{0, 0, 1, 1}, Run(instructions, "3"))
	assert.Equal(t, [...]int{0, 1, 0, 0}, Run(instructions, "4"))
	assert.Equal(t, [...]int{0, 1, 0, 1}, Run(instructions, "5"))
	assert.Equal(t, [...]int{0, 1, 1, 0}, Run(instructions, "6"))
	assert.Equal(t, [...]int{0, 1, 1, 1}, Run(instructions, "7"))
	assert.Equal(t, [...]int{1, 0, 0, 0}, Run(instructions, "8"))
	assert.Equal(t, [...]int{1, 0, 0, 1}, Run(instructions, "9"))
}

func Test_01(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	instructions := ParseInput(reader)

	assert.Equal(t, [...]int{9, 1, 18, 172660766}, Run(instructions, "13579246899999"))

	assert.Equal(t, [...]int{9, 1, 18, 270493644}, Run(instructions, "99999999999999"))
	assert.Equal(t, [...]int{1, 1, 10, 171640596}, Run(instructions, "11111111111111"))
}

//
//func Test_02_example(t *testing.T) {
//	reader, err := os.Open("data-00-example.txt")
//	assert.Nil(t, err)
//
//	image, err := NewImage(reader)
//	assert.Nil(t, err)
//
//	fmt.Println(image.String())
//	fmt.Println("------------------------")
//
//	var enhanced = image
//	for i := 0; i < 50; i++ {
//		enhanced = *enhanced.Enhance()
//	}
//	fmt.Println(enhanced.String())
//
//	assert.Equal(t, 3351, enhanced.LightPixelsCount())
//}
//
//func Test_02(t *testing.T) {
//	reader, err := os.Open("data-01.txt")
//	assert.Nil(t, err)
//
//	image, err := NewImage(reader)
//	assert.Nil(t, err)
//
//	fmt.Println(image.String())
//	fmt.Println("------------------------")
//
//	var enhanced = image
//	for i := 0; i < 50; i++ {
//		enhanced = *enhanced.Enhance()
//	}
//	fmt.Println(enhanced.String())
//
//	assert.Equal(t, 17917, enhanced.LightPixelsCount())
//}
//
//func TestGetPixel(t *testing.T) {
//	reader, err := os.Open("data-00-example.txt")
//	assert.Nil(t, err)
//
//	image, err := NewImage(reader)
//	assert.Nil(t, err)
//
//	assert.Equal(t, '#', image.GetPixel(0, 0))
//	assert.Equal(t, '.', image.GetPixel(1, 0))
//	assert.Equal(t, '.', image.GetPixel(0, 3))
//
//	assert.Equal(t, '.', image.GetPixel(-5, -9))
//	assert.Equal(t, '.', image.GetPixel(200, 50))
//}
