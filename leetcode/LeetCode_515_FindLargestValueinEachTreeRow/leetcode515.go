package LeetCode_515_FindLargestValueinEachTreeRow

import (
	"github.com/Julian-Chu/Leetcode-golang/utils"
)

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
	res := make([]int, 0)
	if root == nil {
		return res
	}
	dfs(root, 0, &res)
	return res
}

func dfs(node *TreeNode, level int, res *[]int) {
	if len(*res) < level+1 {
		*res = append(*res, -1<<31)
	}

	(*res)[level] = max((*res)[level], node.Val)
	if node.Left != nil {
		dfs(node.Left, level+1, res)
	}

	if node.Right != nil {
		dfs(node.Right, level+1, res)
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func largestValues_BFS(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}

	queue := []*TreeNode{root}
	for len(queue) > 0 {
		size := len(queue)
		max := -1 << 31
		for i := 0; i < size; i++ {
			cur := queue[i]
			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
			if max < cur.Val {
				max = cur.Val
			}
		}
		queue = queue[size:]
		res = append(res, max)
	}
	return res
}
