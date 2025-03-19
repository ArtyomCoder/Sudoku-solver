package solver

import "example.com/sudoku-solver/validators"

func SudokuSolverRec(mat [][]int, i, j int, valids []validators.Validator) bool {
	length := len(mat)

	// Дошли до последней клетки последней строки
	if i == length-1 && j == length {
		return true
	}

	// Если дошли до последней клетки в строке перехожу на следующую строку
	if j == length {
		i++
		j = 0
	}

	// Если клетка заполнена, то идем дальше
	if mat[i][j] != 0 {
		return SudokuSolverRec(mat, i, j+1, valids)
	}

	for num := 1; num < length+1; num++ {
		// Проверяем подходит ли число
		isValid := true
		for _, v := range valids {
			if !v.IsValid(i, j, num) {
				isValid = false
				break
			}
		}

		if isValid {
			mat[i][j] = num

			// Обновляем маску
			for _, v := range valids {
				v.MarkUsed(i, j, num)
			}

			if SudokuSolverRec(mat, i, j+1, valids) {
				return true
			}

			// Убираем число из маски
			mat[i][j] = 0
			for _, v := range valids {
				v.UnmarkUsed(i, j, num)
			}
		}
	}

	return false
}

func SolveSudoku(mat [][]int) bool {
	length := len(mat)

	valids := []validators.Validator{
		validators.NewRowValidator(length),
		validators.NewColValidator(length),
		validators.NewBoxValidator(length),
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
