package leetcode222

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

func countNodes(root *TreeNode) int {
	cnt := 0

	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		cnt++
		dfs(node.Left)
		dfs(node.Right)
	}

	dfs(root)
	return cnt
}
