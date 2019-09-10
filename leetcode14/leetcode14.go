package leetcode14

func longestCommonPrefix(strs []string) string {
	short := shortrestStr(strs)
	for i := 0; i < len(short); i++ {
		for j := range strs {
			if strs[j][i] != short[i] {
				return short[:i]
			}
		}
	}
	return short
}

func shortrestStr(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	res := strs[0]
	for _, str := range strs {
		if len(str) < len(res) {
			res = str
		}
	}
	return res
}

//func longestCommonPrefix(strs []string) string {
//	if len(strs) == 0 {
//		return ""
//	}
//	prefixLen := len(strs[0])
//
//	for i := 1; i < len(strs); i++ {
//		if len(strs[i]) < prefixLen {
//			prefixLen = len(strs[i])
//		}
//		for j := 0; j < prefixLen; j++ {
//			if strs[i][j] != strs[i-1][j] {
//				prefixLen = j
//				break
//			}
//		}
//	}
//
//	return strs[0][:prefixLen]
//}

//func longestCommonPrefix(strs []string) string {
//	if len(strs) == 0 {
//		return ""
//	}
//	prefix := strs[0]
//
//	for i := 1; i < len(strs); i++ {
//		prefix = getCommonPrefix(prefix, strs[i])
//	}
//
//	return prefix
//}
//
//func getCommonPrefix(str1, str2 string) string {
//	n1 := len(str1)
//	n2 := len(str2)
//
//	n := n1
//	if n1 > n2 {
//		n = n2
//	}
//
//	prefix := make([]byte, 0)
//	for i := 0; i < n; i++ {
//		if str1[i] != str2[i] {
//			break
//		}
//		prefix = append(prefix, str1[i])
//	}
//	return string(prefix)
//}
