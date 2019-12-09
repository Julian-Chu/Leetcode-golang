package leetcdoe112

import "Leetcode-golang/utils"

type TreeNode = utils.TreeNode

func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	if root.Right == nil && root.Left == nil {
		if root.Val == sum {
			return true
		}
		return false
	}
	sum -= root.Val
	return hasPathSum(root.Left, sum) || hasPathSum(root.Right, sum)
}
