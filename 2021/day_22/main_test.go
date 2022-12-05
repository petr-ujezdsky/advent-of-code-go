package day_22

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_regexCube(t *testing.T) {
	matches := regexCube.FindStringSubmatch("on x=10..12,y=10..12,z=10..12")
	assert.Equal(t, []string{"on x=10..12,y=10..12,z=10..12", "on", "10", "12", "10", "12", "10", "12"}, matches)

	matches = regexCube.FindStringSubmatch("on x=-54112..-39298,y=-85059..-49293,z=-27449..7877")
	assert.Equal(t, []string{"on x=-54112..-39298,y=-85059..-49293,z=-27449..7877", "on", "-54112", "-39298", "-85059", "-49293", "-27449", "7877"}, matches)
}

func Test_Split(t *testing.T) {
	probe := Cube{
		Low:   Vector3i{0, 0, 0},
		High:  Vector3i{3, 3, 3},
		Value: false,
	}

	singleCell := probe.Split()[0].Split()[0]
	expected := Cube{
		Low:   Vector3i{0, 0, 0},
		High:  Vector3i{0, 0, 0},
		Value: false,
	}
	assert.Equal(t, expected, singleCell)
}

func Test_Intersect(t *testing.T) {
	probe := Cube{
		Low:   Vector3i{0, 0, 0},
		High:  Vector3i{0, 0, 0},
		Value: false,
	}

	intersectionType := probe.Intersect(probe)
	assert.Equal(t, Inside, intersectionType)
}

func Test_01_example_1_fast(t *testing.T) {
	reader, err := os.Open("data-00-example-1.txt")
	assert.Nil(t, err)

	cubes := ParseInput2(reader)
	assert.Equal(t, 4, len(cubes))

	count := ExtraFastCount(cubes)
	assert.Equal(t, 39, count)
}

func Test_01_example_2_fast(t *testing.T) {
	reader, err := os.Open("data-00-example-2.txt")
	assert.Nil(t, err)

	cubes := ParseInput2(reader)
	assert.Equal(t, 4, len(cubes))

	count := ExtraFastCount(cubes[:20])
	assert.Equal(t, 590784, count)
}

func Test_01_fast(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	cubes := ParseInput2(reader)
	assert.Equal(t, 420, len(cubes))

	count := ExtraFastCount(cubes[:20])
	assert.Equal(t, 620241, count)
}

func Test_02_example_3_fast(t *testing.T) {
	reader, err := os.Open("data-00-example-3.txt")
	assert.Nil(t, err)

	cubes := ParseInput2(reader)
	assert.Equal(t, 60, len(cubes))

	metric.Enabled = true
	count := ExtraFastCount(cubes)
	assert.Equal(t, 2758514936282235, count)
}

func Test_02_fast(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	cubes := ParseInput2(reader)
	assert.Equal(t, 420, len(cubes))

	count := ExtraFastCount(cubes)
	assert.Equal(t, 1284561759639324, count)
}

func Test_01_example_1(t *testing.T) {
	reader, err := os.Open("data-00-example-1.txt")
	assert.Nil(t, err)

	cubes, _ := ParseInput(reader)
	assert.Equal(t, 4, len(cubes))

	world := NewCubeSymmetric(50, false)

	//count := NaiveCount(world, cubes)
	count := FasterCount(world, cubes)
	assert.Equal(t, 39, count)
}

func Test_01_example_2(t *testing.T) {
	reader, err := os.Open("data-00-example-2.txt")
	assert.Nil(t, err)

	cubes, _ := ParseInput(reader)
	assert.Equal(t, 22, len(cubes))

	world := NewCubeSymmetric(50, false)

	//count := NaiveCount(world, cubes[:20])
	count := FasterCount(world, cubes[:20])
	assert.Equal(t, 590784, count)
}

func Test_01(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	cubes, _ := ParseInput(reader)
	assert.Equal(t, 420, len(cubes))

	world := NewCubeSymmetric(50, false)

	//count := NaiveCount(world, cubes[:20])
	count := FasterCount(world, cubes[:20])
	assert.Equal(t, 620241, count)
}

func Test_02_example_3(t *testing.T) {
	reader, err := os.Open("data-00-example-3.txt")
	assert.Nil(t, err)

	cubes, world := ParseInput(reader)
	assert.Equal(t, 60, len(cubes))

	metric.Enabled = true
	//count := NaiveCount(world, cubes)
	count := FasterCount(world, cubes)
	assert.Equal(t, 2758514936282235, count)
}

// Benchmark_fast-10    	      46	  25 154 525 ns/op
func Benchmark_fast(b *testing.B) {
	reader, err := os.Open("data-00-example-2.txt")
	assert.Nil(b, err)

	cubes, _ := ParseInput(reader)
	assert.Equal(b, 22, len(cubes))

	world := NewCubeSymmetric(50, false)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		count := FasterCount(world, cubes[:20])
		assert.Equal(b, 590784, count)
	}
}

// Benchmark_faster-10    	     556	   2 136 827 ns/op
func Benchmark_faster(b *testing.B) {
	reader, err := os.Open("data-00-example-2.txt")
	assert.Nil(b, err)

	cubes := ParseInput2(reader)
	assert.Equal(b, 22, len(cubes))

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		count := ExtraFastCount(cubes[:20])
		assert.Equal(b, 590784, count)
	}
}
