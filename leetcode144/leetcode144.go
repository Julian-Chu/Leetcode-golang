package leetcode144

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

func preorderTraversal(root *TreeNode) []int {
	var res []int
	cur := root
	var stack []*TreeNode
	for cur != nil || len(stack) != 0 {
		if cur != nil {
			res = append(res, cur.Val)
			if cur.Left == nil && cur.Right == nil {
				cur = nil
				continue
			}
			if cur.Right != nil {
				stack = append(stack, cur.Right)
			}
			if cur.Left != nil {
				cur = cur.Left
			} else {
				cur = nil
			}
		} else if len(stack) != 0 {
			cur = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
	}
	return res
}
