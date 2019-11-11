package leetcode139

import "sort"

func wordBreak(s string, wordDict []string) bool {
	sort.Slice(wordDict, func(i, j int) bool {
		return len(wordDict[i]) > len(wordDict[j])
	})
	if len(s) == 0 {
		return true
	}
	for _, v := range wordDict {
		n := len(v)
		if len(s) < n {
			continue
		}
		if s[0:n] == v {
			str := s[n:]
			if wordBreak(str, wordDict) {
				return true
			}
		}
	}
	return false
}
