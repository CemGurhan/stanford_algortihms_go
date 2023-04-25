package onedimension

func FindClosestPairBrute(oneDimensionalPlane []int) int {
	distanceBetweenPoints := -1
	pairCounter := 0

	for i := 0; i < len(oneDimensionalPlane); i++ {
		if oneDimensionalPlane[i] == 0 {
			continue
		} else if oneDimensionalPlane[i] == 1 {
			pairCounter++
			distanceBetweenPoints++
		}

		if pairCounter == 1 {
			distanceBetweenPoints++
		}

		if pairCounter == 2 {
			pairCounter = 0
		}
	}

	return distanceBetweenPoints
}

// [0 0 0 1 0 0 1 0 1 0 0 0 0 1 0 0 1]
