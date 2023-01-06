package main

import (
	"bufio"
	_ "embed"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"io"
	"regexp"
	"strings"
)

var requiredFields = []string{
	"byr",
	"iyr",
	"eyr",
	"hgt",
	"hcl",
	"ecl",
	"pid",
	//"cid",
}

type Passport struct {
	Fields map[string]string
}

type Validator = func(p Passport) bool

func RequiredFieldsValidator(p Passport) bool {
	for _, field := range requiredFields {
		if _, ok := p.Fields[field]; !ok {
			return false
		}
	}

	return true
}

func numberValidator(s string, interval utils.IntervalI) bool {
	i := utils.ParseInt(s)

	return interval.Contains(i)
}

var regexHeight = regexp.MustCompile(`^(\d+)(cm|in)$`)

func heightValidator(s string) bool {
	parts := regexHeight.FindStringSubmatch(s)

	if len(parts) != 3 {
		return false
	}

	amount := utils.ParseInt(parts[1])
	unit := parts[2]

	switch unit {
	case "cm":
		return 150 <= amount && amount <= 193
	case "in":
		return 59 <= amount && amount <= 76
	}

	return false
}

var regexHairColor = regexp.MustCompile(`^#[0-9a-f]{6}$`)
var regexEyeColor = regexp.MustCompile(`^(amb|blu|brn|gry|grn|hzl|oth)$`)
var regexPassportId = regexp.MustCompile(`^\d{9}$`)

func FieldsContentValidator(p Passport) bool {
	if !RequiredFieldsValidator(p) {
		return false
	}

	return numberValidator(p.Fields["byr"], utils.IntervalI{Low: 1920, High: 2002}) &&
		numberValidator(p.Fields["iyr"], utils.IntervalI{Low: 2010, High: 2020}) &&
		numberValidator(p.Fields["eyr"], utils.IntervalI{Low: 2020, High: 2030}) &&
		heightValidator(p.Fields["hgt"]) &&
		regexHairColor.MatchString(p.Fields["hcl"]) &&
		regexEyeColor.MatchString(p.Fields["ecl"]) &&
		regexPassportId.MatchString(p.Fields["pid"])
}

func ValidatePassports(passports []Passport, validator Validator) int {
	validCount := 0

	for _, passport := range passports {
		if validator(passport) {
			validCount++
		}
	}

	return validCount
}

func ParseInput(r io.Reader) []Passport {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	var passports []Passport
	for scanner.Scan() {
		var lines []string

		for len(scanner.Text()) > 0 {
			lines = append(lines, scanner.Text())
			scanner.Scan()
		}

		fields := make(map[string]string)
		for _, line := range lines {
			pairs := strings.Split(line, " ")
			for _, pair := range pairs {
				parts := strings.Split(pair, ":")
				fields[parts[0]] = parts[1]
			}
		}

		passport := Passport{Fields: fields}

		passports = append(passports, passport)
	}

	return passports
}
