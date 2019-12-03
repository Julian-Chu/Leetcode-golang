package leetcode98

import (
	"Leetcode-golang/utils"
)

type TreeNode = utils.TreeNode

func isValidBST(root *TreeNode) bool {
	return isValidSubBST(root, nil, nil)
}

func isValidSubBST(node *TreeNode, lower *int, upper *int) bool {
	if node == nil {
		return true
	}
	val := node.Val
	if lower != nil && val <= *lower {
		return false
	}

	if upper != nil && val >= *upper {
		return false
	}

	if !isValidSubBST(node.Left, lower, &val) {
		return false
	}

	if !isValidSubBST(node.Right, &val, upper) {
		return false
	}
	return true
}
