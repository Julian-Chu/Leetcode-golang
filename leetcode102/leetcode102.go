package leetcode102

import "Leetcode-golang/utils"

type TreeNode = utils.TreeNode

func levelOrder(root *TreeNode) [][]int {
	level := 0
	res := make([][]int, 0)
	BuildLevelOrderTraversal(root, level, &res)
	return res
}

func BuildLevelOrderTraversal(node *TreeNode, level int, res *[][]int) {
	if node == nil {
		return
	}
	if len(*res) == level {
		*res = append(*res, []int{})
	}

	(*res)[level] = append((*res)[level], node.Val)

	BuildLevelOrderTraversal(node.Left, level+1, res)
	BuildLevelOrderTraversal(node.Right, level+1, res)
}
