package leetcode112

import "github.com/Julian-Chu/Leetcode-golang/utils"

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

func hasPathSum_1(root *TreeNode, targetSum int) bool {

	var dfs func(node *TreeNode, sum int) bool

	dfs = func(node *TreeNode, sum int) bool {
		if node == nil {
			return false
		}

		sum += node.Val
		if node.Left == nil && node.Right == nil {
			if sum == targetSum {
				return true
			}
		}

		return dfs(node.Left, sum) || dfs(node.Right, sum)
	}

	return dfs(root, 0)
}
