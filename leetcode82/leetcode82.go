package leetcode82

import "Leetcode-golang/helper"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode = helper.ListNode

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	preHead := &ListNode{Next: head}

	cur := preHead
	for cur != nil {
		if cur.Next != nil && cur.Next.Next != nil && cur.Next.Val == cur.Next.Next.Val {
			start, end := cur.Next, cur.Next.Next
			for end.Next != nil && end.Next.Val == start.Val {
				end = end.Next
			}
			cur.Next = end.Next
		} else {
			cur = cur.Next
		}
	}

	return preHead.Next
}

func tes(a int,
	b int) {

}
