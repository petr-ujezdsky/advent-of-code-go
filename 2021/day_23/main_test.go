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

func Test_01_example_final_1(t *testing.T) {
	metric.Enabled = true
	building := NewBuilding2(string([]rune{'A', 0}), "BB", "CC", "DD")
	building.Hallway()[9] = 'A'
	building.ConsumedEnergy = 12521 - 8

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

//
//func Test_01(t *testing.T) {
//	reader, err := os.Open("data-01.txt")
//	assert.Nil(t, err)
//
//	cubes := ParseInput(reader)
//	assert.Equal(t, 420, len(cubes))
//
//	count := CountOnCubes(cubes[:20])
//	assert.Equal(t, 620241, count)
//}
//
//func Test_02_example_3(t *testing.T) {
//	reader, err := os.Open("data-00-example-3.txt")
//	assert.Nil(t, err)
//
//	cubes := ParseInput(reader)
//	assert.Equal(t, 60, len(cubes))
//
//	count := CountOnCubes(cubes)
//	assert.Equal(t, 2758514936282235, count)
//}
//
//func Test_02(t *testing.T) {
//	reader, err := os.Open("data-01.txt")
//	assert.Nil(t, err)
//
//	cubes := ParseInput(reader)
//	assert.Equal(t, 420, len(cubes))
//
//	count := CountOnCubes(cubes)
//	assert.Equal(t, 1284561759639324, count)
//}
//
//// Benchmark_CountOnCubes-10    	     556	   2 136 827 ns/op
//func Benchmark_CountOnCubes(b *testing.B) {
//	reader, err := os.Open("data-00-example-2.txt")
//	assert.Nil(b, err)
//
//	cubes := ParseInput(reader)
//	assert.Equal(b, 22, len(cubes))
//
//	b.ResetTimer()
//
//	for i := 0; i < b.N; i++ {
//		count := CountOnCubes(cubes[:20])
//		assert.Equal(b, 590784, count)
//	}
//}
