package leetcode53

func maxSubArray(nums []int) int {

	n := len(nums)
	if n == 0 {
		return 0
	}

	max := nums[0]
	sum := nums[0]
	for i := 1; i < n; i++ {
		sum += nums[i]
		if sum < nums[i] {
			sum = nums[i]
		}

		if sum > max {
			max = sum
		}
	}
	return max
}
