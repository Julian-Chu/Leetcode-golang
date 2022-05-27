package leetcode513

import "github.com/Julian-Chu/Leetcode-golang/utils"

type TreeNode = utils.TreeNode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findBottomLeftValue(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		root = queue[0]
		queue = queue[1:]

		if root.Right != nil {
			queue = append(queue, root.Right)
		}

		if root.Left != nil {
			queue = append(queue, root.Left)
		}
	}
	return root.Val
}
