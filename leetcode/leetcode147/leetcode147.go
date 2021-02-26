package leetcode147

import (
	"Leetcode-golang/utils"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode = utils.ListNode

func insertionSortList(head *ListNode) *ListNode {
	beforeHead := &ListNode{Next: head}
	cur := head
	for cur != nil && cur.Next != nil {
		p := cur.Next
		if cur.Val <= p.Val {
			cur = p
			continue
		}
		cur.Next = p.Next
		prev, next := beforeHead, beforeHead.Next

		for next.Val < p.Val {
			prev = next
			next = next.Next
		}

		prev.Next = p
		p.Next = next
	}

	return beforeHead.Next
}
