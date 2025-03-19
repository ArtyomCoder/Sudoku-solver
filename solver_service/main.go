package main

import (
	"fmt"

	"example.com/sudoku-solver/solver"
)

func main() {
	mat := [][]int{
		{3, 0, 6, 5, 0, 8, 4, 0, 0},
		{5, 2, 0, 0, 0, 0, 0, 0, 0},
		{0, 8, 7, 0, 0, 0, 0, 3, 1},
		{0, 0, 3, 0, 1, 0, 0, 8, 0},
		{9, 0, 0, 8, 6, 3, 0, 0, 5},
		{0, 5, 0, 0, 9, 0, 6, 0, 0},
		{1, 3, 0, 0, 0, 0, 2, 5, 0},
		{0, 0, 0, 0, 0, 0, 0, 7, 4},
		{0, 0, 5, 2, 0, 6, 3, 0, 0},
	}

	fmt.Println("Нерешенное судоку:")
	PrintMat(mat)

	fmt.Println()

	fmt.Println("Решенное судоку:")
	solver.SolveSudoku(mat)
	PrintMat(mat)

	// r := gin.Default()

	// r.POST("/solve", handlers.Solve)

	// r.Run(":8080")
}

func PrintMat(mat [][]int) {
	for _, row := range mat {
		for _, elem := range row {
			fmt.Print(elem, " ")
		}
		fmt.Println()
	}
}
