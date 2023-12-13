package main

import (
	_ "embed"
	"fmt"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/parsers"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/slices"
	"io"
	"regexp"
	"strconv"
	"strings"
)

var regexDots = regexp.MustCompile(`[.]+`)

type Record struct {
	ConditionsRaw string
	Conditions    []rune
	Groups        [][]rune
	GroupSizes    []int
	Unknowns      int
}

type World struct {
	Records []Record
}

func DoWithInputPart01(world World) int {
	sum := 0

	for _, record := range world.Records {
		sum += calculateArrangementsCount(record)
	}

	return sum
}

func calculateArrangementsCount(record Record) int {
	return calculateArrangementsCountMutable(0, record.Conditions, record.GroupSizes)
}

func calculateArrangementsCountMutable(position int, conditions []rune, groupSizes []int) int {
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
		// try '.'
		conditions[position] = '#'
		sum += calculateArrangementsCountMutable(position+1, conditions, groupSizes)

		// try '#'
		conditions[position] = '.'
		sum += calculateArrangementsCountMutable(position+1, conditions, groupSizes)

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

	for i, record := range world.Records {
		sum += calculateArrangementsCountUnfolded(i, record)
	}

	return sum
}

func calculateArrangementsCountUnfolded(i int, record Record) int {
	count := 0
	for _, condition := range record.Conditions {
		if condition == '?' {
			count++
		}
	}

	fmt.Printf("#%d: ?'s count: %d\n", i, record.Unknowns)
	singleCount := calculateArrangementsCount(record)

	unfolded2 := Unfold2(record)
	fmt.Printf("#%d: ?'s count: %d\n\n", i, unfolded2.Unknowns)
	pairCount := calculateArrangementsCount(unfolded2)

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

	unknowns := 0
	for _, condition := range conditionsRaw {
		if condition == '?' {
			unknowns++
		}
	}

	return Record{
		ConditionsRaw: conditionsRaw,
		Conditions:    []rune(conditionsRaw),
		GroupSizes:    groupSizes,
		Unknowns:      unknowns,
	}
}

func ParseInput(r io.Reader) World {
	items := parsers.ParseToObjects(r, ParseRecord)
	return World{Records: items}
}
