package leetcode28

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	idx := -1
	n := len(needle)
	for i := 0; i < len(haystack)-n+1; i++ {
		if !compareStr(needle, haystack, i) {
			continue
		}
		idx = i
		break
	}

	return idx
}

func compareStr(needle string, haystack string, i int) bool {
	for j := range needle {
		if needle[j] != haystack[i+j] {
			return false
		}
	}
	return true
}
