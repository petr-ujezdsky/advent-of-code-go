package day_06

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_01_example_parse(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	timers, err := ParseInput(reader)
	assert.Nil(t, err)

	assert.Equal(t, []int{3, 4, 3, 1, 2}, timers)
}

func Test_01_example_18_days(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	timers, err := ParseInput(reader)
	assert.Nil(t, err)

	timers = PassDays(timers, 18)

	fmt.Println(timers)

	assert.Equal(t, 26, len(timers))
}

func Test_01_example_80_days(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	timers, err := ParseInput(reader)
	assert.Nil(t, err)

	timers = PassDays(timers, 80)

	fmt.Println(timers)

	assert.Equal(t, 5934, len(timers))
}

func Test_01(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	timers, err := ParseInput(reader)
	assert.Nil(t, err)

	timers = PassDays(timers, 80)

	//fmt.Println(timers)

	assert.Equal(t, 349549, len(timers))
}

func Test_02_example_18_days(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	fish, err := ParseFish(reader, 18)
	assert.Nil(t, err)

	count := CountManyFish(fish)

	assert.Equal(t, 26, count)
}

func Test_02_example_256_days(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	fish, err := ParseFish(reader, 256)
	assert.Nil(t, err)

	count := CountManyFish(fish)

	assert.Equal(t, 26984457539, count)
}
