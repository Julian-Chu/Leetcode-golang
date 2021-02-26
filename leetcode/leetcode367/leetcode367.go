package leetcode367

func isPerfectSquare(num int) bool {
	if num == 0 {
		return true
	}
	for i := 1; i*i <= num; i++ {
		if num%i == 0 && num/i == i {
			return true
		}
	}

	return false
}
