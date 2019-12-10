package linearcodes

import (
	"fmt"
	"io"
)

// PrintBooleanMatrix asd
func PrintBooleanMatrix(w io.WriteCloser, matrix [][]bool) {
	_, _ = fmt.Fprintf(w, "%d %d\n", len(matrix[0]), len(matrix))
	for i, _ := range matrix {
		for _, val := range matrix[i] {
			if val {
				_, _ = fmt.Fprintf(w, "1")
			} else {
				_, _ = fmt.Fprint(w, "0")
			}
		}
		_, _ = fmt.Fprintln(w, "")
	}
}

func TransposeMatrix(matrix [][]bool) [][]bool {
	n := len(matrix)
	if n == 0 {
		return matrix
	}
	m := len(matrix[0])

	res := make([][]bool, m)
	for i := 0; i < m; i++ {
		res[i] = make([]bool, n)
		for j := 0; j < n; j++ {
			res[i][j] = matrix[j][i]
		}
	}
	return res
}
