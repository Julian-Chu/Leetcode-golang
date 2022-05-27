package leetcode230

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

func kthSmallest(root *TreeNode, k int) int {
	var stack []*TreeNode
	for root != nil {
		stack = append(stack, root)
		root = root.Left
	}

	root = stack[len(stack)-1]
	stack = stack[:len(stack)-1]

	for i := 0; i < k-1; i++ {
		if root.Right != nil {
			root = root.Right
			for root != nil {
				stack = append(stack, root)
				root = root.Left
			}
			root = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		} else {
			root = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
	}

	return root.Val
}
