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

	world := ParseInput(reader)

	assert.Equal(t, 12, len(world.Numbers))
}

func Test_01_example(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	world := ParseInput(reader)

	result := DoWithInputPart01(world, 3, 2)
	assert.Equal(t, 1, result)
}

func Test_01(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	world := ParseInput(reader)

	result := DoWithInputPart01(world, 25, 6)
	assert.Equal(t, 1862, result)
}

func Test_02_example(t *testing.T) {
	reader, err := os.Open("data-00-example2.txt")
	assert.Nil(t, err)

	world := ParseInput(reader)

	result := DoWithInputPart02(world, 2, 2)
	expected := utils.Msg(`
 #
# `)
	assert.Equal(t, expected, result)
}

func Test_02(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	world := ParseInput(reader)

	result := DoWithInputPart02(world, 25, 6)
	expected := utils.Msg(`
 ##   ##  ###  #  # #    
#  # #  # #  # #  # #    
#    #    #  # #### #    
# ## #    ###  #  # #    
#  # #  # #    #  # #    
 ###  ##  #    #  # #### `)
	assert.Equal(t, expected, result)
}
