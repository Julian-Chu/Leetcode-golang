package leetcode141

import (
	"github.com/Julian-Chu/Leetcode-golang/utils"
)

type ListNode = utils.ListNode

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	slow, fast := head, head.Next
	for slow != nil && fast != nil && fast.Next != nil {
		if slow == fast {
			return true
		}

		slow = slow.Next
		fast = fast.Next.Next
	}

	return false
}
