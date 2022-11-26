package day_16

import (
	"bufio"
	"fmt"
	"io"
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
