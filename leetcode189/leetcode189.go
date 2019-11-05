package leetcode189

func rotate(nums []int, k int) {
	if len(nums) == 1 || k == 0 || k == len(nums) {
		return
	}

	k %= len(nums)
	copy(nums, append(nums[len(nums)-k:], nums[:len(nums)-k]...))
}
