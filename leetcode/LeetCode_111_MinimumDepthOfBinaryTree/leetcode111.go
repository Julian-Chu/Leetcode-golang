package LeetCode_111_MinimumDepthOfBinaryTree

import "github.com/Julian-Chu/Leetcode-golang/utils"

type TreeNode = utils.TreeNode

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil {
		return 1
	}

	l := minDepth(root.Left)
	r := minDepth(root.Right)

	if root.Left == nil || root.Right == nil {
		if l > r {
			return l + 1
		}
		return r + 1
	}

	if l > r {
		return r + 1
	}
	return l + 1
}

func minDepth_BFS(root *TreeNode) int {
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
			if cur.Left == nil && cur.Right == nil {
				return deep
			}

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
