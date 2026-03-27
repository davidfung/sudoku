package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

var iterations int
var depth int
var maxdepth int
var flagGraph bool

var board [9][9]int

func loadBoard(unsolvedFile string) {
	file, err := os.Open(unsolvedFile)
	if err != nil {
		log.Fatal("Error opneing file:", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = ' '
	r := 8
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal("Error reading record:", err)
		}
		for c := range 9 {
			num := 0
			fmt.Sscanf(record[c], "%d", &num)
			board[r][c] = num
		}
		r--
	}
}

func printBoard() {
	for r := 8; r >= 0; r-- {
		for c := range 9 {
			fmt.Printf("%d ", board[r][c])
		}
		fmt.Println()
	}
}

func main() {
	showGraph := flag.Bool("g", false, "show graph")
	flag.Parse()
	if flag.NArg() != 1 {
		fmt.Println("Usage: sudoku [-g] <board>")
		fmt.Println(`where <board> is a space delimited file such as:
1 0 7 0 0 6 4 5 0
0 2 5 3 4 0 0 0 8
0 6 0 0 0 1 0 7 0
0 5 3 0 0 0 0 2 9
6 1 0 0 0 9 8 0 0
0 0 0 6 0 2 0 0 7
0 0 1 0 9 3 2 0 0
0 0 8 0 0 0 0 0 0
0 4 0 0 7 8 5 9 1
		`)
		return
	}
	flagGraph = *showGraph

	fmt.Println("sudoku!")
	loadBoard(flag.Args()[0])
	printBoard()

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
	if flagGraph {
		fmt.Printf("%s\n", strings.Repeat("*", depth))
	}
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
