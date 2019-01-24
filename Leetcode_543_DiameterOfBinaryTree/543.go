package Leetcode_543_DiameterOfBinaryTree

import . "leetcode-golang/helper"

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return getHeight(root.Left) + getHeight(root.Right)
}

func getHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}

	lh := getHeight(root.Left)
	rh := getHeight(root.Right)

	if lh > rh {
		return lh + 1
	} else {
		return rh + 1
	}

}
