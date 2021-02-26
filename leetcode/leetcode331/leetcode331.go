package leetcode331

import "strings"

func isValidSerialization(preorder string) bool {

	parts := strings.Split(preorder, ",")
	return helper(parts, 0) == len(parts)-1
}

func helper(parts []string, i int) int {
	if i >= len(parts) {
		return -1
	}

	if parts[i] == "#" {
		return i
	}

	//l
	next := helper(parts, i+1)
	if next == -1 {
		return -1
	}

	// r
	next = helper(parts, next+1)
	if next == -1 {
		return -1
	}
	return next
}
