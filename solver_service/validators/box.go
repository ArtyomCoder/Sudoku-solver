package validators

var boxLengthInfo = map[int][2]int{
	6:  {2, 3},
	8:  {2, 4},
	9:  {3, 3},
	10: {2, 5},
	12: {3, 4},
	14: {2, 7},
	15: {3, 5},
	16: {4, 4},
	18: {3, 6},
	20: {4, 5},
	21: {3, 7},
	22: {2, 11},
	24: {4, 6},
	25: {5, 5},
}

type BoxValidator struct {
	masks  []int
	length int
}

func NewBoxValidator(length int) *BoxValidator {
	return &BoxValidator{masks: make([]int, length), length: length}
}

func (bv *BoxValidator) IsValid(i, j, num int) bool {
	boxIndex := bv.getBoxIndex(i, j)
	return (bv.masks[boxIndex] & (1 << num)) == 0
}

func (bv *BoxValidator) MarkUsed(i, j, num int) {
	boxIndex := bv.getBoxIndex(i, j)
	bv.masks[boxIndex] |= (1 << num)
}

func (bv *BoxValidator) UnmarkUsed(i, j, num int) {
	boxIndex := bv.getBoxIndex(i, j)
	bv.masks[boxIndex] &^= (1 << num)
}

func (bv *BoxValidator) getBoxIndex(i, j int) int {
	i_step, j_step := boxLengthInfo[bv.length][0], boxLengthInfo[bv.length][1]
	return i/i_step*i_step + j/j_step
}
