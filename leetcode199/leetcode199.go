package leetcode199

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

func rightSideView(root *TreeNode) []int {
	var res []int
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		size := len(queue)
		res = append(res, queue[size-1].Val)
		for i := range queue {
			node := queue[i]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		queue = queue[size:]
	}
	return res
}
