package leetcode148

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

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	l, r := split(head)

	return merge(sortList(l), sortList(r))
}

func merge(node1 *ListNode, node2 *ListNode) *ListNode {
	prehead := &ListNode{}

	cur := prehead
	for node1 != nil && node2 != nil {
		if node1.Val > node2.Val {
			cur.Next = node2
			node2 = node2.Next
		} else {
			cur.Next = node1
			node1 = node1.Next
		}
		cur = cur.Next
	}

	for node1 != nil {
		cur.Next = node1
		node1 = node1.Next
		cur = cur.Next
	}

	for node2 != nil {
		cur.Next = node2
		node2 = node2.Next
		cur = cur.Next
	}

	return prehead.Next
}

func split(head *ListNode) (*ListNode, *ListNode) {
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	tmp := slow
	if slow.Next != nil {
		slow = slow.Next
		tmp.Next = nil
	} else {
		head.Next = nil
	}
	return head, slow
}
