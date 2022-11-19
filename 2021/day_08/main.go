package day_07

import (
	"bufio"
	"fmt"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"io"
	"math"
	"strings"
)

// Segment indexes:
//
//  aaaa
// b    c
// b    c
//  dddd
// e    f
// e    f
//  gggg

var mask2digit = createMask2DigitMapper()

func createMask2DigitMapper() map[int]int {
	mask2digit := make(map[int]int)
	// identity decoder
	//               a, b, c, d, e, f, g
	decoder := []int{0, 1, 2, 3, 4, 5, 6}

	mask2digit[createSegmentMask(decoder, "abcefg")] = 0
	mask2digit[createSegmentMask(decoder, "cf")] = 1
	mask2digit[createSegmentMask(decoder, "acdeg")] = 2
	mask2digit[createSegmentMask(decoder, "acdfg")] = 3
	mask2digit[createSegmentMask(decoder, "bcdf")] = 4
	mask2digit[createSegmentMask(decoder, "abdfg")] = 5
	mask2digit[createSegmentMask(decoder, "abdefg")] = 6
	mask2digit[createSegmentMask(decoder, "acf")] = 7
	mask2digit[createSegmentMask(decoder, "abcdefg")] = 8
	mask2digit[createSegmentMask(decoder, "abcdfg")] = 9

	return mask2digit
}

// Entry
// List of digits and their segments count, * marks unique
// 0 ~ 6 segments
// 1 ~ 2 segments *
// 2 ~ 5 segments
// 3 ~ 5 segments
// 4 ~ 4 segments *
// 5 ~ 5 segments
// 6 ~ 6 segments
// 7 ~ 3 segments *
// 8 ~ 7 segments *
// 9 ~ 6 segments
type Entry struct {
	Digits  []string
	Outputs []string
}

func CountEasyOutputs(entries []Entry) int {
	// easy digit lengths:
	// 1 ~ 2 segments
	// 4 ~ 4 segments
	// 7 ~ 3 segments
	// 8 ~ 7 segments
	count := 0

	for _, entry := range entries {
		for _, output := range entry.Outputs {
			length := len(output)

			if length == 2 || length == 4 || length == 3 || length == 7 {
				count++
			}
		}
	}

	return count
}

func letterToIndex(char rune) int {
	// letter to number, a -> 0, g -> 6
	return int(char) - int('a')
}

// Each value means segment index
// "acb" -> [0, 2, 1]
// a means segment #1 (index 0)
// c means segment #2 (index 1)
// b means segment #3 (index 2)
// ...
// The resulting array (arr) is reverse mapping:
// arr["a"] = arr[0] = 0 (segment #1)
// arr["b"] = arr[1] = 2 (segment #3)
// arr["c"] = arr[2] = 1 (segment #2)
func createSegmentDecoder(mapping []rune) []int {
	mapper := make([]int, 7)

	for i, char := range mapping {
		mapper[letterToIndex(char)] = i
	}

	return mapper
}

func createSegmentMask(decoder []int, segments string) int {
	mask := 0
	for _, segmentEncoded := range []rune(segments) {
		// decode digit to segment index using mapper
		segmentIndex := decoder[letterToIndex(segmentEncoded)]

		// sum segment masks
		mask += 1 << segmentIndex
	}
	return mask
}

func decodeDigits(decoder []int, digits []string) (int, bool) {
	number := 0
	for i, digitEncodedSegments := range digits {
		digitSegmentsMask := createSegmentMask(decoder, digitEncodedSegments)

		// lookup digit by mask
		digit, contains := mask2digit[digitSegmentsMask]
		if !contains {
			return 0, false
		}

		number += digit * int(math.Pow10(len(digits)-i-1))
	}

	return number, true
}

func TryDecodeOutput(mapping []rune, entry Entry) (int, bool) {
	// create mapping from letter to segment index
	decoder := createSegmentDecoder(mapping)

	_, ok := decodeDigits(decoder, entry.Digits)
	if !ok {
		return 0, false
	}

	output, ok := decodeDigits(decoder, entry.Outputs)
	if !ok {
		return 0, false
	}

	return output, true
}

func BruteForceDecode(entry Entry) (int, []rune, int, bool) {
	quit := make(chan interface{})
	initialDecoder := []rune("abcdefg")
	decoders := utils.Permute(quit, initialDecoder)

	i := 0
	for decoder := range decoders {
		output, ok := TryDecodeOutput(decoder, entry)
		if ok {
			close(quit)
			return output, decoder, i, true
		}
		i++
	}

	return 0, nil, i, false
}

func DecodeAndSum(entries []Entry) (int, bool) {
	sum := 0
	for i, entry := range entries {
		output, decoder, iterations, ok := BruteForceDecode(entry)
		if !ok {
			return 0, false
		}
		fmt.Println(i, "decoder:", string(decoder), "iterations:", iterations)

		sum += output
	}

	return sum, true
}

func ParseInput(r io.Reader) ([]Entry, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	var entries []Entry

	// example line
	// be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb | fdgacbe cefdb cefbgd gcbe
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " | ")

		digits := strings.Split(parts[0], " ")
		output := strings.Split(parts[1], " ")

		entries = append(entries, Entry{digits, output})
	}

	return entries, scanner.Err()
}
