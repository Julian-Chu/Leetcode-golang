package leetcode131

func partition(s string) [][]string {
	res := make([][]string, 0)

	var dfs func(int, []string)
	dfs = func(idx int, temp []string) {
		if idx == len(s) {
			cur := make([]string, len(temp))
			copy(cur, temp)
			res = append(res, cur)
			//res = append(res, temp)
			return
		}
		//temp = temp[:len(temp):len(temp)]
		for i := 1; i <= len(s)-idx; i++ {
			subStr := s[idx : idx+i]
			if isPalidrome(subStr) {
				//temp = temp[:len(temp):len(temp)]
				dfs(idx+i, append(temp, subStr))
			}
		}
	}
	dfs(0, []string{})
	return res
}

func isPalidrome(s string) bool {
	n := len(s)
	if n <= 1 {
		return true
	}
	l, r := 0, n-1

	for l < r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}
