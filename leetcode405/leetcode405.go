package leetcode405

import "strings"

func toHex(num int) string {
	if num == 0 {
		return "0"
	}

	idx := []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f'}
	res := make([]byte, 8)
	for i := 7; i >= 0; i-- {
		res[i] = idx[num&15]
		num = num >> 4
	}

	return strings.TrimLeft(string(res), "0")
}
