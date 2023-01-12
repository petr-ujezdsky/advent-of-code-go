package main

import (
	_ "embed"
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
	Ones  uint64
	Zeros uint64
}

type Mem struct {
	Location int
	Value    uint64
}

type MaskOrMem struct {
	IsMask bool
	Mask   Mask
	Mem    Mem
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

	memory := make(map[int]uint64)

	for _, maskOrMem := range items {
		if maskOrMem.IsMask {
			mask = maskOrMem.Mask
			continue
		}

		memory[maskOrMem.Mem.Location] = applyMask(maskOrMem.Mem.Value, mask)
	}

	return int(utils.Sum(maps.Values(memory)))
}

func DoWithInputPart02(items []MaskOrMem) int {
	return len(items)
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
					Ones:  ones.GetMask(),
					Zeros: zeros.GetMask(),
				},
				Mem: Mem{},
			}
		}

		parts := utils.ExtractInts(str, false)

		return MaskOrMem{
			IsMask: false,
			Mask:   Mask{},
			Mem: Mem{
				Location: parts[0],
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
