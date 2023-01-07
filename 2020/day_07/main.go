package main

import (
	"bufio"
	_ "embed"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"io"
	"regexp"
	"strings"
)

type BagRule struct {
	Color        string
	NeededCounts map[string]int
}

func DoWithInput(items []BagRule) int {
	return len(items)
}

// bright white bags contain 1 shiny gold bag.
// light red bags contain 1 bright white bag, 2 muted yellow bags.
var regexRules = regexp.MustCompile(`^(.+) bags contain (.+)\.$`)
var regexRule = regexp.MustCompile(`(\d+) (.+) bags?`)

func ParseInput(r io.Reader) []BagRule {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	var bagRules []BagRule
	for scanner.Scan() {
		parts := regexRules.FindStringSubmatch(scanner.Text())

		color := parts[1]

		rulesParts := strings.Split(parts[2], ", ")

		neededCounts := make(map[string]int)

		for _, rule := range rulesParts {
			ruleParts := regexRule.FindStringSubmatch(rule)
			if len(ruleParts) != 3 {
				continue
			}
			ruleCount := utils.ParseInt(ruleParts[1])
			ruleColor := ruleParts[2]

			neededCounts[ruleColor] = ruleCount
		}

		bagRule := BagRule{
			Color:        color,
			NeededCounts: neededCounts,
		}

		bagRules = append(bagRules, bagRule)
	}

	return bagRules
}
