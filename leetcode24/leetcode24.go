package leetcode24

import "Leetcode-golang/helper"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode = helper.ListNode

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	cur, next, newHead := head, head.Next, head.Next
	var prev *ListNode

	for cur != nil && next != nil {
		cur.Next = next.Next
		next.Next = cur
		cur, next = next, cur
		if prev != nil {
			prev.Next = cur
		}
		if next.Next == nil || next.Next.Next == nil {
			break
		}
		prev = next
		cur = cur.Next.Next
		next = next.Next.Next
	}

	return newHead
}
