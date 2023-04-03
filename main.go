package main

import (
	"fmt"

	"github.com/cemgurhan/algorithms/algorithms"
	"github.com/cemgurhan/algorithms/algorithms/inversions"
	matrixmult "github.com/cemgurhan/algorithms/algorithms/matrix_mult"
)

func main() {
	matrix_one := [][]int{{2, 3}, {4, 5}}
	matrix_two := [][]int{{6, 7}, {8, 9}}
	matrixmult.Strassens_algorithm(matrix_one, matrix_two)
	// fmt.Println(output_matrix)
}

func call_Multiply_n_by_n_matrices_brute() {
	matrix_one := [][]int{{2, 3}, {4, 5}}
	matrix_two := [][]int{{6, 7}, {8, 9}}
	output_matrix := matrixmult.Multiply_n_by_n_matrices_brute(matrix_one, matrix_two)
	fmt.Println(output_matrix)
}

func call_number_of_inversions_efficient() {
	array := []int{3, 5, 7, 8, 2, 1}
	_, inversions := inversions.FindNumberOfInversions(array)
	fmt.Println(inversions)
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
