package LeetCode_151_ReverseWordsInAString

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

func reverseWords1(s string) string {
	if len(s) <= 1 {
		return s
	}

	b := []byte(s)

	slow, fast := 0, 0
	// head
	for fast < len(b) && b[fast] == ' ' {
		fast++
	}

	// between
	for ; fast < len(b); fast++ {
		if fast > 1 && b[fast] == ' ' && b[fast-1] == ' ' {
			continue
		}
		b[slow] = b[fast]
		slow++
	}

	// tail
	if slow > 1 && b[slow-1] == ' ' {
		b = b[:slow-1]
	} else {
		b = b[:slow]
	}

	reverse(b)
	for i := 0; i < len(b); i++ {
		if b[i] == ' ' {
			continue
		}
		start := i
		for i < len(b) && b[i] != ' ' {
			i++
		}
		reverse(b[start:i])
	}
	return string(b)
}

func reverse(b []byte) {
	if len(b) <= 1 {
		return
	}

	l, r := 0, len(b)-1
	for l < r {
		b[l], b[r] = b[r], b[l]
		l++
		r--
	}
}
