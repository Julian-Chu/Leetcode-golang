package leetcode168

func convertToTitle(n int) string {
	res := ""
	for n > 26 {

		rest := n % 26
		n = n / 26
		if rest == 0 {
			res = string('Z') + res
			n--
		} else {
			res = string('A'+rest-1) + res
		}
	}

	res = string('A'+n-1) + res
	return res
}
