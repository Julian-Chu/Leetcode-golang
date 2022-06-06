package leetcode222

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

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return countNodes(root.Right) + countNodes(root.Left) + 1
}

func countNodes_completeBTree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	l := root.Left
	r := root.Right

	lDep := 0
	for l != nil {
		lDep++
		l = l.Left
	}
	rDep := 0
	for r != nil {
		rDep++
		r = r.Right
	}

	// complete binary tree
	if lDep == rDep {
		return 2<<lDep - 1
	}

	return countNodes(root.Left) + countNodes(root.Right) + 1
}
