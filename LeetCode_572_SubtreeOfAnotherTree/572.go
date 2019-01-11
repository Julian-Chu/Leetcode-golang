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
	if s.Val == t.Val {
		return compareTree(s, t)
	}
	if s.Left != nil && isSubtree(s.Left, t) {
		return true
	}
	if s.Right != nil && isSubtree(s.Right, t) {
		return true
	}
	return false
}

func compareTree(sub *TreeNode, t *TreeNode) bool {
	if sub != nil && t != nil && sub.Val != t.Val {
		return false
	}
	if (sub.Left == nil || t.Left == nil) && !(sub.Left == nil && t.Left == nil) || !compareTree(sub.Left, t.Left) {
		return false
	}

	if (sub.Right == nil || t.Right == nil) && !(sub.Right == nil && t.Right == nil) || !compareTree(sub.Right, t.Right) {
		return false
	}
	return true

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
