package solver

func IsSafe(mat [][]int, i, j, num int, row, col, box []int) bool {
	if (row[i]&(1<<num)) != 0 || (col[j]&(1<<num)) != 0 || (box[i/3*3+j/3]&(1<<num)) != 0 {
		return false
	}
	return true
}

func SudokuSolverRec(mat [][]int, i, j int, row, col, box []int) bool {
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
		return SudokuSolverRec(mat, i, j+1, row, col, box)
	}

	for num := 1; num < length+1; num++ {
		// Проверяем подходит ли число
		if IsSafe(mat, i, j, num, row, col, box) {
			mat[i][j] = num

			// Обновляем маску
			row[i] |= (1 << num)
			col[j] |= (1 << num)
			box[i/3*3+j/3] |= (1 << num)

			if SudokuSolverRec(mat, i, j+1, row, col, box) {
				return true
			}

			// Убираем число из маски
			mat[i][j] = 0
			row[i] = row[i] &^ (1 << num)
			col[j] = col[j] &^ (1 << num)
			box[i/3*3+j/3] = box[i/3*3+j/3] &^ (1 << num)
		}
	}

	return false
}

func SolveSudoku(mat [][]int) bool {
	length := len(mat)
	row := make([]int, length)
	col := make([]int, length)
	box := make([]int, length)

	// Создаем битовую маску из чисел
	for i := range length {
		for j := range length {
			if mat[i][j] != 0 {
				row[i] |= (1 << mat[i][j])
				col[j] |= (1 << mat[i][j])
				box[(i/3)*3+j/3] |= (1 << mat[i][j])
			}
		}
	}

	return SudokuSolverRec(mat, 0, 0, row, col, box)
}
