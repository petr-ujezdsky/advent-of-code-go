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

type Packet interface {
	GetVersion() int
	GetTypeID() int
	GetSize() int
}

type LiteralPacket struct {
	Version int
	TypeID  int
	Size    int
	Value   int
}

func (p LiteralPacket) GetVersion() int {
	return p.Version
}

func (p LiteralPacket) GetTypeID() int {
	return p.TypeID
}

func (p LiteralPacket) GetSize() int {
	return p.Size
}

type OperatorPacket struct {
	Version    int
	TypeID     int
	Size       int
	SubPackets []Packet
}

func (p OperatorPacket) GetVersion() int {
	return p.Version
}

func (p OperatorPacket) GetTypeID() int {
	return p.TypeID
}

func (p OperatorPacket) GetSize() int {
	return p.Size
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

		packet := LiteralPacket{
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

		packet := OperatorPacket{
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
		payloadSize += packet.GetSize()
		bits = bitsNew
	}

	packet := OperatorPacket{
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

func ParseInput(r io.Reader) Bits {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	scanner.Scan()
	return HexadecimalStringToBits(scanner.Text())
}
