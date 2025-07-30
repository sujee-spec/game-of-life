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

func (b *Board) countLiveNeighbors(row, col int) int {
	m := len(b.matrix)
	n := len(b.matrix[0])
	liveCount := 0

	for _, dir := range directions {
		newRow := row + dir[0]
		newCol := col + dir[1]

		if newRow >= 0 && newRow < m && newCol >= 0 && newCol < n {
			if b.matrix[newRow][newCol] == 1 || b.matrix[newRow][newCol] == -1 {
				liveCount++
			}
		}
	}
	return liveCount
}

func (b *Board) GameOfLife() {
	m := len(b.matrix)
	if m == 0 {
		return
	}
	n := len(b.matrix[0])

	for i := range m {
		for j := 0; j < n; j++ {
			liveNeighbors := b.countLiveNeighbors(i, j)

			switch b.matrix[i][j] {
			case 1:
				if liveNeighbors < 2 || liveNeighbors > 3 {
					b.matrix[i][j] = -1
				}
			case 0:
				if liveNeighbors == 3 {
					b.matrix[i][j] = 2
				}
			}
		}
	}

	for i := range m {
		for j := 0; j < n; j++ {
			switch b.matrix[i][j] {
			case -1:
				b.matrix[i][j] = 0
			case 2:
				b.matrix[i][j] = 1
			}
		}
	}
}
