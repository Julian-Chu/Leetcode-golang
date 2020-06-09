package leetcode331

import "strings"

func isValidSerialization(preorder string) bool {
	// non null node: 1 indegree, 2 outdegree
	// null node: 1 indegree

	parts := strings.Split(preorder, ",")
	diff := 1

	for _, part := range parts {
		diff--
		if diff < 0 {
			return false // ex. ["#"]
		}

		if part != "#" {
			diff += 2
		}
	}

	return diff == 0
}
