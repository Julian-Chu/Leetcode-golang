package lintcode426

import (
	"strconv"
	"strings"
)

func restoreIpAddresses(s string) []string {
	result := []string{}
	res := make([]string, 0, 4)
	dfs(s, 0, &res, &result)
	return result
}

func dfs(s string, start int, res, result *[]string) {
	if start >= len(s) {
		if len(*res) == 4 {
			*result = append(*result, strings.Join(*res, "."))
		}
		return
	}

	if len(*res) == 4 {
		return
	}

	for i := 1; i < 4; i++ {
		if start+i > len(s) || (i > 1 && s[start] == '0') {
			break
		}

		ip := s[start : start+i]
		if val, _ := strconv.Atoi(ip); val > 255 {
			break
		}

		*res = append(*res, ip)
		dfs(s, start+i, res, result)
		*res = (*res)[:len(*res)-1]
	}

}
