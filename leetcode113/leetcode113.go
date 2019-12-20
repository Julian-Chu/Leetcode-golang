package leetcode113

import "Leetcode-golang/utils"

type TreeNode = utils.TreeNode

func pathSum(root *TreeNode, sum int) [][]int {
	if root == nil {
		return [][]int{}
	}

	var res [][]int
	res = recur(root, sum, res, []int{})
	return res
}

func recur(node *TreeNode, sum int, res [][]int, path []int) [][]int {
	if node == nil {
		return res
	}
	n := len(path)
	path = append(path[:n:n], node.Val)
	if node.Left == nil && node.Right == nil {
		if node.Val == sum {
			res = append(res, path)
		}
		return res
	}
	sum -= node.Val
	if node.Left != nil {
		res = recur(node.Left, sum, res, path)
	}
	if node.Right != nil {
		res = recur(node.Right, sum, res, path)
	}
	return res
}
