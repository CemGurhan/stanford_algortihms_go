package main

import (
	"fmt"

	"github.com/cemgurhan/algorithms/algorithms"
)

func main() {
	array := []int{22, 4, -1, 3, 55}
	sorted_array := algorithms.Merge_sort(array)
	fmt.Println(sorted_array)
}
