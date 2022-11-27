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
	Size    int
	// for literal packet
	Value int
	// for operator packets
	SubPackets []Packet
}

func ParsePacket(bits Bits) (Packet, Bits) {
	version, bits := bits[:3].ToNumber(), bits[3:]
	typeID, bits := bits[:3].ToNumber(), bits[3:]

	// literal packet
	if typeID == 4 {
		var valueBits Bits
		payloadSize := 0

		for true {
			valueBits = append(valueBits, bits[1:5]...)
			payloadSize += 5

			if bits[0] == 0 {
				// "read" remainder
				//remainder := (4 - (3+3+payloadSize)%4) % 4
				remainder := 0
				bits = bits[5+remainder:]
				break
			}

			bits = bits[5:]
		}

		number := valueBits.ToNumber()

		packet := Packet{
			Version: version,
			TypeID:  typeID,
			Size:    3 + 3 + payloadSize,
			Value:   number,
		}

		return packet, bits
	}

	// operator packet
	lengthTypeID, bits := bits[:1].ToNumber(), bits[1:]
	var payloadSize int

	if lengthTypeID == 0 {
		payloadSize, bits = bits[:15].ToNumber(), bits[15:]
		packets, bits := ParsePackets(bits[:payloadSize]), bits[payloadSize:]

		packet := Packet{
			Version:    version,
			TypeID:     typeID,
			Size:       3 + 3 + 1 + payloadSize,
			SubPackets: packets,
		}

		return packet, bits
	}

	packetsCount, bits := bits[:11].ToNumber(), bits[11:]
	var packets []Packet

	for i := 0; i < packetsCount; i++ {
		packet, bitsNew := ParsePacket(bits)
		packets = append(packets, packet)
		payloadSize += packet.Size
		bits = bitsNew
	}

	packet := Packet{
		Version:    version,
		TypeID:     typeID,
		Size:       3 + 3 + 1 + payloadSize,
		SubPackets: packets,
	}

	return packet, bits
}

func ParsePackets(bits Bits) []Packet {
	var packets []Packet
	var current Packet

	for len(bits) > 7 {
		current, bits = ParsePacket(bits)
		packets = append(packets, current)
	}

	return packets
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

func SumVersions(packets []Packet) int {
	sum := 0
	for _, packet := range packets {
		sum += packet.Version + SumVersions(packet.SubPackets)
	}

	return sum
}

func ParseInput(r io.Reader) Bits {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	scanner.Scan()
	return HexadecimalStringToBits(scanner.Text())
}
