package main

import "fmt"
import "strings"

var iterations int
var depth int
var maxdepth int

var board = [9][9]int{
	{1, 0, 7, 0, 0, 6, 4, 5, 0},
	{0, 2, 5, 3, 4, 0, 0, 0, 8},
	{0, 6, 0, 0, 0, 1, 0, 7, 0},
	{0, 5, 3, 0, 0, 0, 0, 2, 9},
	{6, 1, 0, 0, 0, 9, 8, 0, 0},
	{0, 0, 0, 6, 0, 2, 0, 0, 7},
	{0, 0, 1, 0, 9, 3, 2, 0, 0},
	{0, 0, 8, 0, 0, 0, 0, 0, 0},
	{0, 4, 0, 0, 7, 8, 5, 9, 1},
}

func printBoard() {
	for r:=0; r<9; r++ {
		for c:=0; c<9; c++ {
			fmt.Printf("%d ", board[r][c])
		}
		fmt.Println()
	}
}

func main() {
	fmt.Println("sudoku!")
	printBoard()
	iterations = 0

	if solveSudoku() {
		fmt.Println("Solved!")
		printBoard()
	} else {
		fmt.Println("No solution exists.")
	}

	fmt.Printf("Total iterations: %d\n", iterations)
	fmt.Printf("Max depth: %d\n", maxdepth)
}

func solveSudoku() bool {
	iterations++
	depth++
	fmt.Printf("%s\n", strings.Repeat("*", depth))
	if depth > maxdepth {
		maxdepth = depth
	}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == 0 {
				for num := 1; num <= 9; num++ {
					if isValid(i, j, num) {
						board[i][j] = num
						if solveSudoku() {
							depth--
							return true
						}
						board[i][j] = 0
					}
				}
				depth--
				return false
			}
		}
	}
	depth--
	return true
}

func isValid(row int, col int, num int) bool {
	for i := 0; i < 9; i++ {
		if board[row][i] == num || board[i][col] == num || board[row-row%3+i/3][col-col%3+i%3] == num {
			return false
		}
	}
	return true
}
