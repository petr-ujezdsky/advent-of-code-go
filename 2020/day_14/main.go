package main

import (
	_ "embed"
	"fmt"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/collections"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/maps"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/parsers"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/strs"
	"io"
	"math"
	"strings"
)

type Mask struct {
	Trinary string
	Ones    uint64
	Zeros   uint64
}

type Mem struct {
	Location uint64
	Value    uint64
}

type MaskOrMem struct {
	IsMask bool
	Mask   Mask
	Mem    Mem
}

type Trit rune

func (t Trit) CombinationsWith(t2 Trit) int {
	if t == t2 {
		switch t {
		case '1', '0':
			return 1
		case 'X':
			return 4
		}
		panic("Wrong runes")
	}

	str := string(t) + string(t2)
	switch str {
	case "10", "01":
		return 0
	case "1X", "X1", "0X", "X0":
		return 2
	}

	panic("Wrong runes")
}

type Address []Trit

func NewAddress(trinary string) Address {
	address := make(Address, len(trinary))

	for i, char := range trinary {
		address[i] = Trit(char)
	}

	return address
}

func (a Address) Combinations() int {
	combinations := 1
	for _, trit := range a {
		if trit == 'X' {
			combinations *= 2
		}
	}
	return combinations
}

func (a Address) CombinationsWith(a2 Address) int {
	combinations := 1
	for i, trit := range a {
		otherTrit := a2[i]
		combinations *= trit.CombinationsWith(otherTrit)
		if combinations == 0 {
			return 0
		}
	}
	return combinations
}

type Record struct {
	Address Address
	Value   uint64
}

func applyMask(value uint64, mask Mask) uint64 {
	//fmt.Printf("Ones: %b\n", mask.Ones)
	//fmt.Printf("Zeros: %b\n", mask.Zeros)
	//fmt.Printf("Value before : %b (%v)\n", value, value)
	value |= mask.Ones
	value &= mask.Zeros
	//fmt.Printf("Value after : %b (%v)\n", value, value)
	return value
}

func DoWithInputPart01(items []MaskOrMem) int {
	mask := Mask{
		Ones:  0,
		Zeros: math.MaxUint64,
	}

	memory := make(map[uint64]uint64)

	for _, maskOrMem := range items {
		if maskOrMem.IsMask {
			mask = maskOrMem.Mask
			continue
		}

		memory[maskOrMem.Mem.Location] = applyMask(maskOrMem.Mem.Value, mask)
	}

	return int(utils.Sum(maps.Values(memory)))
}

func maskAddress(address uint64, mask Mask) Address {
	trinaryStr := strs.Substring(strs.ToBinary(address), 64-36, 64)
	fmt.Printf("Mask:           %v\n", mask.Trinary)
	fmt.Printf("Address before: %v (%v)\n", trinaryStr, address)

	trinary := []rune(trinaryStr)
	for i, maskChar := range mask.Trinary {
		switch maskChar {
		case '1', 'X':
			trinary[i] = maskChar
		}
	}
	fmt.Printf("Address after:  %v\n", string(trinary))

	return NewAddress(string(trinary))
}

func toRecords(items []MaskOrMem) []Record {
	// empty mask - will be overwritten by first item
	mask := Mask{}

	var records []Record

	for _, maskOrMem := range items {
		if maskOrMem.IsMask {
			mask = maskOrMem.Mask
			continue
		}

		records = append(records, Record{
			Address: maskAddress(maskOrMem.Mem.Location, mask),
			Value:   maskOrMem.Mem.Value,
		})
	}

	return records
}

func DoWithInputPart02(items []MaskOrMem) int {
	records := toRecords(items)

	sum := 0
	for i, record := range records {
		count := record.Address.Combinations()
		for _, otherRecord := range records[i+1:] {
			count -= record.Address.CombinationsWith(otherRecord.Address)
			if count <= 0 {
				count = 0
				break
			}
		}

		sum += count * int(record.Value)
	}

	return sum
}

func ParseInput(r io.Reader) []MaskOrMem {
	parseItem := func(str string) MaskOrMem {
		// mask
		if str[1] == 'a' {
			mask := strings.Split(str, " ")[2]
			ones := collections.NewEmptyBitSet64()
			zeros := collections.NewFullBitSet64()

			for i, char := range strs.ReverseString(mask) {
				switch char {
				case '1':
					ones.Push(i)
				case '0':
					zeros.Remove(i)
				}
			}

			return MaskOrMem{
				IsMask: true,
				Mask: Mask{
					Trinary: mask,
					Ones:    ones.GetMask(),
					Zeros:   zeros.GetMask(),
				},
				Mem: Mem{},
			}
		}

		parts := utils.ExtractInts(str, false)

		return MaskOrMem{
			IsMask: false,
			Mask:   Mask{},
			Mem: Mem{
				Location: uint64(parts[0]),
				Value:    uint64(parts[1]),
			},
		}
	}

	return parsers.ParseToObjects(r, parseItem)
}

//func ParseInput(r io.Reader) []MaskOrMem {
//	scanner := bufio.NewScanner(r)
//	scanner.Split(bufio.ScanLines)
//
//	var items []MaskOrMem
//	for scanner.Scan() {
//		//parts := strings.Split(scanner.Text(), ",")
//		//ints := utils.ExtractInts(scanner.Text(), false)
//
//		item := MaskOrMem{}
//
//		items = append(items, item)
//	}
//
//	return items
//}

//func ParseInput(r io.Reader) utils.Matrix[MaskOrMem] {
//	parseItem := func(char rune) MaskOrMem {
//		return MaskOrMem{}
//	}
//
//	return parsers.ParseToMatrix(r, parseItem)
//}
