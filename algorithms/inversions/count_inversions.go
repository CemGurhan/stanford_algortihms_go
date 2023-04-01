package inversions

func FindNumberOfInversions(array []int) ([]int, int) {
	array_len := len(array)
	halfway_index := array_len / 2
	number_of_inversions := 0

	if array_len <= 1 {
		return array, 0
	}

	left_array, left_number_of_inversions := FindNumberOfInversions(array[0:halfway_index])
	right_array, right_number_of_inversions := FindNumberOfInversions(array[halfway_index:array_len])

	number_of_inversions += left_number_of_inversions
	number_of_inversions += right_number_of_inversions

	sorted_array, returned_number_of_inversions := merge_and_find_inversions(left_array, right_array, number_of_inversions)

	return sorted_array, returned_number_of_inversions
}

func merge_and_find_inversions(left_array []int, right_array []int, number_of_inversions int) ([]int, int) {
	i := 0
	j := 0
	k := 0
	sorted_array := make([]int, len(left_array)+len(right_array))
	for i < len(left_array) && j < len(right_array) {
		if left_array[i] < right_array[j] {
			sorted_array[k] = left_array[i]
			i++
			k++
		} else if right_array[j] < left_array[i] {
			sorted_array[k] = right_array[j]
			number_of_inversions += len(left_array[i:])
			j++
			k++
		}
	}

	if i < len(left_array) {
		copy(sorted_array[k:], left_array[i:])
	}

	if j < len(right_array) {
		copy(sorted_array[k:], right_array[j:])
	}

	return sorted_array, number_of_inversions
}

// [3, 5, 7, 8, 2, 1]
// [3, 5, 7] [8, 2, 1]
// [3] [5, 7] [8] [2, 1]
// [3] [5] [7] [8] [2] [1]
// [3] [5, 7] [8] [1, 2]
// [3, 5, 7] [1, 2, 8]
