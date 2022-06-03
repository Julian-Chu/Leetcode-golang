package LeetCode_28_ImplementstrStr

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

func strStr_KMP(haystack string, needle string) int {
	n := len(needle)
	if n == 0 {
		return 0
	}

	next := make([]int, len(needle))
	getNext(next, needle)

	// prefix table start from 0
	j := 0
	for i := 0; i < len(haystack); i++ {
		// j start from 0, so condition should be j > 0
		for j > 0 && haystack[i] != needle[j] {
			// because prefix table start from, so previous of j should be next[j-1]
			j = next[j-1]
		}
		if haystack[i] == needle[j] {
			j++
		}

		if j == len(needle) {
			return i - n + 1
		}
	}

	return -1
}

func getNext(next []int, s string) {
	// prefix table start from 0
	j := 0
	next[0] = j

	for i := 1; i < len(s); i++ {
		if j > 0 && s[i] != s[j] {
			// because prefix table start from, so previous of j should be next[j-1]
			j = next[j-1]
		}

		if s[i] == s[j] {
			j++
		}
		next[i] = j
	}
}
