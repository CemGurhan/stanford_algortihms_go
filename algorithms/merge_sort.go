package algorithms

import "fmt"

func Merge_sort(array []int) []int {
	array_length := len(array)
	halfway_index := len(array) / 2

	if array_length <= 1 {
		return array
	}

	left_array := array[0:halfway_index]
	right_array := array[halfway_index:array_length]

	left_array_return := Merge_sort(left_array)
	right_array_return := Merge_sort(right_array)

	final_array := merge(left_array_return, right_array_return)

	return final_array
}

func merge(left_array []int, right_array []int) []int {
	i := 0
	j := 0
	k := 0
	final_array := make([]int, len(right_array)+len(left_array))
	for i < len(left_array) && j < len(right_array) {
		if left_array[i] < right_array[j] {
			final_array[k] = left_array[i]
			i++
			k++
		} else if right_array[j] < left_array[i] {
			final_array[k] = right_array[j]
			fmt.Println("Inversion found: [", left_array[i:], ", ", right_array[j], "]")
			j++
			k++
		}
	}

	if i < len(left_array) {
		copy(final_array[k:], left_array[i:])
	}

	if j < len(right_array) {
		copy(final_array[k:], right_array[j:])
	}

	return final_array
}

// [22, 4, -1, 3, 55]
// [4, 22] [-1, 3, 55]

// [3, 5, 7, 8, 2, 1]
// [3, 5, 7] [8, 2, 1]
// [3] [5, 7] [8] [2, 1]
// [3] [5] [7] [8] [2] [1]
// [3] [5, 7] [8] [1, 2]
// [3, 5, 7] [1, 2, 8] if B[j] < A[i] then we can assume the rest of A is greater than B[j] as B and A are sorted
// sorted array
