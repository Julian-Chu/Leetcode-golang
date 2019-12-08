package leetcode111

import "Leetcode-golang/utils"

type TreeNode = utils.TreeNode

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil {
		return 1
	}
	return findMinDepth(root, 1)
}

func findMinDepth(node *TreeNode, cur int) int {
	if node.Left == nil && node.Right == nil {
		return cur
	}
	l := cur
	r := cur
	if node.Left != nil {
		l = findMinDepth(node.Left, cur+1)
	}

	if node.Right != nil {
		r = findMinDepth(node.Right, cur+1)
	}

	if node.Right == nil {
		return l
	}
	if node.Left == nil {
		return r
	}
	if l > r {
		return r
	}
	return l
}
