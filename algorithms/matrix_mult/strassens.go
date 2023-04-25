package matrixmult

import "fmt"

func Strassens_algorithm(matrix_one [][]int, matrix_two [][]int) {
	matrix_map_one := make(map[int]int)
	matrix_map_two := make(map[int]int)
	matrix_product := make([][]int, len(matrix_one))
	quadrant_pointer := 1
	for i := 0; i < len(matrix_one); i++ {
		matrix_product[i] = make([]int, len(matrix_one))
		for j := 0; j < len(matrix_two); j++ {
			matrix_map_one[quadrant_pointer] = matrix_one[i][j]
			matrix_map_two[quadrant_pointer] = matrix_two[i][j]
			quadrant_pointer++
		}
	}
	matrix_product[0][0] = product_five(matrix_map_one, matrix_map_two) + product_four(matrix_map_one, matrix_map_two) -
		product_two(matrix_map_one, matrix_map_two) + product_six(matrix_map_one, matrix_map_two)

	matrix_product[0][1] = product_one(matrix_map_one, matrix_map_two) + product_two(matrix_map_one, matrix_map_two)

	matrix_product[1][0] = product_three(matrix_map_one, matrix_map_two) + product_four(matrix_map_one, matrix_map_two)

	matrix_product[1][1] = product_one(matrix_map_one, matrix_map_two) + product_five(matrix_map_one, matrix_map_two) -
		product_three(matrix_map_one, matrix_map_two) - product_seven(matrix_map_one, matrix_map_two)

	fmt.Println(matrix_product)
}

func product_one(matrix_map_one map[int]int, matrix_map_two map[int]int) int {
	return matrix_map_one[1] * (matrix_map_two[2] - matrix_map_two[4])
}

func product_two(matrix_map_one map[int]int, matrix_map_two map[int]int) int {
	return (matrix_map_one[1] + matrix_map_one[2]) * matrix_map_two[4]
}

func product_three(matrix_map_one map[int]int, matrix_map_two map[int]int) int {
	return (matrix_map_one[3] + matrix_map_one[4]) * matrix_map_two[1]
}

func product_four(matrix_map_one map[int]int, matrix_map_two map[int]int) int {
	return matrix_map_one[4] * (matrix_map_two[3] - matrix_map_two[1])
}

func product_five(matrix_map_one map[int]int, matrix_map_two map[int]int) int {
	return (matrix_map_one[1] + matrix_map_one[4]) * (matrix_map_two[1] + matrix_map_two[4])
}

func product_six(matrix_map_one map[int]int, matrix_map_two map[int]int) int {
	return (matrix_map_one[2] - matrix_map_one[4]) * (matrix_map_two[3] + matrix_map_two[4])
}

func product_seven(matrix_map_one map[int]int, matrix_map_two map[int]int) int {
	return (matrix_map_one[1] - matrix_map_one[3]) * (matrix_map_two[1] + matrix_map_two[2])
}
