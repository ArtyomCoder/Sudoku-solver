package handlers

import (
	"net/http"

	"example.com/sudoku-solver/models"
	"example.com/sudoku-solver/solver"
	"github.com/gin-gonic/gin"
)

func Solve(c *gin.Context) {
	var req models.SudokuRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Input"})
		return
	}

	if solver.SolveSudoku(req.Matrix) {
		c.JSON(http.StatusOK, gin.H{"solved": true, "matrix": req.Matrix})
		return
	}

	c.JSON(http.StatusOK, gin.H{"solved": false, "error": "No solution found"})
}
