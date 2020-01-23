package leetcode19

import "Leetcode-golang/helper"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode = helper.ListNode

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return head
	}

	nodes := make([]*ListNode, 0, n*2)
	node := head
	for node != nil {
		nodes = append(nodes, node)
		node = node.Next
	}

	if len(nodes) == 1 {
		return nil
	}
	i := len(nodes) - n
	switch {
	case i == 0:
		return nodes[1]
	case i == len(nodes)-1:
		nodes[len(nodes)-2].Next = nil
	default:
		nodes[i-1].Next = nodes[i+1]
	}

	return head
}
