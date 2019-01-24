package LeetCode_226_InvertBinaryTree

import . "Leetcode-golang/helper"

func invertTree(root *TreeNode) *TreeNode {
	invertNodes(root)
	return root
}

func invertNodes(root *TreeNode) {
	if root == nil {
		return
	}

	invertNodes(root.Left)
	invertNodes(root.Right)
	root.Left, root.Right = root.Right, root.Left
}
