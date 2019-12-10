package linearcodes

func GetParityCheckMatrix(matrix [][]bool, perm []int) [][]bool {
	n := len(matrix)
	if n == 0 {
		return matrix
	}
	m := len(matrix[0])

	A := make([][]bool, n)
	for i := 0; i < n; i++ {
		A[i] = make([]bool, m-n)
		for j := 0; j < m-n; j++ {
			A[i][j] = matrix[i][n+j]
		}
	}

	AT := TransposeMatrix(A)

	for i, _ := range AT {
		for j := 0; len(AT[i]) < m; j++ {
			AT[i] = append(AT[i], i == j)
		}
	}

	for i, val := range perm {
		swapColumns(AT, i, val-1)
	}

	return AT

}
