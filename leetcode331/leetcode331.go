package leetcode331

import "strings"

func isValidSerialization(preorder string) bool {
	parts := strings.Split(preorder, ",")
	s := make([]string, 0, len(parts))
	for _, part := range parts {
		for part == "#" && len(s) > 0 && s[len(s)-1] == "#" {
			s = s[:len(s)-1]
			if len(s) == 0 {
				return false
			}
			s = s[:len(s)-1]
		}
		s = append(s, part)
	}

	return len(s) == 1 && s[0] == "#"
}
