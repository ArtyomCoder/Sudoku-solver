package validators

type RowValidator struct {
	masks []int
}

func NewRowValidator(length int) *RowValidator {
	return &RowValidator{masks: make([]int, length)}
}

func (rv *RowValidator) IsValid(mat [][]int, i, j, num int) bool {
	return (rv.masks[i] & (1 << num)) == 0
}

func (rv *RowValidator) MarkUsed(i, j, num int) {
	rv.masks[i] |= (1 << num)
}

func (rv *RowValidator) UnmarkUsed(i, j, num int) {
	rv.masks[i] &^= (1 << num)
}
