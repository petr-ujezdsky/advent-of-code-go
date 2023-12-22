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
	//defer measure.Duration(measure.Track(fmt.Sprintf("Optimized calculation for %d unknowns took", record.Unknowns)))
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
	count := 0

	results := utils.ProcessParallel(world.Records, calculateArrangementsCountUnfolded4)
	//results := utils.ProcessSerial(world.Records, calculateArrangementsCountUnfolded3)

	for result := range results {
		sum += result.Value
		count++
		fmt.Printf("Done %d / %d (%.2f%%)\n", count, len(world.Records), float64(100*count)/float64(len(world.Records)))
	}

	return sum
}

func calculateArrangementsCountUnfolded(record Record, i int) int {
	if record.ConditionsRaw[0] != '.' || record.ConditionsRaw[len(record.ConditionsRaw)-1] != '.' {
		return 0
	}

	count1 := calculateArrangementsCount(record)

	unfolded2 := Unfold(record, 2)
	count2 := calculateArrangementsCount(unfolded2)

	unfolded3 := Unfold(record, 3)
	count3 := calculateArrangementsCount(unfolded3)

	unfolded4 := Unfold(record, 4)
	count4 := calculateArrangementsCount(unfolded4)

	//fmt.Println()
	k := count2 / count1
	fivesCount := count2 * k * k * k

	count3Quick := count2 * k
	count4Quick := count2 * k * k

	warning := ""
	if count3 != count3Quick {
		warning = "*"
	}

	//??.?.???.?#?# 1,2,3
	//#  0: unknowns:  8, 1x:    6, 2x:    48, 3x:      408/384, 4x:     3504/3072   *

	//if count2%count1 != 0 {
	//
	//}

	fmt.Printf("#%3d: unknowns: %2d, 1x: %4d, 2x: %5d, 3x: %8d/%d, 4x: %8d/%d%s\n", i, record.Unknowns, count1, count2, count3, count3Quick, count4, count4Quick, warning)
	fmt.Printf("#%3d: 2%%1: %4d, 3%%1x: %5d, 4%%1x: %8d%4s %-30s\n", i, count2%count1, count3%count1, count4%count1, warning, record.ConditionsRaw)

	return fivesCount
}

func calculateArrangementsCountUnfolded3(record Record, i int) int {

	countOrig := calculateArrangementsCount(record)

	countAppended := calculateArrangementsCount(Append(record))
	countPrepended := calculateArrangementsCount(Prepend(record))
	countBoth := calculateArrangementsCount(Append(Prepend(record)))
	count2 := calculateArrangementsCount(Unfold(record, 2))

	warningAppended := ""
	if countOrig != countAppended {
		warningAppended = "*"
	}

	warningPrepended := ""
	if countOrig != countPrepended {
		warningPrepended = "*"
	}

	warningBoth := ""
	if countOrig != countBoth {
		warningBoth = "*"
	}

	warningPair := ""
	if countOrig*countOrig != count2 {
		warningPair = "*"
	}

	extra2 := count2 - countOrig*countOrig

	count3 := calculateArrangementsCount(Unfold(record, 3))
	count3Quick := count2*countOrig + extra2

	extra3 := count3 - count2*countOrig

	if extra2 == 0 && extra3 != 0 {
		fmt.Printf("#%3d: orig: %4d, __?: %4d%1s ?__: %4d%1s ?_?: %4d%1s _?_: %6d%1s extra: %6d 3x: %8d / %-8d extra3: %8d %-30s %v\n", i, countOrig, countAppended, warningAppended, countPrepended, warningPrepended, countBoth, warningBoth, count2, warningPair, extra2, count3, count3Quick, extra3, record.ConditionsRaw, record.GroupSizes)
	}
	return countOrig
}

func calculateArrangementsCountUnfolded4(record Record, i int) int {
	count1 := calculateArrangementsCount(record)
	count2 := calculateArrangementsCount(Unfold(record, 2))

	if count2 == count1*count1 {
		count5 := count1 * count1 * count1 * count1 * count1
		fmt.Printf("#%3d: QUICK count5: %9d\n", i, count5)

		return count5
	}

	count5 := calculateArrangementsCount(Unfold(record, 5))
	fmt.Printf("#%3d: SLOW  count5: %9d\n", i, count5)
	return count5
}

func calculateArrangementsCountUnfolded2(record Record, i int) int {
	unfolded := Unfold(record, 5)
	return calculateArrangementsCount(unfolded)
}

func Unfold(record Record, count int) Record {
	conditionsRawJoined := strings.Join(slices.Repeat([]string{record.ConditionsRaw}, count), "?")
	groupSizes := slices.Repeat(record.GroupSizes, count)

	// convert groupSizes ints to string
	groupSizesStr := slices.Map(groupSizes, strconv.Itoa)
	groupSizesStrJoined := strings.Join(groupSizesStr, ",")

	// join to raw record string
	recordStr := conditionsRawJoined + " " + groupSizesStrJoined

	// parse
	return ParseRecord(recordStr)
}

func Append(record Record) Record {
	conditionsRawJoined := record.ConditionsRaw + "?"

	// convert groupSizes ints to string
	groupSizesStr := slices.Map(record.GroupSizes, strconv.Itoa)
	groupSizesStrJoined := strings.Join(groupSizesStr, ",")

	// join to raw record string
	recordStr := conditionsRawJoined + " " + groupSizesStrJoined

	// parse
	return ParseRecord(recordStr)
}

func Prepend(record Record) Record {
	conditionsRawJoined := "?" + record.ConditionsRaw

	// convert groupSizes ints to string
	groupSizesStr := slices.Map(record.GroupSizes, strconv.Itoa)
	groupSizesStrJoined := strings.Join(groupSizesStr, ",")

	// join to raw record string
	recordStr := conditionsRawJoined + " " + groupSizesStrJoined

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
