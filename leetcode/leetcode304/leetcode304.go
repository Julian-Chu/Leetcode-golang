package leetcode304

type NumMatrix struct {
	dp [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	m := len(matrix)
	if m == 0 {
		return NumMatrix{dp: nil}
	}

	n := len(matrix[0])

	dp := make([][]int, m)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n+1)
	}

	for i := 0; i < m; i++ {
		for j := 1; j <= n; j++ {
			dp[i][j] = dp[i][j-1] + matrix[i][j-1]
		}
	}
	return NumMatrix{dp: dp}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	if this.dp == nil {
		return 0
	}
	sum := 0
	for i := row1; i <= row2; i++ {
		sum += this.dp[i][col2+1] - this.dp[i][col1]
	}

	return sum
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
