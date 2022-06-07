package LeetCode_700_SearchInABinarySearchTree

import "github.com/Julian-Chu/Leetcode-golang/utils"

type TreeNode = utils.TreeNode

func searchBST_iteration(root *TreeNode, val int) *TreeNode {
	node := root

	for node != nil {
		switch {
		case node.Val > val:
			node = node.Left
		case node.Val < val:
			node = node.Right
		default:
			return node
		}
	}

	return node
}

// recursion
func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil || root.Val == val {
		return root
	}

	if root.Val > val {
		return searchBST(root.Left, val)
	}

	if root.Val < val {
		return searchBST(root.Right, val)
	}
	return nil
}
