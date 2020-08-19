package leetcode445

import "Leetcode-golang/utils"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode = utils.ListNode

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	stack1 := make([]*ListNode, 0, 100)
	stack2 := make([]*ListNode, 0, 100)

	for l1 != nil || l2 != nil {
		if l1 != nil {
			stack1 = append(stack1, l1)
			l1 = l1.Next
		}

		if l2 != nil {
			stack2 = append(stack2, l2)
			l2 = l2.Next
		}
	}

	var next *ListNode
	carry := 0
	for len(stack1) != 0 || len(stack2) != 0 || carry != 0 {
		val := 0
		switch {
		case len(stack1) != 0 && len(stack2) != 0:
			val = stack1[len(stack1)-1].Val + stack2[len(stack2)-1].Val
		case len(stack1) != 0:
			val = stack1[len(stack1)-1].Val
		case len(stack2) != 0:
			val = stack2[len(stack2)-1].Val
		}
		val += carry
		val, carry = val%10, val/10
		node := &ListNode{Val: val, Next: next}
		next = node
		if len(stack1) > 0 {
			stack1 = stack1[:len(stack1)-1]
		}
		if len(stack2) > 0 {
			stack2 = stack2[:len(stack2)-1]
		}
	}
	return next
}
