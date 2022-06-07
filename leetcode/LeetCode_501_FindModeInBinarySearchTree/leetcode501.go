package LeetCode_501_FindModeInBinarySearchTree

import "github.com/Julian-Chu/Leetcode-golang/utils"

type TreeNode = utils.TreeNode

func findMode(root *TreeNode) []int {
	count := 1
	max := 1
	var res []int
	var prev *TreeNode
	var traverse func(*TreeNode)
	traverse = func(node *TreeNode) {
		if node == nil {
			return
		}

		traverse(node.Left)

		if prev != nil && prev.Val == node.Val {
			count++
		} else {
			count = 1
		}

		if count >= max {
			if count > max && len(res) > 0 {
				res = res[:0]
			}

			res = append(res, node.Val)
			max = count
		}
		prev = node

		traverse(node.Right)
	}
	traverse(root)
	return res
}
