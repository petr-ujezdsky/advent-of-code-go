package day_14

import (
	"bufio"
	"io"
	"math"
	"strings"
)

const alphabetSize = int('Z' - 'A')

type World struct {
	template []rune
	// index is hash of duo
	rules []rune
}

func growPolymerRecursiveRuneCaching(duo []rune, rules []rune, counts map[rune]int, countsCache []map[rune]int, depth int) {
	if depth > 0 {
		countsHash := duoHeightHash(duo, depth)

		cachedCounts := countsCache[countsHash]

		if cachedCounts == nil {
			cachedCounts = make(map[rune]int)

			hash := duoHash(duo)
			newChar := rules[hash]

			if newChar > 0 {
				cachedCounts[newChar]++
				// left + new
				growPolymerRecursiveRuneCaching([]rune{duo[0], newChar}, rules, cachedCounts, countsCache, depth-1)

				// new + right
				growPolymerRecursiveRuneCaching([]rune{newChar, duo[1]}, rules, cachedCounts, countsCache, depth-1)
			}

			// store cached counts
			countsCache[countsHash] = cachedCounts
		} else {
			//fmt.Println("Cache hit!")
		}

		// merge counts
		mergeCounts(cachedCounts, counts)
	}
}

func GrowPolymerRecursiveRuneCaching(template []rune, rules []rune, stepsCount int) int {
	counts := make(map[rune]int)

	countsCache := make([]map[rune]int, (stepsCount+1)*alphabetSize*alphabetSize)

	for i := 0; i < len(template)-1; i++ {
		duo := template[i : i+2]
		growPolymerRecursiveRuneCaching(duo, rules, counts, countsCache, stepsCount)
	}

	// counts from init template
	for _, char := range template {
		counts[char]++
	}

	return scoreFromCounts(counts)
}

func scoreFromCounts(counts map[rune]int) int {
	minCount := math.MaxInt
	maxCount := math.MinInt
	for _, count := range counts {
		if count > maxCount {
			maxCount = count
		}

		if count < minCount {
			minCount = count
		}
	}

	return maxCount - minCount
}

func mergeCounts(source, target map[rune]int) {
	for char, count := range source {
		target[char] += count
	}
}

func duoHash(duo []rune) int {
	return int(duo[0]-'A')*alphabetSize + int(duo[1]-'A')
}

func duoHeightHash(duo []rune, height int) int {
	return height*alphabetSize*alphabetSize + int(duo[0]-'A')*alphabetSize + int(duo[1]-'A')
}

func ParseInput(r io.Reader) (World, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	scanner.Scan()
	template := []rune(scanner.Text())

	rules := make([]rune, alphabetSize*alphabetSize+1)

	scanner.Scan()

	for scanner.Scan() {
		ruleParts := strings.Split(scanner.Text(), " -> ")

		left := []rune(ruleParts[0])
		right := []rune(ruleParts[1])

		hash := duoHash(left)

		rules[hash] = right[0]
	}

	return World{template, rules}, scanner.Err()
}
