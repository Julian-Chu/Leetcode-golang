package leetcode19

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

func removeNthFromEnd(head *ListNode, n int) *ListNode {

	parent := head
	cursor := head

	for cursor != nil {
		if n < 0 {
			parent = parent.Next
		}
		n--
		cursor = cursor.Next
	}

	if n == 0 {
		return head.Next
	}

	parent.Next = parent.Next.Next

	return head
}
