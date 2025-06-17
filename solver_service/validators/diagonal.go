package validators

type DiagonalValidator struct {
	masks  map[int]int
	length int
}

func NewDiagonalValidator(length int) *DiagonalValidator {
	return &DiagonalValidator{masks: make(map[int]int, length), length: length}
}

func (dv *DiagonalValidator) IsValid(i, j, num int) bool {
	if i == j {
		return (dv.masks[0] & (1 << num)) == 0
	}
	if i+j == dv.length-1 {
		return (dv.masks[i+j] & (1 << num)) == 0
	}
	return true
}

func (dv *DiagonalValidator) MarkUsed(i, j, num int) {
	if i == j {
		dv.masks[0] |= (1 << num)
	}
	if i+j == dv.length-1 {
		dv.masks[i+j] |= (1 << num)
	}
}

func (dv *DiagonalValidator) UnmarkUsed(i, j, num int) {
	if i == j {
		dv.masks[0] &^= (1 << num)
	}
	if i+j == dv.length-1 {
		dv.masks[i+j] &^= (1 << num)
	}
}
