package leetcode299

import "fmt"

func getHint(secret string, guess string) string {
	bull, cow := 0, 0

	count := make(map[rune]int)

	for _, ch := range secret {
		count[ch]++
	}

	for k, v := range guess {
		if count[v] > 0 {
			cow++
			count[v]--
		}

		if secret[k] == byte(v) {
			bull++
			cow--
		}
	}
	return fmt.Sprintf("%dA%dB", bull, cow)
}
