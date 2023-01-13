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

type AddressesMap = map[Address]struct{}

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

func (t Trit) And(t2 Trit) (Trit, bool) {
	if t == t2 {
		return t, true
	}

	if t == 'X' {
		return t2, true
	}

	if t2 == 'X' {
		return t, true
	}

	return 0, false
}

type Address string

func (a Address) Combinations() int {
	combinations := 1
	for _, trit := range a {
		if trit == 'X' {
			combinations *= 2
		}
	}
	return combinations
}

func (a Address) And(a2 Address) Address {
	result := []rune(a)

	for i, trit := range a {
		trit2 := a2[i]
		//if trit == trit2 {
		//	result[i] = trit
		//	continue
		//}
		//
		//if trit == 'X' {
		//	result[i] = trit2
		//	continue
		//}
		//
		//if trit2 == 'X' {
		//	result[i] = trit
		//	continue
		//}

		and, ok := Trit(trit).And(Trit(trit2))
		if !ok {
			return ""
		}
		result[i] = rune(and)
	}

	return Address(result)
}

func (a Address) Intersect(addresses []Address) AddressesMap {
	intersections := make(AddressesMap)

	for _, a2 := range addresses {
		and := a.And(a2)
		if and != "" {
			intersections[and] = struct{}{}
		}
	}

	return intersections
}

func (a Address) Matches(a2 Address) bool {
	for i, trit := range a {
		otherTrit := a2[i]

		if !Trit(trit).Matches(Trit(otherTrit)) {
			return false
		}
	}

	return true
}

type Record struct {
	Address    Address
	AddressStr string
	Value      uint64
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
	//fmt.Printf("Mask:           %v\n", mask.Trinary)
	//fmt.Printf("Address before: %v (%v)\n", trinaryStr, address)

	trinary := []rune(trinaryStr)
	for i, maskChar := range mask.Trinary {
		switch maskChar {
		case '1', 'X':
			trinary[i] = maskChar
		}
	}
	//fmt.Printf("Address after:  %v\n", string(trinary))

	return Address(trinary)
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

		address := maskAddress(maskOrMem.Mem.Location, mask)
		records = append(records, Record{
			Address:    address,
			AddressStr: string(address),
			Value:      maskOrMem.Mem.Value,
		})
	}

	return records
}

func computeIntersections(addresses AddressesMap) AddressesMap {
	addressesSlice := maps.Keys(addresses)
	intersections := make(AddressesMap)

	for i, a1 := range addressesSlice {
		intersection := a1.Intersect(addressesSlice[i+1:])
		for address := range intersection {
			intersections[address] = struct{}{}
		}
	}

	return intersections
}

func countAll(addresses AddressesMap) int {
	count := 0
	for address := range addresses {
		count += address.Combinations()
	}

	return count
}

func countUnique(addresses AddressesMap) int {
	if len(addresses) == 0 {
		return 0
	}

	countAll := countAll(addresses)

	intersections := computeIntersections(addresses)
	countIntersecting := countUnique(intersections)

	return countAll - countIntersecting
}

func DoWithInputPart02(items []MaskOrMem) int {
	records := toRecords(items)

	allAddresses := slices.Map(records, func(r Record) Address { return r.Address })
	sum := 0
	for i, record := range records {
		fmt.Printf("Record #%3d, %v, %v ... ", i, record.AddressStr, record.Value)
		address := record.Address
		totalCount := address.Combinations()

		intersects := address.Intersect(allAddresses[i+1:])
		intersectsCount := countUnique(intersects)

		effectiveCount := totalCount - intersectsCount

		fmt.Printf("count: %v\n", effectiveCount)

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
