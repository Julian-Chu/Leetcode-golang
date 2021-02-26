package leetcode203

import "Leetcode-golang/utils"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode = utils.ListNode

func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}
	preHead := &ListNode{Next: head}
	cur := preHead

	for cur.Next != nil {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return preHead.Next
}
