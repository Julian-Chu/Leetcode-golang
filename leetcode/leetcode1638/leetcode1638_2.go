package leetcode1638

func countSubstrings_2(s string, t string) int {
	dpl, dpr := [101][101]int{}, [101][101]int{}
	ns := len(s)
	nt := len(t)

	for i := 1; i <= ns; i++ {
		for j := 1; j <= nt; j++ {
			if s[i-1] == t[j-1] {
				dpl[i][j] = 1 + dpl[i-1][j-1]
			}
		}
	}

	for i := ns; i > 0; i-- {
		for j := nt; j > 0; j-- {
			if s[i-1] == t[j-1] {
				dpr[i-1][j-1] = 1 + dpr[i][j]
			}
		}
	}
	res := 0
	for i := 0; i < ns; i++ {
		for j := 0; j < nt; j++ {
			if s[i] != t[j] {
				res += (dpl[i][j] + 1) * (dpr[i+1][j+1] + 1)
			}
		}
	}
	return res
}
