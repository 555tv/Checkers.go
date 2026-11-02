package main

import (
	"fmt"
)

const BoardSize = 8

type Board [BoardSize][BoardSize]string

func createBoard() Board {
	var board Board

	for row := 0; row < BoardSize; row++ {
		for col := 0; col < BoardSize; col++ {
			if (row+col)%2 == 0 {
				board[row][col] = "."
			} else {
				board[row][col] = " "
			}
		}
	}

	return board
}

func printBoard(board Board) {
	fmt.Println("    A   B   C   D   E   F   G   H")

	for row := 0; row < BoardSize; row++ {
		fmt.Printf("%d ", row+1)

		for col := 0; col < BoardSize; col++ {
			fmt.Printf("| %s ", board[row][col])

		}
		fmt.Printf("|\n")
	}
	fmt.Println("|")

}

func startPosition(board *Board) {
	for row := 0; row < BoardSize; row++ {
		for col := 0; col < BoardSize; col++ {

			// 1 player
			if row < 3 && board[row][col] == "." {
				board[row][col] = "X"
				fmt.Printf("\n X")
			}

			// 2 player
			if row >= 5 && board[row][col] == "." {
				board[row][col] = "O"
				fmt.Printf("O")
			}
		}
	}
}
func main() {
	board := createBoard()
	printBoard(board)
	startPosition(&board)
}
