package LeetCode_701_InsertintoABinarySearchTree

import "github.com/Julian-Chu/Leetcode-golang/utils"

type TreeNode = utils.TreeNode

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}

	if root.Val > val {
		root.Left = insertIntoBST(root.Left, val)
	} else {
		root.Right = insertIntoBST(root.Right, val)
	}
	return root
}

func insertIntoBST_iteration(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	node := root
	var prev *TreeNode

	for node != nil {
		if val > node.Val {
			prev = node
			node = node.Right
		} else {
			prev = node
			node = node.Left
		}
	}

	if val > prev.Val {
		prev.Right = &TreeNode{Val: val}
	} else {
		prev.Left = &TreeNode{Val: val}
	}
	return root
}
