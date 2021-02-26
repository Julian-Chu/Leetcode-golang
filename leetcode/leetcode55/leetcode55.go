package leetcode55

func canJump(nums []int) bool {
	n := len(nums) - 1
	if n == 0 {
		return true
	}
	var prevJumpPoint func(int) bool

	prevJumpPoint = func(cur int) bool {
		for i := cur - 1; i >= 0; i-- {
			if cur-i <= nums[i] {
				if i == 0 {
					return true
				}
				return prevJumpPoint(i)
			}
		}
		return false
	}
	return prevJumpPoint(n)
}
