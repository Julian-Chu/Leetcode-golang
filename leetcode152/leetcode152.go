package leetcode152

func maxProduct(nums []int) int {
	pos, neg, max := 1, 1, nums[0]
	for i := 0; i < len(nums); i++ {
		switch {
		case nums[i] > 0:
			pos, neg = nums[i]*pos, nums[i]*neg
		case nums[i] < 0:
			pos, neg = nums[i]*neg, nums[i]*pos
		default:
			pos, neg = 0, 1
		}

		if max < pos {
			max = pos
		}
		if pos <= 0 {
			pos = 1
		}
	}
	return max
}
