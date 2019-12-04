package leetcode103

import "Leetcode-golang/utils"

type TreeNode = utils.TreeNode

func zigzagLevelOrder(root *TreeNode) [][]int {
	res := buildZigZagLevelOrder(root, 0, [][]int{})

	for i := 0; i < len(res); i++ {
		if i%2 != 0 {
			for l, r := 0, len(res[i])-1; l < r; l, r = l+1, r-1 {
				res[i][l], res[i][r] = res[i][r], res[i][l]
			}
		}
	}
	return res
}

func buildZigZagLevelOrder(node *TreeNode, level int, res [][]int) [][]int {
	if node == nil {
		return res
	}

	if len(res) == level {
		res = append(res, []int{})
	}

	res[level] = append(res[level], node.Val)

	res = buildZigZagLevelOrder(node.Left, level+1, res)
	res = buildZigZagLevelOrder(node.Right, level+1, res)

	return res
}
