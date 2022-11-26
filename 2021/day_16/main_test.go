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

func Test_ToNumber(t *testing.T) {
	bits := HexadecimalStringToBits("5")
	assert.Equal(t, 5, bits.ToNumber())

	bits = HexadecimalStringToBits("FF")
	assert.Equal(t, 255, bits.ToNumber())

	bits = HexadecimalStringToBits("FFFF")
	assert.Equal(t, 65535, bits.ToNumber())
	// FF
	assert.Equal(t, 255, bits[:len(bits)/2].ToNumber())
	assert.Equal(t, 255, bits[len(bits)/2:].ToNumber())
}

func Test_NewPacket(t *testing.T) {
	bits := HexadecimalStringToBits("D2FE28")
	packet := NewPacket(bits)
	assert.Equal(t, 6, packet.Version)
	assert.Equal(t, 4, packet.TypeID)
	assert.Equal(t, "101111111000101000", packet.Payload.String())
}

func Test_01_example_parse(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	ParseInput(reader)
	assert.Nil(t, err)
}
