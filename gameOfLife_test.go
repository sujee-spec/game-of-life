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
	board := [][]int{
		{0, 1, 0},
		{1, 1, 0},
		{0, 0, 1},
	}

	got := countLiveNeighbors(board, 0, 0)
	want := 3

	if got != want {
		t.Errorf("countLiveNeighbors(0,0) = %d; want %d", got, want)
	}
}

func TestCountLiveNeighborsTopMiddle(t *testing.T) {
	board := [][]int{
		{0, 1, 0},
		{1, 1, 0},
		{0, 0, 1},
	}

	got := countLiveNeighbors(board, 0, 1)
	want := 2

	if got != want {
		t.Errorf("countLiveNeighbors(0,1) = %d; want %d", got, want)
	}
}
