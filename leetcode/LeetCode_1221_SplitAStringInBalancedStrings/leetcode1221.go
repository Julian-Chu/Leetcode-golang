package LeetCode_1221_SplitAStringInBalancedStrings

func balancedStringSplit(s string) int {
	res := 0
	count := 0

	for _, ch := range s {
		if ch == 'R' {
			count++
		} else {
			count--
		}

		if count == 0 {
			res++
		}
	}

	return res
}
