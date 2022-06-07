package LeetCode_235_LowestCommonAncestorOfABinarySearchTree

import "github.com/Julian-Chu/Leetcode-golang/utils"

type TreeNode = utils.TreeNode

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	for root != nil {
		val := root.Val
		if p.Val <= val && val <= q.Val || p.Val >= val && val >= q.Val {
			return root
		}

		if p.Val < val && q.Val < val {
			root = root.Left
		} else {
			root = root.Right
		}
	}

	return nil
}

func lowestCommonAncestor_recursion(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if p.Val > q.Val {
		p, q = q, p
	}

	if root.Val < p.Val && root.Left != nil {
		return lowestCommonAncestor(root.Right, p, q)
	}

	if root.Val > q.Val && root.Right != nil {
		return lowestCommonAncestor(root.Left, p, q)
	}

	return root
}
