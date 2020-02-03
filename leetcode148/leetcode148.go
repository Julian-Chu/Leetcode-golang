package leetcode148

import (
	"Leetcode-golang/helper"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode = helper.ListNode

func sortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	cur := head
	p := &ListNode{Next: head}
	for cur != nil && cur.Next != nil {
		if cur.Val <= cur.Next.Val {
			cur = cur.Next
			continue
		}

		tmp := cur.Next
		cur.Next = cur.Next.Next

		insertAfter := p
		for tmp.Val > insertAfter.Next.Val {
			insertAfter = insertAfter.Next
		}

		tmp.Next = insertAfter.Next
		insertAfter.Next = tmp
	}
	return p.Next
}
