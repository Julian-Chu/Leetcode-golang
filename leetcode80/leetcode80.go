package leetcode80

func removeDuplicates(nums []int) int {
	n := len(nums)
	if len(nums) < 3 {
		return n
	}

	if nums[0] == nums[n-1] {
		return 2
	}

	var moveToTail func(int)
	moveToTail = func(i int) {
		if i == n-1 {
			return
		}

		nums[i], nums[i+1] = nums[i+1], nums[i]
		moveToTail(i + 1)
	}

	i := 0
	cnt := 0
	for i < n {
		if i > 0 && nums[i] < nums[i-1] {
			return i
		}
		if i < 2 {
			i++
			continue
		}

		if (nums[i] == nums[i-1]) && (nums[i] == nums[i-2]) {
			if i == n-1 {
				return i
			}
			if cnt > n-1-i {
				return i
			}
			moveToTail(i)
			cnt++
			continue
		}
		if i < n-1 && nums[i] > nums[i+1] {
			return i + 1
		}
		cnt = 0
		i++
	}

	return n
}
