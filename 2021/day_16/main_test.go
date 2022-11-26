package day_16

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_HexadecimalStringToBits(t *testing.T) {
	bits := HexadecimalStringToBits("D2FE28")
	assert.Equal(t, "110100101111111000101000", fmt.Sprint(bits))

	bits = HexadecimalStringToBits("38006F45291200")
	assert.Equal(t, "00111000000000000110111101000101001010010001001000000000", fmt.Sprint(bits))

	bits = HexadecimalStringToBits("EE00D40C823060")
	assert.Equal(t, "11101110000000001101010000001100100000100011000001100000", fmt.Sprint(bits))
}

func Test_01_example_parse(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	ParseInput(reader)
	assert.Nil(t, err)
}
