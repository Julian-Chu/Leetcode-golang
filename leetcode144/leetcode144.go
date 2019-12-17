package leetcode144

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

func preorderTraversal(root *TreeNode) []int {
	var res []int
	return dfs(root, res)
}

func dfs(node *TreeNode, res []int) []int {
	if node == nil {
		return res
	}
	res = append(res, node.Val)

	if node.Left == nil && node.Right == nil {
		return res
	}
	if node.Left != nil {
		res = dfs(node.Left, res)
	}
	if node.Right != nil {
		res = dfs(node.Right, res)
	}
	return res
}
