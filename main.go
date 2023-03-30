package main

import (
	"fmt"

	"github.com/cemgurhan/algorithms/algorithms"
	"github.com/cemgurhan/algorithms/algorithms/inversions"
)

func main() {
	call_number_of_inversions_efficient()
}

func call_number_of_inversions_efficient() {
	array := []int{3, 5, 7, 8, 2, 1}
	a := inversions.Merge_sort(array)
	fmt.Println(a)
}

func call_number_of_inversions_brute() {
	array := []int{2, 3, 7, 8, 9, 1}
	inversions.Number_of_inversions_brute(array)
}

func call_merge_sort() {
	array := []int{22, 4, -1, 3, 55}
	sorted_array := algorithms.Merge_sort(array)
	fmt.Println(sorted_array)
}
