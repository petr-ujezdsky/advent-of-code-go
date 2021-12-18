package day_02_test

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_01_example(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	commands, err := ParseToCommands(reader)
	assert.Nil(t, err)

	x, y, result := move(commands)

	assert.Equal(t, 15, x)
	assert.Equal(t, 10, y)
	assert.Equal(t, 150, result)
}

func Test_01(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	commands, err := ParseToCommands(reader)
	assert.Nil(t, err)

	x, y, result := move(commands)

	assert.Equal(t, 1968, x)
	assert.Equal(t, 1063, y)
	assert.Equal(t, 2091984, result)
}

func Test_02_example(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	commands, err := ParseToCommands(reader)
	assert.Nil(t, err)

	x, y, result := moveByAim(commands)

	assert.Equal(t, 15, x)
	assert.Equal(t, 60, y)
	assert.Equal(t, 900, result)
}

func Test_02(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	commands, err := ParseToCommands(reader)
	assert.Nil(t, err)

	x, y, result := moveByAim(commands)

	assert.Equal(t, 1968, x)
	assert.Equal(t, 1060092, y)
	assert.Equal(t, 2086261056, result)
}

type Command struct {
	MoveX, MoveY int
}

func ParseToCommands(r io.Reader) ([]Command, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	var result []Command

	for scanner.Scan() {
		row := strings.Split(scanner.Text(), " ")

		commandStr := row[0]

		amount, err := strconv.Atoi(row[1])
		if err != nil {
			return result, err
		}

		var command Command

		switch commandStr {
		case "forward":
			command = Command{amount, 0}
		case "down":
			command = Command{0, amount}
		case "up":
			command = Command{0, -amount}
		default:
			return result, fmt.Errorf("Unknown command %v", commandStr)
		}

		result = append(result, command)
	}

	return result, scanner.Err()
}

func move(commands []Command) (int, int, int) {
	x := 0
	y := 0

	for _, command := range commands {
		x += command.MoveX
		y += command.MoveY
	}

	return x, y, x * y
}

func moveByAim(commands []Command) (int, int, int) {
	x := 0
	y := 0
	aim := 0

	for _, command := range commands {
		x += command.MoveX
		aim += command.MoveY

		if command.MoveX != 0 {
			y += aim * command.MoveX
		}
	}

	return x, y, x * y
}
