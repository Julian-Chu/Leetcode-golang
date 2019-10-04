package leetcode26

func removeDuplicates(nums []int) int {
	l, r := 0, 1
	for ; r < len(nums); r++ {
		if nums[l] == nums[r] {
			continue
		}
		l++
		nums[l], nums[r] = nums[r], nums[l]
	}
	return l + 1
}
