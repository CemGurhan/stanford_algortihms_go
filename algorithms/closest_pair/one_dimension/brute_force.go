package onedimension

func FindClosestPairBrute(oneDimensionalPlane []int) (int, []int) {
	distanceBetweenPointsOne := 0
	distanceBetweenPointsTwo := 0
	closestDistanceBetweenPoints := 0
	closestPointPairOne := make([]int, 2)
	closestPointPairTwo := make([]int, 2)
	finalClosestPair := make([]int, 2)
	pairFound := false
	iterations := 0

	for i := 0; i < len(oneDimensionalPlane); i++ {
		if i+1 == len(oneDimensionalPlane) && distanceBetweenPointsTwo == 0 {
			closestDistanceBetweenPoints = distanceBetweenPointsOne
			finalClosestPair = closestPointPairOne
			break
		}

		if distanceBetweenPointsOne < distanceBetweenPointsTwo && distanceBetweenPointsOne != 0 {
			closestDistanceBetweenPoints = distanceBetweenPointsOne
			finalClosestPair = closestPointPairOne
		} else if distanceBetweenPointsTwo < distanceBetweenPointsOne && distanceBetweenPointsTwo != 0 {
			closestDistanceBetweenPoints = distanceBetweenPointsTwo
			finalClosestPair = closestPointPairTwo
			distanceBetweenPointsOne = distanceBetweenPointsTwo
			closestPointPairOne = []int{closestPointPairTwo[0], closestPointPairTwo[1]}
			distanceBetweenPointsTwo = 0
		}

		if oneDimensionalPlane[i] == 0 {
			continue
		}

		iterations = 0

		for j := i + 1; j < len(oneDimensionalPlane); j++ {
			if oneDimensionalPlane[j] == 1 {
				if !pairFound {
					closestPointPairOne[1] = j
					if distanceBetweenPointsOne == 0 {
						closestDistanceBetweenPoints = 0
						finalClosestPair = []int{j - 1, j}
						closestPointPairOne = []int{j - 1, j}
						distanceBetweenPointsOne++
						break
					}
				} else {
					closestPointPairTwo[1] = j
					if distanceBetweenPointsTwo == 0 {
						closestDistanceBetweenPoints = 0
						distanceBetweenPointsOne = distanceBetweenPointsTwo
						finalClosestPair = []int{j - 1, j}
						closestPointPairOne = []int{j - 1, j}
						break
					}
				}

				pairFound = true
				break
			}
			if !pairFound {
				distanceBetweenPointsOne++
				if iterations == 0 {
					closestPointPairOne[0] = j - 1
					iterations++
				}
			} else {
				distanceBetweenPointsTwo++
				if iterations == 0 {
					closestPointPairTwo[0] = j - 1
					iterations++
				}
			}
		}
	}

	return closestDistanceBetweenPoints, finalClosestPair
}

// [0 0 0 1 0 0 1 0 1 0 0 0 0 1 0 0 1]
// [0 0 1 0 1 0 0 0 0 1 0 0 1]
// [1 0 1 1 0 0 1]
