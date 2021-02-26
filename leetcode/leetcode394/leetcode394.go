package leetcode394

import (
	"strconv"
)

func decodeString(s string) string {
	i := 0
	n := len(s)

	for i < n && (s[i] < '0' || s[i] > '9') {
		i++
	}
	if i == n {
		return s
	}

	j := i + 1
	for s[j] != '[' {
		j++
	}
	k := j
	cnt := 1
	for cnt > 0 {
		k++
		switch s[k] {
		case '[':
			cnt++
		case ']':
			cnt--
		}
	}

	repeat, _ := strconv.Atoi(s[i:j])
	substr := decodeString(s[j+1 : k])
	substrs := ""
	for i := 0; i < repeat; i++ {
		substrs += substr
	}
	return s[:i] + substrs + decodeString(s[k+1:])
}
