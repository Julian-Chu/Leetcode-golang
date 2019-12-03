package leetcode98

import (
	"Leetcode-golang/utils"
)

type TreeNode = utils.TreeNode

func isValidBST(root *TreeNode) bool {
	return isValidSubBST(root, nil, nil)
}

func isValidSubBST(node *TreeNode, minNode *TreeNode, maxNode *TreeNode) bool {
	if node == nil {
		return true
	}

	if minNode != nil && node.Val <= minNode.Val {
		return false
	}

	if maxNode != nil && node.Val >= maxNode.Val {
		return false
	}

	if !isValidSubBST(node.Left, minNode, node) {
		return false
	}

	if !isValidSubBST(node.Right, node, maxNode) {
		return false
	}
	return true
}
