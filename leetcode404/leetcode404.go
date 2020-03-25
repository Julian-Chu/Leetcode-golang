package leetcode404

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

func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}
	r := sumOfLeftLeaves(root.Right)
	if root.Left == nil {
		return r
	}
	if root.Left.Left == nil && root.Left.Right == nil {
		return root.Left.Val + r
	}
	return sumOfLeftLeaves(root.Left) + r
}
