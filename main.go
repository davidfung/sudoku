package main

import "fmt"

var calls int

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

func main() {
	fmt.Println("sudoku!")
	fmt.Println(board)
	calls = 0

	if solveSudoku() {
		fmt.Println("Solved!")
		fmt.Println(board)
	} else {
		fmt.Println("No solution exists.")
	}

	fmt.Printf("Total iterations: %d\n", calls)
}

func solveSudoku() bool {
	calls++
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == 0 {
				for num := 1; num <= 9; num++ {
					if isValid(i, j, num) {
						board[i][j] = num
						if solveSudoku() {
							return true
						}
						board[i][j] = 0
					}
				}
				return false
			}
		}
	}
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
