package leetcode342

func isPowerOfFour(num int) bool {
	if num == 0 {
		return false
	}
	i := num & 715827882
	if i != 0 {
		return false
	}
	com := ^num + 1
	j := num & com
	if j != num {
		return false
	}

	return true
}
