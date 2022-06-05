package LeetCode_144_BinaryTreePreorderTraversal

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

func preorderTraversal_stack(root *TreeNode) []int {
	var rightStack []*TreeNode
	var res []int

	for cur := root; cur != nil; {
		res = append(res, cur.Val)

		if cur.Left != nil {
			if cur.Right != nil {
				rightStack = append(rightStack, cur.Right)
			}
			cur = cur.Left
		} else {
			if cur.Right != nil {
				cur = cur.Right
			} else {
				if len(rightStack) == 0 {
					break
				}
				cur = rightStack[len(rightStack)-1]
				rightStack = rightStack[:len(rightStack)-1]
			}
		}
	}
	return res
}

func preorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	var traverse func(node *TreeNode)
	traverse = func(node *TreeNode) {
		if node == nil {
			return
		}
		res = append(res, node.Val)
		traverse(node.Left)
		traverse(node.Right)
	}

	traverse(root)
	return res
}
