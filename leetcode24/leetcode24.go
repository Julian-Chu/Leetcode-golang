package leetcode24

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

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	preHead := &ListNode{}
	pre := preHead
	cur := head
	next := cur.Next
	pre.Next = cur

	for {
		tmp := next.Next
		cur, next = next, cur
		pre.Next = cur
		cur.Next = next
		next.Next = tmp

		pre = next
		cur = tmp
		if cur == nil {
			break
		}
		next = cur.Next
	}

	return preHead.Next
}
