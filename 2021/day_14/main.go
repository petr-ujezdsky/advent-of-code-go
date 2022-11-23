package day_14

import (
	"bufio"
	"io"
	"math"
	"strings"
)

const alphabetSize = int('Z' - 'A')

type World struct {
	template string
	rules    map[string]string
}

type WorldRunes struct {
	template []rune
	// index is hash of duo
	rules []rune
}

func PolymerScore(polymer string) int {
	counts := make(map[rune]int)

	for _, char := range []rune(polymer) {
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

func GrowPolymerStepIter(template string, rules map[string]string) string {
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

func GrowPolymerIter(template string, rules map[string]string, stepsCount int) string {
	for i := 0; i < stepsCount; i++ {
		//fmt.Printf("Growing polymer #%d\n", i)
		template = GrowPolymerStepIter(template, rules)
	}

	return template
}

func growPolymerRecursive(duo string, rules map[string]string, counts map[rune]int, depth int) {
	if depth > 0 {
		newChar, contains := rules[duo]
		if contains {
			newCharRune := []rune(newChar)[0]

			counts[newCharRune]++
			duoRunes := []rune(duo)
			// left + new
			growPolymerRecursive(string([]rune{duoRunes[0], newCharRune}), rules, counts, depth-1)

			// new + right
			growPolymerRecursive(string([]rune{newCharRune, duoRunes[1]}), rules, counts, depth-1)
		}
	}
}

func GrowPolymerRecursive(template string, rules map[string]string, stepsCount int) int {
	counts := make(map[rune]int)

	// counts from init template
	for _, char := range []rune(template) {
		counts[char]++
	}

	for i := 0; i < len(template)-1; i++ {
		duo := template[i : i+2]
		growPolymerRecursive(duo, rules, counts, stepsCount)
	}

	return scoreFromCounts(counts)
}

func growPolymerRecursiveRune(duo []rune, rules []rune, counts map[rune]int, depth int) {
	if depth > 0 {
		hash := duoHash(duo)
		newChar := rules[hash]

		if newChar > 0 {
			counts[newChar]++
			// left + new
			growPolymerRecursiveRune([]rune{duo[0], newChar}, rules, counts, depth-1)

			// new + right
			growPolymerRecursiveRune([]rune{newChar, duo[1]}, rules, counts, depth-1)
		}
	}
}

func GrowPolymerRecursiveRune(template []rune, rules []rune, stepsCount int) int {
	counts := make(map[rune]int)

	// counts from init template
	for _, char := range template {
		counts[char]++
	}

	for i := 0; i < len(template)-1; i++ {
		duo := template[i : i+2]
		growPolymerRecursiveRune(duo, rules, counts, stepsCount)
	}

	return scoreFromCounts(counts)
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

func duoHash(duo []rune) int {
	return int(duo[0]-'A')*alphabetSize + int(duo[1]-'A')
}

func Runify(world World) WorldRunes {
	template := []rune(world.template)

	rules := make([]rune, alphabetSize*alphabetSize+1)

	for left, right := range world.rules {
		hash := duoHash([]rune(left))

		rules[hash] = rune(right[0])
	}

	return WorldRunes{template, rules}
}