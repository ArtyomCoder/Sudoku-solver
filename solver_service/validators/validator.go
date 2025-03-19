package validators

type Validator interface {
	IsValid(mat [][]int, i, j, num int) bool
	MarkUsed(i, j, num int)
	UnmarkUsed(i, j, num int)
}
