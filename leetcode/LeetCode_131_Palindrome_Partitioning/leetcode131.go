package LeetCode_131_Palindrome_Partitioning

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

func partition_1(s string) [][]string {
	if s == "" {
		return nil
	}

	var res [][]string
	tmps := make([]string, 0, len(s))
	var dfs func(s string, start int, tmps []string)
	dfs = func(s string, start int, tmps []string) {
		if start == len(s) {
			res = append(res, append([]string{}, tmps...))
			return
		}

		for i := start; i < len(s); i++ {
			if isPalindrome_1(s, start, i) {
				tmps = append(tmps, string(s[start:i+1]))
				dfs(s, i+1, tmps)
				tmps = tmps[:len(tmps)-1]
			}
		}
	}
	dfs(s, 0, tmps)
	return res
}

func isPalindrome_1(s string, start, end int) bool {
	i, j := start, end
	for i <= j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}

	return true
}
