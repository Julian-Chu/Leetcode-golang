package LeetCode_450_DeleteNodeInABST

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

func deleteNode(root *TreeNode, key int) *TreeNode {
	found := false
	return preorder(root, key, &found)
}

func preorder(root *TreeNode, key int, found *bool) *TreeNode {
	if root == nil {
		return nil
	}

	if !*found {
		if root.Val == key {
			*found = true
			if root.Right != nil {
				tmp := root.Right
				for tmp.Left != nil {
					tmp = tmp.Left
				}
				tmp.Left = root.Left
				return root.Right
			} else if root.Left != nil {
				tmp := root.Left
				for tmp.Right != nil {
					tmp = tmp.Right
				}
				tmp.Right = root.Right
				return root.Left
			}
			return nil
		} else if key > root.Val {
			root.Right = preorder(root.Right, key, found)
		} else {
			root.Left = preorder(root.Left, key, found)
		}
	}
	return root
}

func deleteNode_1(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val < key {
		root.Right = deleteNode(root.Right, key)
		return root
	} else if root.Val > key {
		root.Left = deleteNode(root.Left, key)
		return root
	}

	if root.Right == nil {
		return root.Left
	}

	if root.Left == nil {
		return root.Right
	}

	minnode := root.Right
	for minnode.Left != nil {
		minnode = minnode.Left
	}
	minnode.Left = root.Left

	root = root.Right
	return root
}
