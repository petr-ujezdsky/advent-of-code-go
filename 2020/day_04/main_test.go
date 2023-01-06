package main

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func Test_01_parse(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	passports := ParseInput(reader)

	assert.Equal(t, 4, len(passports))
}

func Test_01_example(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	passports := ParseInput(reader)

	result := ValidatePassports(passports, RequiredFieldsValidator)
	assert.Equal(t, 2, result)
}

func Test_01(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	passports := ParseInput(reader)

	result := ValidatePassports(passports, RequiredFieldsValidator)
	assert.Equal(t, 260, result)
}

func Test_02_example(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	passports := ParseInput(reader)

	result := ValidatePassports(passports, FieldsContentValidator)
	assert.Equal(t, 2, result)
}

func Test_02(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	passports := ParseInput(reader)

	result := ValidatePassports(passports, FieldsContentValidator)
	assert.Equal(t, 153, result)
}
