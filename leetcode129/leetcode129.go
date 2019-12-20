package leetcode129

import "Leetcode-golang/utils"

type TreeNode = utils.TreeNode

func sumNumbers(root *TreeNode) int {
	return dfs(root, 0)
}

func dfs(node *TreeNode, sum int) int {
	if node == nil {
		return sum
	}
	sum *= 10
	sum += node.Val
	if node.Left == nil && node.Right == nil {
		return sum
	}
	res := 0
	if node.Left != nil {
		res += dfs(node.Left, sum)
	}
	if node.Right != nil {
		res += dfs(node.Right, sum)
	}
	return res
}
