package linearcodes

import (
	"fmt"
	"io"
)

func ReadMatrix(r io.Reader) [][]bool {
	var n, m int
	_, _ = fmt.Fscanf(r, "%d %d", &m, &n)

	matrix := make([][]bool, n)
	for i := 0; i < n; i++ {
		var s string
		_, _ = fmt.Fscanf(r, "%s", &s)
		matrix[i] = make([]bool, m)
		for j, ch := range s {
			matrix[i][j] = ch == '1'
		}

	}

	return matrix
}

func findTrues(matrix [][]bool, col int, from int) []int {
	res := make([]int, 0)
	for i := from; i < len(matrix); i++ {
		if matrix[i][col] {
			res = append(res, i)
		}
	}
	return res
}

func sortedPermutation(m int) []int {
	perm := make([]int, m)
	for i, _ := range perm {
		perm[i] = i
	}
	return perm
}

func makePermutationIndexedOne(perm []int) {
	for i, v := range perm {
		perm[i] = v + 1
	}
}

func swapColumns(matrix [][]bool, col1 int, col2 int) {
	for i := 0; i < len(matrix); i++ {
		matrix[i][col1], matrix[i][col2] = matrix[i][col2], matrix[i][col1]
	}
}

func swapUntilNonzero(matrix [][]bool, col int, perm []int, lastNonZero *int) []int {
	trues := findTrues(matrix, col, col)
	for {
		if len(trues) == 0 {
			swapColumns(matrix, col, *lastNonZero)
			perm[col], perm[*lastNonZero] = perm[*lastNonZero], perm[col]
			*lastNonZero--
			trues = findTrues(matrix, col, col)
		} else {
			break
		}
	}
	return trues
}

func xorRowBy(matrix [][]bool, row1 int, row2 int) {
	for i, _ := range matrix[row1] {
		matrix[row1][i] = matrix[row1][i] != matrix[row2][i]
	}
}

func NormalizeMatrix(matrix [][]bool) ([][]bool, []int) {

	n := len(matrix)
	if n == 0 {
		return matrix, []int{}
	}
	m := len(matrix[0])
	perm := sortedPermutation(m)

	lastNonZero := m - 1

	for col := 0; col < n; col++ {

		trues := swapUntilNonzero(matrix, col, perm, &lastNonZero)

		if !matrix[col][col] {
			xorRowBy(matrix, col, trues[0])
		}

		for _, val := range findTrues(matrix, col, 0) {
			if val != col {
				xorRowBy(matrix, val, col)
			}
		}
	}

	makePermutationIndexedOne(perm)

	return matrix, perm
}
