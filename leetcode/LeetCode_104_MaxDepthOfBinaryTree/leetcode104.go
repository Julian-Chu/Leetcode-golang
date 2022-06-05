package LeetCode_104_MaxDepthOfBinaryTree

import (
	"github.com/Julian-Chu/Leetcode-golang/utils"
)

func maxDepth(root *utils.TreeNode) int {
	if root == nil {
		return 0
	}

	return getMaxDepth(root)
}

func getMaxDepth(root *utils.TreeNode) int {
	if root == nil {
		return 0
	}

	ld, rd := 0, 0
	if root.Left != nil {
		ld = getMaxDepth(root.Left)
	}

	if root.Right != nil {
		rd = getMaxDepth(root.Right)
	}

	dep := ld + 1
	if rd > ld {
		dep = rd + 1
	}

	return dep
}

type TreeNode = utils.TreeNode

func maxDepth_1(root *TreeNode) int {
	if root == nil {
		return 0
	}

	deep := 0
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		deep++
		size := len(queue)
		for i := 0; i < size; i++ {
			cur := queue[i]
			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}

			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
		}
		queue = queue[size:]
	}

	return deep
}
