package inversions

import "fmt"

func Number_of_inversions_brute(array []int) {
	var number_of_inversions [][]int
	k := 0
	for i := 0; i < len(array); i++ {
		for j := i + 1; j < len(array); j++ {
			if array[i] > array[j] {
				number_of_inversions = append(number_of_inversions, []int{array[i], array[j]})
				k++
			}
		}
	}
	fmt.Println("Inverstions: ", number_of_inversions)
}

// [2, 3, 7, 8, 9, 1]
// [[2,1], [3,1], [7,1], ]
