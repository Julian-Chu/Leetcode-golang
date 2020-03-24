package leetcode337

import "Leetcode-golang/utils"

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

	a := includeRoot(root)
	b := notIncludeRoot(root)

	if a > b {
		return a
	}

	return b
}

func notIncludeRoot(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var l1, l2, r1, r2 int
	l1 = includeRoot(root.Left)
	l2 = notIncludeRoot(root.Left)

	r1 = includeRoot(root.Right)
	r2 = notIncludeRoot(root.Right)

	l := l1
	if l2 > l1 {
		l = l2
	}
	r := r1
	if r2 > r1 {
		r = r2
	}
	return l + r
}

func includeRoot(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var l1, r1 int
	l1 = notIncludeRoot(root.Left)
	r1 = notIncludeRoot(root.Right)

	return l1 + r1 + root.Val
}
