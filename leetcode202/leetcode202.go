package leetcode202

func isHappy(n int) bool {
	res := make(map[int]bool)

	num := n

	for !res[num] {
		res[num] = true
		next := 0
		for num > 0 {
			root := num % 10
			num /= 10
			next += root * root
		}
		num = next
	}

	return res[1]
}
