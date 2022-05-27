package leetcode1019

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

func nextLargerNodes(head *ListNode) []int {
	if head == nil {
		return nil
	}

	input := make([]int, 0)
	for head != nil {
		input = append(input, head.Val)
		head = head.Next
	}

	res := make([]int, len(input))
	stack, peek := make([]int, len(input)), -1

	for i := 0; i < len(input); i++ {
		for peek >= 0 && input[stack[peek]] < input[i] {
			res[stack[peek]] = input[i]
			peek--
		}
		peek++
		stack[peek] = i
	}
	return res
}
