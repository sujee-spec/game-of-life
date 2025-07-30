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

var directions = [][]int{
	{-1, -1}, {-1, 0}, {-1, 1},
	{0, -1}, {0, 1},
	{1, -1}, {1, 0}, {1, 1},
}

func countLiveNeighbors(board [][]int, row, col int) int {
	m := len(board)
	n := len(board[0])
	liveCount := 0

	for _, dir := range directions {
		newRow := row + dir[0]
		newCol := col + dir[1]

		if newRow >= 0 && newRow < m && newCol >= 0 && newCol < n {
			if board[newRow][newCol] == 1 || board[newRow][newCol] == -1 {
				liveCount++
			}
		}
	}
	return liveCount
}
