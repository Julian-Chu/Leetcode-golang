package leetcode93

import "strconv"

func restoreIpAddresses(s string) []string {
	if len(s) < 4 || len(s) > 12 {
		return []string{}
	}
	res := make([]string, 0)

	var dfs func(string, int, int)
	dfs = func(temp string, stringIdx, substrIdx int) {
		if stringIdx == len(s) {
			res = append(res, temp)
			return
		}

		for i := 0; i < 3; i++ {
			restLen := len(s[stringIdx:])
			if restLen < i+1 {
				break
			}
			end := stringIdx + i + 1
			subIpStr := s[stringIdx:end]
			if len(subIpStr) == 3 {
				subIp, _ := strconv.Atoi(subIpStr)
				if subIp > 255 {
					break
				}
			}
			newStr := temp + subIpStr + "."
			if substrIdx == 3 {
				if restLen > 3 || restLen == 0 {
					break
				}
				newStr = temp + s[stringIdx:]
			} else {
				if len(s[end:]) == 0 {
					break
				}
			}
			dfs(newStr, end, substrIdx+1)
			if s[stringIdx] == '0' {
				break
			}
		}

	}

	dfs("", 0, 0)
	return res
}
