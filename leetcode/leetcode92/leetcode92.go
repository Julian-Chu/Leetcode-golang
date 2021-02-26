package leetcode92

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

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	m_step := m
	n_step := n

	if head == nil || m_step == n_step {
		return head
	}

	var revStart *ListNode
	preHead, beforeRev, afterRev := &ListNode{Next: head}, &ListNode{Next: head}, &ListNode{Next: head}
	for m_step > 1 {
		beforeRev = beforeRev.Next
		m_step--
	}
	revStart = beforeRev.Next

	for n_step > 0 {
		afterRev = afterRev.Next
		n_step--
	}
	revEnd := afterRev
	afterRev = revEnd.Next
	revEnd.Next = nil

	cur := revStart
	for cur != nil {
		tmp := cur
		cur = cur.Next
		tmp.Next = afterRev
		afterRev = tmp
	}

	if m == 1 {
		preHead.Next = afterRev
	} else {

		beforeRev.Next = afterRev
	}
	return preHead.Next
}
