package main

import (
	_ "embed"
	"fmt"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/collections"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/maps"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/parsers"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/slices"
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
			return 2
		}
		panic("Wrong runes")
	}

	str := string(t) + string(t2)
	switch str {
	case "10", "01":
		return 0
	case "1X", "0X":
		return 1
	case "X1", "X0":
		return 1
	}

	panic("Wrong runes")
}
func (t Trit) Matches(t2 Trit) bool {
	if t == t2 {
		return true
	}

	if t == 'X' || t2 == 'X' {
		return true
	}
	return false
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

func (a Address) Matches(a2 Address) bool {
	for i, trit := range a {
		otherTrit := a2[i]

		if !trit.Matches(otherTrit) {
			return false
		}
	}

	return true
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

func filterMatching(a Address, records []Record) []Record {
	var filtered []Record

	for _, record := range records {
		if a.Matches(record.Address) {
			filtered = append(filtered, record)
		}
	}

	return filtered
}

func merge(records []Record) Address {
	merged := slices.Clone(records[0].Address)

	for i, trit := range merged {
		mergedTrit := trit

		for _, record := range records[1:] {
			otherTrit := record.Address[i]
			if otherTrit != mergedTrit {
				mergedTrit = 'X'
				break
			}
		}
		merged[i] = mergedTrit
	}

	return merged
}

func DoWithInputPart02(items []MaskOrMem) int {
	records := toRecords(items)

	sum := 0
	for i, record := range records {
		address := record.Address
		totalCount := address.Combinations()

		matching := filterMatching(address, records[i+1:])
		commonCount := 0

		if len(matching) > 0 {
			commonCount = 1
			mergedAddress := merge(matching)

			for j, trit := range address {
				if trit != 'X' {
					continue
				}
				// now trit = X

				mergedTrit := mergedAddress[j]
				if trit == mergedTrit {
					// X vs X
					commonCount *= 2
					continue
				}

				// X vs 0 or X vs 1
			}
		}

		effectiveCount := totalCount - commonCount
		if effectiveCount < 0 {
			panic("Whoa")
		}

		sum += effectiveCount * int(record.Value)
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
