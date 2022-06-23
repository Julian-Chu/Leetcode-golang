package leetcode189

func rotate(nums []int, k int) {
	if len(nums) == 1 || k == 0 || k == len(nums) {
		return
	}

	copy(nums, append(nums[len(nums)-k:], nums[:len(nums)-k]...))
}

func rotate_1(nums []int, k int) {
	n := len(nums)
	k = k % n

	reverse(nums)
	reverse(nums[0:k])
	reverse(nums[k:])
}

func reverse(nums []int) {
	l, r := 0, len(nums)-1

	for l < r {
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}
}
