package day_22

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Building_IsSorted(t *testing.T) {
	assert.Equal(t, true, NewBuilding2("AA", "BB", "CC", "DD").IsSorted())
	assert.Equal(t, false, NewBuilding2("AB", "DC", "CB", "AD").IsSorted())
}

func Test_Building_String(t *testing.T) {
	building := NewBuilding2("AB", "DC", "CB", "AD")

	expected := `
#############
#...........#
###B#C#B#D###
  #A#D#C#A#
  #########`
	assert.Equal(t, expected, "\n"+building.String())
}

func Test_01_example_sorted(t *testing.T) {
	metric.Enabled = true
	building := NewBuilding2("AA", "BB", "CC", "DD")
	building.ConsumedEnergy = 12521

	score, _ := Sort(building)
	assert.Equal(t, 12521, score)
}

func Test_01_example_final_2(t *testing.T) {
	metric.Enabled = true
	building := NewBuilding2Full(".....D.D.A.", "A.", "BB", "CC", "..", 12521-8-7000)

	score, _ := Sort(building)
	assert.Equal(t, 12521, score)
}

func Test_01_example_final_1(t *testing.T) {
	metric.Enabled = true
	building := NewBuilding2Full(".........A.", "A.", "BB", "CC", "DD", 12521-8)

	score, _ := Sort(building)
	assert.Equal(t, 12521, score)
}

func Test_01_example(t *testing.T) {
	metric.Enabled = true
	building := NewBuilding2("AB", "DC", "CB", "AD")

	score, winner := Sort(building)
	assert.Equal(t, 12521, score)
	PrintMoves(winner)
}

func Test_01(t *testing.T) {
	metric.Enabled = true
	building := NewBuilding2("CB", "AB", "AD", "CD")

	score, winner := Sort(building)
	assert.Equal(t, 10607, score)
	PrintMoves(winner)
}
