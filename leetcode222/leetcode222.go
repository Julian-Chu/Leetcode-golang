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
	if root == nil {
		return 0
	}

	return countNodes(root.Right) + countNodes(root.Left) + 1
}
