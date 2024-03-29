package leetcode404

import "github.com/Julian-Chu/Leetcode-golang/utils"

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
	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		return root.Left.Val + r
	}
	return sumOfLeftLeaves(root.Left) + r
}

func sumOfLeftLeaves_1(root *TreeNode) int {
	if root == nil {
		return 0
	}

	lsum := sumOfLeftLeaves(root.Left)
	rsum := sumOfLeftLeaves(root.Right)

	midsum := 0
	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		midsum = root.Left.Val
	}

	return lsum + rsum + midsum
}
