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

func findBottomLeftValue_DFS(root *TreeNode) int {
	if root == nil {
		return 0
	}

	maxDepth := 0
	ans := root.Val
	var traverse func(node *TreeNode, depth int)
	traverse = func(node *TreeNode, depth int) {
		if node == nil {
			return
		}
		if node.Left == nil && node.Right == nil {
			if depth > maxDepth {
				maxDepth = depth
				ans = node.Val
			}
			return
		}
		traverse(node.Left, depth+1)
		traverse(node.Right, depth+1)
	}
	traverse(root, 0)
	return ans
}
