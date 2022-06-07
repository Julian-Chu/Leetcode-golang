package leetcode617

import (
	"github.com/Julian-Chu/Leetcode-golang/utils"
)

type TreeNode = utils.TreeNode

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil && root2 == nil {
		return nil
	}

	if root1 == nil {
		return root2
	}

	if root2 == nil {
		return root1
	}

	root1.Val += root2.Val
	root1.Left = mergeTrees(root1.Left, root2.Left)
	root1.Right = mergeTrees(root1.Right, root2.Right)
	return root1
}

func mergeTrees_BFS(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil {
		return root2
	}

	if root2 == nil {
		return root1
	}

	queue := []*TreeNode{root1, root2}
	for len(queue) > 0 {
		node1 := queue[len(queue)-2]
		node2 := queue[len(queue)-1]
		queue = queue[:len(queue)-2]

		node1.Val += node2.Val
		if node1.Left != nil && node2.Left != nil {
			queue = append(queue, node1.Left, node2.Left)
		}

		if node1.Left == nil {
			node1.Left = node2.Left
		}

		if node1.Right != nil && node2.Right != nil {
			queue = append(queue, node1.Right, node2.Right)
		}

		if node1.Right == nil {
			node1.Right = node2.Right
		}
	}
	return root1
}
