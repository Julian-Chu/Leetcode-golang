package leetcode437

import (
	"Leetcode-golang/utils"
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

func pathSum(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}

	res := 0
	iPath(sum, &res, []int{}, root)
	return res
}

func iPath(sum int, res *int, prev []int, root *TreeNode) {
	if root == nil {
		return
	}

	prev = append(prev, root.Val)
	s := 0
	for i := len(prev) - 1; i >= 0; i-- {
		s += prev[i]
		if sum == s {
			*res++
		}
	}
	iPath(sum, res, prev, root.Left)
	iPath(sum, res, prev, root.Right)
}
