package leetcode98

import (
	"Leetcode-golang/utils"
)

type TreeNode = utils.TreeNode

func isValidBST(root *TreeNode) bool {
	min, max := -1<<63, 1<<63-1
	return isValidSubBST(min, max, root)
}

func isValidSubBST(min, max int, root *TreeNode) bool {
	if root == nil {
		return true
	}

	return min < root.Val && root.Val < max && isValidSubBST(min, root.Val, root.Left) && isValidSubBST(root.Val, max, root.Right)
}
