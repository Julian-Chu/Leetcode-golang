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
	var dfs func(node *TreeNode, rest int)
	dfs = func(node *TreeNode, rest int) {
		if node == nil {
			return
		}

		rest -= node.Val
		if rest == 0 {
			res++
		}

		dfs(node.Left, rest)
		dfs(node.Right, rest)
	}
	dfs(root, sum)
	return res + pathSum(root.Left, sum) + pathSum(root.Right, sum)

}
