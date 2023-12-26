package main

import (
	_ "embed"
	"fmt"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/matrix"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/measure"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/parsers"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/slices"
	"io"
	"regexp"
	"strconv"
	"strings"
)

var regexMultiDots = regexp.MustCompile(`[.]+`)

type Record struct {
	Raw                 string
	ConditionsRaw       string
	Conditions          []rune
	ConditionGroups     []string
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
		count := calculateArrangementsCount2(record)
		fmt.Printf("#%3d combinations: %-6d %s\n", i+1, count, record.ConditionsRaw)
		max = utils.Max(max, count)
		sum += count
	}

	fmt.Printf("\nMax combinations: %d\n", max)

	return sum
}

func calculateArrangementsCount(record Record) int {
	return calculateArrangementsCountInner(record.Conditions, record.GroupSizes)
}

func calculateArrangementsCountInner(conditions []rune, groupSizes []int) int {
	//defer measure.Duration(measure.Track(fmt.Sprintf("Optimized calculation for %d unknowns took", record.Unknowns)))
	//defer measure.Duration(measure.Track("Optimized calculation took"))

	sum := 0

	cache := matrix.NewMatrixInt(len(conditions), len(groupSizes))
	cache.SetAll(-1)

	// try to offset first group by N .
	for i := -1; i < len(conditions); i++ {
		// check . separator
		if i >= 0 && conditions[i] == '#' {
			break
		}

		count, canShift := calculateArrangementsCountRecursive(conditions[i+1:], groupSizes, cache)
		sum += count
		if !canShift {
			break
		}
	}

	return sum
}

func calculateArrangementsCountRecursive(conditions []rune, groupSizes []int, cache matrix.MatrixNumber[int]) (int, bool) {
	groupSize := groupSizes[0]
	isLastGroup := len(groupSizes) == 1

	if isLastGroup {
		// not enough remaining conditions for the group
		if groupSize > len(conditions) {
			return 0, false
		}
	} else {
		// not enough remaining conditions for the group
		if groupSize+1 > len(conditions) {
			return 0, false
		}
	}

	// check # group starting at 0
	for i := 0; i < groupSize; i++ {
		if conditions[i] == '.' {
			return 0, true
		}
	}

	sum := 0
	if isLastGroup {
		// expect ending is OK
		sum = 1

		for i := groupSize; i < len(conditions); i++ {
			// check ending . after final group
			if conditions[i] == '#' {
				// ending is not ok
				sum = 0
				break
			}
		}

	} else {
		// try to offset next group by N .
		for i := groupSize; i < len(conditions); i++ {
			// check . separator after group
			if conditions[i] == '#' {
				break
			}

			count, canShift := calculateArrangementsCountRecursive(conditions[i+1:], groupSizes[1:], cache)
			sum += count
			if !canShift {
				break
			}
		}
	}

	return sum, true
}

func calculateArrangementsCount2(record Record) int {
	//countExpected := calculateArrangementsCount(record)
	//
	cache := make(map[key]int)
	countQuick := calculateArrangementsCountRecursive2(record.ConditionGroups, record.GroupSizes, cache)
	//
	//if countExpected != countQuick {
	//	fmt.Printf("      Different, %d != %d\n", countExpected, countQuick)
	//}

	return countQuick
}

func calculateArrangementsCountRecursive2(conditionGroups []string, groupSizes []int, cache map[key]int) int {
	conditions := conditionGroups[0]

	isLastCondition := len(conditionGroups) == 1

	if isLastCondition {
		count := calculateArrangementsCountGroupsCached(conditions, groupSizes, cache)
		return count
	}

	sum := 0
	for groupsCount := 0; groupsCount <= len(groupSizes); groupsCount++ {
		subCount := calculateArrangementsCountGroupsCached(conditions, groupSizes[0:groupsCount], cache)
		if subCount == 0 {
			//if groupsCount == 0 {
			continue
			//}
			//
			//return sum
		}

		k := calculateArrangementsCountRecursive2(conditionGroups[1:], groupSizes[groupsCount:], cache)
		subCount *= k

		sum += subCount
	}

	return sum
}

func calculateArrangementsCountGroupsCached(conditions string, groupSizes []int, cache map[key]int) int {
	k := newKey(conditions, groupSizes)

	if count, ok := cache[k]; ok {
		//fmt.Printf("Cache hit!\n")
		return count
	}

	count := calculateArrangementsCountGroups(conditions, groupSizes)
	cache[k] = count

	return count
}
func calculateArrangementsCountGroups(conditions string, groupSizes []int) int {
	if len(groupSizes) == 0 {
		if OnlyDots(conditions) {
			return 1
		}

		return 0
	}

	return calculateArrangementsCountInner([]rune(conditions), groupSizes)
}

