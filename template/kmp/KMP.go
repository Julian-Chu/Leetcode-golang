package kmp

func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}

	next := getNext(needle)

	j := -1

	for i := 0; i < len(haystack); i++ {
		for j >= 0 && haystack[i] != needle[j+1] {
			j = next[j]
		}

		if haystack[i] == needle[j+1] {
			j++
		}

		if j == len(needle)-1 {
			return i - len(needle) + 1
		}
	}
	return -1
}

// prefix table
func getNext(s string) []int {
	next := make([]int, len(s))

	j := -1
	next[0] = j

	for i := 1; i < len(s); i++ {
		for j >= 0 && s[i] != s[j+1] {
			j = next[j]
		}

		if s[i] == s[j+1] {
			j++
		}

		next[i] = j
	}
	return next
}
