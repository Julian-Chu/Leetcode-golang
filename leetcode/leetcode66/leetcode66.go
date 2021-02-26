package leetcode66

func plusOne(digits []int) []int {
	n := len(digits)
	addOne := true
	for i := n - 1; i >= 0; i-- {
		if addOne {
			if digits[i]+1 > 9 {
				digits[i] = 0
				continue
			}
			digits[i] += 1
			addOne = false
		}
		break
	}

	if addOne {
		digits = append([]int{1}, digits...)
	}
	return digits
}
