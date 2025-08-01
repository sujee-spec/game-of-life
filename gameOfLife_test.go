package gameoflife

import (
	"testing"
)

func TestNewBoard(t *testing.T) {
	input := [][]int{
		{0, 1, 0},
		{1, 0, 1},
		{0, 0, 1},
		{1, 1, 0},
	}
	m := 4
	n := 3
	b := NewBoard(input, m, n)
	if b == nil {
		t.Errorf("Board should be created with %v*%v size", m, n)
	}
}

func TestCountLiveNeighborsTopLeftCorner(t *testing.T) {
	board := &Board{
		matrix: Matrix{
			{0, 1, 0},
			{1, 1, 0},
			{0, 0, 1},
		},
	}

	got := board.countLiveNeighbors(0, 0)
	want := 3

	if got != want {
		t.Errorf("countLiveNeighbors(0,0) = %d; want %d", got, want)
	}
}

func TestCountLiveNeighborsCenter(t *testing.T) {
	board := &Board{
		matrix: Matrix{
			{0, 1, 0},
			{1, 1, 0},
			{0, 0, 1},
		},
	}

	got := board.countLiveNeighbors(1, 1)
	want := 3

	if got != want {
		t.Errorf("countLiveNeighbors(1,1) = %d; want %d", got, want)
	}
}

func TestGameOfLifeWithExpectedMatrix(t *testing.T) {
	initial := Matrix{
		{0, 1, 0},
		{0, 0, 1},
		{1, 1, 1},
		{0, 0, 0},
	}
	expected := Matrix{
		{0, 0, 0},
		{1, 0, 1},
		{0, 1, 1},
		{0, 1, 0},
	}

	board := &Board{matrix: initial}
	board.GameOfLife()

	for i := range expected {
		for j := range expected[i] {
			if board.matrix[i][j] != expected[i][j] {
				t.Errorf("Mismatch at (%d, %d): got %d, want %d", i, j, board.matrix[i][j], expected[i][j])
			}
		}
	}
}
