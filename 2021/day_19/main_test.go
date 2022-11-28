package day_19

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func Test_regexHeader(t *testing.T) {
	matches := regexHeader.FindStringSubmatch("--- scanner 5 ---")
	assert.Equal(t, []string{"--- scanner 5 ---", "5"}, matches)
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
