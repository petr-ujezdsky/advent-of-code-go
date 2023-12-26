package main

import (
	_ "embed"
	"fmt"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/matrix"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/parsers"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/slices"
	"io"
	"regexp"
	"strconv"
	"strings"
)

var regexMultiDots = regexp.MustCompile(`[.]+`)

type Record struct {
	Raw           string
	ConditionsRaw string
	Conditions    []rune
	GroupSizes    []int
}

type World struct {
	Records []Record
}

func DoWithInputPart01(world World) int {
	sum := 0

	max := 0
	for i, record := range world.Records {
		count := calculateArrangementsCount(record)
		fmt.Printf("#%3d combinations: %-6d %s\n", i+1, count, record.ConditionsRaw)
		max = utils.Max(max, count)
		sum += count
	}

	fmt.Printf("\nMax combinations: %d\n", max)

	return sum
}

func calculateArrangementsCount(record Record) int {
	//defer measure.Duration(measure.Track(fmt.Sprintf("Optimized calculation for %d unknowns took", record.Unknowns)))
	//defer measure.Duration(measure.Track("Optimized calculation took"))

	conditions := record.Conditions
	groupSizes := record.GroupSizes

	sum := 0

	cache := matrix.NewMatrix[result](len(conditions)+1, len(groupSizes)+1)
	cache.SetAll(result{-1, false})

	// try to offset first group by N .
	for i := -1; i < len(conditions); i++ {
		// check . separator
		if i >= 0 && conditions[i] == '#' {
			break
		}

		count, canShift := calculateArrangementsCountRecursiveCaching(conditions[i+1:], groupSizes, cache)
		sum += count
		if !canShift {
			break
		}
	}

	return sum
}

type result struct {
	count    int
	canShift bool
}

func calculateArrangementsCountRecursiveCaching(conditions []rune, groupSizes []int, cache matrix.Matrix[result]) (int, bool) {
	cacheKey := utils.Vector2i{X: len(conditions), Y: len(groupSizes)}

	r := cache.GetV(cacheKey)
	if r.count != -1 {
		// cache hit
		return r.count, r.canShift
	}

	// cache miss, compute
	count, canShift := calculateArrangementsCountRecursive(conditions, groupSizes, cache)
	cache.SetV(cacheKey, result{count, canShift})
	return count, canShift
}

func calculateArrangementsCountRecursive(conditions []rune, groupSizes []int, cache matrix.Matrix[result]) (int, bool) {
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

			count, canShift := calculateArrangementsCountRecursiveCaching(conditions[i+1:], groupSizes[1:], cache)
			sum += count
			if !canShift {
				break
			}
		}
	}

	return sum, true
}

func DoWithInputPart02(world World) int {
	sum := 0
	count := 0

	results := utils.ProcessParallel(world.Records, calculateArrangementsCountUnfolded)
	//results := utils.ProcessSerial(world.Records, calculateArrangementsCountUnfolded)

	for result := range results {
		sum += result.Value
		count++
		//fmt.Printf("Done %d / %d (%.2f%%)\n", count, len(world.Records), float64(100*count)/float64(len(world.Records)))
	}

	return sum
}

func calculateArrangementsCountUnfolded(record Record, i int) int {
	unfolded5 := Unfold(record, 5)
	return calculateArrangementsCount(unfolded5)
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

func ParseRecord(str string) Record {
	parts := strings.Split(str, " ")

	conditionsRaw := parts[0]
	groupSizes := utils.ExtractInts(parts[1], false)

	return Record{
		Raw:           str,
		ConditionsRaw: conditionsRaw,
		Conditions:    []rune(conditionsRaw),
		GroupSizes:    groupSizes,
	}
}

func ParseInput(r io.Reader) World {
	items := parsers.ParseToObjects(r, ParseRecord)
	return World{Records: items}
}
