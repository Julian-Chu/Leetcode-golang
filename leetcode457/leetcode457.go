package leetcode457

func circularArrayLoop(nums []int) bool {
	next := func(i int) int {
		return (len(nums) + (i+nums[i])%len(nums)) % len(nums)
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			continue
		}

		slow, fast := i, next(i)

		for nums[fast]*nums[slow] > 0 && nums[next(fast)]*nums[slow] > 0 {
			if fast == slow {
				if next(fast) != fast {
					return true
				}
				break
			}

			fast = next(next(fast))
			slow = next(slow)
		}

		slow = i
		numsi := nums[i]

		// mark slow to disable the route includes the index
		for nums[slow]*numsi > 0 {
			nums[slow], slow = 0, next(slow)
		}
	}
	return false
}
