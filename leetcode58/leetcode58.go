package leetcode58

import "strings"

func lengthOfLastWord(s string) int {
	if s == "" {
		return 0
	}

	arr := strings.Split(s, " ")
	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] == "" {
			continue
		}
		return len(arr[i])
	}

	return 0
}
