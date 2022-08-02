package LeetCode_27_RemoveElement

func removeElement(nums []int, val int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		for left <= right && nums[right] == val {
			right--
		}

		for left <= right && nums[left] != val {
			left++
		}

		if left <= right {
			nums[left], nums[right] = nums[right], nums[left]
			left++
			right--
		}
	}

	return left
}

func removeElement_1(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}

	slow := 0

	for i := range nums {
		if nums[i] != val {
			nums[slow] = nums[i]
			slow++
		}
	}

	return slow
}
