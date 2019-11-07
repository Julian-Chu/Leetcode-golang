package leetcode91

import "strconv"

func numDecodings(s string) int {
	cnt := 0
	rec([]byte(s), &cnt)
	return cnt
}

func rec(restStr []byte, cnt *int) {
	if len(restStr) == 0 {
		*cnt++
		return
	}
	// take 1
	v1, _ := strconv.ParseInt(string(restStr[0]), 10, 32)
	if v1 != 0 {
		rec(append([]byte{}, restStr[1:]...), cnt)
	}
	// take 2
	if len(restStr) < 2 {
		return
	}

	v2, _ := strconv.ParseInt(string(restStr[:2]), 10, 32)
	if v2 > 26 || v2 == 0 || v1 == 0 {
		return
	}
	rec(append([]byte{}, restStr[2:]...), cnt)
}
