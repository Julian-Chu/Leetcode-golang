package leetcode165

import (
	"strconv"
	"strings"
)

func compareVersion(version1 string, version2 string) int {
	v1 := strings.Split(version1, ".")
	v2 := strings.Split(version2, ".")

	v1_num := make([]int, len(v1))
	v2_num := make([]int, len(v2))

	for i := 0; i < len(v1); i++ {
		v, _ := strconv.Atoi(v1[i])
		v1_num[i] = v
	}

	for i := 0; i < len(v2); i++ {
		v, _ := strconv.Atoi(v2[i])
		v2_num[i] = v
	}

	for i := len(v1) - 1; i >= 0; i-- {
		if v1_num[i] == 0 {
			v1_num = v1_num[:i]
			continue
		}
		break
	}

	for i := len(v2) - 1; i >= 0; i-- {
		if v2_num[i] == 0 {
			v2_num = v2_num[:i]
			continue
		}
		break
	}

	l1 := len(v1_num)
	l2 := len(v2_num)
	n := l1
	if n > l2 {
		n = l2
	}

	for i := 0; i < n; i++ {
		switch {
		case v1_num[i] > v2_num[i]:
			return 1
		case v1_num[i] < v2_num[i]:
			return -1
		default:
			continue
		}
	}

	switch {
	case l1 > l2:
		return 1
	case l1 < l2:
		return -1
	default:
		return 0
	}

}
