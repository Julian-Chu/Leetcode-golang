package leetcode687

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

func longestUnivaluePath(root *TreeNode) int {
	maxLen := 0
	return helper(root, &maxLen)

}

func helper(node *TreeNode, maxLen *int) int {
	if node == nil {
		return 0
	}
	l := helper(node.Left, maxLen)
	r := helper(node.Right, maxLen)
	res := 0

	if node.Left != nil && node.Val == node.Left.Val {
		*maxLen = max(*maxLen, l+1)
		res = max(res, l+1)
	}

	if node.Right != nil && node.Val == node.Right.Val {
		*maxLen = max(*maxLen, r+1)
		res = max(res, r+1)
	}

	if node.Right != nil && node.Val == node.Right.Val &&
		node.Left != nil && node.Val == node.Left.Val {
		*maxLen = max(*maxLen, l+r+2)
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
