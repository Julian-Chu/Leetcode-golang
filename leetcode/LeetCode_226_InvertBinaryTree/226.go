package LeetCode_226_InvertBinaryTree

import (
	"github.com/Julian-Chu/Leetcode-golang/utils"
)

func invertTree(root *utils.TreeNode) *utils.TreeNode {
	invertNodes(root)
	return root
}

func invertNodes(root *utils.TreeNode) {
	if root == nil {
		return
	}

	invertNodes(root.Left)
	invertNodes(root.Right)
	root.Left, root.Right = root.Right, root.Left
}

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func invertTree_DFS(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
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
			cur.Left, cur.Right = cur.Right, cur.Left
		}
	}
	return root
}
