package leetcode22

func generateParenthesis(n int) []string {
	res := make([]string, 0)
	bytes := make([]byte, n*2)

	var dfs func(*[]string, []byte, int, int, int)
	dfs = func(res *[]string, bytes []byte, index, l, r int) {
		if l == 0 && r == 0 {
			*res = append(*res, string(bytes))
			return
		}

		if l > 0 {
			bytes[index] = '('
			dfs(res, bytes, index+1, l-1, r)
		}

		if r > 0 && l < r {
			bytes[index] = ')'
			dfs(res, bytes, index+1, l, r-1)
		}
	}

	dfs(&res, bytes, 0, n, n)
	res[0], res[1] = res[1], res[0]
	return res
}
