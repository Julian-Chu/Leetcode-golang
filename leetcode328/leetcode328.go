package leetcode328

import "Leetcode-golang/utils"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode = utils.ListNode

func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	oddTail := head
	evenHead := head.Next
	evenTail := evenHead

	for evenTail != nil && evenTail.Next != nil {
		oddTail.Next = evenTail.Next
		oddTail = oddTail.Next
		evenTail.Next = oddTail.Next
		evenTail = evenTail.Next
		oddTail.Next = evenHead
	}

	return head
}
