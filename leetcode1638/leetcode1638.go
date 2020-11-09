package leetcode1638

func countSubstrings(s string, t string) int {
	ns := len(s)
	nt := len(t)
	res := 0
	for i := 0; i < ns; i++ {
		for j := 0; j < nt; j++ {
			diff := 0
			for k := 0; i+k < ns && j+k < nt; k++ {
				if s[i+k] != t[j+k] {
					diff++
					if diff > 1 {
						break
					}
				}
				res += diff
			}
		}
	}
	return res
}
