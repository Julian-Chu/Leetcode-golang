package leetcode38

import "strconv"

func countAndSay(n int) string {
	if n == 0 {
		return ""
	}
	res := make([]string, n)

	for i := 0; i < n; i++ {
		if i == 0 {
			res[i] = "1"
			continue
		}
		cnt := 1
		v := res[i-1][0]
		for j := 1; j < len(res[i-1]); j++ {
			if res[i-1][j] == v {
				cnt++
				continue
			}
			res[i] += strconv.Itoa(cnt) + string(v)
			v = res[i-1][j]
			cnt = 1
		}

		res[i] += strconv.Itoa(cnt) + string(v)
	}

	return res[n-1]
}
