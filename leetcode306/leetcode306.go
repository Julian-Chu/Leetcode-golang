package leetcode306

import (
	"strconv"
	"strings"
)

func isAdditiveNumber(num string) bool {
	for i := 1; i <= len(num)/2; i++ {
		if byte(num[0]) == '0' && i > 1 {
			return false
		}

		a, _ := strconv.ParseInt(string(num[:i]), 10, 64)
		for j := 1; max(i, j) <= len(num)-i-j; j++ {
			if byte(num[i]) == '0' && j > 1 {
				break
			}
			b, _ := strconv.ParseInt(string(num[i:i+j]), 10, 64)
			if isAdtv(a, b, i+j, num) {
				return true
			}
		}
	}
	return false
}

func isAdtv(a int64, b int64, start int, num string) bool {
	if start == len(num) {
		return true
	}
	sum := strconv.FormatInt(a+b, 10)
	return strings.Contains(num[start:], sum) && isAdtv(b, a+b, start+len(sum), num)
}

func max(i int, j int) int {
	if i > j {
		return i
	}
	return j
}
