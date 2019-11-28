package leetcode71

import "strings"

func simplifyPath(path string) string {
	strs := strings.Split(path, "/")

	res := ""
	moveUpCnt := 0
	for i := len(strs) - 1; i >= 0; i-- {
		if strs[i] == "" || strs[i] == "." {
			continue
		}
		if strs[i] == ".." {
			moveUpCnt++
			continue
		}

		if moveUpCnt != 0 {
			moveUpCnt--
			continue
		}
		if res == "" {
			res = strs[i]
			continue
		}

		res = strs[i] + "/" + res
	}

	return "/" + res
}
