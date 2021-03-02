package leetcode301

func removeInvalidParentheses(s string) []string {
	left, right := leftRightCount(s)
	var result []string

	dfs(s, 0, left, right, &result)
	return result
}

func dfs(s string, start, left, right int, result *[]string) {
	if left == 0 && right == 0 {
		if isValid(s) {
			*result = append(*result, s)
		}
		return
	}

	for i := start; i < len(s); i++ {
		if i > start && s[i] == s[i-1] {
			continue
		}

		if s[i] == '(' && left > 0 {
			dfs(s[:i]+s[i+1:], i, left-1, right, result)
		}

		if s[i] == ')' && right > 0 {
			dfs(s[:i]+s[i+1:], i, left, right-1, result)
		}
	}
}

func isValid(s string) bool {
	left, right := leftRightCount(s)
	return left == 0 && right == 0
}

func leftRightCount(s string) (int, int) {
	l, r := 0, 0
	for _, ch := range s {
		if ch == '(' {
			l++
		}

		if ch == ')' {
			if l > 0 {
				l--
			} else {
				r++
			}
		}
	}
	return l, r
}
