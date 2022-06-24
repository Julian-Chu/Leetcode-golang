package LeetCode_143_ReorderList

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

func reorderList(head *ListNode) {
	if head == nil {
		return
	}

	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	head2 := reverse(slow.Next)
	slow.Next = nil

	idx := 0
	dummy := &ListNode{}
	cur := dummy
	for head != nil && head2 != nil {
		if idx&1 == 0 {
			cur.Next = head
			head = head.Next
		} else {
			cur.Next = head2
			head2 = head2.Next
		}
		cur = cur.Next
		idx++
	}

	if head != nil {
		cur.Next = head
	}

	if head2 != nil {
		cur.Next = head2
	}
	head = dummy.Next

}

func reverse(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	dummy := &ListNode{}
	for head != nil {
		next := dummy.Next
		dummy.Next = head
		head = head.Next
		dummy.Next.Next = next
	}
	return dummy.Next
}

func reorderList_1(head *ListNode) {
	if head == nil {
		return
	}
	nodes := make([]*ListNode, 0)

	for head != nil {
		next := head.Next
		head.Next = nil
		nodes = append(nodes, head)
		head = next
	}

	dummyHead := &ListNode{}

	l, r := 0, len(nodes)-1

	cur := dummyHead
	idx := 0
	for l <= r {
		if idx&1 == 0 {
			cur.Next = nodes[l]
			l++
		} else {
			cur.Next = nodes[r]
			r--
		}
		idx++
		cur = cur.Next
	}

	head = dummyHead.Next
}
