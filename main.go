package main

import "fmt"


func gameOfLife(board [][]int) {
    rowCount := len(board)
    colCount := len(board[0])

    directions := [][]int{
        {-1, -1}, {-1, 0}, {-1, 1},
        {0, -1},           {0, 1},
        {1, -1},  {1, 0},  {1, 1},
    }

    for row := 0; row < rowCount; row++ {
        for col := 0; col < colCount; col++ {
            liveNeighborCount := 0

            for _, direction := range directions {
                neighborRow := row + direction[0]
                neighborCol := col + direction[1]

                if neighborRow >= 0 && neighborRow < rowCount &&
                   neighborCol >= 0 && neighborCol < colCount {

                    if board[neighborRow][neighborCol] == 1 || board[neighborRow][neighborCol] == -1 {
                        liveNeighborCount++
                    }
                }
            }

            if board[row][col] == 1 {
               
                if liveNeighborCount < 2 || liveNeighborCount > 3 {
                    board[row][col] = -1 
                }
            } else {
                
                if liveNeighborCount == 3 {
                    board[row][col] = 2 
                }
            }
        }
    }

    for row := range rowCount {
        for col := range colCount {
            switch board[row][col] {
            case -1: // Alive → Dead
                board[row][col] = 0 
            case 2: // Dead → Alive
                board[row][col] = 1 
            }
        }
    }
}

func printBoard(board [][]int) {
    for _, row := range board {
        fmt.Println(row)
    }
}

func main() {
    board := [][]int{
        {0, 1, 0},
        {0, 0, 1},
        {1, 1, 1},
        {0, 0, 0},
    }

    fmt.Println("Original Board:")
    printBoard(board)

    gameOfLife(board)

    fmt.Println("\nNext State:")
    printBoard(board)
}

