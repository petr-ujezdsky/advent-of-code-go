package day_22

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Building_IsSorted(t *testing.T) {
	assert.Equal(t, true, NewBuilding("AA", "BB", "CC", "DD").IsSorted())
	assert.Equal(t, false, NewBuilding("AB", "DC", "CB", "AD").IsSorted())
}

func Test_Building_String(t *testing.T) {
	building := NewBuilding("AB", "DC", "CB", "AD")

	expected := `
#############
#...........#
###B#C#B#D###
  #A#D#C#A#
  #########`
	assert.Equal(t, expected, "\n"+building.String())

	building = NewBuilding("ADDB", "DBCC", "CABB", "ACAD")

	expected = `
#############
#...........#
###B#C#B#D###
  #D#C#B#A#
  #D#B#A#C#
  #A#D#C#A#
  #########`
	assert.Equal(t, expected, "\n"+building.String())
}

func Test_01_example_sorted(t *testing.T) {
	metrics.Enable()
	building := NewBuilding("AA", "BB", "CC", "DD")
	building.ConsumedEnergy = 12521

	score, _ := Sort(building)
	assert.Equal(t, 12521, score)
}

func Test_01_example_final_2(t *testing.T) {
	metrics.Enable()
	building := NewBuildingFull(".....D.D.A.", "A.", "BB", "CC", "..", 12521-8-7000)

	score, _ := Sort(building)
	assert.Equal(t, 12521, score)
}

func Test_01_example_final_1(t *testing.T) {
	metrics.Enable()
	building := NewBuildingFull(".........A.", "A.", "BB", "CC", "DD", 12521-8)

	score, _ := Sort(building)
	assert.Equal(t, 12521, score)
}

func Test_01_example(t *testing.T) {
	metrics.Enable()
	building := NewBuilding("AB", "DC", "CB", "AD")

	score, winner := Sort(building)
	assert.Equal(t, 12521, score)
	PrintMoves(winner)
}

func Test_01(t *testing.T) {
	metrics.Enable()
	building := NewBuilding("CB", "AB", "AD", "CD")

	score, winner := Sort(building)
	assert.Equal(t, 10607, score)
	PrintMoves(winner)
}

func Test_02_example(t *testing.T) {
	metrics.Enable()
	building := NewBuilding("ADDB", "DBCC", "CABB", "ACAD")

	score, winner := Sort(building)
	assert.Equal(t, 44169, score)
	PrintMoves(winner)
}

func Test_02(t *testing.T) {
	metrics.Enable()
	building := NewBuilding("CDDB", "ABCB", "AABD", "CCAD")

	score, winner := Sort(building)
	assert.Equal(t, 59071, score)
	PrintMoves(winner)
}
