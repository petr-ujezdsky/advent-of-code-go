package day_17_test

import (
	"math"
	"testing"

	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"github.com/stretchr/testify/assert"
)

// target area: x=20..30, y=-10..-5
func Test_01_example(t *testing.T) {
	maxHeight := getMaxHeight(-10)

	assert.Equal(t, 45, maxHeight)
}

// target area: x=179..201, y=-109..-63
func Test_01(t *testing.T) {
	maxHeight := getMaxHeight(-109)

	assert.Equal(t, 5886, maxHeight)
}

// The idea is, that you will *always* reach y=0. Then you make the biggest step directly to lowest border minHeight. This step is +1 bigger due to gravity, so initial velocity is -minHeight - 1.
// Now you only need to sum integers from (-minHeight - 1) [initial velocity] to 0 [velocity at top].
func getMaxHeight(minHeight int) int {
	return utils.SumNtoM(0, utils.Abs(minHeight)-1)
}

func Test_02_example(t *testing.T) {
	velocities := simulateAll(Trench{20, 30, -5, -10})

	assert.Equal(t, 112, len(velocities))
}

func Test_02(t *testing.T) {
	velocities := simulateAll(Trench{179, 201, -63, -109})

	assert.Equal(t, 1806, len(velocities))
}

type Trench struct {
	X0, X1, Y0, Y1 int
}

type Vector2i struct {
	X, Y int
}

// Brute-force all possible initial velocity combinations
func simulateAll(trench Trench) []Vector2i {
	var initialVelocities []Vector2i

	for velocityX := 1; velocityX <= trench.X1; velocityX++ {
		for velocityY := -trench.Y1; velocityY >= trench.Y1; velocityY-- {
			if reachesTrench(velocityX, velocityY, trench) {
				initialVelocities = append(initialVelocities, Vector2i{velocityX, velocityY})
			}
		}
	}

	return initialVelocities
}

// Simulates throw and validates if the probe reaches the trench
func reachesTrench(velocityX, velocityY int, trench Trench) bool {
	for stepsCount := 0; stepsCount < 300; stepsCount++ {
		x := utils.SumNtoM(utils.Clamp(velocityX-stepsCount, 0, math.MaxInt), velocityX)
		y := utils.SumNtoM(velocityY-stepsCount, velocityY)

		if x >= trench.X0 && x <= trench.X1 && y >= trench.Y1 && y <= trench.Y0 {
			// position inside trench
			return true
		}

		if x > trench.X1 || y < trench.Y1 {
			// overshot
			return false
		}
	}

	return false
}
