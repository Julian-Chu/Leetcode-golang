package leetcode515

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

func largestValues(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}
	dfs(root, 0, &res)
	return res
}

func dfs(node *TreeNode, level int, res *[]int) {
	if len(*res) < level+1 {
		*res = append(*res, -1<<31)
	}

	(*res)[level] = max((*res)[level], node.Val)
	if node.Left != nil {
		dfs(node.Left, level+1, res)
	}

	if node.Right != nil {
		dfs(node.Right, level+1, res)
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
