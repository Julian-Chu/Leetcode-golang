package leetcode61

import (
	"github.com/Julian-Chu/Leetcode-golang/utils"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode = utils.ListNode

func rotateRight(head *ListNode, k int) *ListNode {

	if head == nil || k == 0 || head.Next == nil {
		return head
	}

	node := head
	l := 0
	for node != nil {
		l++
		node = node.Next
	}

	k = k % l
	if k == 0 {
		return head
	}
	tail, newTail := head, head
	for tail.Next != nil {
		if k <= 0 {
			newTail = newTail.Next
		}
		k--
		tail = tail.Next
	}
	newHead := newTail.Next
	newTail.Next, tail.Next = nil, head
	return newHead
}
