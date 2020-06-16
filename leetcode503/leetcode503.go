package leetcode503

func nextGreaterElements(nums []int) []int {
	n := len(nums)
	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = -1

		for j := i + 1; ; j++ {
			j %= n
			if j == i {
				break
			}
			if nums[j] > nums[i] {
				res[i] = nums[j]
				break
			}
		}
	}

	return res
}
