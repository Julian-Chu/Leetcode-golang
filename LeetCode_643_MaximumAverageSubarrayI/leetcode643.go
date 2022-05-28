package LeetCode_643_MaximumAverageSubarrayI

func findMaxAverage(nums []int, k int) float64 {
	maxSum := 0
	sum := 0
	for i := 0; i < len(nums); i++ {
		if i < k {
			sum += nums[i]
			maxSum = sum
			continue
		}

		sum = sum + nums[i] - nums[i-k]
		if sum > maxSum {
			maxSum = sum
		}
	}

	return float64(maxSum) / float64(k)
}
