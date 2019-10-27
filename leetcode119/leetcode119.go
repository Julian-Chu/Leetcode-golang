package leetcode119

func getRow(rowIndex int) []int {

	rows := make([][]int, rowIndex+1)
	rows[0] = []int{1}

	if rowIndex == 0 {
		return rows[0]
	}

	for i := 1; i < rowIndex+1; i++ {
		rows[i] = getNextRow(rows[i-1])
	}
	return rows[rowIndex]
}

func getNextRow(ints []int) []int {
	res := make([]int, 0, len(ints)+1)
	res = append(append(res, 0), ints...)
	for i := 0; i < len(res)-1; i++ {
		res[i] = res[i] + res[i+1]
	}
	return res
}
