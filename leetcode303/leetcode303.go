package leetcode303

type NumArray struct {
	dp [][]int
}

func Constructor(nums []int) NumArray {
	n := len(nums)

	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		dp[i][i] = nums[i]
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j || j == 0 {
				continue
			}

			dp[i][j] = dp[i][j-1] + dp[j][j]
		}
	}

	return NumArray{dp: dp}
}

func (this *NumArray) SumRange(i int, j int) int {
	return this.dp[i][j]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(i,j);
 */
