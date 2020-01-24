package leetcode61

import "Leetcode-golang/helper"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode = helper.ListNode

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil || k == 0 {
		return head
	}

	node := head
	var tail, newTail *ListNode
	l := 0
	for node != nil {
		l++
		if node.Next == nil {
			tail = node
		}
		node = node.Next

	}
	offset := (l - k%l) % l
	newHead := head
	if offset == 0 {
		return head
	}
	for offset > 0 {
		newTail = newHead
		newHead = newHead.Next
		offset--
	}

	newTail.Next = nil
	tail.Next = head
	return newHead

}
