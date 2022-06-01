package LeetCode_19_RemoveNthNodeFromEndOfList

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

func removeNthFromEnd1(head *ListNode, n int) *ListNode {
	if head == nil {
		return head
	}
	dummy := &ListNode{Next: head}
	slow, fast := dummy, dummy
	step := 0
	for fast.Next != nil {
		fast = fast.Next
		step++
		if step <= n {
			continue
		}
		slow = slow.Next
	}
	if slow.Next != nil {
		slow.Next = slow.Next.Next
	}
	return dummy.Next
}

func removeNthFromEnd2(head *ListNode, n int) *ListNode {
	dummyHead := &ListNode{Next: head}

	prev, cur := dummyHead, head
	i := 1
	for cur != nil {
		cur = cur.Next
		if i > n {
			prev = prev.Next
		}
		i++
	}
	prev.Next = prev.Next.Next
	return dummyHead.Next
}
