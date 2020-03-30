package leetcode437

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

func pathSum(root *TreeNode, sum int) int {
	return helper(root, 0, sum, map[int]int{0: 1})
}

func helper(node *TreeNode, curSum, target int, m map[int]int) int {
	if node == nil {
		return 0
	}

	curSum += node.Val
	summary := m[curSum-target]
	m[curSum]++
	summary += helper(node.Left, curSum, target, m)
	summary += helper(node.Right, curSum, target, m)
	m[curSum]--
	return summary
}
