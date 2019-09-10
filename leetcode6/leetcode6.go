package leetcode6

import "strings"

func convert(s string, numRows int) string {
	if numRows == 1 || len(s) < numRows {
		return s
	}
	var sb strings.Builder
	offset := numRows*2 - 2

	for i := 0; i < len(s); i += offset {
		sb.WriteByte(s[i])
	}
	for r := 1; r <= numRows-2; r++ {
		sb.WriteByte(s[r])
		for k := offset; k-r < len(s); k += offset {
			sb.WriteByte(s[k-r])
			if k+r < len(s) {
				sb.WriteByte(s[k+r])
			}
		}
	}

	for i := numRows - 1; i < len(s); i += offset {
		sb.WriteByte(s[i])
	}
	return sb.String()
}

//func convert(s string, numRows int) string {
//	var sb strings.Builder
//	if numRows == 1 {
//		return s
//	}
//	for i := 0; i < numRows; i++ {
//		base := numRows*2 - 2
//		for k, v := range s {
//			if k%base == i || k%base == base-i {
//				sb.WriteRune(v)
//			}
//		}
//	}
//
//	return sb.String()
//}
