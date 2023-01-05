package day_24

import (
	"fmt"
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

func Test_01_extract_ABC(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	instructions := ParseInput(reader)

	abcs := extractABC(instructions)
	sb := &strings.Builder{}

	for i, abc := range abcs {
		A := abc.X
		B := abc.Y
		C := abc.Z

		fmt.Printf("Group #%2d: A=%2d, B=%3d, C=%2d\n", i, A, B, C)
		sb.WriteString(fmt.Sprintf("Group #%2d: A=%2d, B=%3d, C=%2d\n", i, A, B, C))
	}

	expected := utils.Msg(`
Group # 0: A= 1, B= 13, C=13
Group # 1: A= 1, B= 11, C=10
Group # 2: A= 1, B= 15, C= 5
Group # 3: A=26, B=-11, C=14
Group # 4: A= 1, B= 14, C= 5
Group # 5: A=26, B=  0, C=15
Group # 6: A= 1, B= 12, C= 4
Group # 7: A= 1, B= 12, C=11
Group # 8: A= 1, B= 14, C= 1
Group # 9: A=26, B= -6, C=15
Group #10: A=26, B=-10, C=12
Group #11: A=26, B=-12, C= 8
Group #12: A=26, B= -3, C=14
Group #13: A=26, B= -5, C= 9
`)

	assert.Equal(t, expected, sb.String())
}

func Test_01(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	instructions := ParseInput(reader)

	modelNumber := BruteForcePossibleValues(instructions, 9, 1)
	assert.Equal(t, "12934998949199", modelNumber)
}

func Test_02(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	instructions := ParseInput(reader)

	modelNumber := BruteForcePossibleValues(instructions, 1, 9)
	assert.Equal(t, "11711691612189", modelNumber)
}
