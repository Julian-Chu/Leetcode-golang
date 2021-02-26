package leetcode58

func lengthOfLastWord(s string) int {
	l := 0
	found := false
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' {
			if found {
				break
			}
			continue
		}
		l++
		found = true
	}
	return l
}
