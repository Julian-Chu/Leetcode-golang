package leetcode237

import "Leetcode-golang/utils"

type ListNode = utils.ListNode

func deleteNode(node *ListNode) {
	next := node.Next
	node.Val = next.Val
	node.Next = next.Next
	next.Next = nil
}
