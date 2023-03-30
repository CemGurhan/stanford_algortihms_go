package inversions

import "fmt"

func Merge_sort(array []int) []int {
	array_len := len(array)
	halfway_index := array_len / 2

	if array_len <= 1 {
		return array
	}

	right_array := array[halfway_index:array_len]
	left_array := array[0:halfway_index]

	right_array_return := Merge_sort(right_array)
	left_array_return := Merge_sort(left_array)

	sorted_array := merge(right_array_return, left_array_return)

	return sorted_array
}

func merge(right_array []int, left_array []int) []int {
	i := 0
	j := 0
	k := 0
	sorted_array := make([]int, len(right_array)+len(left_array))

	for i < len(right_array) && j < len(left_array) {
		if right_array[i] < left_array[j] {
			sorted_array[k] = right_array[i]
			fmt.Println("Inversion found: [", left_array[j:], ", ", right_array[i], "]")
			i++
			k++
		} else if left_array[j] < right_array[i] {
			sorted_array[k] = left_array[j]
			j++
			k++
		}
	}

	if i < len(right_array) {
		copy(sorted_array[k:], right_array[i:])
	}

	if j < len(left_array) {
		copy(sorted_array[k:], left_array[j:])
	}

	return sorted_array
}

// [3, 5, 7, 8, 2, 1]
// [3, 5, 7] [8, 2, 1]
// [3] [5, 7] [8] [2, 1]
// [3] [5] [7] [8] [2] [1]
// [3] [5, 7] [8] [1, 2]
// [3, 5, 7] [1, 2, 8]
