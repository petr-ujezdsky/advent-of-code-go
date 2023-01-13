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

//func (a Address) CombinationsWith(a2 Address) int {
//	combinations := 1
//	for i, trit := range a {
//		otherTrit := a2[i]
//		combinations *= trit.CombinationsWith(otherTrit)
//		if combinations == 0 {
//			return 0
//		}
//	}
//	return combinations
//}

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

func (a Address) Intersect(addresses []Address) []Address {
	var intersections []Address

	for _, a2 := range addresses {
		and := a.And(a2)
		if and != "" {
			intersections = append(intersections, and)
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

//func TritAnd(t Trit, t2 Trit) (Trit, bool) {
//	if t == t2 {
//		return t, true
//	}
//
//	if t == 'X' {
//		return t2, true
//	}
//
//	if t2 == 'X' {
//		return t, true
//	}
//
//	return 0, false
//}

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

func filterMatching(a Address, records []Record) []Record {
	var filtered []Record

	for _, record := range records {
		if a.Matches(record.Address) {
			filtered = append(filtered, record)
		}
	}

	return filtered
}

//func merge(addresses []Address) Address {
//	merged := slices.Clone(addresses[0])
//
//	for i, trit := range merged {
//		mergedTrit := trit
//
//		for _, record := range addresses[1:] {
//			otherTrit := record[i]
//			if otherTrit != mergedTrit {
//				mergedTrit = 'X'
//				break
//			}
//		}
//		merged[i] = mergedTrit
//	}
//
//	return merged
//}

//func andAddresses(address Address, other []Address) []Address {
//	var anded []Address
//
//	for _, address2 := range other {
//		and := address.And(address2)
//		if and != nil {
//			anded = append(anded, and)
//		}
//	}
//
//	return anded
//}

func computeIntersections(addresses []Address) []Address {
	var intersections []Address

	for i, a1 := range addresses {
		intersections = append(intersections, a1.Intersect(addresses[i+1:])...)
	}

	return intersections
}

func countAll(addresses []Address) int {
	count := 0
	for _, address := range addresses {
		count += address.Combinations()
	}

	return count
}

func countUnique(addresses []Address) int {
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
		fmt.Printf("Record #%3d, %v, %v\n", i, record.AddressStr, record.Value)
		address := record.Address
		totalCount := address.Combinations()
		//dividedTotalCount := totalCount

		//matching := filterMatching(address, records[i+1:])
		//adresses := slices.Map(matching, func(r Record) Address { return r.Address })
		//commonCount := 0

		intersects := address.Intersect(allAddresses[i+1:])
		intersectsCount := countUnique(intersects)
		//
		//if len(matching) > 0 {
		//	commonCount = 1
		//	mergedAddress := merge(slices.Map(matching, func(r Record) Address { return r.Address }))
		//
		//	//fmt.Printf("Mask:           %v\n", mask.Trinary)
		//	//fmt.Printf("Address before: %v (%v)\n", trinaryStr, address)
		//
		//	for j, trit := range address {
		//		if trit != 'X' {
		//			continue
		//		}
		//		// now trit = X
		//
		//		mergedTrit := mergedAddress[j]
		//		if trit == mergedTrit {
		//			// X vs X
		//			commonCount *= 2
		//			dividedTotalCount /= 2
		//			continue
		//		}
		//
		//		// X vs 0 or X vs 1
		//	}
		//}
		//
		effectiveCount := totalCount - intersectsCount
		//effectiveCount = dividedTotalCount
		//if effectiveCount < 0 {
		//	panic("Whoa")
		//}

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
