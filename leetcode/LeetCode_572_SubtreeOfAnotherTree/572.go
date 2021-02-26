package leetcode572

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSubtree(s *TreeNode, t *TreeNode) bool {
	if compareTree(s, t) {
		return true
	}
	if s == nil {
		return false
	}
	return isSubtree(s.Left, t) || isSubtree(s.Right, t)
}

func compareTree(sub *TreeNode, t *TreeNode) bool {
	if sub == nil {
		return t == nil
	}

	if t == nil {
		return false
	}

	if sub.Val != t.Val {
		return false
	}
	return compareTree(sub.Left, t.Left) && compareTree(sub.Right, t.Right)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
