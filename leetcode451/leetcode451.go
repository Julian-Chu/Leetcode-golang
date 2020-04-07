package leetcode451

import "sort"

func frequencySort(s string) string {
	r := ['z' + 1]int{}
	for i := range s {
		r[s[i]]++
	}

	ss := make([]string, 0, len(s))
	for i := range r {
		if r[i] == 0 {
			continue
		}
		ss = append(ss, makeString(byte(i), r[i]))
	}

	sort.Sort(segments(ss))

	res := ""
	for _, s := range ss {
		res += s
	}

	return res
}

func makeString(b byte, n int) string {
	bytes := make([]byte, n)
	for i := range bytes {
		bytes[i] = b
	}
	return string(bytes)
}

type segments []string

func (s segments) Len() int {
	return len(s)
}

func (s segments) Less(i, j int) bool {
	if len(s[i]) == len(s[j]) {
		return s[i] < s[j]
	}
	return len(s[i]) > len(s[j])
}

func (s segments) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
