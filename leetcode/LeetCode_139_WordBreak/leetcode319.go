package LeetCode_139_WordBreak

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

func wordBreak_1(s string, wordDict []string) bool {
	wordDictSet := make(map[string]bool)
	for _, word := range wordDict {
		wordDictSet[word] = true
	}

	dp := make([]bool, len(s)+1)
	dp[0] = true

	for i := 1; i < len(s)+1; i++ {
		for j := 0; j < i; j++ {
			if dp[j] && wordDictSet[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(s)]
}

func wordBreak_memoization(s string, wordDict []string) bool {
	wordDictSet := make(map[string]bool)
	for _, word := range wordDict {
		wordDictSet[word] = true
	}

	memo := make(map[string]bool)

	var dfs func(string, map[string]bool) bool
	dfs = func(subStr string, memo map[string]bool) bool {
		if subStr == "" {
			return true
		}

		if v, exist := memo[subStr]; exist {
			return v
		}

		for i := 0; i < len(subStr); i++ {
			if wordDictSet[subStr[:i+1]] {
				if dfs(subStr[i+1:], memo) {
					return true
				}
			}
		}

		memo[subStr] = false
		return false
	}
	return dfs(s, memo)
}
