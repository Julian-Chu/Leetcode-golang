package leetcode1342

func numberOfSteps(num int) int {
	step := 0
	for num > 0 {
		step++
		if num&0x1 == 1 {
			num -= 1
		} else {
			num >>= 1
		}
	}

	return step
}
