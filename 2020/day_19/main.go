package main

import (
	"bufio"
	_ "embed"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"io"
	"strings"
)

type MessageValidator interface {
	Valid(pos int, message string) (bool, int)
}

type AndRule struct {
	Rules []*MessageValidatorHolder
}

func (r AndRule) Valid(pos int, message string) (bool, int) {
	currentPos := pos
	for _, rule := range r.Rules {
		ok, nextPos := rule.Validator.Valid(currentPos, message)
		if !ok {
			return false, nextPos
		}
		currentPos = nextPos
	}

	return true, currentPos
}

type OrRule struct {
	Left, Right MessageValidator
}

func (r OrRule) Valid(pos int, message string) (bool, int) {
	ok, nextPos := r.Left.Valid(pos, message)
	if ok {
		return ok, nextPos
	}

	ok, nextPos = r.Right.Valid(pos, message)
	return ok, nextPos
}

type ValueRule struct {
	Value uint8
}

func (r ValueRule) Valid(pos int, message string) (bool, int) {
	if pos >= len(message) {
		return false, pos + 1
	}

	return message[pos] == r.Value, pos + 1
}

type MessageValidatorHolder struct {
	Validator MessageValidator
}

type Validators = map[int]*MessageValidatorHolder

type World struct {
	Validators Validators
	Messages   []string
}

func IsValid(message string, validator MessageValidator) bool {
	if ok, pos := validator.Valid(0, message); ok && pos == len(message) {
		return true
	}

	return false
}

func DoWithInputPart01(world World) int {
	count := 0

	validator := world.Validators[0].Validator
	for _, message := range world.Messages {
		if IsValid(message, validator) {
			count++
		}
	}

	return count
}

func DoWithInputPart02(world World) int {
	return 0
}

func getOrCreateValidator(id int, validators Validators) *MessageValidatorHolder {
	if validator, ok := validators[id]; ok {
		return validator
	}

	validator := &MessageValidatorHolder{}
	validators[id] = validator
	return validator
}

func parseValidatorPart(str string, validators Validators) MessageValidator {
	// value rule
	if str[0] == '"' {
		return &ValueRule{Value: str[1]}
	}

	// AND rule
	ids := utils.ExtractInts(str, false)
	rules := make([]*MessageValidatorHolder, len(ids))

	for i, id := range ids {
		rules[i] = getOrCreateValidator(id, validators)
	}

	return &AndRule{
		Rules: rules,
	}
}

func parseValidator(id int, str string, validators Validators) *MessageValidatorHolder {
	subRules := strings.Split(str, " | ")
	validator := getOrCreateValidator(id, validators)

	if len(subRules) == 1 {
		// value rule or AND rule
		validator.Validator = parseValidatorPart(str, validators)
		return validator
	}

	// OR rule
	left := parseValidatorPart(subRules[0], validators)
	right := parseValidatorPart(subRules[1], validators)

	validator.Validator = &OrRule{
		Left:  left,
		Right: right,
	}

	return validator
}

func ParseInput(r io.Reader) World {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	validators := make(Validators)
	for scanner.Scan() && scanner.Text() != "" {
		mainParts := strings.Split(scanner.Text(), ": ")
		ruleId := utils.ParseInt(mainParts[0])

		parseValidator(ruleId, mainParts[1], validators)
	}

	var messages []string
	for scanner.Scan() {
		messages = append(messages, scanner.Text())
	}

	return World{
		Validators: validators,
		Messages:   messages,
	}
}
