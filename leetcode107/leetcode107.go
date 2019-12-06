package leetcode107

import "Leetcode-golang/utils"

type TreeNode = utils.TreeNode

func levelOrderBottom(root *TreeNode) [][]int {
	var res [][]int

	res = rec(root, 0, res)
	for i := 0; i < len(res)/2; i++ {
		res[i], res[len(res)-1-i] = res[len(res)-1-i], res[i]
	}
	return res
}

func rec(node *TreeNode, level int, res [][]int) [][]int {
	if node == nil {
		return res
	}

	if len(res) <= level {
		res = append(res, []int{node.Val})
	} else {
		res[level] = append(res[level], node.Val)
	}

	res = rec(node.Left, level+1, res)
	res = rec(node.Right, level+1, res)

	return res
}
