package leetcode1638

func countSubstrings_3(s string, t string) int {
	ns := len(s)
	nt := len(t)
	res := 0

	var helper func(int, int) int
	helper = func(i int, j int) int {
		pre, cur, res := 0, 0, 0
		for i < ns && j < nt {
			cur++
			if s[i] != t[j] {
				pre = cur
				cur = 0
			}
			res += pre
			i++
			j++
		}
		return res
	}

	for i := 0; i < ns; i++ {
		res += helper(i, 0)
	}

	for j := 1; j < nt; j++ {
		res += helper(0, j)
	}
	return res
}
