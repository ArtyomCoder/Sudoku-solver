package models

type SudokuRequest struct {
	Matrix [][]int `json:"matrix"`
}