type key struct {
	conditions     string
	groupSizesHash uint64
}

func newKey(conditions string, groupSizes []int) key {
	hash := uint64(0)
	for _, size := range groupSizes {
		// size max value is 15, so move by 2^4 = 16 (max 16 group sizes)
		// max 6 group sizes in original -> max 5*6=30 in total
		hash <<= 4
		hash += uint64(size)
	}

	return key{
		conditions:     conditions,
		groupSizesHash: hash,
	}
}

func OnlyDots(conditions string) bool {
	for _, condition := range conditions {
		if condition == '#' {
			return false
		}
	}

	return true
}

func OnlySharps(conditions string) bool {
	for _, condition := range conditions {
		if condition == '.' {
			return false
		}
	}

	return true
}

func DoWithInputPart02(world World) int {
	sum := 0
	count := 0

	results := utils.ProcessParallel(world.Records, calculateArrangementsCountUnfolded5)
	//results := utils.ProcessSerial(world.Records, calculateArrangementsCountUnfolded5)

	for result := range results {
		sum += result.Value
		count++
		//fmt.Printf("Done %d / %d (%.2f%%)\n", count, len(world.Records), float64(100*count)/float64(len(world.Records)))
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
	defer measure.Duration(measure.Track(fmt.Sprintf("#%3d (%30s)", i+1, record.Raw)))

	count1 := calculateArrangementsCount(record)
	count2 := calculateArrangementsCount(Unfold(record, 2))

	if count2 == count1*count1 {
		count5 := count1 * count1 * count1 * count1 * count1
		fmt.Printf("#%3d: QUICK count5: %9d\n", i+1, count5)

		return count5
	}

	count5 := calculateArrangementsCount2(Unfold(record, 5))
	fmt.Printf("#%3d: SLOW  count5: %9d\n", i+1, count5)
	return count5
}

//Unfold 3x
//#145: 658.952417ms
//#287: 792.577042ms
//#408: 2.262375208s
//#775: 767.091333ms

// #323: 201.515667ms (.?.?????#?????.???? 1,6,1,2,1)
// #663: 205.578209ms (?????#????#??????? 5,1,1,1,1,1)
// #776: 215.4115ms (?????????.???? 1,1,3,1)
// #288: 314.184958ms (??.??????.??#?#????? 1,2,1,5,1,1)
// #146: 331.222709ms (.???????.#.????????? 4,1,1,1,3,1)
// #409: 721.529916ms (????????????? 1,1,1,2)
func calculateArrangementsCountUnfolded5(record Record, i int) int {
	//start := time.Now()
	//count := calculateArrangementsCount(Unfold(record, 3))
	//
	//duration := time.Since(start)
	//if duration.Milliseconds() > 200 {
	//	fmt.Printf("#%3d: %v (%v)\n", i+1, duration, record.Raw)
	//}
	//return count

	record = Unfold(record, 3)

	countExpected := calculateArrangementsCount(record)
	countActual := calculateArrangementsCount(record)

	if countActual != countExpected {
		fmt.Printf("#%3d: %9d / %-9d\n", i+1, countExpected, countActual)
	}

	return countActual
}

func calculateArrangementsCountUnfolded2(record Record, i int) int {
	unfolded := Unfold(record, 5)
	return calculateArrangementsCount(unfolded)
}

func Unfold(record Record, count int) Record {
	if count == 1 {
		return record
	}

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

	// trim '.'
	conditionsRawTrimmed := strings.Trim(conditionsRaw, ".")
	// replace multiple '.' by single '.'
	conditionsRawSimplified := regexMultiDots.ReplaceAllLiteralString(conditionsRawTrimmed, ".")

	conditionGroups := strings.Split(conditionsRawSimplified, ".")
	//conditionGroups := slices.Map(conditionGroupsStr, func(s string) []rune { return []rune(s) })

	//fmt.Printf("%-20v -> %-20v | %v\n", conditionsRaw, conditionsRawSimplified, conditionGroupsStr)

	return Record{
		Raw:                 str,
		ConditionsRaw:       conditionsRaw,
		Conditions:          []rune(conditionsRaw),
		ConditionGroups:     conditionGroups,
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
