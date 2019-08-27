package leetcode712

func minimumDeleteSum(s1 string, s2 string) int {
	n1, n2 := len(s1), len(s2)

	res := make([][]int, n1+1)
	for index := range res {
		res[index] = make([]int, n2+1)
	}
	for i := n1 - 1; i >= 0; i-- {
		res[i][n2] = res[i+1][n2] + int(s1[i])
	}

	for j := n2 - 1; j >= 0; j-- {
		res[n1][j] = res[n1][j+1] + int(s2[j])
	}
	for i := n1 - 1; i >= 0; i-- {
		for j := n2 - 1; j >= 0; j-- {
			if s1[i] == s2[j] {
				res[i][j] = res[i+1][j+1]
			} else {
				res[i][j] = min(res[i+1][j]+int(s1[i]), res[i][j+1]+int(s2[j]))
			}
		}
	}
	return res[0][0]
}

func min(i, i2 int) int {
	if i > i2 {
		return i2
	}
	return i
}
