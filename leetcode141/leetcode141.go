package leetcode141

import "Leetcode-golang/helper"

type ListNode = helper.ListNode

func hasCycle(head *ListNode) bool {
	m := make(map[*ListNode]bool)

	for head != nil {
		if _, ok := m[head]; ok {
			return true
		}
		m[head] = true
		head = head.Next
	}

	return false
}
