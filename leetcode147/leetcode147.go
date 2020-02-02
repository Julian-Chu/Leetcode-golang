package leetcode147

import "Leetcode-golang/helper"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode = helper.ListNode

func insertionSortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	start, end, prev, cur := head, head, head, head.Next

	for cur != nil {
		if cur.Val <= start.Val {
			tmp := cur
			cur = cur.Next
			prev.Next = cur
			tmp.Next = start
			start = tmp
			continue
		} else if cur.Val >= end.Val {
			end = cur
			cur = cur.Next
			prev = prev.Next
			continue
		} else {
			tmp := cur
			cur = cur.Next
			prev.Next = cur
			insertAfter := start
			for insertAfter.Next.Val < tmp.Val {
				insertAfter = insertAfter.Next
			}
			insertBefore := insertAfter.Next
			insertAfter.Next = tmp
			tmp.Next = insertBefore
		}
	}

	return start
}
