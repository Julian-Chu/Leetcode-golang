package LeetCode_101_SymmetricTree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return compareSymmetric(root.Left, root.Right)
}

func compareSymmetric(Left *TreeNode, Right *TreeNode) bool {
	if Left == nil && Right == nil {
		return true
	}
	if Left == nil || Right == nil {
		return false
	}

	return Left.Val == Right.Val && compareSymmetric(Left.Left, Right.Right) && compareSymmetric(Left.Right, Right.Left)

}
