package main

import (
	"fmt"

	"github.com/cemgurhan/algorithms/algorithms"
	"github.com/cemgurhan/algorithms/algorithms/inversions"
)

func main() {
	array := []int{2, 3, 7, 8, 9, 1}
	inversions.Number_of_inversions_brute(array)
}

func call_merge_sort() {
	array := []int{22, 4, -1, 3, 55}
	sorted_array := algorithms.Merge_sort(array)
	fmt.Println(sorted_array)
}
