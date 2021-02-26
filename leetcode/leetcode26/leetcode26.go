package leetcode26

func removeDuplicates(nums []int) int {
	size := len(nums)
	cur := 0
	for i := 1; i < size; i++ {
		if nums[cur] != nums[i] {
			cur++
			nums[cur] = nums[i]
			continue
		}

	}
	return cur + 1
}
