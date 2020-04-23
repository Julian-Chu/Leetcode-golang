package leetcode402

import (
	"strconv"
	"strings"
)

func removeKdigits(num string, k int) string {
	if len(num) == 0 {
		return "0"
	}
	if k == 0 {
		return num
	}

	v := num
	for j := 0; j < k; j++ {
		min := v[1:]
		for i := 1; i < len(v); i++ {
			bs := []byte(v)
			//bs[0], bs[i] =   bs[i], bs[0]
			copy(bs[1:], bs[0:i])

			minNum, _ := strconv.Atoi(min)
			target, _ := strconv.Atoi(string(bs[1:]))

			if target < minNum {
				min = string(bs[1:])
			}
		}
		v = min
	}
	val, _ := strconv.Atoi(v)
	if v == "" || val == 0 {
		return "0"
	}
	return strings.TrimLeft(v, "0")
}
