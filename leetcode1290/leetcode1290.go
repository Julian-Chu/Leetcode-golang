package leetcode1290

import "Leetcode-golang/utils"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode = utils.ListNode

func getDecimalValue(head *ListNode) int {
	res := 0

	for head != nil {
		res <<= 1
		res += head.Val
		head = head.Next
	}

	return res
}
