package onedimension

import (
	"testing"

	"github.com/stretchy/testify/assert"
)

func TestFindClosestPairBrute(t *testing.T) {
	oneDimensionalPlane := []int{0, 0, 0, 1, 0, 0, 1, 0, 1, 0, 0, 0, 0, 1, 0, 0, 1}

	expectedDistanceBetweenPoints := FindClosestPairBrute(oneDimensionalPlane)

	assert.Equal(expectedDistanceBetweenPoints, 2)
}
