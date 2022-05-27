package LeetCode_226_InvertBinaryTree

import (
	"github.com/Julian-Chu/Leetcode-golang/utils"
)

func invertTree(root *utils.TreeNode) *utils.TreeNode {
	invertNodes(root)
	return root
}

func invertNodes(root *utils.TreeNode) {
	if root == nil {
		return
	}

	invertNodes(root.Left)
	invertNodes(root.Right)
	root.Left, root.Right = root.Right, root.Left
}
