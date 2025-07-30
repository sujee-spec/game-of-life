package gameoflife

type Matrix [][]int

type Board struct {
	matrix Matrix
}

func NewBoard(input [][]int, m, n int) *Board {
	mat := make(Matrix, m)
	for i := 0; i < m; i++ {
		mat[i] = make([]int, n)
		for j := 0; j < n; j++ {
			mat[i][j] = input[i][j]
		}
	}
	return &Board{matrix: mat}
}
