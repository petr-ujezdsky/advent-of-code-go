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
	defer measure.Duration(measure.Track(fmt.Sprintf("Calculation for %d unknowns took", record.Unknowns)))
	return calculateArrangementsCountMutable(0, record.Conditions, record.GroupSizes, record.DamagedCountInitial, record.DamagedCountTotal)
}

func calculateArrangementsCountMutable(position int, conditions []rune, groupSizes []int, damagedCountCurrent, damagedCountTarget int) int {
	// too many damaged springs
	if damagedCountCurrent > damagedCountTarget {
		return 0
	}

	// find next '?'
	found := false
	for i := position; i < len(conditions); i++ {
		if conditions[i] == '?' {
			position = i
			found = true
			break
		}
	}

	if found {
		sum := 0
		// try '#'
		conditions[position] = '#'
		sum += calculateArrangementsCountMutable(position+1, conditions, groupSizes, damagedCountCurrent+1, damagedCountTarget)

		// try '.'
		conditions[position] = '.'
		sum += calculateArrangementsCountMutable(position+1, conditions, groupSizes, damagedCountCurrent, damagedCountTarget)

		// revert
		conditions[position] = '?'

		return sum
	}

	// no '?' found -> check validity
	if isValid(conditions, groupSizes) {
		//fmt.Printf("%v\n", string(conditions))
		return 1
	}

	// invalid
	return 0
}

func isValid(conditions []rune, groupSizes []int) bool {
	currentGroupSize := 0
	currentGroup := 0
	previous := '.'

	for i, current := range conditions {
		// increase group
		if current == '#' {
			currentGroupSize++
		}

		// group end
		last := i == len(conditions)-1
		if previous == '#' && (current == '.') || current == '#' && last {
			// found more groups
			if currentGroup >= len(groupSizes) {
				return false
			}

			// group size is different -> not valid
			if groupSizes[currentGroup] != currentGroupSize {
				return false
			}

			// group size is same -> continue
			currentGroup++
			currentGroupSize = 0
		}

		previous = current
	}

	// found less groups
	if currentGroup != len(groupSizes) {
		return false
	}

	return true
}

func DoWithInputPart02(world World) int {
	sum := 0

	results := utils.ProcessParallel(world.Records, calculateArrangementsCountUnfolded)

	for result := range results {
		sum += result.Value
	}

	return sum
}

func calculateArrangementsCountUnfolded(record Record, i int) int {
	fmt.Printf("#%d: ?'s count: %d\n", i, record.Unknowns)
	singleCount := calculateArrangementsCount(record)

	unfolded2 := Unfold2(record)
	fmt.Printf("#%d: ?'s count: %d\n", i, unfolded2.Unknowns)
	pairCount := calculateArrangementsCount(unfolded2)

	fmt.Println()
	k := pairCount / singleCount
	return pairCount * k * k * k
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
