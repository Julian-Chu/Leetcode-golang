package leetcode303

type NumArray struct {
	dp []int
}

func Constructor(nums []int) NumArray {
	dp := make([]int, len(nums)+1)
	dp[0] = 0

	for i := 1; i <= len(nums); i++ {
		dp[i] = dp[i-1] + nums[i-1]
	}
	return NumArray{dp: dp}
}

func (this *NumArray) SumRange(i int, j int) int {
	return this.dp[j+1] - this.dp[i]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(i,j);
 */
