
//Problema de las N Reinas sin concurrencia

package main
import (
	"fmt"
)

func isValid(board []int, row int, col int) bool {
	for row_i := 0; row_i < row; row_i++ {
		col_i := board[row_i]
		delta := row - row_i
		if col == col_i || col == col_i+delta || col == col_i-delta {
			return false
		}
	}
	return true
}

func nqueens(board []int, row int) {
	n := len(board)

	if row == n {
		draw(board)
	} else {
		for col := 0; col < n; col++ {
			if isValid(board, row, col) {
				board[row] = col
				nqueens(board, row+1)
			}
		}
	}
}

func draw(board []int) {
	n := len(board)
	for row := 0; row < n; row++ {
		rowString := ""
		for col := 0; col < n; col++ {
			if board[row] == col {
				rowString += "Q"
			} else {
				rowString += "."
			}
		}
		fmt.Println(rowString)
	}
	fmt.Println()
}

func main() {
	n := 8
	board := make([]int, n)
	nqueens(board, 0)
}