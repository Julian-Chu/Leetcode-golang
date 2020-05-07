package leetcode515

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

func largestValues(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	var res []int
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		max := -1 << 31
		nextq := make([]*TreeNode, 0, 2*len(queue))
		for _, node := range queue {
			if node.Val > max {
				max = node.Val
			}

			if node.Left != nil {
				nextq = append(nextq, node.Left)
			}
			if node.Right != nil {
				nextq = append(nextq, node.Right)
			}
		}

		res = append(res, max)
		queue = nextq
	}
	return res
}
