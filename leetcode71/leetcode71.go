package leetcode71

import "strings"

func simplifyPath(path string) string {
	splitStrArr := strings.Split(path, "/")

	res := ""
	moveUpCnt := 0
	for i := len(splitStrArr) - 1; i >= 0; i-- {
		if splitStrArr[i] == "" || splitStrArr[i] == "." {
			continue
		}
		if splitStrArr[i] == ".." {
			moveUpCnt++
			continue
		}

		if moveUpCnt != 0 {
			moveUpCnt--
			continue
		}
		if res == "" {
			res = splitStrArr[i]
			continue
		}

		res = splitStrArr[i] + "/" + res
	}

	return "/" + res
}
