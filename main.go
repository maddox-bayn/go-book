//Create a 3×3 tic-tac-toe board using var board [3][3]rune and let the middle be x

package main

import (
	"fmt"
)

func main() {
	fmt.Println("Tic-Tac-Toe Board:")

	var board [3][3]rune

	// Initialize board: place 'X' in the center, spaces elsewhere
	for i := range board {
		for j := range board[i] {
			if i == 1 && j == 1 {
				board[i][j] = 'X'
			} else {
				board[i][j] = ' '
			}
		}
	}

	// Print rows
	for _, row := range board {
		fmt.Printf("|%c|%c|%c|\n", row[0], row[1], row[2])
	}
}
