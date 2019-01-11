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
	if s.Val == t.Val && compareTree(s, t) {
		return true
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

	if (sub == nil && t == nil) || (sub.Val == t.Val && sub.Left == nil && sub.Right == nil && t.Left == nil && t.Right == nil) {
		return true
	}
	if sub.Val != t.Val {
		return false
	}
	if !(sub.Left == nil && t.Left == nil) && (sub.Left == nil || t.Left == nil) {
		return false
	}
	if !(sub.Right == nil && t.Right == nil) && (sub.Right == nil || t.Right == nil) {
		return false
	}

	return compareTree(sub.Left, t.Left) && compareTree(sub.Right, t.Right)

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
