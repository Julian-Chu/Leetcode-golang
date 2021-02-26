package leetcode409

func longestPalindrome(s string) int {
	m := make(map[rune]int)

	for _, ch := range s {
		m[ch]++
	}

	pairs := 0
	single := 0

	for _, cnt := range m {
		if cnt%2 == 1 {
			single = 1
			cnt -= 1
		}
		pairs += cnt / 2
	}

	return pairs*2 + single
}
