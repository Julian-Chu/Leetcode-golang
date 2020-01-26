package leetcode86

import "Leetcode-golang/helper"

type ListNode = helper.ListNode

func partition(head *ListNode, x int) *ListNode {
	q := make([]*ListNode, 0)

	cur := head
	var prev, insertBefore *ListNode

	for cur != nil {
		if insertBefore == nil && cur.Val >= x {
			insertBefore = cur
		}
		if cur.Val < x {
			q = append(q, cur)
			if prev != nil {
				prev.Next = cur.Next
			}

		} else {
			prev = cur
		}
		cur = cur.Next
	}

	if len(q) == 0 {
		return head
	}
	newHead := q[0]

	for i := 1; i < len(q); i++ {
		q[i-1].Next = q[i]
	}
	q[len(q)-1].Next = insertBefore

	return newHead
}
