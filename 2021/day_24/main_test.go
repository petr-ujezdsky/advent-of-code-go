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

func Test_01(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	instructions := ParseInput(reader)

	assert.Equal(t, [...]int{9, 1, 18, 172660766}, Run(instructions, "13579246899999"))

	assert.Equal(t, [...]int{9, 1, 18, 270493644}, Run(instructions, "99999999999999"))
	assert.Equal(t, [...]int{1, 1, 10, 171640596}, Run(instructions, "11111111111111"))

	// A @ 4
	// B @ 5
	// C @ 15
	A, B, C := 0, 0, 0
	iGroup := 0
	//formula := "(z / A) * 26 + i + C"
	for i, instruction := range instructions {
		if i%18 == 4 {
			//fmt.Printf("%v, %v\n", instruction.Name, instruction.VRight)
			A = instruction.VRight
		}

		if i%18 == 5 {
			//fmt.Printf("%v, %v\n", instruction.Name, instruction.VRight)
			B = instruction.VRight
		}

		if i%18 == 15 {
			//fmt.Printf("%v, %v\n", instruction.Name, instruction.VRight)
			C = instruction.VRight

			//fmt.Printf("Group #%2d: A=%2d, B=%3d, C=%2d\n", iGroup+1, A, B, C)
			//fmt.Printf("Group #%2d: A=%2d, B=%3d, C=%2d   z = (z / %v) * 26 + i + %v\n", iGroup+1, A, B, C, A, C)
			fmt.Printf("Group #%2d: z = (z / %v) * 26 + i + %v\n", iGroup+1, A, C)
			iGroup++
		}
	}
	_ = B
}

func Test_01_Run(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	instructions := ParseInput(reader)

	//input := "13579246899999"
	//input := "11111111111111"
	input := "99999999999999"

	//..................1111
	//........01234567890123
	//input := "99934999949999"
	//input := "99934999949999"

	//.................1111
	//.......01234567890123
	//input = "99934999944499"
	//input = "99994999944399"
	groups := groupInstructions(instructions)
	//instructions = deGroupInstructions(groups, 8, 9, 10, 11, 12, 13)
	instructions = deGroupInstructions(groups, 8, 9, 10, 11, 12, 13)
	registers := Registers{}
	//registers[3] = 6623276
	registers[3] = 6623281
	fmt.Printf("r=%v\n", RunRegisters(registers, instructions, input))
	//assert.Equal(t, RunRegisters(registers, instructions, input)[3], RunDecompiledRegister(registers[3], instructions, input))
	//assert.Equal(t, RunRegisters(registers, instructions, input)[3], RunDecompiledRegister(registers[3], instructions, input))
}

func Test_01_decompiled(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	instructions := ParseInput(reader)

	//input := "13579246899999"
	//input := "11111111111111"
	input := "99999999999999"

	//..................1111
	//........01234567890123
	//input := "99934999949999"
	//input := "99934999949999"

	//.................1111
	//.......01234567890123
	//input = "99934999944499"
	//input = "99994999944399"
	assert.Equal(t, Run(instructions, input)[3], RunDecompiled(instructions, input))
}

func Test_01_decompiled_cycle(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	instructions := ParseInput(reader)

	index := 0
	for i := 0; i < 9; i++ {
		//.........................1111
		//...............01234567890123
		input := []rune("99999999999999")
		input[index] = rune('1' + i)

		RunDecompiled(instructions, string(input))
		fmt.Println("======================")
	}
}

func Test_01_decompiled_cycle2(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	instructions := ParseInput(reader)

	groups := groupInstructions(instructions)

	iGroup := 8
	for z0 := 432; z0 < 432; z0++ {
		for i := 0; i < 9; i++ {
			//.........................1111
			//...............01234567890123
			inputStr := string(rune('1' + i))

			registers := Registers{}
			registers[3] = z0
			r := RunRegisters(registers, groups[iGroup], inputStr)
			fmt.Printf("Group #%v, z=%v, input=%v, r=%v\n", iGroup, z0, inputStr, r)
			assert.Equal(t, r[3], RunDecompiledRegister(z0, groups[iGroup], inputStr))
		}
		fmt.Println("======================")
	}
}

func Test_01_decompiled_cycle3(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	instructions := ParseInput(reader)

	abcs := utils.Reverse(extractABC(instructions))

	zDesiredFrom, zDesiredTo := 0, 0
	for i, abc := range abcs {
		A := abc.X
		B := abc.Y
		//C := abc.Z

		zFrom := zDesiredFrom*A + 1 - B
		zTo := zDesiredTo*A + 9 - B

		fmt.Printf("z%2d = [%10d,%10d]  =>  z%2d =[%10d,%10d]\n", len(abcs)-i-1, zFrom, zTo, len(abcs)-i, zDesiredFrom, zDesiredTo)
		zDesiredFrom = zFrom
		zDesiredTo = zTo
	}
}

func Test_01_decompiled_subgroups(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	instructions := ParseInput(reader)

	//input := "13579246899999"
	//input := "11111111111111"
	input := "99999999999999"

	//..................1111
	//........01234567890123
	//input := "99934999949999"
	//input := "99934999949999"

	//.................1111
	//.......01234567890123
	//input = "99934999944499"
	//input = "99994999944399"
	groups := groupInstructions(instructions)
	//instructions = deGroupInstructions(groups, 8, 9, 10, 11, 12, 13)
	instructions = deGroupInstructions(groups, 8, 9, 10, 11, 12, 13)
	instructions = deGroupInstructions(groups, 8)
	registers := Registers{}
	//registers[3] = 6623276
	registers[3] = 6623281
	registers[3] = (100 - 9 - 1) / 26
	fmt.Printf("%v -> r=%v\n", registers, RunRegisters(registers, instructions, input))
	//assert.Equal(t, RunRegisters(registers, instructions, input)[3], RunDecompiledRegister(registers[3], instructions, input))
	//assert.Equal(t, RunRegisters(registers, instructions, input)[3], RunDecompiledRegister(registers[3], instructions, input))
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
