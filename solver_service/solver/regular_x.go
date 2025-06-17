package solver

import "example.com/sudoku-solver/validators"

type RegularXSudoku struct{}

func (rs *RegularXSudoku) Solve(mat [][]int) bool{
	length := len(mat)

	valids := []validators.Validator{
		validators.NewRowValidator(length),
		validators.NewColValidator(length),
		validators.NewBoxValidator(length),
		validators.NewDiagonalValidator(length),
	}

	// Создаем битовую маску из чисел
	for i := range length {
		for j := range length {
			if mat[i][j] != 0 {
				for _, v := range valids {
					v.MarkUsed(i, j, mat[i][j])
				}
			}
		}
	}

	return SudokuSolverRec(mat, 0, 0, valids)
}