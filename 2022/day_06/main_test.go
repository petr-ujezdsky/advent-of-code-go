package main

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func Test_01_example(t *testing.T) {
	assert.Equal(t, 7, FindPacketStart("mjqjpqmgbljsphdztnvjfqwrcgsmlb", 4))
	assert.Equal(t, 5, FindPacketStart("bvwbjplbgvbhsrlpgdmjqwftvncz", 4))
	assert.Equal(t, 6, FindPacketStart("nppdvjthqldpwncqszvftbrmjlhg", 4))
	assert.Equal(t, 10, FindPacketStart("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 4))
	assert.Equal(t, 11, FindPacketStart("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 4))
}

func Test_01(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	str := ParseInput(reader)

	pos := FindPacketStart(str, 4)
	assert.Equal(t, 1531, pos)
}

func Test_02_example(t *testing.T) {
	assert.Equal(t, 19, FindPacketStart("mjqjpqmgbljsphdztnvjfqwrcgsmlb", 14))
	assert.Equal(t, 23, FindPacketStart("bvwbjplbgvbhsrlpgdmjqwftvncz", 14))
	assert.Equal(t, 23, FindPacketStart("nppdvjthqldpwncqszvftbrmjlhg", 14))
	assert.Equal(t, 29, FindPacketStart("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 14))
	assert.Equal(t, 26, FindPacketStart("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 14))
}

func Test_02(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	str := ParseInput(reader)

	pos := FindPacketStart(str, 14)
	assert.Equal(t, 2518, pos)
}
