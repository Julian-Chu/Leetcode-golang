package LeetCode_104_MaxDepthOfBinaryTree

import (
	"github.com/Julian-Chu/Leetcode-golang/utils"
)

func maxDepth(root *utils.TreeNode) int {
	if root == nil {
		return 0
	}

	return getMaxDepth(root)
}

func getMaxDepth(root *utils.TreeNode) int {
	if root == nil {
		return 0
	}

	ld, rd := 0, 0
	if root.Left != nil {
		ld = getMaxDepth(root.Left)
	}

	if root.Right != nil {
		rd = getMaxDepth(root.Right)
	}

	dep := ld + 1
	if rd > ld {
		dep = rd + 1
	}

	return dep
}
