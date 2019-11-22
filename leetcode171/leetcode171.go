package leetcode171

func titleToNumber(s string) int {
	sum := 0
	for _, v := range s {
		v := v - 'A' + 1
		sum *= 26
		sum += int(v)
	}

	return sum
}
