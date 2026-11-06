package main

import (
	"checkers/validations"
	"fmt"
	_ "strconv"
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

}

func startPosition(board *Board) {
	for row := 0; row < BoardSize; row++ {

		for col := 0; col < BoardSize; col++ {

			// 1 player
			if row < 3 && board[row][col] == "." {
				board[row][col] = "X"

			}

			// 2 player
			if row >= 5 && board[row][col] == "." {
				board[row][col] = "O"

			}
		}
	}
}

func appendMove(board *Board, fromRow, fromCol, toRow, toCol int, player string) bool {
	// fmt.Println(fromRow, fromCol, toRow, toCol)
	if fromRow < 0 || fromRow >= BoardSize ||
		fromCol < 0 || fromCol >= BoardSize ||
		toRow < 0 || toRow >= BoardSize ||
		toCol < 0 || toCol >= BoardSize {
		return false
	}

	//  Проверяем наличие своей шашки
	if board[fromRow][fromCol] != player {
		fmt.Println("Ошибка перемещения: Неправильный ход. На этом поле отсутствует шашка")
		return false
	}

	fmt.Println(fromRow, fromCol, toRow, toCol)
	//  Проверяем, свободна ли конечная клетка
	if board[toRow][toCol] != "." {
		fmt.Println("Ошибка перемещения: Неправильный ход. Перемещаемое поле не пустое")
		return false
	}

	//  Определяем расстояние перемещения
	rowDiff := toRow - fromRow
	colDiff := toCol - fromCol

	if absInt(rowDiff) != 1 || absInt(colDiff) != 1 {
		return false
	}

	//  Проверяем направление движения
	switch player {
	case "X":
		if rowDiff != 1 {
			return false
		}

	case "O":
		if rowDiff != -1 {
			return false
		}

	default:
		return false
	}
	applyMove(board, fromRow, fromCol, toRow, toCol)
	return true
}

func absInt(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func applyMove(board *Board, fromRow, fromCol, toRow, toCol int) {
	board[toRow][toCol] = board[fromRow][fromCol]
	board[fromRow][fromCol] = "."
}

func main() {

	// Slice := []string{"X", "O"}
	board := createBoard()
	startPosition(&board)
	printBoard(board)
	ok, secondvalue, firstvalue, fourthvalue, thirdvalue := validations.ValidateMove()

	if ok != true {
		fmt.Println("Ошибка!")
	}

	boola := appendMove(&board, secondvalue-1, firstvalue-1, fourthvalue-1, thirdvalue-1, "X")
	if boola == true {

		printBoard(board)

	}

}
