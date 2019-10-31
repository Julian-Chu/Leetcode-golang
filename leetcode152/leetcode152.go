package leetcode152

func maxProduct(nums []int) int {
	var maxInts = func(nums []int) int {
		max := nums[0]
		for i := 1; i < len(nums); i++ {
			if nums[i] > max {
				max = nums[i]
			}
		}
		return max
	}

	var minInts = func(nums []int) int {
		min := nums[0]
		for i := 1; i < len(nums); i++ {
			if nums[i] < min {
				min = nums[i]
			}
		}
		return min
	}
	n := len(nums)
	max := make([]int, n, n)
	min := make([]int, n, n)

	max[0] = nums[0]
	min[0] = nums[0]
	ans := nums[0]

	for i := 1; i < n; i++ {
		v1 := max[i-1] * nums[i]
		v2 := min[i-1] * nums[i]

		max[i] = maxInts([]int{v1, v2, nums[i]})
		min[i] = minInts([]int{v1, v2, nums[i]})

		if ans < max[i] {
			ans = max[i]
		}
	}

	return ans
}
