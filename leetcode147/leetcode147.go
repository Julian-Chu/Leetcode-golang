package leetcode147

import (
	"Leetcode-golang/helper"
	"math"
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
	p := &ListNode{Next: head, Val: math.MinInt64}
	for cur != nil && cur.Next != nil {
		if cur.Val <= cur.Next.Val {
			cur = cur.Next
			continue
		}

		// remove cur.Next from Linked list
		tmp := cur.Next
		cur.Next = cur.Next.Next

		insertAfter := p
		for insertAfter != nil && insertAfter.Next != nil {
			if insertAfter.Val <= tmp.Val && tmp.Val < insertAfter.Next.Val {
				tmp.Next = insertAfter.Next
				insertAfter.Next = tmp
				break
			}

			insertAfter = insertAfter.Next
		}

	}
	return p.Next
}
