package leetcode1638

func countSubstrings_1(s string, t string) int {
	ns := len(s)
	nt := len(t)
	res := 0
	for i := 0; i < ns; i++ {
		for j := 0; j < nt; j++ {
			if s[i] != t[j] {
				l, r := 1, 1

				for i-l >= 0 && j-l >= 0 && s[i-l] == t[j-l] {
					l++
				}

				for i+r < ns && j+r < ns && s[i+r] == t[j+r] {
					r++
				}
				res += l * r
			}
		}
	}
	return res
}
