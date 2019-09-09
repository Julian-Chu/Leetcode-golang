package leetcode8

import (
	"math"
	"strconv"
	"strings"
)

func myAtoi(str string) int {
	trimStr := strings.TrimSpace(str)
	if len(trimStr) == 0 {
		return 0
	}

	i := strings.IndexFunc(trimStr, func(r rune) bool {
		return r > 57 || r < 48 && r != 45 && r != 43
	})
	if i == 0 {
		return 0
	}
	if i != -1 {
		trimStr = trimStr[:i]
	}
	for {
		j := strings.LastIndexFunc(trimStr, func(r rune) bool {
			return r > 57 || r < 48
		})
		if i == j || j <= 0 {
			break
		}
		if j > 0 {
			trimStr = trimStr[:j]
		}
	}
	if trimStr[0] == 48 {
		i := strings.IndexFunc(trimStr, func(r rune) bool {
			return r != 48
		})
		if i == -1 {
			return 0
		}
		if trimStr[i] < 48 || trimStr[i] > 57 {
			return 0
		}
		trimStr = trimStr[i:]
	} else if len(trimStr) > 1 && trimStr[0] == 45 && trimStr[1] == 48 {
		i := strings.IndexFunc(trimStr, func(r rune) bool {
			return r != 48 && r != 45
		})
		if i == -1 {
			return 0
		}
		trimStr = trimStr[0:1] + trimStr[i:]
	}
	max := strconv.Itoa(math.MaxInt32)
	min := strconv.Itoa(math.MinInt32)
	if trimStr[0] == 45 {
		if len(trimStr) >= len(min) {
			if len(trimStr) > len(min) || trimStr >= min {
				return math.MinInt32
			}
		}
	} else {
		if trimStr[0] == 43 {
			trimStr = trimStr[1:]
		}
		if len(trimStr) >= len(max) {
			if len(trimStr) > len(max) || trimStr >= max {
				return math.MaxInt32
			}
		}
	}
	v, _ := strconv.Atoi(trimStr)
	return v
}
