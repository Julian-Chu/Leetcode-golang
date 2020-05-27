package leetcode424

func characterReplacement(s string, k int) int {

	maxLen := 0
	for ch := byte('A'); ch <= byte('Z'); ch++ {
		l := 0
		kLeft := k
		for r := 0; r < len(s); r++ {
			if s[r] != ch {
				if kLeft > 0 {
					kLeft--
				} else {
					maxLen = Max(maxLen, r-l)
					for s[l] == ch {
						l++
					}
					l++
				}
			}
		}
		maxLen = Max(maxLen, len(s)-l)
	}
	return maxLen
}

func Max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
