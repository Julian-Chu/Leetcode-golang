package LeetCode_129_SumRootToLeafNumbers

import "github.com/Julian-Chu/Leetcode-golang/utils"

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

func sumNumbers_1(root *TreeNode) int {
	if root == nil {
		return 0
	}

	sub := make([]int, 0, 10)
	sum := 0
	var dfs func(*TreeNode, []int)
	dfs = func(node *TreeNode, sub []int) {
		sub = append(sub, node.Val)
		if node.Left == nil && node.Right == nil {
			num := 0

			for i := range sub {
				num = num*10 + sub[i]
			}
			sum += num
			return
		}

		if node.Left != nil {
			dfs(node.Left, sub)
		}

		if node.Right != nil {
			dfs(node.Right, sub)
		}
	}

	dfs(root, sub)
	return sum
}
