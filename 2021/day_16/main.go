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

type LiteralPacket struct {
	Version int
	TypeID  int
	Value   int
}

func NewLiteralPacket(bits Bits) LiteralPacket {
	version, bits := bits[:3].ToNumber(), bits[3:]
	typeID, bits := bits[:3].ToNumber(), bits[3:]

	var valueBits Bits
	for len(bits) > 0 {
		valueBits = append(valueBits, bits[1:5]...)

		if bits[0] == 0 {
			break
		}

		bits = bits[5:]
	}

	number := valueBits.ToNumber()

	return LiteralPacket{
		Version: version,
		TypeID:  typeID,
		Value:   number,
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
