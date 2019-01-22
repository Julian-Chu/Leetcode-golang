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
	res := compareSymmetric(root.Left, root.Right)
	return res
}

func compareSymmetric(Left *TreeNode, Right *TreeNode) bool {

	if Left == nil {
		return Right == nil
	}
	if Right == nil {
		return false
	}
	if Left.Val != Right.Val {
		return false
	}

	return compareSymmetric(Left.Left, Right.Right) && compareSymmetric(Left.Right, Right.Left)

}
