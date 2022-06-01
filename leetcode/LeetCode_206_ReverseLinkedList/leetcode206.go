package LeetCode_206_ReverseLinkedList

import "github.com/Julian-Chu/Leetcode-golang/utils"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode = utils.ListNode

//func reverseList(head *ListNode) *ListNode {
//	var prev *ListNode
//	cur := head
//	// alternative:
//	//prev, cur:= (*ListNode)(nil), head
//	for cur != nil {
//		prev, cur, cur.Next = cur, cur.Next, prev
//	}
//
//	return prev
//}

func reverseList(head *ListNode) *ListNode {
	return help(nil, head)
}

func help(prev, cur *ListNode) *ListNode {
	if cur == nil {
		return prev
	}

	next := cur.Next
	cur.Next = prev

	return help(cur, next)
}
