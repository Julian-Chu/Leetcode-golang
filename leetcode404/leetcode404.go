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
	var l, r int
	if root.Left != nil {
		if root.Left.Left == nil && root.Left.Right == nil {
			l = root.Left.Val
		} else {
			l = sumOfLeftLeaves(root.Left)
		}
	}

	if root.Right != nil {
		r = sumOfLeftLeaves(root.Right)
	}

	return l + r
}
