package leetcode206

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
	var pre *ListNode
	cur := head
	for cur != nil {
		pre, cur, cur.Next = cur, cur.Next, pre
	}

	return pre
}
