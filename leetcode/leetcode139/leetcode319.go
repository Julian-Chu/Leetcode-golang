package leetcode139

import "sort"

func wordBreak(s string, wordDict []string) bool {
	if len(wordDict) == 0 {
		return false
	}

	dict := make(map[string]bool, len(wordDict))
	length := make(map[int]bool, len(wordDict))
	for _, w := range wordDict {
		dict[w] = true
		length[len(w)] = true
	}

	sizes := make([]int, 0, len(length))
	for l := range length {
		sizes = append(sizes, l)
	}

	sort.Ints(sizes)
	dp := make([]bool, len(s)+1)
	dp[0] = true
	n := len(s)
	for i := 0; i <= n; i++ {
		if !dp[i] {
			continue
		}

		for _, size := range sizes {
			if i+size <= n {
				dp[i+size] = dp[i+size] || dict[s[i:i+size]]
			}
		}
	}
	return dp[n]
}
