package main

import (
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func Test_01_parse(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	commands := ParseInput(reader)

	assert.Equal(t, 10, len(commands))
	assert.Equal(t, "7214296 k", commands[9].Output[3])
}

func Test_01_example(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	commands := ParseInput(reader)

	root := ReplayCommands(commands)

	expected := utils.Msg(`
- / (dir)
  - a (dir)
    - e (dir)
      - i (file, size=584)
    - f (file, size=29116)
    - g (file, size=2557)
    - h.lst (file, size=62596)
  - b.txt (file, size=14848514)
  - c.dat (file, size=8504156)
  - d (dir)
    - j (file, size=4060174)
    - d.log (file, size=8033020)
    - d.ext (file, size=5626152)
    - k (file, size=7214296)
`)

	assert.Equal(t, expected, root.String())

	sum := Filter100k(root)
	assert.Equal(t, 95437, sum)
}

func Test_01(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	commands := ParseInput(reader)

	root := ReplayCommands(commands)

	sum := Filter100k(root)
	assert.Equal(t, 1783610, sum)
}

//
//func Test_02_example(t *testing.T) {
//	reader, err := os.Open("data-00-example.txt")
//	assert.Nil(t, err)
//
//	commands := ParseInput(reader)
//
//	root := ReplayCommands(commands)
//	assert.Equal(t, 0, root)
//}
//
//func Test_02(t *testing.T) {
//	reader, err := os.Open("data-01.txt")
//	assert.Nil(t, err)
//
//	commands := ParseInput(reader)
//
//	root := ReplayCommands(commands)
//	assert.Equal(t, 0, root)
//}
