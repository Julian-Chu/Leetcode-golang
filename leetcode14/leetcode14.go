package leetcode14

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	prefix := strs[0]

	for i := 1; i < len(strs); i++ {
		prefix = getCommonPrefix(prefix, strs[i])
	}

	return prefix
}

func getCommonPrefix(str1, str2 string) string {
	n1 := len(str1)
	n2 := len(str2)

	n := n1
	if n1 > n2 {
		n = n2
	}

	prefix := make([]byte, 0)
	for i := 0; i < n; i++ {
		if str1[i] != str2[i] {
			break
		}
		prefix = append(prefix, str1[i])
	}
	return string(prefix)
}
