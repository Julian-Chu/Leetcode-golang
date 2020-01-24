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

	nodes := make([]*ListNode, 0)
	for head != nil {
		nodes = append(nodes, head)
		head = head.Next
	}

	l := len(nodes)
	offset := l - k%l

	if offset == l {
		return nodes[0]
	}
	if offset > 0 {
		nodes[offset-1].Next = nil
		nodes[l-1].Next = nodes[0]
	}

	return nodes[offset]

}
