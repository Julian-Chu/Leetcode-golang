package leetcode93

func restoreIpAddresses(s string) []string {
	if len(s) == 0 {
		return []string{}
	}
	// submask which starts with zero has only digit
	// submask max 3 digits, less than or equal to 255
	res := make([]string, 0)

	var dfs func([]byte, int, int)
	dfs = func(temp []byte, idx, pointCount int) {
		if len(temp) == len(s)+3 {
			res = append(res, string(temp))
			return
		}
		temp = temp[:len(temp):len(temp)]
		for i := 1; i <= 3; i++ {
			pointCount++
			newTemp := s[idx : idx+i]
			newTemp
			if pointCount == 3 {
				dfs()
			} else {
				dfs(append(temp, newTemp...), idx+i, pointCount)
			}
		}
	}

	dfs([]byte{}, 0, 0)
	return []string{}
}
