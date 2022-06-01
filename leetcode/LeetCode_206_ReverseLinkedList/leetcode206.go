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

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	cur := head
	// alternative:
	//prev, cur:= (*ListNode)(nil), head
	for cur != nil {
		prev, cur, cur.Next = cur, cur.Next, prev
	}

	return prev
}
