package leetcode139

import "sort"

func wordBreak(s string, wordDict []string) bool {
	sort.Slice(wordDict, func(i, j int) bool {
		return len(wordDict[i]) > len(wordDict[j])
	})
	strLen := len(s)
	var recur func(int) bool

	recur = func(start int) bool {
		for _, v := range wordDict {
			if start == strLen {
				return true
			}
			n := len(v)
			end := start + n
			if end > strLen {
				continue
			}
			if s[start:end] == v {
				if recur(end) {
					return true
				}
			}
		}
		return false
	}

	return recur(0)
}
