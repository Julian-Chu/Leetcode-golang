package leetcode477

func totalHammingDistance(nums []int) int {
	res := 0
	n := len(nums)

	for i := 0; i < 32; i++ {
		k := 0
		for j := 0; j < n; j++ {
			k += (nums[j] >> uint(i)) & 1
		}

		res += k * (n - k)
	}
	return res
}
