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
