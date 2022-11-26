package day_16

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type Bits []byte

func (bits Bits) String() string {
	sb := strings.Builder{}

	for _, bit := range bits {
		sb.WriteByte('0' + bit)
	}

	return sb.String()
}

func (bits Bits) ToNumber() int {
	length := len(bits)

	if length > strconv.IntSize {
		panic("Too many bits " + strconv.Itoa(length))
	}

	number := 0
	for i, bit := range bits {
		number += int(bit) << (length - i - 1)
	}

	return number
}

type Packet struct {
	Version int
	TypeID  int
	Payload Bits
}

func NewPacket(bits Bits) Packet {
	version, bits := bits[:3].ToNumber(), bits[3:]
	typeID, bits := bits[:3].ToNumber(), bits[3:]

	return Packet{
		Version: version,
		TypeID:  typeID,
		Payload: bits,
	}
}

func HexadecimalStringToBits(text string) Bits {
	var bits Bits

	for _, char := range text {
		var number int32

		if '0' <= char && char <= '9' {
			number = char - '0'
		} else if 'A' <= char && char <= 'F' {
			number = char - 'A' + 10
		}

		for _, zeroOne := range fmt.Sprintf("%04b", number) {
			bits = append(bits, byte(zeroOne-'0'))
		}
	}

	return bits
}

func ParseInput(r io.Reader) Bits {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	scanner.Scan()
	return HexadecimalStringToBits(scanner.Text())
}
