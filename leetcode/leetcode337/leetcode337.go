package leetcode337

import "github.com/Julian-Chu/Leetcode-golang/utils"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode = utils.TreeNode

func rob(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var dfs func(*TreeNode) (int, int)
	dfs = func(node *TreeNode) (withRoot int, woRoot int) {
		if node == nil {
			return 0, 0
		}

		lwr, lwor := dfs(node.Left)
		rwr, rwor := dfs(node.Right)

		return node.Val + lwor + rwor, max(lwr, lwor) + max(rwr, rwor)
	}

	return max(dfs(root))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func rob_dp(root *TreeNode) int {
	vals := robTree(root)
	return max(vals[0], vals[1])
}

func robTree(node *TreeNode) []int {
	if node == nil {
		return []int{0, 0}
	}

	left := robTree(node.Left)
	right := robTree(node.Right)

	robCur := node.Val + left[0] + right[0]
	notRobCur := max(left[0], left[1]) + max(right[0], right[1])
	return []int{notRobCur, robCur}
}

func rob_dfs(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil {
		return root.Val
	}

	val := root.Val
	if root.Left != nil {
		val += rob(root.Left.Left) + rob(root.Left.Right)
	}
	if root.Right != nil {
		val += rob(root.Right.Left) + rob(root.Right.Right)
	}

	val2 := rob(root.Left) + rob(root.Right)
	return max(val, val2)
}

func rob_memo(root *TreeNode) int {
	if val, ok := memo[root]; ok {
		return val
	}
	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil {
		return root.Val
	}

	val := root.Val
	if root.Left != nil {
		val += rob(root.Left.Left) + rob(root.Left.Right)
	}
	if root.Right != nil {
		val += rob(root.Right.Left) + rob(root.Right.Right)
	}

	val2 := rob(root.Left) + rob(root.Right)
	memo[root] = max(val, val2)
	return memo[root]
}

var memo = make(map[*TreeNode]int)
