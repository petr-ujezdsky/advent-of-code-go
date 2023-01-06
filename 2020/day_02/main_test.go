package main

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func Test_01_parse(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	passwordRecords := ParseInput(reader)

	assert.Equal(t, 3, len(passwordRecords))
}

func Test_01_example(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	passwordRecords := ParseInput(reader)

	result := ValidatePasswords1(passwordRecords)
	assert.Equal(t, 2, result)
}

func Test_01(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	passwordRecords := ParseInput(reader)

	result := ValidatePasswords1(passwordRecords)
	assert.Equal(t, 483, result)
}

func Test_02_example(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	passwordRecords := ParseInput(reader)

	result := ValidatePasswords2(passwordRecords)
	assert.Equal(t, 1, result)
}

func Test_02(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	passwordRecords := ParseInput(reader)

	result := ValidatePasswords2(passwordRecords)
	assert.Equal(t, 482, result)
}
