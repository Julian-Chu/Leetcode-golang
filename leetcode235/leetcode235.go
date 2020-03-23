package leetcode235

import "Leetcode-golang/utils"

type TreeNode = utils.TreeNode

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	switch {
	case root.Val < p.Val && root.Val < q.Val:
		return lowestCommonAncestor(root.Right, p, q)
	case root.Val > p.Val && root.Val > q.Val:
		return lowestCommonAncestor(root.Left, p, q)
	default:
		return root
	}
}
