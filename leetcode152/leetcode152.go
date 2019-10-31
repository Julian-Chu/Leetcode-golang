package leetcode152

func maxProduct(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	max := 0
	for i := 0; i < len(nums); i++ {
		temp := 1
		for j := i; j < len(nums); j++ {
			temp *= nums[j]
			if temp > max {
				max = temp
			}
			if temp == 0 {
				temp = 1
			}
		}
	}
	return max
}
