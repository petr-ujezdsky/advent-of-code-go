package main

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func Test_01_parseSNAFU(t *testing.T) {
	assert.Equal(t, 0, ParseSNAFU("0"))
	assert.Equal(t, 1, ParseSNAFU("1"))
	assert.Equal(t, 2, ParseSNAFU("2"))
	assert.Equal(t, 3, ParseSNAFU("1="))
	assert.Equal(t, 4, ParseSNAFU("1-"))
	assert.Equal(t, 5, ParseSNAFU("10"))
	assert.Equal(t, 6, ParseSNAFU("11"))
	assert.Equal(t, 7, ParseSNAFU("12"))
	assert.Equal(t, 8, ParseSNAFU("2="))
	assert.Equal(t, 9, ParseSNAFU("2-"))
	assert.Equal(t, 10, ParseSNAFU("20"))

	assert.Equal(t, 15, ParseSNAFU("1=0"))
	assert.Equal(t, 20, ParseSNAFU("1-0"))
	assert.Equal(t, 2022, ParseSNAFU("1=11-2"))
	assert.Equal(t, 12345, ParseSNAFU("1-0---0"))
	assert.Equal(t, 314159265, ParseSNAFU("1121-1110-1=0"))
}
func Test_01_parse(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	snafuNumbers := ParseInput(reader)

	assert.Equal(t, 0, len(snafuNumbers))
}

func Test_01_example(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	snafuNumbers := ParseInput(reader)

	result := DoWithInput(snafuNumbers)
	assert.Equal(t, 0, result)
}

func Test_01(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	snafuNumbers := ParseInput(reader)

	result := DoWithInput(snafuNumbers)
	assert.Equal(t, 0, result)
}

func Test_02_example(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	snafuNumbers := ParseInput(reader)

	result := DoWithInput(snafuNumbers)
	assert.Equal(t, 0, result)
}

func Test_02(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	snafuNumbers := ParseInput(reader)

	result := DoWithInput(snafuNumbers)
	assert.Equal(t, 0, result)
}
