package matrixmult

func Multiply_n_by_n_matrices(matrix_one [][]int, matrix_two [][]int) [][]int {
	output_matrix := make([][]int, len(matrix_one))
	for i := 0; i < len(matrix_one); i++ {
		output_matrix[i] = make([]int, len(matrix_two[0])) // initialize slice rows
		for j := 0; j < len(matrix_one[i]); j++ {
			for k := 0; k < len(matrix_two); k++ {
				output_matrix[i][j] = output_matrix[i][j] + (matrix_one[i][k] * matrix_two[k][j])
			}
		}
	}
	return output_matrix
}

// [[2, 3], [[6, 7],
// [4, 5]]  [8, 9]]

// (A[0][0] * B[0][0]) + (A[0][1] * B[1][0])
// (A[0][0] * B[0][1]) + (A[0][1] * B[1][1])
// (A[1][0] * B[0][0]) + (A[1][1] * B[1][0])
