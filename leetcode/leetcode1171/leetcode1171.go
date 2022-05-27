package leetcode1171

import "github.com/Julian-Chu/Leetcode-golang/utils"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode = utils.ListNode

func removeZeroSumSublists(head *ListNode) *ListNode {
	pre := &ListNode{
		Val:  1000 * 1000 * 2,
		Next: head,
	}

	m := make(map[int]*ListNode)

	stack, top := make([]int, 1000), -1

	sum := 0
	head = pre
	for head != nil {
		sum += head.Val
		if m[sum] == nil {
			m[sum] = head
			top++
			stack[top] = sum
		} else {
			m[sum].Next = head.Next
			for stack[top] != sum {
				m[stack[top]] = nil
				top--
			}
		}
		head = head.Next
	}
	return pre.Next
}
