package leetcode8

import (
	"math"
	"strconv"
	"strings"
)

func myAtoi(str string) int {
	s := strings.TrimSpace(str)
	if s == "" {
		return 0
	}
	var sign int
	var abs string
	switch first := s[0]; {
	case first >= '0' && first <= '9':
		sign, abs = 1, s
	case first == '+':
		sign, abs = 1, s[1:]
	case first == '-':
		sign, abs = -1, s[1:]
	default:
		sign, abs = 0, ""
	}
	if sign == 0 || len(abs) == 0 {
		return 0
	}

	for i, b := range abs {
		if b < '0' || b > '9' {
			abs = abs[:i]
			break
		}
	}

	if len(abs) == 0 {
		return 0
	}
	if abs[0] == '0' {
		n := len(abs)
		lastZeroIndex := 0
		for k, v := range abs {
			if v != '0' {
				abs = abs[k:]
				break
			}
			lastZeroIndex = k
		}
		if lastZeroIndex == n-1 {
			return 0
		}
	}

	maxStr := strconv.Itoa(math.MaxInt32)
	minStr := strconv.Itoa(math.MinInt32)
	if sign == 1 {
		if len(abs) > len(maxStr) || (len(abs) == len(maxStr) && abs > maxStr) {
			return math.MaxInt32
		}
	} else if sign == -1 {
		if len(abs)+1 > len(minStr) || (len(abs)+1 == len(minStr) && abs > minStr[1:]) {
			return math.MinInt32
		}
	}

	v, _ := strconv.Atoi(abs)

	return sign * v

}
