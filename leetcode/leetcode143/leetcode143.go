package leetcode143

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
	if head == nil || head.Next == nil {
		return
	}

	slow, fast := head, head

	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	tmp := slow.Next
	slow.Next = nil
	var tmp_rev *ListNode

	for tmp != nil {
		node := tmp
		tmp = tmp.Next
		node.Next = tmp_rev
		tmp_rev = node
	}
	preHead := &ListNode{}
	start := head
	cur := preHead
	for start != nil || tmp_rev != nil {
		if start != nil {
			cur.Next = start
			start = start.Next
			cur = cur.Next
		}

		if tmp_rev != nil {
			cur.Next = tmp_rev
			tmp_rev = tmp_rev.Next
			cur = cur.Next
		}
	}
	head = preHead.Next
}
