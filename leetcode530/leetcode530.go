package leetcode530

import (
	"Leetcode-golang/utils"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode = utils.TreeNode

func getMinimumDifference(root *TreeNode) int {
	preVal := -1 << 31
	res := 1<<31 - 1

	var helper func(node *TreeNode)
	helper = func(node *TreeNode) {
		if node.Left != nil {
			helper(node.Left)
		}
		res = min(res, node.Val-preVal)
		preVal = node.Val
		if node.Right != nil {
			helper(node.Right)
		}
	}

	helper(root)
	return res
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
