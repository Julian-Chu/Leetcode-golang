package leetcode142

import (
	"github.com/Julian-Chu/Leetcode-golang/utils"
)

type ListNode = utils.ListNode

func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	slow, fast := head.Next, head.Next.Next
	for fast != nil && fast.Next != nil && slow != fast {
		slow = slow.Next
		fast = fast.Next.Next
	}

	if slow != fast {
		return nil
	}

	for slow != head {
		slow, head = slow.Next, head.Next
	}

	return slow
}
