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
	var pre, cur *TreeNode

	cur = root

	for cur.Val != key {
		if key < cur.Val {
			if cur.Left == nil {
				return root
			}
			pre = cur
			cur = cur.Left
			continue
		}

		if key > cur.Val {
			if cur.Right == nil {
				return root
			}
			pre = cur
			cur = cur.Right
			continue
		}
	}
	if cur.Left == nil && cur.Right == nil {
		if pre == nil {
			return nil
		}

		if key > pre.Val {
			pre.Right = nil
		} else {
			pre.Left = nil
		}
		return root
	}

	var newNode *TreeNode
	if cur.Left != nil {
		newNode = cur.Left
		for newNode.Right != nil {
			if newNode.Right.Right == nil {
				tmp := newNode
				newNode = newNode.Right
				if newNode.Left != nil {
					tmp.Right = newNode.Left
					newNode.Left = nil
				}
			} else {
				newNode = newNode.Right
			}
		}
	} else {
		newNode = cur.Right
		for newNode.Left != nil {
			if newNode.Left.Left == nil {
				tmp := newNode
				newNode = newNode.Left
				if newNode.Right != nil {
					tmp.Left = newNode.Right
					newNode.Right = nil
				}
			} else {
				newNode = newNode.Left
			}
		}
	}

	if cur.Left != nil && cur.Left.Val != newNode.Val {
		newNode.Left = cur.Left
	}
	if cur.Right != nil && cur.Right.Val != newNode.Val {
		newNode.Right = cur.Right
	}

	if pre != nil {
		if pre.Val > newNode.Val {
			pre.Left = newNode
		} else {
			pre.Right = newNode
		}
	}

	if root.Val == key {
		return newNode
	} else {
		return root
	}

}
