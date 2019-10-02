package leetcode131

func partition(s string) [][]string {
	res := make([][]string, 0)

	var dfs func(int, []string)
	dfs = func(idx int, temp []string) {
		if idx == len(s) {
			res = append(res, temp)
			return
		}
		temp = temp[:len(temp):len(temp)]
		for i := 1; i <= len(s)-idx; i++ {
			subStr := s[idx : idx+i]
			if isPalidrome(subStr) {
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
	l, r := n/2, n/2

	if len(s)%2 == 0 {
		l = r - 1
	}

	for l >= 0 {
		if s[l] != s[r] {
			return false
		}
		l--
		r++
	}
	return true
}
