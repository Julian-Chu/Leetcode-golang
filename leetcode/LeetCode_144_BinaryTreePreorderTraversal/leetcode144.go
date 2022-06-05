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

func preorderTraversal_stack2(root *TreeNode) []int {
	res := make([]int, 0)
	stack := make([]*TreeNode, 0)

	if root == nil {
		return res
	}

	stack = append(stack, root)
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, node.Val)
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}
	return res
}

func preorderTraversal_template(root *TreeNode) []int {
	res := make([]int, 0)

	if root == nil {
		return res
	}

	stack := []*TreeNode{root}

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if node != nil {
			if node.Right != nil {
				stack = append(stack, node.Right)
			}
			if node.Left != nil {
				stack = append(stack, node.Left)
			}

			stack = append(stack, node)
			stack = append(stack, nil)
		} else {
			cur := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res = append(res, cur.Val)
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
