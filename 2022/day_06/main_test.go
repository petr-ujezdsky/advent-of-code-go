package main

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func Test_01_example(t *testing.T) {
	assert.Equal(t, 7, FindPacketStart("mjqjpqmgbljsphdztnvjfqwrcgsmlb"))
	assert.Equal(t, 5, FindPacketStart("bvwbjplbgvbhsrlpgdmjqwftvncz"))
	assert.Equal(t, 6, FindPacketStart("nppdvjthqldpwncqszvftbrmjlhg"))
	assert.Equal(t, 10, FindPacketStart("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg"))
	assert.Equal(t, 11, FindPacketStart("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"))
}

func Test_01(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	str := ParseInput(reader)

	pos := FindPacketStart(str)
	assert.Equal(t, 1531, pos)
}

//
//func Test_02_example(t *testing.T) {
//	reader, err := os.Open("data-00-example.txt")
//	assert.Nil(t, err)
//
//	stacks, ops := ParseInput(reader)
//
//	topCrates := MoveCratesByOps(stacks, ops, true)
//	assert.Equal(t, "MCD", topCrates)
//}
//
//func Test_02(t *testing.T) {
//	reader, err := os.Open("data-01.txt")
//	assert.Nil(t, err)
//
//	stacks, ops := ParseInput(reader)
//
//	topCrates := MoveCratesByOps(stacks, ops, true)
//	assert.Equal(t, "CJVLJQPHS", topCrates)
//}
