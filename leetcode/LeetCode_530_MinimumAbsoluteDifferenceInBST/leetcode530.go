package LeetCode_530_MinimumAbsoluteDifferenceInBST

import (
	"github.com/Julian-Chu/Leetcode-golang/utils"
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

func getMinimumDifference(root *TreeNode) int {
	preVal := -1 << 31
	res := 1<<31 - 1

	var helper func(node *TreeNode)
	helper = func(node *TreeNode) {
		if node.Left != nil {
			helper(node.Left)
		}
		res = min(res, node.Val-preVal)
		preVal = node.Val
		if node.Right != nil {
			helper(node.Right)
		}
	}

	helper(root)
	return res
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func getMinimumDifference_arr(root *TreeNode) int {
	var nums []int

	var traverse func(node *TreeNode)
	traverse = func(node *TreeNode) {
		if node == nil {
			return
		}

		traverse(node.Left)
		nums = append(nums, node.Val)
		traverse(node.Right)
	}
	traverse(root)

	ans := 1<<31 - 1
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			continue
		}
		diff := nums[i] - nums[i-1]
		if ans > diff {
			ans = diff
		}
	}
	return ans
}

func getMinimumDifference_inorder_traverse(root *TreeNode) int {
	var prev *TreeNode
	ans := 1<<31 - 1
	var traverse func(*TreeNode)
	traverse = func(node *TreeNode) {
		if node == nil {
			return
		}

		traverse(node.Left)
		if prev != nil {
			diff := node.Val - prev.Val
			if diff < ans {
				ans = diff
			}
		}
		prev = node
		traverse(node.Right)
	}
	traverse(root)
	return ans
}
