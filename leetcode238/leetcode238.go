package leetcode238

func productExceptSelf(nums []int) []int {
	p := 1
	zeros := 0
	for _, v := range nums {
		if v == 0 {
			zeros++
			continue
		}
		p *= v
	}

	for i := range nums {
		switch zeros {
		case 0:
			nums[i] = p / nums[i]
		case 1:
			if nums[i] == 0 {
				nums[i] = p
				continue
			}
			nums[i] = 0
		default:
			nums[i] = 0
		}
	}

	return nums
}
