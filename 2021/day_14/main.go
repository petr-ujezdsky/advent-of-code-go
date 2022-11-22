package day_14

import (
	"bufio"
	"io"
	"math"
	"strings"
)

type World struct {
	template string
	rules    map[string]string
}

func PolymerScore(polymer string) int {
	counts := make(map[rune]int)

	for _, char := range []rune(polymer) {
		counts[char]++
	}

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

func GrowPolymerStep(template string, rules map[string]string) string {
	var polymer strings.Builder

	for i := 0; i < len(template)-1; i++ {
		duo := template[i : i+2]

		polymer.WriteString(string(duo[0]))

		newChar, contains := rules[duo]
		if contains {
			polymer.WriteString(newChar)
		} else {
			polymer.WriteString(string(duo[1]))
		}
	}

	// last char
	polymer.WriteString(string(template[len(template)-1]))

	return polymer.String()
}

func GrowPolymer(template string, rules map[string]string, stepsCount int) string {
	for i := 0; i < stepsCount; i++ {
		template = GrowPolymerStep(template, rules)
	}

	return template
}

func ParseInput(r io.Reader) (World, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	scanner.Scan()
	template := scanner.Text()

	rules := make(map[string]string)
	scanner.Scan()

	for scanner.Scan() {
		ruleParts := strings.Split(scanner.Text(), " -> ")
		rules[ruleParts[0]] = ruleParts[1]
	}

	return World{template, rules}, scanner.Err()
}
