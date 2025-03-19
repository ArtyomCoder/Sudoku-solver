package validators

type BoxValidator struct {
	masks []int
}

func NewBoxValidator(length int) *BoxValidator {
	return &BoxValidator{masks: make([]int, length)}
}

func (bv *BoxValidator) IsValid(mat [][]int, i, j, num int) bool {
	return (bv.masks[i/3*3+j/3] & (1 << num)) == 0
}

func (bv *BoxValidator) MarkUsed(i, j, num int) {
	bv.masks[i/3*3+j/3] |= (1 << num)
}

func (bv *BoxValidator) UnmarkUsed(i, j, num int) {
	bv.masks[i/3*3+j/3] &^= (1 << num)
}
