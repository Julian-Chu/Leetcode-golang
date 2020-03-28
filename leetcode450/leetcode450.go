package leetcode450

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

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}

	if key < root.Val {
		root.Left = deleteNode(root.Left, key)
	} else if key > root.Val {
		root.Right = deleteNode(root.Right, key)
	} else {
		if root.Left == nil {
			return root.Right
		} else if root.Right == nil {
			return root.Left
		}

		minNode := findMax(root.Left)
		root.Val = minNode.Val
		root.Left = deleteNode(root.Left, root.Val)
	}

	return root
}

func findMax(root *TreeNode) *TreeNode {
	for root.Right != nil {
		root = root.Right
	}
	return root
}

//func deleteNode(root *TreeNode, key int) *TreeNode {
//	if root == nil {
//		return nil
//	}
//
//	if key < root.Val {
//		root.Left = deleteNode(root.Left, key)
//	} else if key > root.Val {
//		root.Right = deleteNode(root.Right, key)
//	} else {
//		if root.Left == nil {
//			return root.Right
//		} else if root.Right == nil {
//			return root.Left
//		}
//
//		minNode := root.Right
//		for minNode.Left != nil {
//			minNode = minNode.Left
//		}
//
//		root.Val = minNode.Val
//		root.Right = deleteNode(root.Right, key)
//
//	}
//
//	return root
//}
