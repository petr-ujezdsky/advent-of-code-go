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

func Test_ParsePackets_LiteralPacket(t *testing.T) {
	bits := HexadecimalStringToBits("D2FE28")
	packets := ParsePackets(bits)
	assert.Equal(t, 1, len(packets))

	packet := packets[0]
	assert.Equal(t, 6, packet.Version)
	assert.Equal(t, 4, packet.TypeID)
	//assert.Equal(t, 2021, packet.Value)
}

func Test_ParsePackets_OperatorPacket_1(t *testing.T) {
	bits := HexadecimalStringToBits("38006F45291200")
	packets := ParsePackets(bits)
	assert.Equal(t, 1, len(packets))

	packet := packets[0]
	assert.Equal(t, 1, packet.Version)
	assert.Equal(t, 6, packet.TypeID)
	//assert.Equal(t, 2021, packet.Value)
}

func Test_ParsePackets_OperatorPacket_2(t *testing.T) {
	bits := HexadecimalStringToBits("EE00D40C823060")
	packets := ParsePackets(bits)
	assert.Equal(t, 1, len(packets))

	packet := packets[0]
	assert.Equal(t, 7, packet.Version)
	assert.Equal(t, 3, packet.TypeID)
	//assert.Equal(t, 2021, packet.Value)
}

func Test_01_example_parse(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	ParseInput(reader)
	assert.Nil(t, err)
}

func Test_01_examples(t *testing.T) {
	assert.Equal(t, 16, SumVersions(ParsePackets(HexadecimalStringToBits("8A004A801A8002F478"))))
	assert.Equal(t, 12, SumVersions(ParsePackets(HexadecimalStringToBits("620080001611562C8802118E34"))))
	assert.Equal(t, 23, SumVersions(ParsePackets(HexadecimalStringToBits("C0015000016115A2E0802F182340"))))
	assert.Equal(t, 31, SumVersions(ParsePackets(HexadecimalStringToBits("A0016C880162017C3686B18A3D4780"))))
}

func Test_01(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	bits := ParseInput(reader)
	assert.Nil(t, err)

	sum := SumVersions(ParsePackets(bits))
	assert.Equal(t, 984, sum)
}
