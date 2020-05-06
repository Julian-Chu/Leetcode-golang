package leetcode513

import "Leetcode-golang/utils"

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
	var lastQueue []*TreeNode
	for len(queue) > 0 {
		nextQueue := make([]*TreeNode, 0, len(queue)*2)
		for _, node := range queue {
			if node.Left != nil {
				nextQueue = append(nextQueue, node.Left)
			}

			if node.Right != nil {
				nextQueue = append(nextQueue, node.Right)
			}
		}
		lastQueue = queue
		queue = nextQueue
	}
	return lastQueue[0].Val
}
