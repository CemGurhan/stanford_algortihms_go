package algorithms

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

// [22, 33, 4, 5, 6]
// [22, 33]
// [22] [33]
// [22, 33]
// [4, 5, 6, 22, 33]
