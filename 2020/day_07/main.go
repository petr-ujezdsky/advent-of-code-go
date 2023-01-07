package main

import (
	_ "embed"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/parsers"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/slices"
	"io"
	"regexp"
	"strings"
)

type BagRule struct {
	Color        string
	NeededCounts map[string]int
}

type ExpandedNeededCounts = map[string]int

func multiplyCounts(amount int, neededCounts, expandedCounts ExpandedNeededCounts) {
	for color := range neededCounts {
		expandedCounts[color] += neededCounts[color] * amount
	}
}

func expandRules(color string, rules map[string]BagRule, cache map[string]ExpandedNeededCounts) ExpandedNeededCounts {
	if expandedRule, ok := cache[color]; ok {
		return expandedRule
	}

	rule := rules[color]
	expandedCounts := make(ExpandedNeededCounts)
	for subColor, neededCount := range rule.NeededCounts {
		subRuleNeededCounts := expandRules(subColor, rules, cache)

		multiplyCounts(neededCount, subRuleNeededCounts, expandedCounts)
		expandedCounts[subColor] += neededCount
	}

	cache[color] = expandedCounts

	return expandedCounts
}

func ExpandRules(color string, rules map[string]BagRule) ExpandedNeededCounts {
	expandedNeededCounts := expandRules(color, rules, make(map[string]ExpandedNeededCounts))
	return expandedNeededCounts
}

func ExpandAllRules(rules map[string]BagRule) map[string]ExpandedNeededCounts {
	allExpandedNeededCounts := make(map[string]ExpandedNeededCounts)

	for color := range rules {
		expandRules(color, rules, allExpandedNeededCounts)
	}

	return allExpandedNeededCounts
}

func ContainableByBagsCount(bagRules map[string]BagRule) int {
	allExpandedNeededCounts := ExpandAllRules(bagRules)

	count := 0
	for color := range bagRules {
		expandedNeededCounts := allExpandedNeededCounts[color]
		if neededCount, ok := expandedNeededCounts["shiny gold"]; ok && neededCount >= 1 {
			count++
		}
	}

	return count
}

func TotalBagsCount(bagRules map[string]BagRule) int {
	allExpandedNeededCounts := ExpandAllRules(bagRules)

	count := 0
	expandedNeededCounts := allExpandedNeededCounts["shiny gold"]
	for _, neededCount := range expandedNeededCounts {
		count += neededCount
	}

	return count
}

// bright white bags contain 1 shiny gold bag.
// light red bags contain 1 bright white bag, 2 muted yellow bags.
var regexRules = regexp.MustCompile(`^(.+) bags contain (.+)\.$`)
var regexRule = regexp.MustCompile(`(\d+) (.+) bags?`)

func ParseRule(str string) BagRule {
	parts := regexRules.FindStringSubmatch(str)

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

	return BagRule{
		Color:        color,
		NeededCounts: neededCounts,
	}
}

func ParseInput(r io.Reader) map[string]BagRule {
	bagRules := parsers.ParseToObjects(r, ParseRule)

	return slices.ToMap[BagRule](bagRules, func(barRule BagRule) string { return barRule.Color })
}
