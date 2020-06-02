package leetcode457

func circularArrayLoop(nums []int) bool {
	if len(nums) < 2 {
		return false
	}
	for i := range nums {
		if checkCircle(nums, i) {
			return true
		}
	}
	return false
}

func checkCircle(nums []int, idx int) bool {
	slow, fast := idx, idx
	slow = getIdx(nums, slow)
	fast = getIdx(nums, fast)
	fast = getIdx(nums, fast)
	for slow != fast {
		slow = getIdx(nums, slow)
		fast = getIdx(nums, fast)
		fast = getIdx(nums, fast)
	}

	steps := make(map[int]bool)
	for _, ok := steps[slow]; !ok; _, ok = steps[slow] {
		steps[slow] = true
		pre := slow
		slow = getIdx(nums, slow)
		if (nums[pre] < 0 && nums[slow] > 0) || (nums[pre] > 0 && nums[slow] < 0) {
			return false
		}
	}

	if len(steps) > 1 {
		return true
	}

	return false
}

func getIdx(nums []int, cur int) int {
	n := len(nums)
	cur += nums[cur]
	cur %= n
	if cur < 0 {
		cur += n
	}

	return cur
}
