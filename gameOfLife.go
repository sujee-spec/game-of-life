package gameoflife

type board struct {
	matrix matrix
}

type matrix [][]int

func NewBoard(m, n int) *board {
	matrix := make([][]int, m)

	for i := range matrix {
		matrix[i] = make([]int, n)
	}
	return &board{matrix}
}
