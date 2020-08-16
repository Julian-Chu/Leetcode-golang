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
	preHead := &ListNode{}
	preHead.Next = head
	pre := preHead
	cur := head

	for cur != nil {
		if cur.Val != val {
			pre = cur
			cur = cur.Next
			continue
		}
		tmp := cur.Next
		pre.Next = tmp
		cur = tmp
	}
	return preHead.Next
}
