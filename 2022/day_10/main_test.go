package main

import (
	"fmt"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/matrix"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func Test_01_parse(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	instructions := ParseInput(reader)

	assert.Equal(t, 146, len(instructions))
}

func Test_01_example(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	instructions := ParseInput(reader)

	result, _ := DoWithInput(instructions)
	assert.Equal(t, 13140, result)
}

func Test_01(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	instructions := ParseInput(reader)

	result, _ := DoWithInput(instructions)
	assert.Equal(t, 13920, result)
}

func Test_02_example(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	instructions := ParseInput(reader)

	_, pixels := DoWithInput(instructions)

	fmt.Printf("%v\n", pixels.StringFmt(matrix.FmtBoolean[bool]))
}

func Test_02(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	instructions := ParseInput(reader)

	_, pixels := DoWithInput(instructions)

	fmt.Printf("%v\n", pixels.StringFmt(matrix.FmtBoolean[bool]))
}
