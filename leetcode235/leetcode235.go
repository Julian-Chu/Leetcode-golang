package leetcode235

import "Leetcode-golang/utils"

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
