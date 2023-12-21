package main

import (
	_ "embed"
	"fmt"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/measure"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/parsers"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/slices"
	"io"
	"regexp"
	"strconv"
	"strings"
)

var regexDots = regexp.MustCompile(`[.]+`)

type Record struct {
	ConditionsRaw       string
	Conditions          []rune
	Groups              [][]rune
	GroupSizes          []int
	DamagedCountTotal   int
	DamagedCountInitial int
	Unknowns            int
}

type World struct {
	Records []Record
}

func DoWithInputPart01(world World) int {
	sum := 0

	max := 0
	for i, record := range world.Records {
		count := calculateArrangementsCount(record)
		fmt.Printf("#%d combinations: %d\n", i, count)
		max = utils.Max(max, count)
		sum += count
	}

	fmt.Printf("\nMax combinations: %d\n", max)

	return sum
}

func calculateArrangementsCount(record Record) int {
	defer measure.Duration(measure.Track(fmt.Sprintf("Optimized calculation for %d unknowns took", record.Unknowns)))
	return calculateArrangementsCountMutable(0, record.Conditions, '.', 0, 0, record.GroupSizes)
}

func calculateArrangementsCountMutable(position int, conditions []rune, previous rune, currentGroupIndex, currentGroupSize int, groupSizes []int) int {
	for i := position; i < len(conditions); i++ {
		current := conditions[i]

		// increase group
		if current == '#' {
			currentGroupSize++

			if currentGroupIndex >= len(groupSizes) {
				return 0
			}

			if currentGroupSize > groupSizes[currentGroupIndex] {
				return 0
			}
		}

		// group end
		last := i == len(conditions)-1
		if previous == '#' && (current == '.') || current == '#' && last {
			// found more groups
			if currentGroupIndex >= len(groupSizes) {
				return 0
			}

			// group size is different -> not valid
			if groupSizes[currentGroupIndex] != currentGroupSize {
				return 0
			}

			// group size is same -> continue
			currentGroupIndex++
			currentGroupSize = 0
		}

		if current == '?' {
			position = i
			sum := 0

			// try '#'
			conditions[position] = '#'
			sum += calculateArrangementsCountMutable(position, conditions, previous, currentGroupIndex, currentGroupSize, groupSizes)

			conditions[position] = '.'
			// try '.'
			sum += calculateArrangementsCountMutable(position, conditions, previous, currentGroupIndex, currentGroupSize, groupSizes)

			// revert
			conditions[position] = '?'

			return sum
		}

		previous = current
	}

	// found less groups
	if currentGroupIndex != len(groupSizes) {
		return 0
	}

	//fmt.Printf("Valid: %v\n", string(conditions))
	// valid
	return 1
}

func DoWithInputPart02(world World) int {
	sum := 0

	results := utils.ProcessParallel(world.Records, calculateArrangementsCountUnfolded2)

	for result := range results {
		sum += result.Value
	}

	return sum
}

func calculateArrangementsCountUnfolded(record Record, i int) int {
	//fmt.Printf("#%d: ?'s count: %d\n", i, record.Unknowns)
	singleCount := calculateArrangementsCount(record)

	unfolded2 := Unfold2(record)
	//fmt.Printf("#%d: ?'s count: %d\n", i, unfolded2.Unknowns)
	pairCount := calculateArrangementsCount(unfolded2)

	fmt.Println()
	k := pairCount / singleCount
	return pairCount * k * k * k
}
func calculateArrangementsCountUnfolded2(record Record, i int) int {
	unfolded := Unfold(record)
	return calculateArrangementsCount(unfolded)
}

func Unfold(record Record) Record {
	// multiply data 5x
	conditionsRaw := record.ConditionsRaw + "?" + record.ConditionsRaw + "?" + record.ConditionsRaw + "?" + record.ConditionsRaw + "?" + record.ConditionsRaw
	groupSizes := slices.Repeat(record.GroupSizes, 5)

	// convert groupSizes ints to string
	groupSizesStr := slices.Map(groupSizes, strconv.Itoa)
	groupSizesStrJoined := strings.Join(groupSizesStr, ",")

	// join to raw record string
	recordStr := conditionsRaw + " " + groupSizesStrJoined

	// parse
	return ParseRecord(recordStr)
}

func Unfold2(record Record) Record {
	// multiply data 2x
	conditionsRaw := record.ConditionsRaw + "?" + record.ConditionsRaw
	groupSizes := slices.Repeat(record.GroupSizes, 2)

	// convert groupSizes ints to string
	groupSizesStr := slices.Map(groupSizes, strconv.Itoa)
	groupSizesStrJoined := strings.Join(groupSizesStr, ",")

	// join to raw record string
	recordStr := conditionsRaw + " " + groupSizesStrJoined

	// parse
	return ParseRecord(recordStr)
}

func ParseRecord(str string) Record {
	parts := strings.Split(str, " ")

	conditionsRaw := parts[0]
	groupSizes := utils.ExtractInts(parts[1], false)

	groupsSum := utils.Sum(groupSizes)

	unknowns := 0
	for _, condition := range conditionsRaw {
		if condition == '?' {
			unknowns++
		}
	}

	damaged := 0
	for _, condition := range conditionsRaw {
		if condition == '#' {
			damaged++
		}
	}

	return Record{
		ConditionsRaw:       conditionsRaw,
		Conditions:          []rune(conditionsRaw),
		GroupSizes:          groupSizes,
		DamagedCountTotal:   groupsSum,
		DamagedCountInitial: damaged,
		Unknowns:            unknowns,
	}
}

func ParseInput(r io.Reader) World {
	items := parsers.ParseToObjects(r, ParseRecord)
	return World{Records: items}
}
