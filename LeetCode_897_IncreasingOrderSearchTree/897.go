package leetcode897

import "Leetcode-golang/helper"

type TreeNode = helper.TreeNode

func increasingBST(root *TreeNode) *TreeNode {
	head := &TreeNode{Val: -1}
	tail := head
	reconnect(root, head)
	return tail
}

func reconnect(root, head *TreeNode) {
	if root == nil {
		return
	}

	if root.Left != nil {
		reconnect(root.Left, head)
	}

	if head.Val == -1 {
		head.Val = root.Val
	} else {
		for head.Right != nil {
			head = head.Right
		}
		head.Right = &TreeNode{
			Val: root.Val,
		}
		head = head.Right
	}
	if root.Right != nil {
		reconnect(root.Right, head)
	}
}
