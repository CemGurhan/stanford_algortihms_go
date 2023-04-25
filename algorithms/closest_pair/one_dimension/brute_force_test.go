package onedimension

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindClosestPairBrute_One(t *testing.T) {
	oneDimensionalPlane := []int{0, 0, 0, 1, 0, 0, 1, 0, 1, 0, 0, 0, 0, 1, 0, 0, 1}

	actualDistanceBetweenPoints, actualPair := FindClosestPairBrute(oneDimensionalPlane)

	assert.Equal(t, 1, actualDistanceBetweenPoints)
	assert.Equal(t, []int{6, 8}, actualPair)

}

func TestFindClosestPairBrute_Two(t *testing.T) {
	oneDimensionalPlane := []int{0, 0, 1, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 1}

	actualDistanceBetweenPoints, actualPair := FindClosestPairBrute(oneDimensionalPlane)

	assert.Equal(t, 2, actualDistanceBetweenPoints)
	assert.Equal(t, []int{2, 5}, actualPair)

}

func TestFindClosestPairBrute_Three(t *testing.T) {
	oneDimensionalPlane := []int{1, 0, 1, 1, 0, 0, 1}

	actualDistanceBetweenPoints, actualPair := FindClosestPairBrute(oneDimensionalPlane)

	assert.Equal(t, 0, actualDistanceBetweenPoints)
	assert.Equal(t, []int{2, 3}, actualPair)

}

func TestFindClosestPairBrute_Four(t *testing.T) {
	oneDimensionalPlane := []int{1, 1, 0, 1, 0, 0, 1}

	actualDistanceBetweenPoints, actualPair := FindClosestPairBrute(oneDimensionalPlane)

	assert.Equal(t, 0, actualDistanceBetweenPoints)
	assert.Equal(t, []int{0, 1}, actualPair)

}

func TestFindClosestPairBrute_Five(t *testing.T) {
	oneDimensionalPlane := []int{1, 0, 0, 0, 0, 0, 1}

	actualDistanceBetweenPoints, actualPair := FindClosestPairBrute(oneDimensionalPlane)

	assert.Equal(t, 5, actualDistanceBetweenPoints)
	assert.Equal(t, []int{0, 6}, actualPair)
}
