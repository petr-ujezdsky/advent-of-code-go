package main

import (
	"bufio"
	_ "embed"
	"io"
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
