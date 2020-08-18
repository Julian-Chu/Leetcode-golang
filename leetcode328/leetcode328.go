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
	tail := head
	length := 1

	for tail.Next != nil {
		tail = tail.Next
		length++
	}

	if length == 1 {
		return head
	}

	cur := head

	for length > 1 {
		if cur.Next == nil {
			break
		}

		tail.Next = cur.Next
		cur.Next = cur.Next.Next
		cur = cur.Next
		tail = tail.Next
		tail.Next = nil
		length -= 2
	}

	return head

}
