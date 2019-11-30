package leetcode151

import "strings"

func reverseWords(s string) string {
	strs := strings.Split(s, " ")
	res := make([]string, 0)
	for i := len(strs) - 1; i >= 0; i-- {
		if strs[i] == "" {
			continue
		}
		res = append(res, strs[i])
	}

	return strings.Join(res, " ")
}
