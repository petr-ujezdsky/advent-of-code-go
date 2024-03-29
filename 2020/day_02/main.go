package main

import (
	"bufio"
	_ "embed"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"io"
	"strings"
)

type PasswordRecord struct {
	Counts   utils.IntervalI
	Char     rune
	Password string
}

func ValidatePasswords1(passwordRecords []PasswordRecord) int {
	validCount := 0

	for _, passwordRecord := range passwordRecords {
		occurance := strings.Count(passwordRecord.Password, string(passwordRecord.Char))
		if passwordRecord.Counts.Contains(occurance) {
			validCount++
		}
	}

	return validCount
}

func ValidatePasswords2(passwordRecords []PasswordRecord) int {
	validCount := 0

	for _, passwordRecord := range passwordRecords {
		ch1 := rune(passwordRecord.Password[passwordRecord.Counts.Low-1])
		ch2 := rune(passwordRecord.Password[passwordRecord.Counts.High-1])

		if ch1 != ch2 && (ch1 == passwordRecord.Char || ch2 == passwordRecord.Char) {
			validCount++
		}
	}

	return validCount
}

func ParseInput(r io.Reader) []PasswordRecord {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	var passwordRecords []PasswordRecord
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " ")

		ints := utils.ExtractInts(parts[0], false)
		counts := utils.IntervalI{Low: ints[0], High: ints[1]}

		char := rune(parts[1][0])

		password := parts[2]

		passwordRecord := PasswordRecord{
			Counts:   counts,
			Char:     char,
			Password: password,
		}

		passwordRecords = append(passwordRecords, passwordRecord)
	}

	return passwordRecords
}
