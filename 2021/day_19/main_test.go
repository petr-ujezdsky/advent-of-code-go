package day_19

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func Test_regexHeader(t *testing.T) {
	matches := regexHeader.FindStringSubmatch("--- scanner 5 ---")
	assert.Equal(t, []string{"--- scanner 5 ---", "5"}, matches)
}

func Test_printAllRotations(t *testing.T) {
	rots := allRotations([]Vector3i{{686, 422, 578}})
	for i, rot := range rots {
		if rot[0] == (Vector3i{686, 422, 578}) {
			fmt.Printf("* %0d %s\n", i, rot)
		} else {
			fmt.Printf("  %0d %s\n", i, rot)
		}
	}
}

func Test_01_example_parse(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	scanners, err := ParseInput(reader)
	assert.Nil(t, err)

	assert.Equal(t, 0, scanners[0].Id)
	assert.Equal(t, Vector3i{404, -588, -901}, scanners[0].RotatedBeacons[0][0])
	assert.Equal(t, Vector3i{459, -707, 401}, scanners[0].RotatedBeacons[0][24])

	assert.Equal(t, 4, scanners[4].Id)
	assert.Equal(t, Vector3i{727, 592, 562}, scanners[4].RotatedBeacons[0][0])
	assert.Equal(t, Vector3i{30, -46, -14}, scanners[4].RotatedBeacons[0][25])
}

func Test_01_example(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	scanners, err := ParseInput(reader)
	assert.Nil(t, err)

	mainScanner := SearchAndConsume(scanners)

	assert.Equal(t, 79, len(mainScanner.UniqueBeacons))
}

func Test_01(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	scanners, err := ParseInput(reader)
	assert.Nil(t, err)

	mainScanner := SearchAndConsume(scanners)

	assert.Equal(t, 432, len(mainScanner.UniqueBeacons))
}
