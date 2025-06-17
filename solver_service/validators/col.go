package validators

type ColValidator struct {
	masks []int
}

func NewColValidator(length int) *ColValidator {
	return &ColValidator{masks: make([]int, length)}
}

func (cv *ColValidator) IsValid(i, j, num int) bool {
	return (cv.masks[j] & (1 << num)) == 0
}

func (cv *ColValidator) MarkUsed(i, j, num int) {
	cv.masks[j] |= (1 << num)
}

func (cv *ColValidator) UnmarkUsed(i, j, num int) {
	cv.masks[j] &^= (1 << num)
}
