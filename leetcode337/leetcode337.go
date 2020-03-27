package leetcode337

import "Leetcode-golang/utils"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode = utils.TreeNode

func rob(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var dfs func(*TreeNode) (int, int)
	dfs = func(node *TreeNode) (withRoot int, woRoot int) {
		if node == nil {
			return 0, 0
		}

		lwr, lwor := dfs(node.Left)
		rwr, rwor := dfs(node.Right)

		return node.Val + lwor + rwor, max(lwr, lwor) + max(rwr, rwor)
	}

	return max(dfs(root))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
