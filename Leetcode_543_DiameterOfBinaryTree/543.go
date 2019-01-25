package Leetcode_543_DiameterOfBinaryTree

import (
	. "leetcode-golang/helper"
)

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	_, res := getLenAndDia(root)
	return res
}

func getLenAndDia(root *TreeNode) (length, diameter int) {
	if root == nil {
		return
	}

	lh, ld := getLenAndDia(root.Left)
	rh, rd := getLenAndDia(root.Right)

	diameter = rd
	if ld > rd {
		diameter = ld
	}
	if lh+rh > diameter {
		diameter = lh + rh
	}

	if lh > rh {
		return lh + 1, diameter
	}
	return rh + 1, diameter

}
