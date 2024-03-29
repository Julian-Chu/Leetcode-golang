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

func reverseWords_2(s string) string {
	words := strings.Fields(s)
	reverseStrsSlice(words)

	return strings.Join(words, " ")
}

func reverseStrsSlice(s []string) {
	if len(s) <= 1 {
		return
	}

	l, r := 0, len(s)-1
	for l < r {
		s[l], s[r] = s[r], s[l]
		l++
		r--
	}
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

func reverseWords_3(s string) string {
	s = strings.TrimSpace(s)
	words := strings.Split(s, " ")

	l, r := 0, len(words)-1
	for l <= r {
		words[l], words[r] = words[r], words[l]
		l++
		r--
	}

	slow := 0

	for i := range words {
		if words[i] != "" {
			words[slow] = words[i]
			slow++
		}
	}

	return strings.Join(words[:slow], " ")
}
