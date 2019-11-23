package leetcode28

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	n := len(needle)
	for i := 0; i < len(haystack)-n+1; i++ {
		if haystack[i:i+n] == needle {
			return i
		}
	}

	return -1
}
