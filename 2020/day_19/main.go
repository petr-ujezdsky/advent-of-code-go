package main

import (
	"bufio"
	_ "embed"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/strs"
	"io"
	"strings"
)

type VariantVisitor = func(variant string)

type MessageValidator interface {
	Valid(pos int, last bool, message string) (bool, int)
	Print(prefix string, last bool, visitor VariantVisitor) []string
}

type AndRule struct {
	Rules []*MessageValidatorHolder
}

func (r AndRule) Valid(pos int, last bool, message string) (bool, int) {
	currentPos := pos
	for i, rule := range r.Rules {
		end := i == len(r.Rules)-1
		ok, nextPos := rule.Validator.Valid(currentPos, last && end, message)
		if !ok {
			return false, nextPos
		}
		currentPos = nextPos
	}

	return true, currentPos
}

func (r AndRule) Print(prefix string, last bool, visitor VariantVisitor) []string {
	variants := []string{prefix}
	var nextVariants []string

	for i, rule := range r.Rules {
		end := i == len(r.Rules)-1
		for _, variant := range variants {
			for _, nextVariant := range rule.Validator.Print(variant, last && end, visitor) {
				nextVariants = append(nextVariants, nextVariant)
			}
		}
		variants = nextVariants
	}

	return variants
}

type OrRule struct {
	Left, Right MessageValidator
}

func (r OrRule) Valid(pos int, last bool, message string) (bool, int) {
	ok, nextPos := r.Left.Valid(pos, last, message)
	if ok {
		return ok, nextPos
	}

	ok, nextPos = r.Right.Valid(pos, last, message)
	return ok, nextPos
}

func (r OrRule) Print(prefix string, last bool, visitor VariantVisitor) []string {
	var variants []string
	for _, variant := range r.Left.Print(prefix, last, visitor) {
		variants = append(variants, variant)
	}

	for _, variant := range r.Right.Print(prefix, last, visitor) {
		variants = append(variants, variant)
	}

	return variants
}

type ValueRule struct {
	Value uint8
}

func (r ValueRule) Valid(pos int, last bool, message string) (bool, int) {
	if pos >= len(message) {
		return false, pos + 1
	}

	if last && pos != len(message)-1 {
		//fmt.Println("Last but actually not")
		return false, pos + 1
	}

	return message[pos] == r.Value, pos + 1
}

func (r ValueRule) Print(prefix string, last bool, visitor VariantVisitor) []string {
	variant := prefix + string(r.Value)

	if last {
		visitor(variant)
	}

	return []string{variant}
}

type MessageValidatorHolder struct {
	Id        int
	Validator MessageValidator
}

type Validators = map[int]*MessageValidatorHolder

type World struct {
	Validators Validators
	Validator  MessageValidator
	Messages   []string
}

func IsValid(message string, validator MessageValidator) bool {
	if ok, pos := validator.Valid(0, true, message); ok {
		if pos != len(message) {
			//fmt.Printf("Matched but not whole: %v @ %v/%v\n", message, pos, len(message))
			return false
		}
		return true
	}

	return false
}

func IsValidWithCycles(message string, validator MessageValidator, validator8 MessageValidator) bool {
	for i := range message {
		if i == 0 {
			continue
		}

		if ok, _ := validator.Valid(i, true, message); ok {
			//fmt.Printf("Matched right part from index %v/%v\n", i, len(message))

			if ok, _ := validator8.Valid(0, true, strs.Substring(message, 0, i)); ok {
				//fmt.Println("Also matched left part!")
				return true
			}

			//fmt.Println("Did not match left part")
		}
	}

	return false
}

func DoWithInputPart01(world World) int {
	count := 0

	for _, message := range world.Messages {
		if IsValid(message, world.Validator) {
			count++
		}
	}

	return count
}

func DoWithInputPart02(world World) int {
	count := 0
	validator8 := world.Validators[8].Validator

	for _, message := range world.Messages {
		if IsValidWithCycles(message, world.Validator, validator8) {
			count++
		}
	}

	return count
}

func getOrCreateValidator(id int, validators Validators) *MessageValidatorHolder {
	if validator, ok := validators[id]; ok {
		return validator
	}

	validator := &MessageValidatorHolder{
		Id: id,
	}

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
		Validator:  validators[0].Validator,
		Messages:   messages,
	}
}
