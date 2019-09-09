package leetcode8

import (
	"math"
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

	absV := 0
	for _, b := range abs {
		absV = 10*absV + int(b-'0')
		switch sign {
		case 1:
			if absV > math.MaxInt32 {
				return math.MaxInt32
			}
		case -1:
			if absV*-1 < math.MinInt32 {
				return math.MinInt32
			}
		}
	}

	return sign * absV

}
