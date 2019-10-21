package leetcode80

func removeDuplicates(nums []int) int {
	n := len(nums)

	i := 2
	for j := i; j < n; j++ {
		if nums[i-2] != nums[j] {
			nums[i] = nums[j]
			i++
		}
	}
	return i
}
