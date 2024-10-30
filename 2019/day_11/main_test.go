package main

import (
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func Test_01_parse(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	world := ParseInput(reader)

	assert.Equal(t, 647, len(world.Program))
}

func Test_01(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	world := ParseInput(reader)

	result := DoWithInputPart01(world)
	assert.Equal(t, 2219, result)
}

func Test_02(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	world := ParseInput(reader)

	result := DoWithInputPart02(world)
	expected := utils.Msg(`
 #  #  ##  #### #  # #     ##  ###  ####   
 #  # #  # #    #  # #    #  # #  # #      
 #### #  # ###  #  # #    #  # #  # ###    
 #  # #### #    #  # #    #### ###  #      
 #  # #  # #    #  # #    #  # #    #      
 #  # #  # #     ##  #### #  # #    ####   `)
	assert.Equal(t, expected, result)
}
